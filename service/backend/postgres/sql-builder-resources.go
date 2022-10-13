package postgres

import (
	"data-handler/stub/model"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

type QueryRunner interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

func resourceSetupTables(runner QueryRunner) error {
	_, err := runner.Exec(`
		create table if not exists public.resource (
		  name character varying(64) primary key not null,
		  workspace character varying(64) not null,
		  type smallint not null,
		  source_data_source character varying(64) not null,
		  source_mapping character varying(64) not null,
		  read_only_records boolean not null,
		  unique_record boolean not null,
		  keep_history boolean not null,
		  created_on timestamp without time zone not null,
		  updated_on timestamp without time zone,
		  created_by character varying(64) not null,
		  updated_by character varying(64),
		  version integer not null
		);
		
		create table if not exists public.resource_property (
		  resource_name character varying(64) not null,
		  property_name character varying(64) not null,
		  type smallint,
		  source_type smallint,
		  source_mapping character varying(64),
		  required boolean,
		  length integer,
		  primary key (resource_name, property_name)
		);
`)

	return err
}

func resourceCountsByName(runner QueryRunner, resourceName string) (int, error) {
	res := runner.QueryRow("select count(*) as count from resource where name = $1", resourceName)

	var count = new(int)
	err := res.Scan(count)

	return *count, err
}

func resourceCreateTable(runner QueryRunner, resource *model.Resource) error {
	builder := sqlbuilder.CreateTable(resource.SourceConfig.Mapping)

	builder.IfNotExists()

	builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")

	for _, property := range resource.Properties {
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
	_, err := runner.Exec(sqlQuery)

	return err
}

func resourceInsert(runner QueryRunner, resource *model.Resource) error {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
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
		resource.Name,
		resource.Workspace,
		resource.Type.Number(),
		resource.SourceConfig.DataSource,
		resource.SourceConfig.Mapping,
		resource.Flags.ReadOnlyRecords,
		resource.Flags.UniqueRecord,
		resource.Flags.KeepHistory,
		time.Now(),
		nil,
		"test-usr",
		nil,
		1,
	)

	sqlQuery, args := insertBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	return err
}

func resourceLoadDetails(runner QueryRunner, resource *model.Resource, name string) error {
	selectBuilder := sqlbuilder.Select(
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
		"version",
	).
		From("resource").
		Where("name='" + name + "'")

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, _ := selectBuilder.Build()

	row := runner.QueryRow(sqlQuery)

	if row.Err() != nil {
		return row.Err()
	}

	resource.SourceConfig = &model.ResourceSourceConfig{}
	resource.Flags = &model.ResourceFlags{}
	resource.AuditData = &model.AuditData{}

	var createdOn = new(time.Time)
	var updatedOn = new(*time.Time)
	var updatedBy = new(*string)

	err := row.Scan(
		&resource.Name,
		&resource.Workspace,
		&resource.Type,
		&resource.SourceConfig.DataSource,
		&resource.SourceConfig.Mapping,
		&resource.Flags.ReadOnlyRecords,
		&resource.Flags.UniqueRecord,
		&resource.Flags.KeepHistory,
		createdOn,
		updatedOn,
		&resource.AuditData.CreatedBy,
		updatedBy,
		&resource.Version,
	)

	if err != nil {
		return err
	}

	resource.AuditData.CreatedOn = timestamppb.New(*createdOn)
	if *updatedOn != nil {
		resource.AuditData.UpdatedOn = timestamppb.New(**updatedOn)
	}
	if *updatedBy != nil {
		resource.AuditData.UpdatedBy = **updatedBy
	}

	return nil
}

func resourceInsertProperties(runner QueryRunner, resource *model.Resource) error {
	for _, property := range resource.Properties {
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
			resource.Name,
			property.Name,
			property.Type,
			sourceType,
			property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping,
			property.Required,
			property.Length,
		)

		sql, args := propertyInsertBuilder.Build()

		_, err := runner.Exec(sql, args...)

		if err != nil {
			return err
		}
	}

	return nil
}

func resourceLoadProperties(runner QueryRunner, resource *model.Resource, name string) error {
	selectBuilder := sqlbuilder.Select(
		"resource_name",
		"property_name",
		"type",
		"source_type",
		"source_mapping",
		"required",
		"length",
	).
		From("resource_property").
		Where("resource_name='" + name + "'")

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, _ := selectBuilder.Build()

	rows, err := runner.Query(sqlQuery)

	if err != nil {
		return err
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	for rows.Next() {
		resourceProperty := new(model.ResourceProperty)

		var sourceType = new(int)
		var sourceMapping = new(string)
		err = rows.Scan(
			&resource.Name,
			&resourceProperty.Name,
			&resourceProperty.Type,
			&sourceType,
			&sourceMapping,
			&resourceProperty.Required,
			&resourceProperty.Length,
		)

		if *sourceType == 0 {
			resourceProperty.SourceConfig = &model.ResourceProperty_Mapping{
				Mapping: *sourceMapping,
			}
		}

		if err != nil {
			return err
		}

		resource.Properties = append(resource.Properties, resourceProperty)
	}

	return nil
}

func dereferenceProperty(value interface{}, propertyType model.ResourcePropertyType, required bool) interface{} {
	switch propertyType {
	case model.ResourcePropertyType_INT32:
		return *value.(*int32)
	case model.ResourcePropertyType_STRING:
		return *value.(*string)
	default:
		panic("unknown property type")
	}
}

func getPropertyPointer(propertyType model.ResourcePropertyType, required bool) interface{} {
	switch propertyType {
	case model.ResourcePropertyType_INT32:
		return new(int32)
	case model.ResourcePropertyType_STRING:
		return new(string)
	default:
		panic("unknown property type")
	}
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
