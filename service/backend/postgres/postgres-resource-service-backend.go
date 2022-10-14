package postgres

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub"
	"data-handler/stub/model"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionMap     map[string]*sql.DB
	systemBackend     backend.DataSourceBackend
	dataSourceService backend.DataSourceLocator
}

func (p *postgresResourceServiceBackend) InjectDataSourceService(dataSourceService backend.DataSourceLocator) {
	p.dataSourceService = dataSourceService
}

func (p *postgresResourceServiceBackend) DestroyDataSource(dataSourceId string) {
	if p.connectionMap[dataSourceId] != nil {
		err := p.connectionMap[dataSourceId].Close()

		if err != nil {
			log.Println("Cannot Close destroyed datasource connection: ", err)
		}

		delete(p.connectionMap, dataSourceId)
	}
}

func (p *postgresResourceServiceBackend) Init() {
	p.systemBackend = p.dataSourceService.GetSystemDataSourceBackend()

	err := p.withBackend(p.systemBackend.GetDataSourceId(), func(tx *sql.Tx) error {
		return resourceSetupTables(tx)
	})

	if err != nil {
		panic(err)
	}
}

func (p *postgresResourceServiceBackend) GetStatus(dataSourceId string) (result *stub.StatusResponse, err error) {
	result = new(stub.StatusResponse)

	result.ConnectionAlreadyInitiated = p.connectionMap[dataSourceId] != nil

	conn, err := p.acquireConnection(dataSourceId)

	if err != nil {
		return
	}

	err = conn.Ping()

	result.TestConnection = err == nil

	return
}

func (p *postgresResourceServiceBackend) GetResourceByName(resourceName string) (resource *model.Resource, err error) {
	resource = new(model.Resource)

	err = p.withSystemBackend(func(tx *sql.Tx) error {
		if err = resourceLoadDetails(tx, resource, resourceName); err != nil {
			log.Error("Unable to load resource details", err)
			return err
		}

		if err := resourceLoadProperties(tx, resource, resourceName); err != nil {
			log.Error("Unable to load resource properties", err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Error("Unable load resource", err)
		return nil, err
	}

	return resource, nil
}

func (p *postgresResourceServiceBackend) AddResource(params backend.AddResourceParams) (*model.Resource, error) {
	err := p.withBackend(params.Resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		// check if resource exists

		if existingCount, err := resourceCountsByName(tx, params.Resource.Name); err != nil {
			return err
		} else if existingCount > 0 {
			if params.IgnoreIfExists {
				return nil
			}
			return errors.New("resource is already exists")
		}

		if params.Migrate {
			if err := resourceCreateTable(tx, params.Resource); err != nil {
				log.Error("Unable to create resource table", err)
				return err
			}
		}

		if err := resourceInsert(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource", err)
			return err
		}

		if err := resourceInsertProperties(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource properties", err)
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return params.Resource, nil
}

func (p *postgresResourceServiceBackend) acquireConnection(dataSourceId string) (*sql.DB, error) {
	if p.connectionMap[dataSourceId] == nil {
		dsBck, err := p.dataSourceService.GetDataSourceBackendById(dataSourceId)

		if err != nil {
			return nil, err
		}

		bck := dsBck.(*postgresDataSourceBackend)

		connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", bck.Options.Username, bck.Options.Password, bck.Options.Host, bck.Options.Port, bck.Options.DbName)
		// Connect to database
		conn, err := sql.Open("postgres", connStr)

		if err != nil {
			return nil, err
		}

		p.connectionMap[dataSourceId] = conn

		log.Info("Connected to Datasource: ", dataSourceId)
	}

	return p.connectionMap[dataSourceId], nil
}

func NewPostgresResourceServiceBackend() backend.ResourceServiceBackend {
	return &postgresResourceServiceBackend{
		connectionMap: make(map[string]*sql.DB),
	}
}

func (p *postgresResourceServiceBackend) withSystemBackend(fn func(tx *sql.Tx) error) error {
	return p.withBackend(p.systemBackend.GetDataSourceId(), fn)
}

func (p *postgresResourceServiceBackend) withBackend2(dataSourceId string, fn func(tx *sql.Tx)) {
	_ = p.withBackend(dataSourceId, func(tx *sql.Tx) error {
		fn(tx)

		return nil
	})
}

func (p *postgresResourceServiceBackend) withBackend(dataSourceId string, fn func(tx *sql.Tx) error) error {
	conn, err := p.acquireConnection(dataSourceId)

	if err != nil {
		log.Error("Unable to acquire connection", err, dataSourceId)
		return err
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{})

	if err != nil {
		log.Error("Unable to begin transaction", err, dataSourceId)
		return err
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	err = fn(tx)

	if err != nil {
		log.Error("Unable to execute code inside transaction", err)

		return err
	}

	return tx.Commit()
}
