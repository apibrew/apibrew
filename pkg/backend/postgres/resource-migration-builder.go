package postgres

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
	options        postgreSqlBackendOptions
	handleDbError  func(ctx context.Context, err error) errors.ServiceError
	execs          []func() errors.ServiceError
	tableName      string
	schema         *abs.Schema
}

func (r *resourceMigrationBuilder) prepareIndexDef(index *model.ResourceIndex, params abs.UpgradeResourceParams, resource *model.Resource) (string, errors.ServiceError) {
	var uniqueStr = ""

	if index.Unique {
		uniqueStr = "UNIQUE"
	}

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

		colsEscaped = append(colsEscaped, r.options.Quote(prop.Name))
	}

	indexName := r.prepareIndexName(index, resource)

	sql := fmt.Sprintf("CREATE %s INDEX %s ON %s(%s)", uniqueStr, r.options.Quote(indexName), r.options.GetFullTableName(resource.SourceConfig), strings.Join(colsEscaped, ","))
	return sql, nil
}

func (r *resourceMigrationBuilder) prepareIndexName(index *model.ResourceIndex, resource *model.Resource) string {
	if annotations.Get(index, annotations.SourceIdentity) != "" {
		return annotations.Get(index, annotations.SourceIdentity)
	}

	var cols []string

	for _, indexProp := range index.Properties {
		var prop *model.ResourceProperty
		for _, prop = range resource.Properties {
			if prop.Name == indexProp.Name {
				break
			}
		}

		cols = append(cols, prop.Name)
	}

	var indexName = resource.SourceConfig.Entity + "_" + strings.Join(cols, "_")

	if index.Unique {
		indexName = indexName + "_uniq_idx"
	} else {
		indexName = indexName + "_idx"
	}
	return indexName
}

func (r *resourceMigrationBuilder) prepareResourceTableColumnDefinition(resource *model.Resource, property *model.ResourceProperty, schema abs.Schema) (string, errors.ServiceError) {
	uniqModifier := ""
	nullModifier := "NULL"
	if property.Required {
		nullModifier = "NOT NULL"
	}
	if property.Unique {
		uniqModifier = "UNIQUE"
	}
	sqlType := r.options.GetSqlTypeFromProperty(property.Type, property.Length)

	if property.Annotations != nil && property.Annotations[annotations.SQLType] != "" {
		sqlType = property.Annotations[annotations.SQLType]
	}

	var def = []string{r.options.Quote(property.Name), sqlType, nullModifier, uniqModifier}

	if property.Type == model.ResourceProperty_REFERENCE {
		if property.Reference != nil {
			referenceNamespace := property.Reference.Namespace
			if referenceNamespace == "" {
				referenceNamespace = resource.Namespace
			}
			referencedResource := schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]

			var refClause = ""
			if property.Reference.Cascade {
				refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
			}

			if referencedResource == nil {
				return "", errors.LogicalError.WithDetails("Referenced resource not exists with name: " + referenceNamespace + "/" + property.Reference.Resource)
			}

			if !referencedResource.Virtual {
				def = append(def, fmt.Sprintf(" CONSTRAINT %s REFERENCES %s (%s) %s", r.options.Quote(resource.SourceConfig.Entity+"_"+property.Name+"_fk"), r.options.Quote(referencedResource.SourceConfig.Entity), "id", refClause))
			}
		}
	}

	if property.DefaultValue != nil && property.DefaultValue.AsInterface() != nil {
		propertyType := types.ByResourcePropertyType(property.Type)
		val, _ := propertyType.UnPack(property.DefaultValue)

		def = append(def, fmt.Sprintf("DEFAULT '%v'", val))
	}

	return strings.Join(def, " "), nil
}

func (r *resourceMigrationBuilder) definePrimaryKeyColumn(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) {
	var pk []string
	for _, prop := range resource.Properties {
		if prop.Primary {
			var typ = r.options.GetSqlTypeFromProperty(prop.Type, prop.Length)

			if prop.Annotations != nil && prop.Annotations[annotations.SQLType] != "" {
				typ = prop.Annotations[annotations.SQLType]
			}

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

func (r *resourceMigrationBuilder) resourceCreateTable(resource *model.Resource) errors.ServiceError {
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
	r.execs = append(r.execs, func() errors.ServiceError {
		return r.resourceCreateTable(resource)
	})

	return r
}

func (r *resourceMigrationBuilder) UpdateResource(existing, updated *model.Resource) helper.ResourceMigrationBuilder {
	panic("not implemented")
}

func (r *resourceMigrationBuilder) DeleteResource(resource *model.Resource) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() errors.ServiceError {
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
	r.execs = append(r.execs, func() errors.ServiceError {
		refPart, serviceErr := r.prepareResourceTableColumnDefinition(r.params.MigrationPlan.CurrentResource, prop, *r.schema)

		if serviceErr != nil {
			return serviceErr
		}

		sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", r.options.GetFullTableName(r.params.MigrationPlan.CurrentResource.SourceConfig), refPart)

		_, err := r.runner.ExecContext(r.ctx, sql)

		return r.options.handleDbError(r.ctx, err)
	})

	return r
}

func (r *resourceMigrationBuilder) UpdateProperty(resource *model.Resource, prevProperty, property *model.ResourceProperty) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() errors.ServiceError {
		var sqlPrefix = fmt.Sprintf("ALTER TABLE %s ", r.tableName)
		var sqlParts []string
		changes := 0
		if r.options.GetSqlTypeFromProperty(prevProperty.Type, property.Length) != r.options.GetSqlTypeFromProperty(property.Type, property.Length) {

			sqlType := r.options.GetSqlTypeFromProperty(property.Type, property.Length)

			if property.Annotations != nil && property.Annotations[annotations.SQLType] != "" {
				sqlType = property.Annotations[annotations.SQLType]
			}

			sqlParts = append(sqlParts, fmt.Sprintf("ALTER COLUMN %s TYPE %s", r.options.Quote(property.Name), sqlType))
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
			sqlParts = append(sqlParts, fmt.Sprintf("DROP CONSTRAINT IF EXISTS %s", r.options.Quote(property.Name+"_uniq")))
			changes++
		}

		if !prevProperty.Unique && property.Unique {
			sqlParts = append(sqlParts, fmt.Sprintf("ADD CONSTRAINT %s UNIQUE (%s)", r.options.Quote(r.params.MigrationPlan.CurrentResource.SourceConfig.Entity+"_"+property.Name+"_uniq"), r.options.Quote(property.Name)))
			changes++
		}

		// fixme Default Value Modification logic

		if property.Type == model.ResourceProperty_REFERENCE {
			if prevProperty.Reference == nil && property.Reference != nil {
				referenceNamespace := property.Reference.Namespace
				if referenceNamespace == "" {
					referenceNamespace = resource.Namespace
				}
				referencedResource := r.schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]
				var refClause = ""
				if property.Reference.Cascade {
					refClause = "ON UPDATE CASCADE ON DELETE CASCADE"
				}

				if referencedResource == nil {
					return errors.ReferenceViolation.WithDetails("Referenced resource not exists with name: " + referenceNamespace + "/" + property.Reference.Resource)
				}

				if !referencedResource.Virtual {
					sqlParts = append(sqlParts, fmt.Sprintf("ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s) "+refClause, r.options.Quote(r.params.MigrationPlan.CurrentResource.SourceConfig.Entity+"_"+property.Name+"_fk"), r.options.Quote(property.Name), r.options.Quote(referencedResource.SourceConfig.Entity), r.options.Quote("id")))
					changes++
				}
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
	r.execs = append(r.execs, func() errors.ServiceError {
		sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN \"%s\"", r.tableName, prop.Name)

		_, sqlError := r.runner.ExecContext(r.ctx, sql)

		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) AddIndex(prop *model.ResourceIndex) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() errors.ServiceError {
		var err errors.ServiceError
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

func (r *resourceMigrationBuilder) DeleteIndex(index *model.ResourceIndex) helper.ResourceMigrationBuilder {
	r.execs = append(r.execs, func() errors.ServiceError {
		var indexName = r.prepareIndexName(index, r.params.MigrationPlan.CurrentResource)

		sql := fmt.Sprintf("DROP INDEX %s(\"%s\")", r.tableName, indexName)

		_, sqlError := r.runner.ExecContext(r.ctx, sql)
		return r.options.handleDbError(r.ctx, sqlError)
	})

	return r
}

func (r *resourceMigrationBuilder) Exec() errors.ServiceError {
	for _, exec := range r.execs {
		err := exec()

		if err != nil {
			return err
		}
	}

	return nil
}
