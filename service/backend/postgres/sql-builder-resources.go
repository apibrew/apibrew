package postgres

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"strings"
)

type QueryResultScanner interface {
	Scan(dest ...any) error
}

type QueryRunner interface {
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func resourceCreateTable(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig.Mapping, false))

	builder.IfNotExists()

	if !resource.Flags.DoPrimaryKeyLookup {
		builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")
	}

	prepareCreateTableQuery(resource, builder)

	// audit
	if !resource.Flags.DisableAudit {
		builder.Define("created_on", "timestamp", "NOT NULL")
		builder.Define("updated_on", "timestamp", "NULL")
		builder.Define("created_by", DbNameType, "NOT NULL")
		builder.Define("updated_by", DbNameType, "NULL")
		// version
		builder.Define("version", "int2", "NOT NULL")
	}

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return handleDbError(err)
}

type ReferenceLocalDetails struct {
	sourceTableName       string
	fkConstraintName      string
	sourceTableColumn     string
	referencedTable       string
	referencedTableColumn string
	joinAlias             string
}

func prepareCreateTableQuery(resource *model.Resource, builder *sqlbuilder.CreateTableBuilder) {
	if resource.Flags == nil {
		resource.Flags = new(model.ResourceFlags)
	}

	var primaryKeys []string
	for _, property := range resource.Properties {
		columnDef := prepareResourceTableColumnDefinition(property)

		if columnDef != "" {
			builder.Define(columnDef)
		}

		if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			if property.Primary {
				primaryKeys = append(primaryKeys, sourceConfig.Mapping.Mapping)
			}
		}
	}
	if len(primaryKeys) > 0 {
		builder.Define("PRIMARY KEY (" + strings.Join(primaryKeys, ",") + ")")
	}
}

func prepareResourceTableColumnDefinition(property *model.ResourceProperty) string {
	if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
		uniqModifier := ""
		nullModifier := "NULL"
		if property.Required {
			nullModifier = "NOT NULL"
		}
		if property.Unique {
			uniqModifier = "UNIQUE"
		}
		sqlType := getPsqlTypeFromProperty(property.Type, property.Length)

		var def = []string{sourceConfig.Mapping.Mapping, sqlType, nullModifier, uniqModifier}

		return strings.Join(def, " ")
	}

	return ""
}

func resourceCreateHistoryTable(runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig.Mapping, true))

	builder.IfNotExists()

	builder.Define("id", "uuid", "NOT NULL")

	prepareCreateTableQuery(resource, builder)

	builder.Define("created_on", "timestamp", "NOT NULL")
	builder.Define("updated_on", "timestamp", "NULL")
	builder.Define("created_by", DbNameType, "NOT NULL")
	builder.Define("updated_by", DbNameType, "NULL")
	// version
	builder.Define("version", "int2", "NOT NULL")

	builder.Define("PRIMARY KEY (id, version)")

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	return handleDbError(err)
}

func resourceDropTable(runner QueryRunner, mapping string) errors.ServiceError {
	_, err := runner.Exec("DROP TABLE " + mapping)

	return handleDbError(err)
}

func resourceListEntities(ctx context.Context, runner QueryRunner) (result []string, err errors.ServiceError) {
	rows, sqlErr := runner.QueryContext(ctx, `select table_schema || '.' || table_name from information_schema.tables`)
	err = handleDbError(sqlErr)

	if err != nil {
		return
	}

	for rows.Next() {
		var entityName = new(string)

		err = handleDbError(rows.Scan(entityName))

		if err != nil {
			return
		}

		result = append(result, *entityName)
	}

	return
}

func isAuditColumn(column string) bool {
	return column == "created_on" || column == "updated_on" || column == "created_by" || column == "updated_by" || column == "version"
}
