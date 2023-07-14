package common

import (
	"context"
	"database/sql"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"sync"
)

type sqlBackend struct {
	connection     *sql.DB
	transactionMap map[string]*txData
	dataSourceName string
	options        SqlBackendOptions
	schema         *abs.Schema
	mu             sync.Mutex
}

func (p *sqlBackend) SetSchema(schema *abs.Schema) {
	p.schema = schema
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

func NewSqlBackend(dataSource abs.DataSource, options SqlBackendOptions) abs.Backend {
	backend := &sqlBackend{
		options:        options,
		dataSourceName: dataSource.GetName(),
		transactionMap: make(map[string]*txData),
	}

	options.UseDbHandleError(backend.handleDbError)

	return backend
}
