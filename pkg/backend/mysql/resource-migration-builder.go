package mysql

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"strings"
)

type resourceMigrationBuilder struct {
	forceMigration bool
	params         abs.UpgradeResourceParams
	runner         helper.QueryRunner
	ctx            context.Context
	options        mysqlBackendOptions
	handleDbError  func(ctx context.Context, err error) error
	execs          []func() error
	tableName      string
	schema         *abs.Schema
}

func (r *resourceMigrationBuilder) prepareIndexDef(index *model.ResourceIndex, params abs.UpgradeResourceParams, resource *model.Resource) (string, error) {
	var uniqueStr = ""

	if index.Unique {
		uniqueStr = "unique"
	}

	var cols []string
	var colsEscaped []string

	for _, indexProp := range index.Properties {
		var prop *model.ResourceProperty
		for _, prop = range params.MigrationPlan.CurrentResource.Properties {
			if prop.Name == indexProp.Name {
				break
			}
		}

		if prop == nil {
			return "", errors.LogicalError.WithDetails("Property not found with name: " + prop.Name)
		}

		cols = append(cols, prop.Name)
		colsEscaped = append(colsEscaped, r.options.Quote(prop.Name))
	}

	var indexName = resource.SourceConfig.Entity + "_" + strings.Join(cols, "_")

	if index.Unique {
		indexName = indexName + "_uniq_idx"
	} else {
		indexName = indexName + "_idx"
	}

	sql := fmt.Sprintf("create %s index %s on %s(%s)", uniqueStr, r.options.Quote(indexName), r.options.Quote(resource.SourceConfig.Entity), strings.Join(colsEscaped, ","))
	return sql, nil
}

func (r *resourceMigrationBuilder) prepareResourceTableColumnDefinition(resource *model.Resource, property *model.ResourceProperty, schema abs.Schema) string {
	uniqModifier := ""
	nullModifier := "NULL"
	if property.Required {
		nullModifier = "NOT NULL "
	}
	if property.Unique {
		uniqModifier = "UNIQUE"
	}

	sqlType := r.options.GetSqlTypeFromProperty(property.Type, property.Length)

	var def = []string{r.options.Quote(property.Name), sqlType, nullModifier, uniqModifier}

	if property.Type == model.ResourceProperty_REFERENCE {
		if property.Reference != nil {
			referenceNamespace := property.Reference.Namespace
			if referenceNamespace == "" {
				referenceNamespace = resource.Namespace
			}
			referencedResource := schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]
			var refClause = ""
			if annotations.IsEnabled(property, annotations.CascadeReference) {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}
			def = append(def, fmt.Sprintf(" CONSTRAINT %s REFERENCES %s (%s) %s", r.options.Quote(resource.SourceConfig.Entity+"_"+property.Name+"_fk"), r.options.Quote(referencedResource.SourceConfig.Entity), "id", refClause))

		}
	}

	if property.Required {
		propertyType := types.ByResourcePropertyType(property.Type)
		def = append(def, fmt.Sprintf("DEFAULT '%s'", propertyType.Default()))
	}

	return strings.Join(def, " ")
}

func (r *resourceMigrationBuilder) definePrimaryKeyColumn(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) {
	var pk []string
	for _, prop := range resource.Properties {
		if prop.Primary {
			var typ = r.options.GetSqlTypeFromProperty(prop.Type, prop.Length)

			if annotations.IsEnabled(prop, annotations.Identity) {
				if typ == "INT" {
					typ = "SERIAL"
				} else {
					typ = "BIGSERIAL"
				}
			}

			builder.Define(r.options.Quote(prop.Name), typ, "NOT NULL")
			pk = append(pk, r.options.Quote(prop.Name))
		}
	}

	if len(pk) > 0 {
		builder.Define("Primary Key(", strings.Join(pk, ","), ")")
	}
}

func (r *resourceMigrationBuilder) resourceCreateTable(resource *model.Resource) error {
	if resource.SourceConfig.Catalog != "" {
		_, err := r.runner.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", r.options.Quote(resource.SourceConfig.Catalog)))

		if err != nil {
			return r.handleDbError(r.ctx, err)
		}
	}

	builder := sqlbuilder.CreateTable(r.options.GetFullTableName(resource.SourceConfig))

	builder.IfNotExists()

	r.definePrimaryKeyColumn(resource, builder)

	sqlQuery, _ := builder.Build()
	_, err := r.runner.Exec(sqlQuery)

	return r.handleDbError(r.ctx, err)
}

func (r *resourceMigrationBuilder) AddResource(resource *model.Resource) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		return r.resourceCreateTable(resource)
	})

	return r
}

func (r *resourceMigrationBuilder) UpdateResource(existing, updated *model.Resource) helper.ResourceMigrationBuilder {
	panic("not implemented")
}

func (r *resourceMigrationBuilder) DeleteResource(resource *model.Resource) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		s := "DROP TABLE " + r.options.GetFullTableName(r.params.MigrationPlan.CurrentResource.SourceConfig)

		if r.forceMigration {
			s += " CASCADE;"
		}

		_, err := r.runner.Exec(s)

		return r.handleDbError(r.ctx, err)
	})

	return r
}

func (r *resourceMigrationBuilder) AddProperty(prop *model.ResourceProperty) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", r.options.GetFullTableName(r.params.MigrationPlan.CurrentResource.SourceConfig), r.prepareResourceTableColumnDefinition(r.params.MigrationPlan.CurrentResource, prop, *r.schema))

		_, err := r.runner.ExecContext(r.ctx, sql)

		return r.options.handleDbError(r.ctx, err)
	})

	return r
}

func (r *resourceMigrationBuilder) UpdateProperty(resource *model.Resource, prevProperty, property *model.ResourceProperty) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		var sqlPrefix = fmt.Sprintf("ALTER TABLE %s ", r.tableName)
		var sqlParts []string
		changes := 0
		if r.options.GetSqlTypeFromProperty(prevProperty.Type, property.Length) != r.options.GetSqlTypeFromProperty(property.Type, property.Length) {
			sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN %s TYPE %s", r.options.Quote(property.Name), r.options.GetSqlTypeFromProperty(property.Type, property.Length)))
			changes++
		}

		if prevProperty.Required && !property.Required {
			sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN %s DROP NOT NULL", r.options.Quote(property.Name)))
			changes++
		}

		if !prevProperty.Required && property.Required {
			sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN %s SET NOT NULL", r.options.Quote(property.Name)))
			changes++
		}

		if prevProperty.Unique && !property.Unique {
			sqlParts = append(sqlParts, fmt.Sprintf("DROP CONSTRAINT %s", r.options.Quote(property.Name+"_uniq")))
			changes++
		}

		if !prevProperty.Unique && property.Unique {
			sqlParts = append(sqlParts, fmt.Sprintf("ADD CONSTRAINT %s UNIQUE (%s)", r.options.Quote(r.params.MigrationPlan.CurrentResource.SourceConfig.Entity+"_"+property.Name+"_uniq"), r.options.Quote(property.Name)))
			changes++
		}

		// fixme Default Value Modification logic

		if property.Type == model.ResourceProperty_REFERENCE {
			prevCascace := annotations.IsEnabled(prevProperty, annotations.CascadeReference)
			cascade := annotations.IsEnabled(property, annotations.CascadeReference)

			if prevProperty.Reference != property.Reference || prevCascace != cascade {
				referenceNamespace := property.Reference.Namespace
				if referenceNamespace == "" {
					referenceNamespace = resource.Namespace
				}
				referencedResource := r.schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]
				var refClause = ""
				if annotations.IsEnabled(property, annotations.CascadeReference) {
					refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
				}

				var constraintName = r.options.Quote(r.params.MigrationPlan.CurrentResource.SourceConfig.Entity + "_" + property.Name + "_fk")

				sqlParts = append(sqlParts, fmt.Sprintf("DROP CONSTRAINT IF EXISTS %s, ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s) "+refClause, constraintName, constraintName, r.options.Quote(property.Name), r.options.Quote(referencedResource.SourceConfig.Entity), r.options.Quote("id")))
				changes++
			}
		}

		if changes == 0 {
			return nil
		}

		sql := sqlPrefix + "\n" + strings.Join(sqlParts, ",\n")

		_, sqlError := r.runner.ExecContext(r.ctx, sql)

		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) DeleteProperty(prop *model.ResourceProperty) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", r.options.Quote(r.tableName), r.options.Quote(prop.Name))

		_, sqlError := r.runner.ExecContext(r.ctx, sql)

		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) AddIndex(prop *model.ResourceIndex) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		var err error
		var sql string
		if annotations.Get(prop, annotations.SourceDef) != "" {
			sql = annotations.Get(prop, annotations.SourceDef)
		} else {
			sql, err = r.prepareIndexDef(prop, r.params, r.params.MigrationPlan.CurrentResource)
			if err != nil {
				return err
			}
		}

		_, sqlError := r.runner.ExecContext(r.ctx, sql)
		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) DeleteIndex(prop *model.ResourceIndex) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() error {
		sql := fmt.Sprintf("DROP INDEX %s", annotations.Get(prop, annotations.SourceIdentity))

		_, sqlError := r.runner.ExecContext(r.ctx, sql)
		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) Exec() error {
	for _, exec := range r.execs {
		err := exec()

		if err != nil {
			return err
		}
	}

	return nil
}
