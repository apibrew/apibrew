package common

import (
	"context"
	"fmt"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/util"
	"strconv"
	"strings"
)

func (p *sqlBackend) resourceMigrateTable(ctx context.Context, runner helper.QueryRunner, params abs.UpgradeResourceParams, history bool) errors.ServiceError {
	hp := p.options.GetResourceMigrationBuilderConstructor()(ctx, runner, params, history, params.ForceMigration)

	return helper.ResourceMigrateTableViaResourceMigrationBuilder(hp, params.MigrationPlan, history, params.ForceMigration)
}

func (p *sqlBackend) resourcePrepareResourceFromEntity(ctx context.Context, runner helper.QueryRunner, catalog, entity string) (resource *model.Resource, err errors.ServiceError) {
	if catalog == "" {
		catalog = p.options.GetDefaultCatalog()
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

	if catalog != p.options.GetDefaultCatalog() && catalog != "" {
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

func (p *sqlBackend) resourcePrepareProperties(ctx context.Context, runner helper.QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
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

func (p *sqlBackend) resourcePrepareIndexes(ctx context.Context, runner helper.QueryRunner, catalog string, entity string, resource *model.Resource) errors.ServiceError {
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
