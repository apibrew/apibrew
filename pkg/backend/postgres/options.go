package postgres

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
)

type postgreSqlBackendOptions struct {
	connectionDetails *model.DataSource_PostgresqlParams
}

func (p postgreSqlBackendOptions) HandleError(err error) (errors.ServiceError, bool) {
	if pqErr, ok := err.(*pq.Error); ok {
		return p.handlePqErr(pqErr), true
	}

	return nil, false
}

func (p postgreSqlBackendOptions) Quote(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}

func (p postgreSqlBackendOptions) GetFlavor() sqlbuilder.Flavor {
	return sqlbuilder.PostgreSQL
}

func (p postgreSqlBackendOptions) GetDefaultCatalog() string {
	return "public"
}

func (p postgreSqlBackendOptions) handlePqErr(err *pq.Error) errors.ServiceError {
	switch err.Code {
	case "28000":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "28P01":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message)
	case "23505":
		return errors.UniqueViolation.WithDetails(err.Message)
	case "23503":
		return errors.ReferenceViolation.WithDetails(err.Message)
	default:
		return errors.InternalError.WithMessage(err.Message)
	}
}

func (p postgreSqlBackendOptions) GetDriverName() string {
	return "postgres"
}

func (p postgreSqlBackendOptions) GetConnectionString() string {
	params := p.connectionDetails.PostgresqlParams
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", params.GetUsername(), params.GetPassword(), params.GetHost(), params.GetPort(), params.GetDbName())
}
