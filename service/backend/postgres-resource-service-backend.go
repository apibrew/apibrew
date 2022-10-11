package backend

import (
	"context"
	"data-handler/stub/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq" // add this
	"strconv"
	"time"
)

const DbNameType = "VARCHAR(64)"

type postgresResourceServiceBackend struct {
	connectionMap map[string]*sql.DB
}

func (p *postgresResourceServiceBackend) Init(backend DataSourceBackend) {
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

func (p *postgresResourceServiceBackend) AddResource(params AddResourceParams) (*model.Resource, error) {
	err := p.withBackend(params.Backend, func(tx *sql.Tx) error {
		// check if resource exists
		res := tx.QueryRow("select count(*) as count from resource where name = $1", params.Resource.Name)

		var count = new(int)
		res.Scan(count)

		if *count > 0 {
			return errors.New("resource is already exists")
		}

		builder := sqlbuilder.CreateTable(params.Resource.SourceConfig.Mapping)

		builder.IfNotExists()

		builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")

		for _, property := range params.Resource.Properties {
			if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
				nullModifier := "NULL"
				if property.Required {
					nullModifier = "NOT NULL"
				}
				sqlType := getPsqlType(property.Type, property.Length)
				builder.Define(sourceConfig.Mapping, sqlType, nullModifier)
			}
		}

		// audit
		builder.Define("created_on", "timestamp", "NOT NULL")
		builder.Define("updated_on", "timestamp", "NULL")
		builder.Define("created_by", DbNameType, "NOT NULL")
		builder.Define("updated_by", DbNameType, "NULL")
		// version
		builder.Define("version", "int2", "NOT NULL")

		sqlQuery, _ := builder.Build()
		_, err := tx.Exec(sqlQuery)

		if err != nil {
			return err
		}

		if params.Resource.Flags == nil {
			params.Resource.Flags = &model.ResourceFlags{}
		}

		insertBuilder := sqlbuilder.InsertInto("resource")
		insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
		insertBuilder.Cols(
			"name",
			"workspace",
			"type",
			"source_data_source",
			"source_mapping",
			"read_only_records",
			"unique_record",
			"keep_history",
			"created_on",
			"updated_on",
			"created_by",
			"updated_by",
			"version")
		insertBuilder.Values(
			params.Resource.Name,
			params.Resource.Workspace,
			params.Resource.Type.Number(),
			params.Resource.SourceConfig.DataSource,
			params.Resource.SourceConfig.Mapping,
			params.Resource.Flags.ReadOnlyRecords,
			params.Resource.Flags.UniqueRecord,
			params.Resource.Flags.KeepHistory,
			time.Now(),
			nil,
			"test-usr",
			nil,
			1,
		)

		sql, args := insertBuilder.Build()

		_, err = tx.Exec(sql, args...)

		if err != nil {
			return err
		}

		for _, property := range params.Resource.Properties {
			propertyInsertBuilder := sqlbuilder.InsertInto("resource_property")
			propertyInsertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
			propertyInsertBuilder.Cols(
				"resource_name",
				"property_name",
				"type",
				"source_type",
				"source_mapping",
				"required",
				"length",
			)
			sourceType := 0
			propertyInsertBuilder.Values(
				params.Resource.Name,
				property.Name,
				property.Type,
				sourceType,
				property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping,
				property.Required,
				property.Length,
			)

			sql, args := propertyInsertBuilder.Build()

			_, err = tx.Exec(sql, args...)

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

func getPsqlType(propertyType model.ResourcePropertyType, length uint32) string {
	switch propertyType {
	case model.ResourcePropertyType_INT32:
		return "INT"
	case model.ResourcePropertyType_STRING:
		return "VARCHAR(" + strconv.Itoa(int(length)) + ")"
	default:

		panic("unknown property type")
	}
}

func (p *postgresResourceServiceBackend) acquireConnection(backend DataSourceBackend) (*sql.DB, error) {
	if p.connectionMap[backend.getDataSourceId()] == nil {
		bck := backend.(*postgresDataSourceBackend)

		connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", bck.Options.Username, bck.Options.Password, bck.Options.Host, bck.Options.Port, bck.Options.DbName)
		// Connect to database
		conn, err := sql.Open("postgres", connStr)

		if err != nil {
			return nil, err
		}

		p.connectionMap[backend.getDataSourceId()] = conn
	}

	return p.connectionMap[backend.getDataSourceId()], nil
}

func NewPostgresResourceServiceBackend() ResourceServiceBackend {
	return &postgresResourceServiceBackend{
		connectionMap: make(map[string]*sql.DB),
	}
}

func (p *postgresResourceServiceBackend) withBackend(backend DataSourceBackend, fn func(tx *sql.Tx) error) error {
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
