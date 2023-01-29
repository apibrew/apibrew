package postgres

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations2 "github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"strconv"
	"strings"
)

func resourceMigrateTable(ctx context.Context, runner QueryRunner, params backend.UpgradeResourceParams, history bool) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	var err errors.ServiceError
	var existingResource *model.Resource

	entityName := params.Resource.SourceConfig.Entity
	if history {
		entityName = entityName + "_h"
	}

	if existingResource, err = resourcePrepareResourceFromEntity(ctx, runner, params.Resource.SourceConfig.Catalog, entityName); err != nil {
		logger.Error("Unable to load resource details: ", err)
		return err
	}

	var notChangedColumns = make(map[string]bool)
	var changedColumns = make(map[string]bool)
	var newPrevMap = make(map[*model.ResourceProperty]*model.ResourceProperty)
	var removedColumns = make(map[string]bool)
	var newColumns = make(map[string]bool)
	var changesCount = 0

	// check left
	for _, existingProperty := range existingResource.Properties {
		existingColName := getPropertyColumnName(existingProperty)

		found := false
		for _, newProperty := range params.Resource.Properties {
			newColName := getPropertyColumnName(newProperty)
			if existingColName == newColName {
				if util.IsSameResourceProperty(existingProperty, newProperty) {
					notChangedColumns[existingColName] = true
				} else {
					changedColumns[existingColName] = true
					newPrevMap[newProperty] = existingProperty
				}
				found = true
			}
		}
		if !found {
			removedColumns[existingColName] = true
		}
	}

	// check right
	for _, newProperty := range params.Resource.Properties {
		newColName := getPropertyColumnName(newProperty)

		found := false
		for _, existingProperty := range existingResource.Properties {
			existingColName := getPropertyColumnName(existingProperty)
			if existingColName == newColName {
				found = true
			}
		}
		if !found {
			newColumns[newColName] = true
		}
	}

	if len(changedColumns) == 0 && len(newColumns) == 0 && (!params.ForceMigration || len(removedColumns) == 0) {
		// no need to migration
		return nil
	}

	// create new properties
	var alterTableQuery = fmt.Sprintf(`ALTER TABLE %s`, getTableName(params.Resource.GetSourceConfig(), history))

	var alterTableQueryDefs []string

	for _, property := range params.Resource.Properties {
		colName := getPropertyColumnName(property)
		if !newColumns[colName] {
			continue
		}

		alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("ADD COLUMN %s", prepareResourceTableColumnDefinition(property)))
		changesCount++
	}

	// delete properties (IF FORCE MIGRATION)
	if params.ForceMigration {
		for _, property := range existingResource.Properties {
			colName := getPropertyColumnName(property)
			if !removedColumns[colName] {
				continue
			}

			alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("DROP COLUMN \"%s\"", colName))
			changesCount++
		}
	}

	// change updated columns
	for _, property := range params.Resource.Properties {
		colName := getPropertyColumnName(property)
		if !changedColumns[colName] {
			continue
		}
		prevProperty := newPrevMap[property]

		if prevProperty.Type != property.Type {
			alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("ALTER COLUMN \"%s\" TYPE %s", colName, getPsqlTypeFromProperty(property.Type, property.Length)))
			changesCount++
		}

		if prevProperty.Required && !property.Required {
			alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("ALTER COLUMN \"%s\" DROP NOT NULL", colName))
			changesCount++
		}

		if !prevProperty.Required && property.Required {
			alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("ALTER COLUMN \"%s\" SET NOT NULL", colName))
			changesCount++
		}
	}

	if params.ReferenceMap != nil {

		for _, existingFk := range existingResource.References {
			found := false
			for _, newFk := range params.Resource.References {
				if existingFk.PropertyName == newFk.PropertyName {
					found = true
					break
				}
			}
			if !found {
				// not implemented
				//alterTableQuery += " " + fmt.Sprintf("ADD CONSTRAINT FOREIGN KEY (%s) REFERENCES %s (%s)")
			}
		}

		for _, newFk := range params.Resource.References {
			// locating property
			var sourceProperty *model.ResourceProperty

			for _, property := range params.Resource.Properties {
				if property.Name == newFk.PropertyName {
					sourceProperty = property
				}
			}

			if sourceProperty == nil {
				return errors.LogicalError.WithDetails("Source property could not be found: " + newFk.ReferencedResource)
			}

			var sourceColumn = sourceProperty.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.Mapping

			found := false
			for _, existingFk := range existingResource.References {
				if existingFk.PropertyName == newFk.PropertyName || existingFk.PropertyName == "["+sourceColumn+"]" {
					found = true
					break
				}
			}
			if !found {

				targetTable := params.ReferenceMap[newFk.ReferencedResource]

				if targetTable.Entity == "" {
					return errors.LogicalError.WithDetails("Reference map not found for: " + newFk.ReferencedResource)
				}

				if targetTable.Catalog == "" {
					targetTable.Catalog = "public"
				}

				sqlPart := fmt.Sprintf("ADD CONSTRAINT \"fk_%s_%s\" FOREIGN KEY (\"%s\") REFERENCES \"%s\".\"%s\" (\"%s\")", sourceColumn, targetTable.Entity, sourceColumn, targetTable.Catalog, targetTable.Entity, targetTable.IdColumn)
				if newFk.Cascade {
					sqlPart += " on update cascade on delete cascade"
				}
				alterTableQueryDefs = append(alterTableQueryDefs, sqlPart)
				changesCount++
			}
		}
	}

	alterTableQuery += "\n" + strings.Join(alterTableQueryDefs, ",")

	if changesCount == 0 {
		return nil
	}

	_, sqlError := runner.Exec(alterTableQuery)

	logger.Trace("SqlQuery: " + alterTableQuery)

	if sqlError != nil {
		return handleDbError(ctx, sqlError)
	}

	return nil
}

func resourcePrepareResourceFromEntity(ctx context.Context, runner QueryRunner, catalog, entity string) (resource *model.Resource, err errors.ServiceError) {
	if catalog == "" {
		catalog = "public"
	}
	// check if entity exists
	row := runner.QueryRowContext(ctx, `select count(*) from information_schema.tables where table_type = 'BASE TABLE' and tables.table_schema = $1 and tables.table_name = $2`, catalog, entity)

	if row.Err() != nil {
		return nil, handleDbError(ctx, row.Err())
	}

	var count = new(int)

	err = handleDbError(ctx, row.Scan(&count))

	if err != nil {
		return
	}

	if *count == 0 {
		err = errors.RecordNotFoundError
		return
	}

	resource = new(model.Resource)
	resource.Annotations = make(map[string]string)
	annotations2.Enable(resource, annotations2.AutoCreated, annotations2.DisableMigration, annotations2.DisableAudit)
	resource.AuditData = new(model.AuditData)
	resource.DataType = model.DataType_USER
	resource.Name = strings.Replace(entity, ".", "_", -1)
	resource.Namespace = "default"
	resource.SourceConfig = &model.ResourceSourceConfig{
		Catalog: catalog,
		Entity:  entity,
	}

	// properties

	err = resourcePrepareProperties(ctx, runner, catalog, entity, resource)
	if err != nil {
		return
	}

	// references

	err = resourcePrepareReferences(ctx, runner, catalog, entity, resource)
	if err != nil {
		return
	}

	doResourceCleanup(resource)

	return
}

func resourcePrepareReferences(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
	rows, sqlErr := runner.QueryContext(ctx, `SELECT
       (SELECT a.attname
        FROM pg_attribute a
        WHERE a.attrelid = m.oid AND a.attnum = o.conkey[1] AND a.attisdropped = false)  AS source_column,
       (SELECT nspname FROM pg_namespace WHERE oid = f.relnamespace)                     AS target_schema,
       f.relname                                                                         AS target_table,
       (SELECT a.attname
        FROM pg_attribute a
        WHERE a.attrelid = f.oid AND a.attnum = o.confkey[1] AND a.attisdropped = false) AS target_column
FROM pg_constraint o
         LEFT JOIN pg_class f ON f.oid = o.confrelid
         LEFT JOIN pg_class m ON m.oid = o.conrelid
WHERE o.contype = 'f'
  AND o.conrelid IN (SELECT oid FROM pg_class c WHERE c.relkind = 'r')
and (SELECT nspname FROM pg_namespace WHERE oid = m.relnamespace) = $1
and m.relname   = $2`, catalog, entity)

	err := handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	for rows.Next() {
		sourceColumn := new(string)
		targetSchema := new(string)
		targetTable := new(string)
		targetColumn := new(string)

		sqlErr = rows.Scan(sourceColumn, targetSchema, targetTable, targetColumn)

		if sqlErr != nil {
			return handleDbError(ctx, sqlErr)
		}

		// locating source property
		var sourceProperty *model.ResourceProperty
		for _, property := range resource.Properties {
			if property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.Mapping == *sourceColumn {
				sourceProperty = property
				break
			}
		}

		if sourceProperty == nil {
			return errors.LogicalError.WithDetails("Source property cannot be located")
		}

		resource.References = append(resource.References, &model.ResourceReference{
			PropertyName:       "[" + sourceProperty.Name + "]",
			ReferencedResource: "[" + *targetTable + "]",
			Cascade:            true,
		})
	}

	return nil
}

func resourcePrepareProperties(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
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
where columns.table_schema = $1 and columns.table_name = $2 order by columns.ordinal_position`, catalog, entity)
	err := handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	primaryCount := 0
	for rows.Next() {
		var columnName = new(string)
		var columnType = new(string)
		var columnLength = new(*int)
		var isNullable = new(bool)
		var isPrimary = new(bool)

		err = handleDbError(ctx, rows.Scan(columnName, columnType, columnLength, isNullable, isPrimary))

		if err != nil {
			return err
		}

		var sourceDef = *columnType

		if *columnLength != nil {
			sourceDef = *columnType + "(" + strconv.Itoa(**columnLength) + ")"
		} else {
			*columnLength = new(int)
			**columnLength = 0
		}

		if *isPrimary && !annotations2.IsEnabled(resource, annotations2.DoPrimaryKeyLookup) {
			primaryCount++

			if primaryCount > 1 {
				annotations2.Enable(resource, annotations2.DoPrimaryKeyLookup)
			}

			if *columnName != "id" {
				annotations2.Enable(resource, annotations2.DoPrimaryKeyLookup)
			}

			if *columnType != "uuid" {
				annotations2.Enable(resource, annotations2.DoPrimaryKeyLookup)
			}
		}

		typ := getPropertyTypeFromPsql(*columnType)

		if typ == model.ResourcePropertyType_TYPE_STRING && uint32(**columnLength) == 0 {
			**columnLength = 256
		}

		property := &model.ResourceProperty{
			Name: *columnName,
			Type: typ,
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
	return err
}

func doResourceCleanup(resource *model.Resource) {
	createdOnDetected := false
	updatedOnDetected := false
	createdByDetected := false
	updatedByDetected := false
	versionDetected := false
	for _, property := range resource.Properties {
		if property.Name == "created_on" && property.Type == model.ResourcePropertyType_TYPE_TIMESTAMP {
			createdOnDetected = true
		}
		if property.Name == "updated_on" && property.Type == model.ResourcePropertyType_TYPE_TIMESTAMP {
			updatedOnDetected = true
		}
		if property.Name == "created_by" && property.Type == model.ResourcePropertyType_TYPE_STRING {
			createdByDetected = true
		}
		if property.Name == "updated_by" && property.Type == model.ResourcePropertyType_TYPE_STRING {
			updatedByDetected = true
		}
		if property.Name == "version" && property.Type == model.ResourcePropertyType_TYPE_INT32 {
			versionDetected = true
		}
	}
	enableAudit := createdOnDetected && updatedOnDetected && createdByDetected && updatedByDetected && versionDetected

	var newColumns []*model.ResourceProperty

	for _, property := range resource.Properties {
		if enableAudit && isAuditColumn(property.Name) {
			continue
		}

		if property.Primary && !annotations2.IsEnabled(resource, annotations2.DoPrimaryKeyLookup) {
			// ignore id column if it is same as standard id column
			continue
		}

		newColumns = append(newColumns, property)
	}

	resource.Properties = newColumns

	if enableAudit {
		annotations2.Disable(resource, annotations2.DisableAudit)
	}
}

func getPropertyColumnName(property *model.ResourceProperty) string {
	if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
		return sourceConfig.Mapping.Mapping
	} else {
		return ""
	}
}
