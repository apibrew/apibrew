package postgres

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations "github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"strconv"
	"strings"
)

func resourceMigrateTable(ctx context.Context, runner QueryRunner, params abs.UpgradeResourceParams, history bool) errors.ServiceError {
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

	// fixing references
	for _, prop := range existingResource.Properties {
		if prop.Type == model.ResourcePropertyType_TYPE_REFERENCE {
			for _, resource := range params.Schema.Resources {
				if prop.Reference.ReferencedResource == "["+resource.SourceConfig.Entity+"]" {
					prop.Reference.ReferencedResource = resource.Name
				}
			}
		}
	}

	var tableName = getFullTableName(params.Resource.GetSourceConfig(), history)

	err = arrayDiffer(existingResource.Properties,
		params.Resource.Properties,
		util.IsSameIdentifiedResourceProperty,
		util.IsSameResourceProperty,
		func(property *model.ResourceProperty) errors.ServiceError {
			sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, prepareResourceTableColumnDefinition(params.Resource, property, *params.Schema))

			logger.Info("DB Migrate Sql: " + sql)
			_, sqlError := runner.ExecContext(ctx, sql)
			return handleDbError(ctx, sqlError)
		}, func(prevProperty, property *model.ResourceProperty) errors.ServiceError {
			return migrateResourceColumn(prevProperty, property, tableName, existingResource, logger, runner, ctx, *params.Schema)
		}, func(property *model.ResourceProperty) errors.ServiceError {
			if params.ForceMigration {
				sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, property.Mapping)

				logger.Info("DB Migrate Sql: " + sql)
				_, sqlError := runner.ExecContext(ctx, sql)
				return handleDbError(ctx, sqlError)
			} else {
				return nil
			}
		})

	if err != nil {
		return err
	}

	if !history {
		err = arrayDiffer(existingResource.Indexes,
			params.Resource.Indexes,
			util.IsSameIdentifiedResourceIndex,
			util.IsSameResourceIndex,
			func(index *model.ResourceIndex) errors.ServiceError {
				var sql = ""
				if annotations.Get(index, annotations.SourceDef) != "" {
					sql = annotations.Get(index, annotations.SourceDef)
				} else {
					sql, err = prepareIndexDef(index, params, params.Resource, sql)
					if err != nil {
						return err
					}
				}

				logger.Info("DB Migrate Sql: " + sql)
				_, sqlError := runner.ExecContext(ctx, sql)
				return handleDbError(ctx, sqlError)
			}, func(prevProperty, property *model.ResourceIndex) errors.ServiceError {
				return errors.LogicalError.WithDetails("Cannot alter index")
			}, func(index *model.ResourceIndex) errors.ServiceError {
				if params.ForceMigration {
					sql := fmt.Sprintf("DROP INDEX %s", annotations.Get(index, annotations.SourceIdentity))

					logger.Info("DB Migrate Sql: " + sql)
					_, sqlError := runner.ExecContext(ctx, sql)
					return handleDbError(ctx, sqlError)
				} else {
					return nil
				}
			})
	}

	if err != nil {
		return err
	}

	return nil
}

func prepareIndexDef(index *model.ResourceIndex, params abs.UpgradeResourceParams, resource *model.Resource, sql string) (string, errors.ServiceError) {
	var uniqueStr = ""

	if index.Unique {
		uniqueStr = "unique"
	}

	var cols []string
	var colsEscaped []string

	for _, indexProp := range index.Properties {
		var prop *model.ResourceProperty
		for _, prop = range params.Resource.Properties {
			if prop.Name == indexProp.Name {
				break
			}
		}

		if prop == nil {
			return "", errors.LogicalError.WithDetails("Property not found with name: " + prop.Name)
		}

		cols = append(cols, prop.Mapping)
		colsEscaped = append(colsEscaped, fmt.Sprintf("\"%s\"", prop.Mapping))
	}

	var indexName = resource.SourceConfig.Entity + "_" + strings.Join(cols, "_")

	if index.Unique {
		indexName = indexName + "_uniq_idx"
	} else {
		indexName = indexName + "_idx"
	}

	sql = fmt.Sprintf("create %s index \"%s\" on %s(%s)", uniqueStr, indexName, resource.SourceConfig.Entity, strings.Join(colsEscaped, ","))
	return sql, nil
}

func migrateResourceColumn(prevProperty *model.ResourceProperty, property *model.ResourceProperty, tableName string, existingResource *model.Resource, logger *log.Entry, runner QueryRunner, ctx context.Context, schema abs.Schema) errors.ServiceError {
	var sql = fmt.Sprintf("ALTER TABLE %s ", tableName)
	changes := 0
	if getPsqlTypeFromProperty(prevProperty.Type, property.Length) != getPsqlTypeFromProperty(property.Type, property.Length) {
		sql = sql + fmt.Sprintf("\n ALTER COLUMN \"%s\" TYPE %s", property.Mapping, getPsqlTypeFromProperty(property.Type, property.Length))
		changes++
	}

	if prevProperty.Required && !property.Required {
		sql = sql + fmt.Sprintf("\n ALTER COLUMN \"%s\" DROP NOT NULL", property.Mapping)
		changes++
	}

	if !prevProperty.Required && property.Required {
		sql = sql + fmt.Sprintf("\n ALTER COLUMN \"%s\" SET NOT NULL", property.Mapping)
		changes++
	}

	if prevProperty.Unique && !property.Unique {
		sql = sql + fmt.Sprintf("\n DROP CONSTRAINT \"%s\"", property.Mapping+"_uniq")
		changes++
	}

	if !prevProperty.Unique && property.Unique {
		sql = sql + fmt.Sprintf("\n ADD CONSTRAINT \"%s\" UNIQUE (\"%s\")", existingResource.SourceConfig.Entity+"_"+property.Mapping+"_uniq", property.Mapping)
		changes++
	}

	if property.Type == model.ResourcePropertyType_TYPE_REFERENCE {
		if prevProperty.Reference == nil && property.Reference != nil {
			referencedResource := schema.ResourceByNamespaceSlashName["default"+"/"+property.Reference.ReferencedResource]
			var refClause = ""
			if property.Reference.Cascade {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}

			sql = sql + fmt.Sprintf("\n ADD CONSTRAINT \"%s\" FOREIGN KEY (\"%s\") REFERENCES \"%s\" (\"%s\") "+refClause, existingResource.SourceConfig.Entity+"_"+property.Mapping+"_fk", property.Mapping, referencedResource.SourceConfig.Entity, "id")
			changes++
		}

		if (prevProperty.Reference == nil) != (property.Reference == nil) {
			log.Print("a")
		} else if prevProperty.Reference.ReferencedResource != property.Reference.ReferencedResource {
			log.Print("b")
		} else if prevProperty.Reference.Cascade != property.Reference.Cascade {
			log.Print("c")
		} else {
			panic("Unknown condition")
		}
	}

	if changes == 0 {
		return nil
	}

	logger.Info("DB Migrate Sql: " + sql)
	_, sqlError := runner.ExecContext(ctx, sql)
	return handleDbError(ctx, sqlError)
}

func arrayDiffer[T interface{}](existing []T, updated []T, hasSameId func(a, b T) bool, isEqual func(a, b T) bool, onNew func(rec T) errors.ServiceError, onUpdate func(e, u T) errors.ServiceError, onDelete func(rec T) errors.ServiceError) errors.ServiceError {
	for _, e := range existing {
		found := false
		for _, u := range updated {
			if hasSameId(e, u) {
				if !isEqual(e, u) {
					err := onUpdate(e, u)

					if err != nil {
						return err
					}
				}

				found = true
				break
			}
		}

		if !found {
			err := onDelete(e)

			if err != nil {
				return err
			}
		}
	}

	for _, u := range updated {
		found := false
		for _, e := range existing {

			if hasSameId(e, u) {
				found = true
			}
		}

		if !found {
			err := onNew(u)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func resourcePrepareResourceFromEntity(ctx context.Context, runner QueryRunner, catalog, entity string) (resource *model.Resource, err errors.ServiceError) {
	if catalog == "" {
		catalog = "public"
	}
	// check if entity exists
	row := runner.QueryRowContext(ctx, `
select sum(count) from (
select count(*) as count from information_schema.tables where table_type = 'BASE TABLE' and tables.table_schema = $1 and tables.table_name = $2
union all
select count(*) as count from information_schema.views where views.table_schema = $1 and views.table_name = $2
)_
`, catalog, entity)

	if row.Err() != nil {
		return nil, handleDbError(ctx, row.Err())
	}

	var count = new(int)

	err = handleDbError(ctx, row.Scan(&count))

	if err != nil {
		return
	}

	if *count == 0 {
		err = errors.RecordNotFoundError.WithMessage(fmt.Sprintf("Entity not found: %s/%s", catalog, entity))
		return
	}

	resource = new(model.Resource)
	resource.Annotations = make(map[string]string)
	annotations.Enable(resource, annotations.AutoCreated, annotations.DisableMigration, annotations.DisableAudit)
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

	// indexes

	err = resourcePrepareIndexes(ctx, runner, catalog, entity, resource)
	if err != nil {
		return
	}

	// references

	doResourceCleanup(resource)

	return
}

func resourcePrepareProperties(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
	rows, sqlErr := runner.QueryContext(ctx, `

select columns.column_name,
       columns.udt_name as column_type,
       columns.character_maximum_length as length,
       columns.is_nullable = 'YES' as is_nullable,
       column_pkey.constraint_def is not null as is_primary,
       column_ukey.constraint_def is not null as is_unique,
       column_fkey.constraint_def is not null as is_referenced,
       column_fkey.target_schema,
       column_fkey.target_table,
       column_fkey.target_column
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
                    WHERE contype = 'p') column_pkey  on column_pkey.conname = key_column_usage.constraint_name
         left join (SELECT nspname                     as table_schema,
                           conname,
                           contype,
                           pg_get_constraintdef(c.oid) as constraint_def
                    FROM pg_constraint c
                             JOIN pg_namespace n ON n.oid = c.connamespace
                    WHERE contype = 'u') column_ukey  on column_ukey.conname = key_column_usage.constraint_name
         left join (SELECT nspname                     as table_schema,
                           conname,
                           contype,
                           pg_get_constraintdef(c.oid) as constraint_def,
                           (SELECT nspname FROM pg_namespace WHERE oid = f.relnamespace)                     AS target_schema,
                           f.relname                                                                         AS target_table,
                           (SELECT a.attname
                            FROM pg_attribute a
                            WHERE a.attrelid = f.oid AND a.attnum = c.confkey[1] AND a.attisdropped = false) AS target_column
                    FROM pg_constraint c
                             JOIN pg_namespace n ON n.oid = c.connamespace
                             LEFT JOIN pg_class f ON f.oid = c.confrelid
                             LEFT JOIN pg_class m ON m.oid = c.conrelid
                    WHERE contype = 'f' and c.conrelid IN (SELECT oid FROM pg_class c WHERE c.relkind = 'r')) column_fkey  on column_fkey.conname = key_column_usage.constraint_name
where columns.table_schema = $1 and columns.table_name = $2 order by columns.ordinal_position

`, catalog, entity)
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
		var isUnique = new(bool)
		var isReferenced = new(bool)
		var targetSchema = new(*string)
		var targetTable = new(*string)
		var targetColumn = new(*string)

		err = handleDbError(ctx, rows.Scan(columnName, columnType, columnLength, isNullable, isPrimary, isUnique, isReferenced, targetSchema, targetTable, targetColumn))

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

		if *isPrimary && !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
			primaryCount++

			if primaryCount > 1 {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
			}

			if *columnName != "id" {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
			}

			if *columnType != "uuid" {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
			}
		}

		typ := getPropertyTypeFromPsql(*columnType)

		if typ == model.ResourcePropertyType_TYPE_STRING && uint32(**columnLength) == 0 {
			**columnLength = 256
		}

		property := &model.ResourceProperty{
			Name: *columnName,
			Type: typ,

			Mapping:     *columnName,
			Primary:     *isPrimary,
			Required:    !*isNullable,
			Unique:      *isUnique,
			Length:      uint32(**columnLength),
			Annotations: make(map[string]string),
		}

		if *isReferenced {
			property.Type = model.ResourcePropertyType_TYPE_REFERENCE
			property.Reference = &model.Reference{
				ReferencedResource: fmt.Sprintf("[%s]", **targetTable),
				Cascade:            false,
			}
		}

		annotations.Set(property, annotations.SourceDef, sourceDef)

		resource.Properties = append(resource.Properties, property)
	}
	return err
}

func resourcePrepareIndexes(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
	rows, sqlErr := runner.QueryContext(ctx, `

select i.relname as index_name,
       indisunique,
       ixs.indexdef,
       string_agg(a.attname, ',') 
from pg_class t,
     pg_class i,
     pg_index ix,
     pg_indexes ixs,
     pg_attribute a
where t.oid = ix.indrelid
  and i.oid = ix.indexrelid
  and a.attrelid = t.oid
  and a.attnum = ANY (ix.indkey)
  and t.relkind = 'r'
  and ixs.indexname = i.relname
  and ixs.schemaname = $1
  and ixs.tablename = $2
group by 1, 2, 3
having count(distinct a.attname) > 1

`, catalog, entity)
	err := handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	for rows.Next() {
		var indexName = new(string)
		var unique = new(bool)
		var indexDef = new(string)
		var colsStr = new(string)

		err = handleDbError(ctx, rows.Scan(indexName, unique, indexDef, colsStr))

		var cols = strings.Split(*colsStr, ",")

		if err != nil {
			return err
		}

		var properties []*model.ResourceIndexProperty

		for _, col := range cols {
			var prop *model.ResourceProperty
			for _, prop = range resource.Properties {
				if prop.Mapping == col {
					break
				}
			}

			if prop == nil {
				return errors.LogicalError.WithDetails("Property not found with col: " + col)
			}

			properties = append(properties, &model.ResourceIndexProperty{
				Name:  prop.Name,
				Order: model.Order_ORDER_UNKNOWN,
			})
		}

		var resourceIndex = &model.ResourceIndex{
			Properties:  properties,
			IndexType:   model.ResourceIndexType_BTREE,
			Unique:      *unique,
			Annotations: map[string]string{},
		}

		annotations.Set(resourceIndex, annotations.SourceDef, *indexDef)
		annotations.Set(resourceIndex, annotations.SourceIdentity, *indexName)

		// check for duplicate index

		isDuplicate := false

		for _, existingIndex := range resource.Indexes {
			if util.IsSameResourceIndex(existingIndex, resourceIndex) {
				isDuplicate = true
			}
		}

		if isDuplicate {
			continue
		}

		resource.Indexes = append(resource.Indexes, resourceIndex)
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

		if property.Primary && !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
			// ignore id column if it is same as standard id column
			continue
		}

		newColumns = append(newColumns, property)
	}

	resource.Properties = newColumns

	if enableAudit {
		annotations.Disable(resource, annotations.DisableAudit)
	}
}
