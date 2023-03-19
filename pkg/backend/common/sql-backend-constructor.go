package common

import (
	"context"
	"database/sql"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

type sqlBackend struct {
	connection     *sql.DB
	transactionMap map[string]*txData
	dataSourceName string
	options        SqlBackendOptions
}

type SqlBackendOptions interface {
	UseDbHandleError(func(ctx context.Context, err error) errors.ServiceError)
	GetConnectionString() string
	GetSql(s string) string
	GetDriverName() string
	HandleError(err error) (errors.ServiceError, bool)
	GetSqlTypeFromProperty(propertyType model.ResourceProperty_Type, length uint32) string
	GetPropertyTypeFromPsql(columnType string) model.ResourceProperty_Type
	Quote(str string) string
	GetFlavor() sqlbuilder.Flavor
	GetDefaultCatalog() string
	GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor
	GetFullTableName(config *model.ResourceSourceConfig) string
	DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, errors.ServiceError)
	TypeModifier(propertyType model.ResourceProperty_Type) types.PropertyType
}

func NewSqlBackend(dataSource *model.DataSource, options SqlBackendOptions) abs.Backend {
	backend := &sqlBackend{
		options:        options,
		dataSourceName: dataSource.Name,
		transactionMap: make(map[string]*txData),
	}

	options.UseDbHandleError(backend.handleDbError)

	return backend
}
