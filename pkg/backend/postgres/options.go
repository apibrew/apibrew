package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/backend/helper"
	"github.com/tislib/apibrew/pkg/backend/sqlbuilder"
	"github.com/tislib/apibrew/pkg/errors"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

type postgreSqlBackendOptions struct {
	connectionDetails *model.DataSource_PostgresqlParams
	handleDbError     func(ctx context.Context, err error) errors.ServiceError
}

func (p *postgreSqlBackendOptions) UseDbHandleError(f func(ctx context.Context, err error) errors.ServiceError) {
	p.handleDbError = f
}

func (p postgreSqlBackendOptions) GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor {
	return func(ctx context.Context, runner helper.QueryRunner, schema *abs.Schema, params abs.UpgradeResourceParams, forceMigration bool) helper.ResourceMigrationBuilder {
		return &resourceMigrationBuilder{handleDbError: p.handleDbError, options: p, ctx: ctx, runner: runner, params: params, forceMigration: forceMigration, tableName: p.GetFullTableName(params.MigrationPlan.CurrentResource.SourceConfig), schema: schema}
	}
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

func (p *postgreSqlBackendOptions) GetFullTableName(sourceConfig *model.ResourceSourceConfig) string {
	var tableName = sourceConfig.Entity
	def := ""
	if sourceConfig.Catalog != "" {
		def = fmt.Sprintf("%s.%s", p.Quote(sourceConfig.Catalog), p.Quote(tableName))
	} else {
		def = p.Quote(sourceConfig.Entity)
	}

	return def
}

func (p *postgreSqlBackendOptions) DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, errors.ServiceError) {
	propertyType := types.ByResourcePropertyType(property.Type)
	var val interface{}

	if property.Type == model.ResourceProperty_OBJECT || property.Type == model.ResourceProperty_STRUCT || property.Type == model.ResourceProperty_ENUM || property.Type == model.ResourceProperty_MAP || property.Type == model.ResourceProperty_LIST {
		var err error
		val, err = json.Marshal(packedVal.AsInterface())

		if err != nil {
			return nil, errors.InternalError.WithDetails(err.Error())
		}
		val = string(val.([]byte))
	} else {
		var err error
		val, err = propertyType.UnPack(packedVal)

		if err != nil {
			return nil, errors.InternalError.WithDetails(err.Error())
		}
	}
	return val, nil
}
