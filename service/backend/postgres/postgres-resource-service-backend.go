package postgres

import (
	"context"
	"data-handler/service/backend"
	"data-handler/stub/model"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq" // add this
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionMap map[string]*sql.DB
}

func (p *postgresResourceServiceBackend) Init(backend backend.DataSourceBackend) {
	//err := p.withBackend(backend, func(tx *sql.Tx) error {
	//	builder := sqlbuilder.CreateTable("resources")
	//
	//	builder.IfNotExists()
	//
	//	// basic properties
	//	builder.Define("name", DbNameType, "NOT NULL", "PRIMARY KEY")
	//	builder.Define("workspace", DbNameType, "NOT NULL")
	//	builder.Define("type", "int2", "NOT NULL")
	//	builder.Define("source_data_source", DbNameType, "NOT NULL")
	//	builder.Define("source_mapping", DbNameType, "NOT NULL")
	//	// flags
	//	builder.Define("read_only_records", "bool", "NOT NULL")
	//	builder.Define("unique_record", "bool", "NOT NULL")
	//	builder.Define("keep_history", "bool", "NOT NULL")
	//	// audit
	//	builder.Define("created_on", "timestamp", "NOT NULL")
	//	builder.Define("updated_on", "timestamp", "NULL")
	//	builder.Define("created_by", DbNameType, "NOT NULL")
	//	builder.Define("updated_by", DbNameType, "NULL")
	//	// version
	//	builder.Define("version", "int2", "NOT NULL")
	//
	//	_, err := tx.Exec(builder.Build())
	//
	//	return err
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
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
				return err
			}
		}

		if err := resourceInsert(tx, params.Resource); err != nil {
			return err
		}

		if err := resourceInsertProperties(tx, params.Resource); err != nil {
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
		return err
	}

	tx, err := conn.BeginTx(context.TODO(), &sql.TxOptions{})

	if err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	err = fn(tx)

	if err != nil {
		return err
	}

	return tx.Commit()
}
