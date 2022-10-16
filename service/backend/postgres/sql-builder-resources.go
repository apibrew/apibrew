package postgres

import (
	"context"
	"data-handler/stub/model"
	"data-handler/util"
	"database/sql"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
	"time"
)

type QueryResultScanner interface {
	Scan(dest ...any) error
}

type QueryRunner interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

var resourceColumns = []string{
	"name",
	"workspace",
	"type",
	"source_data_source",
	"source_mapping",
	"read_only_records",
	"unique_record",
	"keep_history",
	"auto_created",
	"disable_migration",
	"disable_audit",
	"do_primary_key_lookup",
	"created_on",
	"updated_on",
	"created_by",
	"updated_by",
	"version",
}

var resourcePropertyColumns = []string{
	"resource_name",
	"property_name",
	"type",
	"source_type",
	"source_mapping",
	"source_def",
	"source_primary",
	"source_auto_generation",
	"required",
	"length",
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
		  auto_created boolean not null,
		  disable_migration boolean not null,
		  disable_audit boolean not null,
		  do_primary_key_lookup boolean not null,
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
		  source_def character varying(64),
		  source_primary bool,
		  source_auto_generation smallint,
		  required boolean,
		  length integer,
		  primary key (resource_name, property_name),
		  foreign key (resource_name) references public.resource (name) match simple on update cascade on delete cascade
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
			sqlType := getPsqlTypeFromProperty(property.Type, property.Length)
			builder.Define(sourceConfig.Mapping.Mapping, sqlType, nullModifier)
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

func resourcePrepareResourceFromEntity(ctx context.Context, runner QueryRunner, entity string) (resource *model.Resource, err error) {
	matchEntityName := func(ref string) string { return ref + `.table_schema || '.' || ` + ref + `.table_name = $1 ` }
	// check if entity exists
	row := runner.QueryRow(`select count(*) from information_schema.tables where table_type = 'BASE TABLE' and `+matchEntityName("tables"), entity)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var count = new(int)

	err = row.Scan(&count)

	if err != nil {
		return
	}

	if *count == 0 {
		err = errors.New("")
		return
	}

	resource = new(model.Resource)
	resource.Flags = new(model.ResourceFlags)
	resource.Flags.AutoCreated = true
	resource.Flags.DisableMigration = true
	resource.Flags.DisableAudit = true
	resource.AuditData = new(model.AuditData)
	resource.Type = model.DataType_USER
	resource.Name = strings.Replace(entity, ".", "_", -1)

	// properties

	rows, err := runner.Query(`select columns.column_name,
       columns.udt_name as column_type,
       columns.character_maximum_length as length,
       columns.is_nullable = 'YES' as is_nullable,
       column_key.constraint_def is not null as is_primary
from information_schema.columns
         left join information_schema.key_column_usage on key_column_usage.table_name = columns.table_name and
                                                          key_column_usage.table_schema = columns.table_schema and
                                                          key_column_usage.column_name = columns.column_name
         left join (SELECT nspname                     as table_schema,
                           conname,
                           contype,
                           pg_get_constraintdef(c.oid) as constraint_def
                    FROM pg_constraint c
                             JOIN pg_namespace n ON n.oid = c.connamespace
                    WHERE contype = 'p') column_key
                   on column_key.conname = key_column_usage.constraint_name
where `+matchEntityName("columns")+` order by columns.ordinal_position`, entity)

	if err != nil {
		return
	}

	primaryCount := 0
	for rows.Next() {
		var columnName = new(string)
		var columnType = new(string)
		var columnLength = new(*int)
		var isNullable = new(bool)
		var isPrimary = new(bool)

		err = rows.Scan(columnName, columnType, columnLength, isNullable, isPrimary)

		if err != nil {
			return
		}

		var sourceDef = *columnType

		if *columnLength != nil {
			sourceDef = *columnType + "(" + strconv.Itoa(**columnLength) + ")"
		} else {
			*columnLength = new(int)
			**columnLength = 0
		}

		if *isPrimary && !resource.Flags.DoPrimaryKeyLookup {
			primaryCount++

			if primaryCount > 1 {
				resource.Flags.DoPrimaryKeyLookup = true
			}

			if *columnName != "id" {
				resource.Flags.DoPrimaryKeyLookup = true
			}

			if *columnType != "uuid" {
				resource.Flags.DoPrimaryKeyLookup = true
			}
		}

		property := &model.ResourceProperty{
			Name: *columnName,
			Type: getPropertyTypeFromPsql(*columnType),
			SourceConfig: &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:   *columnName,
					SourceDef: sourceDef,
				},
			},
			Primary:  *isPrimary,
			Required: !*isNullable,
			Length:   uint32(**columnLength),
		}

		resource.Properties = append(resource.Properties, property)
	}

	return
}

func resourceInsert(runner QueryRunner, resource *model.Resource) error {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	insertBuilder := sqlbuilder.InsertInto("resource")
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	insertBuilder.Cols(resourceColumns...)
	insertBuilder.Values(
		resource.Name,
		resource.Workspace,
		resource.Type.Number(),
		resource.SourceConfig.DataSource,
		resource.SourceConfig.Mapping,
		resource.Flags.ReadOnlyRecords,
		resource.Flags.UniqueRecord,
		resource.Flags.KeepHistory,
		resource.Flags.AutoCreated,
		resource.Flags.DisableMigration,
		resource.Flags.DisableAudit,
		resource.Flags.DoPrimaryKeyLookup,
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
	selectBuilder := sqlbuilder.Select(resourceColumns...).
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
		&resource.Flags.AutoCreated,
		&resource.Flags.DisableMigration,
		&resource.Flags.DisableAudit,
		&resource.Flags.DoPrimaryKeyLookup,
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
		propertyInsertBuilder.Cols(resourcePropertyColumns...)
		sourceType := 0
		propertyInsertBuilder.Values(
			resource.Name,
			property.Name,
			property.Type,
			sourceType,
			property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.Mapping,
			property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.SourceDef,
			property.Primary,
			property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.AutoGeneration,
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

func resourceDelete(ctx context.Context, runner QueryRunner, ids []string) error {
	deleteBuilder := sqlbuilder.DeleteFrom("resource")
	deleteBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	deleteBuilder.Where(deleteBuilder.In("name", util.ArrayMapToInterface(ids)...))

	sqlQuery, args := deleteBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return err
	}

	return nil
}

func resourceLoadProperties(runner QueryRunner, resource *model.Resource, name string) error {
	selectBuilder := sqlbuilder.Select(resourcePropertyColumns...).From("resource_property")

	selectBuilder.Where(selectBuilder.Equal("resource_name", name))

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, args := selectBuilder.Build()

	rows, err := runner.Query(sqlQuery, args...)

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
		var sourceDef = new(string)
		var autoGeneration = new(int)
		err = rows.Scan(
			&resource.Name,
			&resourceProperty.Name,
			&resourceProperty.Type,
			&sourceType,
			&sourceMapping,
			&sourceDef,
			&resourceProperty.Primary,
			&autoGeneration,
			&resourceProperty.Required,
			&resourceProperty.Length,
		)

		if *sourceType == 0 {
			resourceProperty.SourceConfig = &model.ResourceProperty_Mapping{
				Mapping: &model.ResourcePropertyMappingConfig{
					Mapping:        *sourceMapping,
					SourceDef:      *sourceDef,
					AutoGeneration: model.AutoGenerationType(*autoGeneration),
				},
			}
		}

		if err != nil {
			return err
		}

		resource.Properties = append(resource.Properties, resourceProperty)
	}

	return nil
}
