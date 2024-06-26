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
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/protobuf/types/known/structpb"
)

type mysqlBackendOptions struct {
	dataSource    *resource_model.DataSource
	handleDbError func(ctx context.Context, err error) error
}

func (p *mysqlBackendOptions) UseDbHandleError(f func(ctx context.Context, err error) error) {
	p.handleDbError = f
}

func (p mysqlBackendOptions) GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor {
	return func(ctx context.Context, runner helper.QueryRunner, schema *abs.Schema, params abs.UpgradeResourceParams, forceMigration bool) helper.ResourceMigrationBuilder {
		return &resourceMigrationBuilder{handleDbError: p.handleDbError, schema: schema, options: p, ctx: ctx, runner: runner, params: params, forceMigration: forceMigration, tableName: p.GetFullTableName(params.MigrationPlan.CurrentResource.SourceConfig)}
	}
}

func (p mysqlBackendOptions) HandleError(err error) (error, bool) {
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

func (p mysqlBackendOptions) handlePqErr(err *mysql.MySQLError) error {
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
	username := p.dataSource.Options["username"]
	password := p.dataSource.Options["password"]
	host := p.dataSource.Options["host"]
	port := p.dataSource.Options["port"]
	dbName := p.dataSource.Options["db_name"]

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)
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

func (p *mysqlBackendOptions) DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, error) {
	if packedVal == nil {
		return nil, nil
	}

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
