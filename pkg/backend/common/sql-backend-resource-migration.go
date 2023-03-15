package common

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"strconv"
	"strings"
)

func (p *sqlBackend) resourceMigrateTable(ctx context.Context, runner QueryRunner, params abs.UpgradeResourceParams, history bool) errors.ServiceError {
	var currentPropertyMap = util.GetNamedMap(params.MigrationPlan.CurrentResource.Properties)
	var existingPropertyMap = util.GetNamedMap(params.MigrationPlan.ExistingResource.Properties)

	logger := log.WithFields(logging.CtxFields(ctx))

	var tableName = p.getFullTableName(params.Resource.GetSourceConfig(), history)

	for _, step := range params.MigrationPlan.Steps {
		switch sk := step.Kind.(type) {
		case *model.ResourceMigrationStep_CreateResource:
			if err := p.resourceCreateTable(ctx, runner, params.Resource); err != nil {
				return err
			}
		case *model.ResourceMigrationStep_DeleteResource:
			if params.ForceMigration {
				if err := p.resourceDropTable(ctx, runner, params.Resource, false, true); err != nil {
					return err
				}

				if err := p.resourceDropTable(ctx, runner, params.Resource, true, true); err != nil {
					return err
				}
			}
		case *model.ResourceMigrationStep_CreateProperty:
			property := currentPropertyMap[sk.CreateProperty.Property]
			if property.Primary {
				continue
			}
			sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, p.prepareResourceTableColumnDefinition(params.Resource, property, *params.Schema))

			_, sqlError := runner.ExecContext(ctx, sql)

			if sqlError != nil {
				return p.handleDbError(ctx, sqlError)
			}
		case *model.ResourceMigrationStep_UpdateProperty:
			err := p.migrateResourceColumn(existingPropertyMap[sk.UpdateProperty.ExistingProperty], currentPropertyMap[sk.UpdateProperty.Property], tableName, params.MigrationPlan.ExistingResource, logger, runner, ctx, *params.Schema)

			if err != nil {
				return err
			}
		case *model.ResourceMigrationStep_DeleteProperty:
			if params.ForceMigration {
				sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, existingPropertyMap[sk.DeleteProperty.ExistingProperty].Mapping)

				_, sqlError := runner.ExecContext(ctx, sql)

				if sqlError != nil {
					return p.handleDbError(ctx, sqlError)
				}
			}
		}
	}

	if !history {
		for _, step := range params.MigrationPlan.Steps {
			switch sk := step.Kind.(type) {
			case *model.ResourceMigrationStep_CreateIndex:
				var err errors.ServiceError
				var sql string
				if annotations.Get(params.MigrationPlan.CurrentResource.Indexes[sk.CreateIndex.Index], annotations.SourceDef) != "" {
					sql = annotations.Get(params.MigrationPlan.CurrentResource.Indexes[sk.CreateIndex.Index], annotations.SourceDef)
				} else {
					sql, err = p.prepareIndexDef(params.MigrationPlan.CurrentResource.Indexes[sk.CreateIndex.Index], params, params.Resource)
					if err != nil {
						return err
					}
				}

				_, sqlError := runner.ExecContext(ctx, sql)
				return p.handleDbError(ctx, sqlError)
			case *model.ResourceMigrationStep_DeleteIndex:
				sql := fmt.Sprintf("DROP INDEX %s", annotations.Get(params.MigrationPlan.ExistingResource.Indexes[sk.DeleteIndex.ExistingIndex], annotations.SourceIdentity))

				_, sqlError := runner.ExecContext(ctx, sql)
				return p.handleDbError(ctx, sqlError)
			}
		}
	}

	return nil
}

func (p *sqlBackend) prepareIndexDef(index *model.ResourceIndex, params abs.UpgradeResourceParams, resource *model.Resource) (string, errors.ServiceError) {
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

	sql := fmt.Sprintf("create %s index \"%s\" on %s(%s)", uniqueStr, indexName, resource.SourceConfig.Entity, strings.Join(colsEscaped, ","))
	return sql, nil
}

func (p *sqlBackend) migrateResourceColumn(prevProperty *model.ResourceProperty, property *model.ResourceProperty, tableName string, existingResource *model.Resource, logger *log.Entry, runner QueryRunner, ctx context.Context, schema abs.Schema) errors.ServiceError {
	var sqlPrefix = fmt.Sprintf("ALTER TABLE %s ", tableName)
	var sqlParts []string
	changes := 0
	if p.options.GetSqlTypeFromProperty(prevProperty.Type, property.Length) != p.options.GetSqlTypeFromProperty(property.Type, property.Length) {
		sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN \"%s\" TYPE %s", property.Mapping, p.options.GetSqlTypeFromProperty(property.Type, property.Length)))
		changes++
	}

	if prevProperty.Required && !property.Required {
		sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN \"%s\" DROP NOT NULL", property.Mapping))
		changes++
	}

	if !prevProperty.Required && property.Required {
		sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN \"%s\" SET NOT NULL", property.Mapping))
		changes++
	}

	if prevProperty.Unique && !property.Unique {
		sqlParts = append(sqlParts, fmt.Sprintf("DROP CONSTRAINT \"%s\"", property.Mapping+"_uniq"))
		changes++
	}

	if !prevProperty.Unique && property.Unique {
		sqlParts = append(sqlParts, fmt.Sprintf("ADD CONSTRAINT \"%s\" UNIQUE (\"%s\")", existingResource.SourceConfig.Entity+"_"+property.Mapping+"_uniq", property.Mapping))
		changes++
	}

	// fixme Default Value Modification logic

	if property.Type == model.ResourceProperty_REFERENCE {
		if prevProperty.Reference == nil && property.Reference != nil {
			referencedResource := schema.ResourceByNamespaceSlashName["default"+"/"+property.Reference.ReferencedResource]
			var refClause = ""
			if property.Reference.Cascade {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}

			sqlParts = append(sqlParts, fmt.Sprintf("ADD CONSTRAINT \"%s\" FOREIGN KEY (\"%s\") REFERENCES \"%s\" (\"%s\") "+refClause, existingResource.SourceConfig.Entity+"_"+property.Mapping+"_fk", property.Mapping, referencedResource.SourceConfig.Entity, "id"))
			changes++
		}
	}

	if changes == 0 {
		return nil
	}

	sql := sqlPrefix + "\n" + strings.Join(sqlParts, ",\n")

	_, sqlError := runner.ExecContext(ctx, sql)
	return p.handleDbError(ctx, sqlError)
}

func (p *sqlBackend) resourcePrepareResourceFromEntity(ctx context.Context, runner QueryRunner, catalog, entity string) (resource *model.Resource, err errors.ServiceError) {
	if catalog == "" {
		catalog = "public"
	}
	// check if entity exists
	row := runner.QueryRowContext(ctx, p.options.GetSql("entity-exists"), catalog, entity)

	if row.Err() != nil {
		return nil, p.handleDbError(ctx, row.Err())
	}

	var count = new(int)

	err = p.handleDbError(ctx, row.Scan(&count))

	if err != nil {
		return
	}

	if *count == 0 {
		err = errors.RecordNotFoundError.WithMessage(fmt.Sprintf("Entity not found: %s/%s", catalog, entity))
		return
	}

	resource = new(model.Resource)
	resource.Annotations = make(map[string]string)
	annotations.Enable(resource, annotations.AutoCreated)
	resource.AuditData = new(model.AuditData)
	resource.Name = strings.Replace(entity, ".", "_", -1)

	if catalog != "public" && catalog != "" {
		resource.Name = catalog + "_" + resource.Name
	}

	resource.Namespace = "default"
	resource.SourceConfig = &model.ResourceSourceConfig{
		Catalog: catalog,
		Entity:  entity,
	}

	// properties

	err = p.resourcePrepareProperties(ctx, runner, catalog, entity, resource)
	if err != nil {
		return
	}

	// indexes

	err = p.resourcePrepareIndexes(ctx, runner, catalog, entity, resource)
	if err != nil {
		return
	}

	// references

	p.doResourceCleanup(resource)

	return
}

func (p *sqlBackend) resourcePrepareProperties(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
	rows, sqlErr := runner.QueryContext(ctx, p.options.GetSql("prepare-properties"), catalog, entity)
	err := p.handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	primaryCount := 0
	primaryFound := false

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

		err = p.handleDbError(ctx, rows.Scan(columnName, columnType, columnLength, isNullable, isPrimary, isUnique, isReferenced, targetSchema, targetTable, targetColumn))

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

		var isIdentity = false

		if *isPrimary && !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
			primaryCount++

			if primaryCount > 1 {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
			}

			if *columnName != "id" {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
				isIdentity = true
			}

			if *columnType != "uuid" {
				annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
				isIdentity = true
			}
		}

		if *isPrimary {
			primaryFound = true
		}

		typ := p.options.GetPropertyTypeFromPsql(*columnType)

		if typ == model.ResourceProperty_STRING && uint32(**columnLength) == 0 {
			**columnLength = 256
		}

		property := &model.ResourceProperty{
			Name: util.SnakeCaseToCamelCase(*columnName),
			Type: typ,

			Mapping:     *columnName,
			Primary:     *isPrimary,
			Required:    !*isNullable,
			Unique:      *isUnique,
			Length:      uint32(**columnLength),
			Annotations: make(map[string]string),
		}

		if *isReferenced {
			property.Type = model.ResourceProperty_REFERENCE
			property.Reference = &model.Reference{
				ReferencedResource: fmt.Sprintf("[%s]", **targetTable),
				Cascade:            false,
			}
		}

		if isIdentity {
			annotations.Enable(property, annotations.Identity)
		}

		annotations.Set(property, annotations.SourceDef, sourceDef)

		resource.Properties = append(resource.Properties, property)
	}

	if !primaryFound {
		annotations.Enable(resource, annotations.DoPrimaryKeyLookup)
	}

	return err
}

func (p *sqlBackend) resourcePrepareIndexes(ctx context.Context, runner QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
	rows, sqlErr := runner.QueryContext(ctx, p.options.GetSql("prepare-indexes"), catalog, entity)
	err := p.handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	for rows.Next() {
		var indexName = new(string)
		var unique = new(bool)
		var indexDef = new(string)
		var colsStr = new(string)

		err = p.handleDbError(ctx, rows.Scan(indexName, unique, indexDef, colsStr))

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

func (p *sqlBackend) doResourceCleanup(resource *model.Resource) {
	createdOnDetected := false
	updatedOnDetected := false
	createdByDetected := false
	updatedByDetected := false
	versionDetected := false
	for _, property := range resource.Properties {
		if property.Mapping == "created_on" && property.Type == model.ResourceProperty_TIMESTAMP {
			createdOnDetected = true
		}
		if property.Mapping == "updated_on" && property.Type == model.ResourceProperty_TIMESTAMP {
			updatedOnDetected = true
		}
		if property.Mapping == "created_by" && property.Type == model.ResourceProperty_STRING {
			createdByDetected = true
		}
		if property.Mapping == "updated_by" && property.Type == model.ResourceProperty_STRING {
			updatedByDetected = true
		}
		if property.Mapping == "version" && property.Type == model.ResourceProperty_INT32 {
			versionDetected = true
		}
	}
	enableAudit := createdOnDetected && updatedOnDetected && createdByDetected && updatedByDetected

	var newColumns []*model.ResourceProperty

	for _, property := range resource.Properties {
		if enableAudit && p.isAuditColumn(property.Mapping) {
			continue
		}

		if property.Primary && !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
			// ignore id column if it is same as standard id column
			continue
		}

		newColumns = append(newColumns, property)
	}

	resource.Properties = newColumns

	if !enableAudit {
		annotations.Enable(resource, annotations.DisableAudit)
	}

	if !versionDetected {
		annotations.Enable(resource, annotations.DisableVersion)
	}
}
