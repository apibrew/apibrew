package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations2 "github.com/tislib/data-handler/pkg/service/annotations"
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

func resourceCreateTable(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig, false))

	builder.IfNotExists()

	if !annotations2.IsEnabled(resource, annotations2.DoPrimaryKeyLookup) {
		builder.Define("id", "uuid", "NOT NULL", "PRIMARY KEY")
	}

	prepareCreateTableQuery(resource, builder)

	// audit
	if !annotations2.IsEnabled(resource, annotations2.DisableAudit) {
		builder.Define("created_on", "timestamp", "NOT NULL")
		builder.Define("updated_on", "timestamp", "NULL")
		builder.Define("created_by", DbNameType, "NOT NULL")
		builder.Define("updated_by", DbNameType, "NULL")
		// version
		builder.Define("version", "int2", "NOT NULL")
	}

	sqlQuery, _ := builder.Build()
	_, err := runner.Exec(sqlQuery)

	logger.Trace("sqlQuery: ", sqlQuery)

	return handleDbError(ctx, err)
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

		var def = []string{fmt.Sprintf("\"%s\"", sourceConfig.Mapping.Mapping), sqlType, nullModifier, uniqModifier}

		return strings.Join(def, " ")
	}

	return ""
}

func resourceCreateHistoryTable(ctx context.Context, runner QueryRunner, resource *model.Resource) errors.ServiceError {
	builder := sqlbuilder.CreateTable(getTableName(resource.SourceConfig, true))

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

	return handleDbError(ctx, err)
}

func resourceDropTable(ctx context.Context, runner QueryRunner, mapping string) errors.ServiceError {
	_, err := runner.Exec("DROP TABLE " + mapping)

	return handleDbError(ctx, err)
}

func resourceListEntities(ctx context.Context, runner QueryRunner) (result []string, err errors.ServiceError) {
	rows, sqlErr := runner.QueryContext(ctx, `select table_schema || '.' || table_name from information_schema.tables`)
	err = handleDbError(ctx, sqlErr)

	if err != nil {
		return
	}

	for rows.Next() {
		var entityName = new(string)

		err = handleDbError(ctx, rows.Scan(entityName))

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
