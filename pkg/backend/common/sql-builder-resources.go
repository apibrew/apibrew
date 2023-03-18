package common

import (
	"context"
	"fmt"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"strings"
)

func (p *sqlBackend) prepareResourceTableColumnDefinition(resource *model.Resource, property *model.ResourceProperty, schema abs.Schema) string {
	uniqModifier := ""
	nullModifier := "NULL"
	if property.Required {
		nullModifier = "NOT NULL"
	}
	if property.Unique {
		uniqModifier = "UNIQUE"
	}
	sqlType := p.options.GetSqlTypeFromProperty(property.Type, property.Length)

	var def = []string{fmt.Sprintf("%s", p.options.Quote(property.Mapping)), sqlType, nullModifier, uniqModifier}

	if property.Type == model.ResourceProperty_REFERENCE {
		if property.Reference != nil {
			referencedResource := schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+property.Reference.ReferencedResource]
			var refClause = ""
			if property.Reference.Cascade {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}
			def = append(def, fmt.Sprintf(" CONSTRAINT %s REFERENCES %s (%s) %s", p.options.Quote(resource.SourceConfig.Entity+"_"+property.Mapping+"_fk"), p.options.Quote(referencedResource.SourceConfig.Entity), "id", refClause))

		}
	}

	if property.DefaultValue != nil && property.DefaultValue.AsInterface() != nil {
		propertyType := types.ByResourcePropertyType(property.Type)
		val, _ := propertyType.UnPack(property.DefaultValue)

		def = append(def, fmt.Sprintf("DEFAULT '%s'", val))
	}

	return strings.Join(def, " ")
}

func (p *sqlBackend) resourceListEntities(ctx context.Context, runner helper.QueryRunner) (result []*model.DataSourceCatalog, err errors.ServiceError) {
	rows, sqlErr := runner.QueryContext(ctx, `
select table_schema, table_name, false
from information_schema.tables
where table_schema not in ('information_schema', 'pg_catalog')
union all
select table_schema, table_name, true
from information_schema.views
where table_schema not in ('information_schema', 'pg_catalog')
order by table_schema
`)
	err = p.handleDbError(ctx, sqlErr)

	if err != nil {
		return
	}

	var catalog *model.DataSourceCatalog

	for rows.Next() {
		var catalogName = new(string)
		var entityName = new(string)
		var readOnly = new(bool)

		err = p.handleDbError(ctx, rows.Scan(catalogName, entityName, readOnly))

		if err != nil {
			return
		}

		if catalog == nil || catalog.Name != *catalogName {
			if catalog != nil {
				result = append(result, catalog)
			}
			catalog = &model.DataSourceCatalog{Name: *catalogName}
		}

		catalog.Entities = append(catalog.Entities, &model.DataSourceEntity{
			Name:     *entityName,
			ReadOnly: *readOnly,
		})
	}

	return
}
