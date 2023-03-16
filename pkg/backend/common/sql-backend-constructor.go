package common

import (
	"database/sql"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

type sqlBackend struct {
	connection     *sql.DB
	transactionMap map[string]*txData
	dataSourceName string
	options        SqlBackendOptions
}

type SqlBackendOptions interface {
	GetConnectionString() string
	GetSql(s string) string
	GetDriverName() string
	HandleError(err error) (errors.ServiceError, bool)
	GetSqlTypeFromProperty(propertyType model.ResourceProperty_Type, length uint32) string
	GetPropertyTypeFromPsql(columnType string) model.ResourceProperty_Type
	Quote(str string) string
	GetFlavor() sqlbuilder.Flavor
	GetDefaultCatalog() string
}

func NewSqlBackend(dataSource *model.DataSource, options SqlBackendOptions) abs.Backend {
	return &sqlBackend{
		options:        options,
		dataSourceName: dataSource.Name,
		transactionMap: make(map[string]*txData),
	}
}
