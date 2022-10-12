package postgres

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub/model"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionMap map[string]*sql.DB
	systemBackend backend.DataSourceBackend
}

func (p *postgresResourceServiceBackend) Init(systemBackend backend.DataSourceBackend) {
	p.systemBackend = systemBackend
}

func (p *postgresResourceServiceBackend) GetResourceByName(resourceName string) (*model.Resource, error) {
	var resource = new(model.Resource)

	err := p.withBackend(p.systemBackend, func(tx *sql.Tx) error {
		if err := resourceLoadDetails(tx, resource, resourceName); err != nil {
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
	err := p.withBackend(params.Backend, func(tx *sql.Tx) error {
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

func (p *postgresResourceServiceBackend) acquireConnection(backend backend.DataSourceBackend) (*sql.DB, error) {
	if p.connectionMap[backend.GetDataSourceId()] == nil {
		bck := backend.(*postgresDataSourceBackend)

		connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", bck.Options.Username, bck.Options.Password, bck.Options.Host, bck.Options.Port, bck.Options.DbName)
		// Connect to database
		conn, err := sql.Open("postgres", connStr)

		if err != nil {
			return nil, err
		}

		p.connectionMap[backend.GetDataSourceId()] = conn

		log.Info("Connected to Datasource: ", backend.GetDataSourceId())
	}

	return p.connectionMap[backend.GetDataSourceId()], nil
}

func NewPostgresResourceServiceBackend() backend.ResourceServiceBackend {
	return &postgresResourceServiceBackend{
		connectionMap: make(map[string]*sql.DB),
	}
}

func (p *postgresResourceServiceBackend) withBackend(backend backend.DataSourceBackend, fn func(tx *sql.Tx) error) error {
	conn, err := p.acquireConnection(backend)

	if err != nil {
		log.Error("Unable to acquire connection", err, backend.GetDataSourceId())
		return err
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{})

	if err != nil {
		log.Error("Unable to begin transaction", err, backend.GetDataSourceId())
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
