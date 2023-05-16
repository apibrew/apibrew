package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/structpb"
)

type mysqlBackendOptions struct {
	connectionDetails *model.DataSource_MysqlParams
	dataSource        *model.DataSource
	handleDbError     func(ctx context.Context, err error) errors.ServiceError
}

func (p *mysqlBackendOptions) UseDbHandleError(f func(ctx context.Context, err error) errors.ServiceError) {
	p.handleDbError = f
}

func (p mysqlBackendOptions) GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor {
	return func(ctx context.Context, runner helper.QueryRunner, schema *abs.Schema, params abs.UpgradeResourceParams, forceMigration bool) helper.ResourceMigrationBuilder {
		return &resourceMigrationBuilder{handleDbError: p.handleDbError, schema: schema, options: p, ctx: ctx, runner: runner, params: params, forceMigration: forceMigration, tableName: p.GetFullTableName(params.MigrationPlan.CurrentResource.SourceConfig)}
	}
}

func (p mysqlBackendOptions) HandleError(err error) (errors.ServiceError, bool) {
	if pqErr, ok := err.(*mysql.MySQLError); ok {
		return p.handlePqErr(pqErr), true
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
	return "public"
}

func (p mysqlBackendOptions) handlePqErr(err *mysql.MySQLError) errors.ServiceError {
	switch err.Error() {
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

func (p mysqlBackendOptions) GetDriverName() string {
	return "mysql"
}

func (p mysqlBackendOptions) GetConnectionString() string {
	params := p.connectionDetails.MysqlParams
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", params.GetUsername(), params.GetPassword(), params.GetHost(), params.GetPort(), params.GetDbName())
}

func (p *mysqlBackendOptions) GetFullTableName(sourceConfig *model.ResourceSourceConfig) string {
	var tableName = sourceConfig.Entity
	def := ""
	if sourceConfig.Catalog != "" {
		def = fmt.Sprintf("%s.%s", p.Quote(sourceConfig.Catalog), p.Quote(tableName))
	} else {
		def = p.Quote(sourceConfig.Entity)
	}

	return def
}

func (p *mysqlBackendOptions) DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, errors.ServiceError) {
	propertyType := types.ByResourcePropertyType(property.Type)
	var val interface{}

	v, err := propertyType.UnPack(packedVal)
	if err != nil {
		return nil, errors.LogicalError.WithDetails(err.Error())
	}

	if property.Type == model.ResourceProperty_TIME {
		return propertyType.String(v), nil
	} else if property.Type == model.ResourceProperty_STRUCT || property.Type == model.ResourceProperty_OBJECT || property.Type == model.ResourceProperty_ENUM || property.Type == model.ResourceProperty_MAP || property.Type == model.ResourceProperty_LIST {
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
