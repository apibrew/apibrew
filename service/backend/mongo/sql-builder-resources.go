package mongo

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/util"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	"id",
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
	case "id":
		return resource.Id
	case "name":
		return resource.Name
	case "workspace":
		return resource.Workspace
	case "type":
		return resource.DataType
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

var resourceReferenceColumns = []string{
	"workspace",
	"resource_name",
	"property_name",
	"referenced_resource",
	"\"cascade\"",
}

var resourceReferenceColumnMapFn = func(column string, resource *model.Resource, property *model.ResourceReference) interface{} {
	switch column {
	case "workspace":
		return resource.Workspace
	case "resource_name":
		return resource.Name
	case "property_name":
		return property.PropertyName
	case "referenced_resource":
		return property.ReferencedResource
	case "\"cascade\"":
		return property.Cascade
	default:
		panic("Unknown column: " + column)
	}
}

func valuesFromCol[T interface{}](columns []string, data T, mapperFn func(col string, data T) interface{}) []interface{} {
	return util.ArrayMap[string, interface{}](columns, func(col string) interface{} {
		return mapperFn(col, data)
	})
}

func resourceCountsByName(runner QueryRunner, workspace, resourceName string) (int, errors.ServiceError) {
	res := runner.QueryRow("select count(*) as count from resource where name = $1 and workspace = $2", resourceName, workspace)

	var count = new(int)
	err := res.Scan(count)

	return *count, handleDbError(err)
}

func resourceCreateTable(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig.Mapping, false))

	builder.IfNotExists()

	if !resource.Flags.DoPrimaryKeyLookup {
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

type ReferenceLocalDetails struct {
	sourceTableName       string
	fkConstraintName      string
	sourceTableColumn     string
	referencedTable       string
	referencedTableColumn string
	joinAlias             string
}

func resolveReferenceDetails(runner QueryRunner, resource *model.Resource, reference *model.ResourceReference) (*ReferenceLocalDetails, errors.ServiceError) {
	var err errors.ServiceError

	// locate referenced resource table name
	var referencedResource = new(model.Resource)
	err = resourceLoadDetailsByName(runner, referencedResource, resource.Workspace, reference.ReferencedResource)
	if err != nil {
		return nil, handleDbError(err)
	}
	if referencedResource.Flags.DoPrimaryKeyLookup {
		return nil, errors.LogicalError.WithMessage("referenced resource is not allowed for non standard primary keys")
	}
	referencedTable := referencedResource.SourceConfig.Mapping

	// locate referenced property name
	property := locatePropertyByName(resource, reference.PropertyName)
	if property == nil {
		return nil, errors.PropertyNotFoundError
	}

	var sourceTableColumn string = ""
	if propertySource, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
		sourceTableColumn = propertySource.Mapping.Mapping
	} else {
		return nil, errors.PropertyNotFoundError.WithDetails("Property is not with mapping type")
	}

	referencedTableColumn := "id"
	sourceTableName := resource.SourceConfig.Mapping
	fkConstraintName := "fk_" + reference.PropertyName
	joinAlias := "l_" + referencedTable

	return &ReferenceLocalDetails{
		joinAlias:             joinAlias,
		sourceTableName:       sourceTableName,
		fkConstraintName:      fkConstraintName,
		sourceTableColumn:     sourceTableColumn,
		referencedTable:       referencedTable,
		referencedTableColumn: referencedTableColumn,
	}, nil
}

func resourceCreateForeignKey(runner QueryRunner, resource *model.Resource, reference *model.ResourceReference) errors.ServiceError {
	referenceLocalDetails, err := resolveReferenceDetails(runner, resource, reference)

	if err != nil {
		return nil
	}

	query := fmt.Sprintf(`ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)`, referenceLocalDetails.sourceTableName, referenceLocalDetails.fkConstraintName, referenceLocalDetails.sourceTableColumn, referenceLocalDetails.referencedTable, referenceLocalDetails.referencedTableColumn)

	_, sqlErr := runner.Exec(query)

	return handleDbError(sqlErr)
}

func prepareCreateTableQuery(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) {
	if resource.Flags == nil {
		resource.Flags = new(model.ResourceFlags)
	}

	var primaryKeys []string
	for _, property := range resource.Properties {
		columnDef := prepareResourceTableColumnDefinition(property)

		if columnDef != "" {
			builder.Define(columnDef)
		}

		if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			if property.Primary {
				primaryKeys = append(primaryKeys, sourceConfig.Mapping.Mapping)
			}
		}
	}
	if len(primaryKeys) > 0 {
		builder.Define("PRIMARY KEY (" + strings.Join(primaryKeys, ",") + ")")
	}
}

func prepareResourceTableColumnDefinition(property *model.ResourceProperty) string {
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

		var def = []string{sourceConfig.Mapping.Mapping, sqlType, nullModifier, uniqModifier}

		return strings.Join(def, " ")
	}

	return ""
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

func isAuditColumn(column string) bool {
	return column == "created_on" || column == "updated_on" || column == "created_by" || column == "updated_by" || column == "version"
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

func resourceLoadDetails(runner QueryRunner, resource *model.Resource, workspace string, id string) errors.ServiceError {
	selectBuilder := sqlbuilder.Select(resourceColumns...).
		From("resource")
	selectBuilder.Where(selectBuilder.And(
		selectBuilder.Equal("id", id),
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

func resourceLoadDetailsByName(runner QueryRunner, resource *model.Resource, workspace string, resourceName string) errors.ServiceError {
	selectBuilder := sqlbuilder.Select(resourceColumns...).
		From("resource")
	selectBuilder.Where(selectBuilder.And(
		selectBuilder.Equal("name", resourceName),
		selectBuilder.Equal("workspace", workspace),
	))

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, args := selectBuilder.Build()

	log.Tracef("SQL: %s", sqlQuery)

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
		&resource.Id,
		&resource.Name,
		&resource.Workspace,
		&resource.DataType,
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

func resourceUpsertProperties(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	for _, property := range resource.Properties {
		propertyInsertBuilder := sqlbuilder.InsertInto("resource_property")
		propertyInsertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
		propertyInsertBuilder.Cols(resourcePropertyColumns...)
		propertyInsertBuilder.Values(util.ArrayMap(resourcePropertyColumns, func(col string) interface{} {
			return resourcePropertyColumnMapFn(col, resource, property)
		})...)

		propertyInsertBuilder.SQL("ON CONFLICT(workspace, resource_name, property_name) DO UPDATE SET")
		var updates []string
		for _, col := range resourcePropertyColumns {
			updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
		}
		propertyInsertBuilder.SQL(strings.Join(updates, ","))

		sqlQuery, args := propertyInsertBuilder.Build()

		log.Tracef("SQL: %s", sqlQuery)

		_, err := runner.Exec(sqlQuery, args...)

		if err != nil {
			return handleDbError(err)
		}
	}

	return nil
}

func resourceUpsertReferences(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	for _, property := range resource.References {
		propertyInsertBuilder := sqlbuilder.InsertInto("resource_reference")
		propertyInsertBuilder.SetFlavor(sqlbuilder.PostgreSQL)
		propertyInsertBuilder.Cols(resourceReferenceColumns...)
		propertyInsertBuilder.Values(util.ArrayMap(resourceReferenceColumns, func(col string) interface{} {
			return resourceReferenceColumnMapFn(col, resource, property)
		})...)

		propertyInsertBuilder.SQL("ON CONFLICT(workspace, resource_name, property_name) DO UPDATE SET")
		var updates []string
		for _, col := range resourceReferenceColumns {
			updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
		}
		propertyInsertBuilder.SQL(strings.Join(updates, ","))

		sqlQuery, args := propertyInsertBuilder.Build()

		log.Tracef("SQL: %s", sqlQuery)

		_, err := runner.Exec(sqlQuery, args...)

		if err != nil {
			return handleDbError(err)
		}
	}

	deleteQueryBuilder := sqlbuilder.DeleteFrom("resource_property")
	deleteQueryBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	deleteQueryBuilder.Where(deleteQueryBuilder.And(
		deleteQueryBuilder.Equal("workspace", resource.GetWorkspace()),
		deleteQueryBuilder.Equal("resource_name", resource.GetName()),
		deleteQueryBuilder.NotIn("property_name", util.ArrayMapToInterface(util.ArrayMap(resource.GetProperties(), func(item *model.ResourceProperty) string { return item.Name }))...),
	))

	deleteQuery, args := deleteQueryBuilder.Build()

	_, err := runner.Exec(deleteQuery, args...)

	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func resourceUpdate(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	updateBuilder := sqlbuilder.Update("resource")
	updateBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	if resource.AuditData == nil {
		resource.AuditData = &model.AuditData{}
	}

	updateBuilder.Where(updateBuilder.And(
		updateBuilder.Equal("id", resource.Id),
		updateBuilder.Equal("workspace", resource.Workspace),
	))

	for _, col := range resourceColumns {
		if col == "created_on" {
			continue
		}

		if col == "created_by" {
			continue
		}
		updateBuilder.SetMore(updateBuilder.Equal(col, resourceColumnMapFn(col, resource)))
	}

	sqlQuery, args := updateBuilder.Build()

	log.Tracef("SQL: %s", sqlQuery)

	_, err := runner.ExecContext(ctx, sqlQuery, args...)

	return handleDbError(err)
}

func resourceDelete(ctx context.Context, runner QueryRunner, ids []string) errors.ServiceError {
	deleteBuilder := sqlbuilder.DeleteFrom("resource")
	deleteBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	deleteBuilder.Where(deleteBuilder.In("id", util.ArrayMapToInterface(ids)...))

	sqlQuery, args := deleteBuilder.Build()

	log.Tracef("SQL: %s", sqlQuery)

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

	log.Tracef("SQL: %s", sqlQuery)

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

func resourceLoadReferences(runner QueryRunner, resource *model.Resource, workspace string, resourceName string) errors.ServiceError {
	selectBuilder := sqlbuilder.Select(resourceReferenceColumns...).From("resource_reference")

	selectBuilder.Where(selectBuilder.And(
		selectBuilder.Equal("resource_name", resourceName),
		selectBuilder.Equal("workspace", workspace),
	))

	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	sqlQuery, args := selectBuilder.Build()

	log.Tracef("SQL: %s", sqlQuery)

	rows, err := runner.Query(sqlQuery, args...)

	if err != nil {
		return handleDbError(err)
	}

	if rows.Err() != nil {
		return handleDbError(rows.Err())
	}

	for rows.Next() {
		resourceReference := new(model.ResourceReference)

		err = rows.Scan(
			&resource.Workspace,
			&resource.Name,
			&resourceReference.PropertyName,
			&resourceReference.ReferencedResource,
			&resourceReference.Cascade,
		)

		if err != nil {
			return handleDbError(err)
		}

		resource.References = append(resource.References, resourceReference)
	}

	return nil
}
