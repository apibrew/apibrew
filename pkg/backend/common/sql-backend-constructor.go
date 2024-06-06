package common

import (
	"context"
	"database/sql"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
)

type sqlBackend struct {
	connection     *sql.DB
	dataSourceName string
	options        SqlBackendOptions
	schema         *abs.Schema
}

func (p *sqlBackend) SetSchema(schema *abs.Schema) {
	p.schema = schema
}

type SqlBackendOptions interface {
	UseDbHandleError(func(ctx context.Context, err error) error)
	GetConnectionString() string
	GetSql(s string) string
	GetDriverName() string
	HandleError(err error) (error, bool)
	GetSqlTypeFromProperty(propertyType model.ResourceProperty_Type, length uint32) string
	GetPropertyTypeFromPsql(columnType string) model.ResourceProperty_Type
	Quote(str string) string
	GetFlavor() sqlbuilder.Flavor
	GetDefaultCatalog() string
	GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor
	GetFullTableName(config *model.ResourceSourceConfig) string
	DbEncode(property *model.ResourceProperty, packedVal interface{}) (interface{}, error)
	TypeModifier(propertyType model.ResourceProperty_Type) types.PropertyType
}

func NewSqlBackend(dataSource abs.DataSource, options SqlBackendOptions) abs.Backend {
	backend := &sqlBackend{
		options:        options,
		dataSourceName: dataSource.GetName(),
	}

	options.UseDbHandleError(backend.handleDbError)

	return backend
}
