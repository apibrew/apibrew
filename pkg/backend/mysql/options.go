package mysql

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

type mysqlBackendOptions struct {
	connectionDetails *model.DataSource_MysqlParams
	dataSource        *model.DataSource
}

func (p mysqlBackendOptions) HandleError(err error) (errors.ServiceError, bool) {
	if pqErr, ok := err.(*mysql.MySQLError); ok {
		return p.handleMysqlErr(pqErr), true
	}

	return nil, false
}

func (p mysqlBackendOptions) Quote(str string) string {
	return fmt.Sprintf("`%s`", str)
}

func (p mysqlBackendOptions) GetFlavor() sqlbuilder.Flavor {
	return sqlbuilder.MySQL
}

func (p mysqlBackendOptions) GetDefaultCatalog() string {
	return p.dataSource.Options.(*model.DataSource_MysqlParams).MysqlParams.DbName
}

func (p mysqlBackendOptions) handleMysqlErr(err *mysql.MySQLError) errors.ServiceError {
	switch err.Number {
	//case "28000":
	//	return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	//case "28P01":
	//	return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	//case "23505":
	//	return errors.UniqueViolation.WithDetails(err.Message)
	//case "23503":
	//	return errors.ReferenceViolation.WithDetails(err.Message)
	default:
		return errors.InternalError.WithMessage(err.Message)
	}
}

func (p mysqlBackendOptions) GetDriverName() string {
	return "mysql"
}

func (p mysqlBackendOptions) GetConnectionString() string {
	//return fmt.Sprintf("mysql://%s%s@tcp(%s:%d)/%s", params.GetUsername(), params.GetPassword(), params.GetHost(), params.GetPort(), params.GetDbName())
	return "root@tcp(localhost:3306)/dh_test"
}
