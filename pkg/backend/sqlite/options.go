package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	sqliteDdriver "github.com/mattn/go-sqlite3"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

type sqliteBackendOptions struct {
	connectionDetails *model.DataSource_SqliteParams
	dataSource        *model.DataSource
	handleDbError     func(ctx context.Context, err error) errors.ServiceError
}

func (p *sqliteBackendOptions) UseDbHandleError(f func(ctx context.Context, err error) errors.ServiceError) {
	p.handleDbError = f
}

func (p sqliteBackendOptions) GetResourceMigrationBuilderConstructor() helper.ResourceMigrationBuilderConstructor {
	return func(ctx context.Context, runner helper.QueryRunner, params abs.UpgradeResourceParams, forceMigration bool) helper.ResourceMigrationBuilder {
		return &resourceMigrationBuilder{handleDbError: p.handleDbError, options: p, ctx: ctx, runner: runner, params: params, forceMigration: forceMigration, tableName: p.GetFullTableName(params.MigrationPlan.CurrentResource.SourceConfig)}
	}
}

func (p sqliteBackendOptions) HandleError(err error) (errors.ServiceError, bool) {
	if pqErr, ok := err.(*sqliteDdriver.Error); ok {
		return p.handlePqErr(pqErr), true
	}

	return nil, false
}

func (p sqliteBackendOptions) Quote(str string) string {
	return fmt.Sprintf("`%s`", str)
}

func (p sqliteBackendOptions) GetFlavor() sqlbuilder.Flavor {
	return sqlbuilder.MySQL
}

func (p sqliteBackendOptions) GetDefaultCatalog() string {
	return "public"
}

func (p sqliteBackendOptions) handlePqErr(err *sqliteDdriver.Error) errors.ServiceError {
	switch err.Error() {
	case "28000":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Error())
	case "28P01":
		return errors.BackendConnectionAuthenticationError.WithMessage(err.Error())
	case "23505":
		return errors.UniqueViolation.WithDetails(err.Error())
	case "23503":
		return errors.ReferenceViolation.WithDetails(err.Error())
	default:
		return errors.InternalError.WithMessage(err.Error())
	}
}

func (p sqliteBackendOptions) GetDriverName() string {
	return "sqlite"
}

func (p sqliteBackendOptions) GetConnectionString() string {
	params := p.connectionDetails.SqliteParams
	return params.Path
}

func (p *sqliteBackendOptions) GetFullTableName(sourceConfig *model.ResourceSourceConfig) string {
	var tableName = sourceConfig.Entity
	def := ""
	if sourceConfig.Catalog != "" {
		def = fmt.Sprintf("%s.%s", p.Quote(sourceConfig.Catalog), p.Quote(tableName))
	} else {
		def = p.Quote(sourceConfig.Entity)
	}

	return def
}

func (p *sqliteBackendOptions) DbEncode(property *model.ResourceProperty, packedVal *structpb.Value) (interface{}, errors.ServiceError) {
	propertyType := types.ByResourcePropertyType(property.Type)
	var val interface{}

	v, err := propertyType.UnPack(packedVal)
	if err != nil {
		return nil, errors.LogicalError.WithDetails(err.Error())
	}

	if property.Type == model.ResourceProperty_TIME {
		return propertyType.String(v), nil
	} else if property.Type == model.ResourceProperty_OBJECT || property.Type == model.ResourceProperty_ENUM || property.Type == model.ResourceProperty_MAP || property.Type == model.ResourceProperty_LIST {
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
