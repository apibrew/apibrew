package postgres

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/util"
	"database/sql"
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
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
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

var resourceColumnMapFn = func(column string, resource *model.Resource) interface{} {
	switch column {
	case "name":
		return resource.Name
	case "workspace":
		return resource.Workspace
	case "type":
		return resource.Type
	case "source_data_source":
		return resource.SourceConfig.DataSource
	case "source_mapping":
		return resource.SourceConfig.Mapping
	case "read_only_records":
		return resource.Flags.ReadOnlyRecords
	case "unique_record":
		return resource.Flags.UniqueRecord
	case "keep_history":
		return resource.Flags.KeepHistory
	case "auto_created":
		return resource.Flags.AutoCreated
	case "disable_migration":
		return resource.Flags.DisableMigration
	case "disable_audit":
		return resource.Flags.DisableAudit
	case "do_primary_key_lookup":
		return resource.Flags.DoPrimaryKeyLookup
	case "created_on":
		return resource.AuditData.CreatedOn.AsTime()
	case "updated_on":
		if resource.AuditData.UpdatedOn == nil {
			return nil
		}
		return resource.AuditData.UpdatedOn.AsTime()
	case "created_by":
		return resource.AuditData.CreatedBy
	case "updated_by":
		return resource.AuditData.UpdatedBy
	case "version":
		return resource.Version
	default:
		panic("Unknown column: " + column)
	}
}

var resourcePropertyColumns = []string{
	"workspace",
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
	"\"unique\"",
}

var resourcePropertyColumnMapFn = func(column string, resource *model.Resource, property *model.ResourceProperty) interface{} {
	sourceType := 0

	switch column {
	case "workspace":
		return resource.Workspace
	case "resource_name":
		return resource.Name
	case "property_name":
		return property.Name
	case "type":
		return property.Type
	case "source_type":
		return sourceType
	case "source_mapping":
		return property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.Mapping
	case "source_def":
		return property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.SourceDef
	case "source_primary":
		return property.Primary
	case "source_auto_generation":
		return property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.AutoGeneration
	case "required":
		return property.Required
	case "length":
		return property.Length
	case "\"unique\"":
		return property.Unique
	default:
		panic("Unknown column: " + column)
	}
}

func valuesFromCol[T interface{}](columns []string, data T, mapperFn func(col string, data T) interface{}) []interface{} {
	return util.ArrayMap[string, interface{}](columns, func(col string) interface{} {
		return mapperFn(col, data)
	})
}

func resourceSetupTables(runner QueryRunner) errors.ServiceError {
	_, err := runner.Exec(`
		create table if not exists public.resource (
		  name character varying(64) not null,
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
		  version integer not null,
		  PRIMARY KEY(name, workspace)
		);
		
		create table if not exists public.resource_property (
		  workspace     character varying(64) not null,
		  resource_name character varying(64) not null,
		  property_name character varying(64) not null,
		  type smallint,
		  source_type smallint,
		  source_mapping character varying(64),
		  source_def character varying(64),
		  source_primary bool,
		  source_auto_generation smallint,
		  required boolean,
		  "unique" boolean,
		  length integer,
		  primary key (workspace, resource_name, property_name),
		  foreign key (workspace, resource_name) references public.resource (workspace, name) match simple on update cascade on delete cascade
		);
`)

	return handleDbError(err)
}

func resourceCountsByName(runner QueryRunner, workspace, resourceName string) (int, errors.ServiceError) {
	res := runner.QueryRow("select count(*) as count from resource where name = $1 and workspace = $2", resourceName, workspace)

	var count = new(int)
	err := res.Scan(count)

	return *count, handleDbError(err)
}

func resourceAlterTable(ctx context.Context, runner QueryRunner, existingResource *model.Resource, resource *model.Resource) errors.ServiceError {
	panic("not implemented yet")
}

func resourceCreateTable(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig.Mapping, false))

	builder.IfNotExists()

	if !resource.Flags.AutoCreated && !resource.Flags.DoPrimaryKeyLookup {
		builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")
	}

	prepareCreateTableQuery(resource, builder)

	// audit
	if !resource.Flags.DisableAudit {
		builder.Define("created_on", "timestamp", "NOT NULL")
		builder.Define("updated_on", "timestamp", "NULL")
		builder.Define("created_by", DbNameType, "NOT NULL")
		builder.Define("updated_by", DbNameType, "NULL")
		// version
		builder.Define("version", "int2", "NOT NULL")
	}

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return handleDbError(err)
}

func prepareCreateTableQuery(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) {
	if resource.Flags == nil {
		resource.Flags = new(model.ResourceFlags)
	}

	var primaryKeys []string
	for _, property := range resource.Properties {
		if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			uniqModifier := ""
			nullModifier := "NULL"
			if property.Required {
				nullModifier = "NOT NULL"
			}
			if property.Unique {
				uniqModifier = "UNIQUE"
			}
			sqlType := getPsqlTypeFromProperty(property.Type, property.Length)
			builder.Define(sourceConfig.Mapping.Mapping, sqlType, nullModifier, uniqModifier)

			if property.Primary {
				primaryKeys = append(primaryKeys, sourceConfig.Mapping.Mapping)
			}
		}
	}
	if len(primaryKeys) > 0 {
		builder.Define("PRIMARY KEY (" + strings.Join(primaryKeys, ",") + ")")
	}
}

func resourceCreateHistoryTable(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig.Mapping, true))

	builder.IfNotExists()

	builder.Define("id", "uuid", "NOT NULL")

	prepareCreateTableQuery(resource, builder)

	builder.Define("created_on", "timestamp", "NOT NULL")
	builder.Define("updated_on", "timestamp", "NULL")
	builder.Define("created_by", DbNameType, "NOT NULL")
	builder.Define("updated_by", DbNameType, "NULL")
	// version
	builder.Define("version", "int2", "NOT NULL")

	builder.Define("PRIMARY KEY (id, version)")

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return handleDbError(err)
}

func resourceDropTable(runner QueryRunner, mapping string) errors.ServiceError {
	_, err := runner.Exec("DROP TABLE " + mapping)

	return handleDbError(err)
}

func resourceListEntities(ctx context.Context, runner QueryRunner) (result []string, err errors.ServiceError) {
	rows, sqlErr := runner.QueryContext(ctx, `select table_schema || '.' || table_name from information_schema.tables`)
	err = handleDbError(sqlErr)

	if err != nil {
		return
	}

	for rows.Next() {
		var entityName = new(string)

		err = handleDbError(rows.Scan(entityName))

		if err != nil {
			return
		}

		result = append(result, *entityName)
	}

	return
}

func resourceList(ctx context.Context, runner QueryRunner) (result []*model.Resource, err errors.ServiceError) {
	selectBuilder := sqlbuilder.Select(resourceColumns...).From("resource")

	sqlQuery, args := selectBuilder.Build()

	rows, sqlErr := runner.QueryContext(ctx, sqlQuery, args...)

	err = handleDbError(sqlErr)

	if err != nil {
		return
	}

	for rows.Next() {
		resource := new(model.Resource)
		err = ScanResource(resource, rows)

		if err != nil {
			return nil, err
		}

		result = append(result, resource)
	}

	return
}

func resourcePrepareResourceFromEntity(ctx context.Context, runner QueryRunner, entity string) (resource *model.Resource, err errors.ServiceError) {
	matchEntityName := func(ref string) string { return ref + `.table_schema || '.' || ` + ref + `.table_name = $1 ` }
	// check if entity exists
	row := runner.QueryRowContext(ctx, `select count(*) from information_schema.tables where table_type = 'BASE TABLE' and `+matchEntityName("tables"), entity)

	if row.Err() != nil {
		return nil, handleDbError(row.Err())
	}

	var count = new(int)

	err = handleDbError(row.Scan(&count))

	if err != nil {
		return
	}

	if *count == 0 {
		err = errors.NotFoundError
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
	resource.Workspace = "default"

	// properties

	rows, sqlErr := runner.QueryContext(ctx, `select columns.column_name,
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
	err = handleDbError(sqlErr)

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

		err = handleDbError(rows.Scan(columnName, columnType, columnLength, isNullable, isPrimary))

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

func resourceInsert(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	if resource.AuditData == nil {
		resource.AuditData = &model.AuditData{}
	}
	resource.AuditData.CreatedOn = timestamppb.New(time.Now())
	resource.AuditData.CreatedBy = "test-usr"

	insertBuilder := sqlbuilder.InsertInto("resource")
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	insertBuilder.Cols(resourceColumns...)
	insertBuilder.Values(valuesFromCol(resourceColumns, resource, resourceColumnMapFn)...)

	sqlQuery, args := insertBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	return handleDbError(err)
}

func resourceLoadDetails(runner QueryRunner, resource *model.Resource, workspace string, resourceName string) errors.ServiceError {
	selectBuilder := sqlbuilder.Select(resourceColumns...).
		From("resource")
	selectBuilder.Where(selectBuilder.And(
		selectBuilder.Equal("name", resourceName),
		selectBuilder.Equal("workspace", workspace),
	))

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, args := selectBuilder.Build()

	row := runner.QueryRow(sqlQuery, args...)

	if row.Err() != nil {
		return handleDbError(row.Err())
	}

	err := ScanResource(resource, row)

	if err != nil {
		return err
	}

	return nil
}

func ScanResource(resource *model.Resource, row QueryResultScanner) errors.ServiceError {
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
		return handleDbError(err)
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

func resourceInsertProperties(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	for _, property := range resource.Properties {
		propertyInsertBuilder := sqlbuilder.InsertInto("resource_property")
		propertyInsertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
		propertyInsertBuilder.Cols(resourcePropertyColumns...)
		propertyInsertBuilder.Values(util.ArrayMap(resourcePropertyColumns, func(col string) interface{} {
			return resourcePropertyColumnMapFn(col, resource, property)
		})...)

		sqlQuery, args := propertyInsertBuilder.Build()

		_, err := runner.Exec(sqlQuery, args...)

		if err != nil {
			return handleDbError(err)
		}
	}

	return nil
}

func resourceUpdateProperties(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	for _, property := range resource.Properties {
		propertyInsertBuilder := sqlbuilder.Update("resource_property")
		propertyInsertBuilder.SetFlavor(sqlbuilder.PostgreSQL)

		for _, col := range resourcePropertyColumns {
			propertyInsertBuilder.SetMore(propertyInsertBuilder.Equal(col, resourcePropertyColumnMapFn(col, resource, property)))
		}

		sqlQuery, args := propertyInsertBuilder.Build()

		_, err := runner.ExecContext(ctx, sqlQuery, args...)

		if err != nil {
			return handleDbError(err)
		}
	}

	return nil
}

func resourceUpdate(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	updateBuilder := sqlbuilder.Update("resource")
	updateBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	const ResourceFieldName = "name"

	updateBuilder.Where(updateBuilder.Equal(ResourceFieldName, resource.Name))

	for _, col := range resourceColumns {
		updateBuilder.SetMore(updateBuilder.Equal(col, resourceColumnMapFn(col, resource)))
	}

	sqlQuery, args := updateBuilder.Build()

	_, err := runner.ExecContext(ctx, sqlQuery, args...)

	return handleDbError(err)
}

func resourceDelete(ctx context.Context, runner QueryRunner, ids []string) errors.ServiceError {
	deleteBuilder := sqlbuilder.DeleteFrom("resource")
	deleteBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	deleteBuilder.Where(deleteBuilder.In("name", util.ArrayMapToInterface(ids)...))

	sqlQuery, args := deleteBuilder.Build()

	_, err := runner.ExecContext(ctx, sqlQuery, args...)

	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func resourceLoadProperties(runner QueryRunner, resource *model.Resource, workspace string, resourceName string) errors.ServiceError {
	selectBuilder := sqlbuilder.Select(resourcePropertyColumns...).From("resource_property")

	selectBuilder.Where(selectBuilder.And(
		selectBuilder.Equal("resource_name", resourceName),
		selectBuilder.Equal("workspace", workspace),
	))

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, args := selectBuilder.Build()

	rows, err := runner.Query(sqlQuery, args...)

	if err != nil {
		return handleDbError(err)
	}

	if rows.Err() != nil {
		return handleDbError(rows.Err())
	}

	for rows.Next() {
		resourceProperty := new(model.ResourceProperty)

		var sourceType = new(int)
		var sourceMapping = new(string)
		var sourceDef = new(string)
		var autoGeneration = new(int)
		err = rows.Scan(
			&resource.Workspace,
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
			&resourceProperty.Unique,
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
			return handleDbError(err)
		}

		resource.Properties = append(resource.Properties, resourceProperty)
	}

	return nil
}
