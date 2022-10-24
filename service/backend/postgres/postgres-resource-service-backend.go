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

func (p *postgresResourceServiceBackend) ListResources(ctx context.Context) (result []*model.Resource, err error) {
	err = p.withSystemBackend(true, func(tx *sql.Tx) error {
		result, err = resourceList(ctx, tx)

		return err
	})

	return
}

func (p *postgresResourceServiceBackend) ListEntities(ctx context.Context, dataSourceId string) (result []string, err error) {
	err = p.withBackend(dataSourceId, true, func(tx *sql.Tx) error {
		result, err = resourceListEntities(ctx, tx)

		return err
	})

	return
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

	err := p.withBackend(p.systemBackend.GetDataSourceId(), false, func(tx *sql.Tx) error {
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

func (p *postgresResourceServiceBackend) PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (resource *model.Resource, err error) {
	err = p.withBackend(dataSourceId, false, func(tx *sql.Tx) error {
		if resource, err = resourcePrepareResourceFromEntity(ctx, tx, entity); err != nil {
			log.Error("Unable to load resource details", err)
			return err
		}

		resource.SourceConfig = &model.ResourceSourceConfig{
			DataSource: dataSourceId,
			Mapping:    entity,
		}

		return nil
	})

	if err != nil {
		log.Error("Unable load resource", err)
		return nil, err
	}

	return resource, nil
}

func (p *postgresResourceServiceBackend) GetResourceByName(ctx context.Context, workspace string, resourceName string) (resource *model.Resource, err error) {
	resource = new(model.Resource)

	err = p.withSystemBackend(true, func(tx *sql.Tx) error {
		if err = resourceLoadDetails(tx, resource, workspace, resourceName); err != nil {
			if err == sql.ErrNoRows {
				resource = nil
				return nil
			}

			log.Error("Unable to load resource details", err)
			return err
		}

		if err = resourceLoadProperties(tx, resource, workspace, resourceName); err != nil {
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
	if params.Resource.Workspace == "" {
		params.Resource.Workspace = "default"
	}

	err := p.withSystemBackend(false, func(tx *sql.Tx) error {
		// check if resource exists

		if existingCount, err := resourceCountsByName(tx, params.Resource.Workspace, params.Resource.Name); err != nil {
			return err
		} else if existingCount > 0 {
			if params.IgnoreIfExists {
				return nil
			}
			return errors.New("resource is already exists")
		}

		if err := resourceInsert(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource", err)
			return err
		}

		if err := resourceInsertProperties(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource properties", err)
			return err
		}

		if params.Migrate {
			err := p.withBackend(params.Resource.SourceConfig.DataSource, false, func(tx *sql.Tx) error {
				err := resourceCreateTable(tx, params.Resource)
				if err != nil {
					return err
				}

				if params.Resource.Flags.KeepHistory {
					if params.Resource.Flags.DisableAudit {
						return errors.New("history cannot created while audit is disabled")
					}
					err = resourceCreateHistoryTable(tx, params.Resource)
					if err != nil {
						return err
					}
				}
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return params.Resource, nil
}

func (p *postgresResourceServiceBackend) UpdateResource(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) error {
	existingResource, err := p.GetResourceByName(nil, resource.Name, resource.Workspace)

	if err != nil {
		return err
	}

	return p.withSystemBackend(false, func(tx *sql.Tx) error {
		if err = resourceUpdate(ctx, tx, resource); err != nil {
			return err
		}

		if err = resourceUpdateProperties(ctx, tx, resource); err != nil {
			return err
		}

		if doMigration {
			err = p.withBackend(existingResource.SourceConfig.DataSource, false, func(tx *sql.Tx) error {
				return resourceAlterTable(ctx, tx, existingResource, resource)
			})

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *postgresResourceServiceBackend) DeleteResources(ctx context.Context, workspace string, ids []string, doMigration bool, forceMigration bool) error {
	var sources []*model.ResourceSourceConfig

	if workspace == "" {
		workspace = "default"
	}

	for _, id := range ids {
		resource, err := p.GetResourceByName(nil, workspace, id)

		if err != nil {
			return err
		}

		if resource == nil {
			return errors.New("resource not found:" + workspace + " " + id)
		}

		sources = append(sources, resource.SourceConfig)
	}
	err := p.withSystemBackend(false, func(tx *sql.Tx) error {
		return resourceDelete(ctx, tx, ids)
	})

	if err != nil {
		return err
	}

	if doMigration {
		for _, source := range sources {
			err = p.withBackend(source.DataSource, false, func(tx *sql.Tx) error {
				return resourceDropTable(tx, source.Mapping)
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
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

func (p *postgresResourceServiceBackend) withSystemBackend(readOnly bool, fn func(tx *sql.Tx) error) error {
	return p.withBackend(p.systemBackend.GetDataSourceId(), readOnly, fn)
}

func (p *postgresResourceServiceBackend) withBackend(dataSourceId string, readOnly bool, fn func(tx *sql.Tx) error) error {
	conn, err := p.acquireConnection(dataSourceId)

	if err != nil {
		log.Error("Unable to acquire connection", err, dataSourceId)
		return err
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		log.Error("Unable to begin transaction", err, dataSourceId)
		return err
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	err = fn(tx)

	if err == sql.ErrNoRows {
		return err
	}

	if err != nil {
		log.Error("Rollback; Error: ", err)
		//debug.PrintStack()
		return err
	}

	return tx.Commit()
}
