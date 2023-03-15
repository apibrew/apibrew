package common

import (
	"context"
	"fmt"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"strings"
)

func (p *sqlBackend) resourceCreateTable(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	if resource.SourceConfig.Catalog != "" {
		_, err := runner.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\"", resource.SourceConfig.Catalog))

		if err != nil {
			return p.handleDbError(ctx, err)
		}
	}

	builder := sqlbuilder.CreateTable(p.getFullTableName(resource.SourceConfig, false))

	builder.IfNotExists()

	serviceError := p.definePrimaryKeyColumn(resource, builder)
	if serviceError != nil {
		return serviceError
	}

	// audit
	if !annotations.IsEnabled(resource, annotations.DisableAudit) {
		builder.Define("created_on", "timestamp", "NOT NULL")
		builder.Define("updated_on", "timestamp", "NULL")
		builder.Define("created_by", DbNameType, "NOT NULL")
		builder.Define("updated_by", DbNameType, "NULL")
	}

	// version
	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		builder.Define("version", "int2", "NOT NULL")
	}

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) definePrimaryKeyColumn(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) errors.ServiceError {
	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")
	} else {
		for _, prop := range resource.Properties {
			if prop.Primary {
				var typ = p.getPsqlTypeFromProperty(prop.Type, prop.Length)

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
	sqlType := p.getPsqlTypeFromProperty(property.Type, property.Length)

	var def = []string{fmt.Sprintf("\"%s\"", property.Mapping), sqlType, nullModifier, uniqModifier}

	if property.Type == model.ResourceProperty_REFERENCE {
		if property.Reference != nil {
			referencedResource := schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+property.Reference.ReferencedResource]
			var refClause = ""
			if property.Reference.Cascade {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}
			def = append(def, fmt.Sprintf(" CONSTRAINT \"%s\" REFERENCES \"%s\" (\"%s\") %s", resource.SourceConfig.Entity+"_"+property.Mapping+"_fk", referencedResource.SourceConfig.Entity, "id", refClause))

		}
	}

	if property.DefaultValue != nil && property.DefaultValue.AsInterface() != nil {
		propertyType := types.ByResourcePropertyType(property.Type)
		val, _ := propertyType.UnPack(property.DefaultValue)

		def = append(def, fmt.Sprintf("DEFAULT '%s'", val))
	}

	return strings.Join(def, " ")
}

func (p *sqlBackend) resourceCreateHistoryTable(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(p.getFullTableName(resource.SourceConfig, true))

	builder.IfNotExists()

	serviceError := p.definePrimaryKeyColumn(resource, builder)
	if serviceError != nil {
		return serviceError
	}

	builder.Define("created_on", "timestamp", "NOT NULL")
	builder.Define("updated_on", "timestamp", "NULL")
	builder.Define("created_by", DbNameType, "NOT NULL")
	builder.Define("updated_by", DbNameType, "NULL")
	// version
	builder.Define("version", "int2", "NOT NULL")

	builder.Define("PRIMARY KEY (id, version)")

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) resourceDropTable(ctx context.Context, runner QueryRunner, resource *model.Resource, history bool, forceMigration bool) errors.ServiceError {
	s := "DROP TABLE " + p.getFullTableName(resource.SourceConfig, history)

	if forceMigration {
		s += " CASCADE;"
	}

	_, err := runner.Exec(s)

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) resourceListEntities(ctx context.Context, runner QueryRunner) (result []*model.DataSourceCatalog, err errors.ServiceError) {
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
