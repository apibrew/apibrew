package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/structpb"
)

type postgreSqlBackendOptions struct {
	handleDbError func(ctx context.Context, err error) errors.ServiceError
	dataSource    *resource_model.DataSource
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
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message).WithDetails(err.Detail)
	case "28P01":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Message).WithDetails(err.Detail)
	case "23505":
		return errors.UniqueViolation.WithDetails(err.Message).WithDetails(err.Detail)
	case "23503":
		return errors.ReferenceViolation.WithDetails(err.Message).WithDetails(err.Detail)
	default:
		return errors.InternalError.WithMessage(err.Message).WithDetails(err.Detail)
	}
}

func (p postgreSqlBackendOptions) GetDriverName() string {
	return "postgres"
}

func (p postgreSqlBackendOptions) GetConnectionString() string {
	username := p.dataSource.Options["username"]
	password := p.dataSource.Options["password"]
	host := p.dataSource.Options["host"]
	port := p.dataSource.Options["port"]
	dbName := p.dataSource.Options["db_name"]

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)
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
	if packedVal == nil || packedVal.AsInterface() == nil {
		return nil, nil
	}

	propertyType := types.ByResourcePropertyType(property.Type)
	var val interface{}

	if property.Type == model.ResourceProperty_OBJECT || property.Type == model.ResourceProperty_STRUCT || property.Type == model.ResourceProperty_MAP || property.Type == model.ResourceProperty_LIST {
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
