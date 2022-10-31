package postgres

import (
	"context"
	"data-handler/model"
	"data-handler/service/backend"
	"data-handler/service/errors"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionMap     map[string]*sql.DB
	systemBackend     backend.DataSourceBackend
	dataSourceService backend.DataSourceLocator
}

func (p *postgresResourceServiceBackend) ListResources(ctx context.Context) (result []*model.Resource, err errors.ServiceError) {
	err = p.withSystemBackend(true, func(tx *sql.Tx) errors.ServiceError {
		result, err = resourceList(ctx, tx)

		return err
	})

	return
}

func (p *postgresResourceServiceBackend) ListEntities(ctx context.Context, dataSourceId string) (result []string, err errors.ServiceError) {
	err = p.withBackend(dataSourceId, true, func(tx *sql.Tx) errors.ServiceError {
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

	err := p.withBackend(p.systemBackend.GetDataSourceId(), false, func(tx *sql.Tx) errors.ServiceError {
		return resourceSetupTables(tx)
	})

	if err != nil {
		panic(err)
	}
}

func (p *postgresResourceServiceBackend) GetStatus(dataSourceId string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	connectionAlreadyInitiated = p.connectionMap[dataSourceId] != nil

	conn, err := p.acquireConnection(dataSourceId)

	if err != nil {
		return
	}

	err = handleDbError(conn.Ping())

	testConnection = err == nil

	return
}

func (p *postgresResourceServiceBackend) PrepareResourceFromEntity(ctx context.Context, dataSourceId string, entity string) (resource *model.Resource, err errors.ServiceError) {
	err = p.withBackend(dataSourceId, false, func(tx *sql.Tx) errors.ServiceError {
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

func (p *postgresResourceServiceBackend) GetResourceByName(ctx context.Context, workspace string, resourceName string) (resource *model.Resource, err errors.ServiceError) {
	resource = new(model.Resource)

	err = p.withSystemBackend(true, func(tx *sql.Tx) errors.ServiceError {
		if err = resourceLoadDetails(tx, resource, workspace, resourceName); err != nil {
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

func (p *postgresResourceServiceBackend) AddResource(params backend.AddResourceParams) (*model.Resource, errors.ServiceError) {
	if params.Resource.Workspace == "" {
		params.Resource.Workspace = "default"
	}

	err := p.withSystemBackend(false, func(tx *sql.Tx) errors.ServiceError {
		// check if resource exists

		if existingCount, err := resourceCountsByName(tx, params.Resource.Workspace, params.Resource.Name); err != nil {
			return err
		} else if existingCount > 0 {
			if params.IgnoreIfExists {
				return nil
			}
			return errors.AlreadyExistsError
		}

		if err := resourceInsert(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource", err)
			return err
		}

		if err := resourceUpsertProperties(tx, params.Resource); err != nil {
			log.Error("Unable to insert resource properties", err)
			return err
		}

		if params.Migrate {
			err := p.withBackend(params.Resource.SourceConfig.DataSource, false, func(tx *sql.Tx) errors.ServiceError {
				err := resourceCreateTable(tx, params.Resource)
				if err != nil {
					return err
				}

				if params.Resource.Flags.KeepHistory {
					if params.Resource.Flags.DisableAudit {
						return errors.LogicalError.WithMessage("history cannot created while audit is disabled")
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

func (p *postgresResourceServiceBackend) UpdateResource(ctx context.Context, resource *model.Resource, doMigration bool, forceMigration bool) errors.ServiceError {
	var err errors.ServiceError
	return p.withSystemBackend(false, func(tx *sql.Tx) errors.ServiceError {
		if err = resourceUpdate(ctx, tx, resource); err != nil {
			return err
		}

		doResourceCleanup(resource)

		if err = resourceUpsertProperties(tx, resource); err != nil {
			return err
		}

		if doMigration {
			err = p.withBackend(resource.SourceConfig.DataSource, false, func(tx *sql.Tx) errors.ServiceError {
				if err = resourceCreateTable(tx, resource); err != nil {
					return err
				}

				if err = resourceMigrateTable(ctx, tx, resource, forceMigration); err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *postgresResourceServiceBackend) DeleteResources(ctx context.Context, workspace string, ids []string, doMigration bool, forceMigration bool) errors.ServiceError {
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
			return errors.NotFoundError
		}

		sources = append(sources, resource.SourceConfig)
	}
	err := p.withSystemBackend(false, func(tx *sql.Tx) errors.ServiceError {
		return resourceDelete(ctx, tx, ids)
	})

	if err != nil {
		return err
	}

	if doMigration {
		for _, source := range sources {
			err = p.withBackend(source.DataSource, false, func(tx *sql.Tx) errors.ServiceError {
				return resourceDropTable(tx, source.Mapping)
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *postgresResourceServiceBackend) acquireConnection(dataSourceId string) (*sql.DB, errors.ServiceError) {
	if p.connectionMap[dataSourceId] == nil {
		dsBck, err := p.dataSourceService.GetDataSourceBackendById(dataSourceId)

		if err != nil {
			return nil, err
		}

		bck := dsBck.(*postgresDataSourceBackend)

		connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", bck.Options.Username, bck.Options.Password, bck.Options.Host, bck.Options.Port, bck.Options.DbName)
		// Connect to database
		conn, sqlErr := sql.Open("postgres", connStr)
		err = handleDbError(sqlErr)

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

func (p *postgresResourceServiceBackend) withSystemBackend(readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	return p.withBackend(p.systemBackend.GetDataSourceId(), readOnly, fn)
}

func (p *postgresResourceServiceBackend) withBackend(dataSourceId string, readOnly bool, fn func(tx *sql.Tx) errors.ServiceError) errors.ServiceError {
	log.Tracef("begin transaction: %s, readonly=%v", dataSourceId, readOnly)
	conn, serviceErr := p.acquireConnection(dataSourceId)

	if serviceErr != nil {
		return serviceErr
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{
		ReadOnly: readOnly,
	})

	if err != nil {
		log.Errorf("Unable to begin transaction: %s %s", err, dataSourceId)
		return handleDbError(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	serviceErr = fn(tx)

	if serviceErr != nil {
		log.Errorf("Rollback: %s", serviceErr)
		return serviceErr
	}

	serviceErr = handleDbError(tx.Commit())
	log.Tracef("end transaction: %s, readonly=%v", dataSourceId, readOnly)

	return serviceErr
}

func handleDbError(err error) errors.ServiceError {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return errors.NotFoundError
	}

	if err == sql.ErrTxDone {
		log.Panic("Illegal situation")
	}

	if _, ok := err.(errors.ServiceError); ok {
		log.Panic("database error is expected: ", err)
	}

	if pqErr, ok := err.(*pq.Error); ok {
		return handlePqErr(pqErr)
	}

	if netErr, ok := err.(*net.OpError); ok {
		return errors.InternalError.WithDetails(netErr.Error())
	}

	panic("Unhandled situation")
}

func handlePqErr(err *pq.Error) errors.ServiceError {
	switch err.Code {
	case "28000":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "28P01":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "23505":
		return errors.UniqueViolation.WithDetails(err.Message)
	default:
		return errors.InternalError.WithMessage(err.Message)
	}
}
