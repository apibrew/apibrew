package postgres

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func resourceMigrateTable(ctx context.Context, runner QueryRunner, resource *model.Resource, forceMigration bool, history bool) errors.ServiceError {
	var err errors.ServiceError
	var existingResource *model.Resource
	entityName := resource.SourceConfig.Mapping
	if !strings.Contains(entityName, ".") {
		entityName = "public." + entityName
	}

	if history {
		entityName = entityName + "_h"
	}

	if existingResource, err = resourcePrepareResourceFromEntity(ctx, runner, entityName); err != nil {
		log.Error("Unable to load resource details: ", err)
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
		for _, newProperty := range resource.Properties {
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
	for _, newProperty := range resource.Properties {
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

	if len(changedColumns) == 0 && len(newColumns) == 0 && (!forceMigration || len(removedColumns) == 0) {
		// no need to migration
		return nil
	}

	// create new properties
	var alterTableQuery = fmt.Sprintf(`ALTER TABLE %s`, getTableName(resource.GetSourceConfig().GetMapping(), history))

	var alterTableQueryDefs []string

	for _, property := range resource.Properties {
		colName := getPropertyColumnName(property)
		if !newColumns[colName] {
			continue
		}

		alterTableQueryDefs = append(alterTableQueryDefs, fmt.Sprintf("ADD COLUMN \"%s\"", prepareResourceTableColumnDefinition(property)))
		changesCount++
	}

	// delete properties (IF FORCE MIGRATION)
	if forceMigration {
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
	for _, property := range resource.Properties {
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

	if changesCount == 0 {
		return nil
	}

	alterTableQuery += " " + strings.Join(alterTableQueryDefs, ",")

	_, sqlError := runner.Exec(alterTableQuery)

	log.Trace("SqlQuery: " + alterTableQuery)

	if sqlError != nil {
		return handleDbError(sqlError)
	}
	return nil
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
	resource.DataType = model.DataType_USER
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

	doResourceCleanup(resource)

	return
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

		if property.Primary && !resource.Flags.DoPrimaryKeyLookup {
			// ignore id column if it is same as standard id column
			continue
		}

		newColumns = append(newColumns, property)
	}

	resource.Properties = newColumns

	if enableAudit {
		resource.Flags.DisableAudit = false
	}
}

func getPropertyColumnName(property *model.ResourceProperty) string {
	if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
		return sourceConfig.Mapping.Mapping
	} else {
		return ""
	}
}
