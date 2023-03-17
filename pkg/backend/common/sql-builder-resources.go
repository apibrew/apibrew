package common

import (
	"context"
	"fmt"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"strings"
)

func (p *sqlBackend) definePrimaryKeyColumn(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) errors.ServiceError {
	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		builder.Define("id", p.options.GetSqlTypeFromProperty(model.ResourceProperty_UUID, 0), "NOT NULL", "PRIMARY KEY")
	} else {
		for _, prop := range resource.Properties {
			if prop.Primary {
				var typ = p.options.GetSqlTypeFromProperty(prop.Type, prop.Length)

				if annotations.IsEnabled(prop, annotations.Identity) {
					if typ == "INT" {
						typ = "SERIAL"
					} else {
						typ = "BIGSERIAL"
					}
				}

				builder.Define(prop.Mapping, typ, "NOT NULL", "PRIMARY KEY")
				break
			}
		}
	}
	return nil
}

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

func (p *sqlBackend) isAuditColumn(column string) bool {
	return column == "created_on" || column == "updated_on" || column == "created_by" || column == "updated_by" || column == "version"
}
