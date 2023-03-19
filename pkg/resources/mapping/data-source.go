package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func DataSourceToRecord(dataSource *model.DataSource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(dataSource.Name)
	properties["description"] = structpb.NewStringValue(dataSource.Description)
	properties["backend"] = structpb.NewNumberValue(float64(dataSource.Backend))

	if options, ok := dataSource.Params.(*model.DataSource_PostgresqlParams); ok {
		properties["options_postgres_username"] = structpb.NewStringValue(options.PostgresqlParams.Username)
		properties["options_postgres_password"] = structpb.NewStringValue(options.PostgresqlParams.Password)
		properties["options_postgres_host"] = structpb.NewStringValue(options.PostgresqlParams.Host)
		properties["options_postgres_port"] = structpb.NewNumberValue(float64(options.PostgresqlParams.Port))
		properties["options_postgres_db_name"] = structpb.NewStringValue(options.PostgresqlParams.DbName)
		properties["options_postgres_default_schema"] = structpb.NewStringValue(options.PostgresqlParams.DefaultSchema)
	}

	if options, ok := dataSource.Params.(*model.DataSource_MysqlParams); ok {
		properties["options_mysql_username"] = structpb.NewStringValue(options.MysqlParams.Username)
		properties["options_mysql_password"] = structpb.NewStringValue(options.MysqlParams.Password)
		properties["options_mysql_host"] = structpb.NewStringValue(options.MysqlParams.Host)
		properties["options_mysql_port"] = structpb.NewNumberValue(float64(options.MysqlParams.Port))
		properties["options_mysql_db_name"] = structpb.NewStringValue(options.MysqlParams.DbName)
		properties["options_mysql_default_schema"] = structpb.NewStringValue(options.MysqlParams.DefaultSchema)
	}

	if options, ok := dataSource.Params.(*model.DataSource_RedisParams); ok {
		properties["options_redis_addr"] = structpb.NewStringValue(options.RedisParams.Addr)
		properties["options_redis_password"] = structpb.NewStringValue(options.RedisParams.Password)
		properties["options_redis_db"] = structpb.NewNumberValue(float64(options.RedisParams.Db))
	}

	if options, ok := dataSource.Params.(*model.DataSource_MongoParams); ok {
		properties["options_mongo_uri"] = structpb.NewStringValue(options.MongoParams.Uri)
		properties["options_mongo_db_name"] = structpb.NewStringValue(options.MongoParams.DbName)
	}

	return &model.Record{
		Id:         dataSource.Id,
		Properties: properties,
		AuditData:  dataSource.AuditData,
		Version:    dataSource.Version,
	}
}

func DataSourceFromRecord(record *model.Record) *model.DataSource {
	if record == nil {
		return nil
	}

	backendNumber := record.Properties["backend"].GetNumberValue()

	result := &model.DataSource{
		Id:          record.Id,
		Backend:     model.DataSourceBackendType(backendNumber),
		Name:        record.Properties["name"].GetStringValue(),
		Description: record.Properties["description"].GetStringValue(),
		AuditData:   record.AuditData,
		Version:     record.Version,
	}

	if result.Backend == model.DataSourceBackendType_POSTGRESQL {
		options := new(model.DataSource_PostgresqlParams)

		options.PostgresqlParams = &model.PostgresqlParams{
			Username:      record.Properties["options_postgres_username"].GetStringValue(),
			Password:      record.Properties["options_postgres_password"].GetStringValue(),
			Host:          record.Properties["options_postgres_host"].GetStringValue(),
			Port:          uint32(record.Properties["options_postgres_port"].GetNumberValue()),
			DbName:        record.Properties["options_postgres_db_name"].GetStringValue(),
			DefaultSchema: record.Properties["options_postgres_default_schema"].GetStringValue(),
		}

		result.Params = options
	}

	if result.Backend == model.DataSourceBackendType_MYSQL {
		options := new(model.DataSource_MysqlParams)

		options.MysqlParams = &model.MysqlParams{
			Username:      record.Properties["options_mysql_username"].GetStringValue(),
			Password:      record.Properties["options_mysql_password"].GetStringValue(),
			Host:          record.Properties["options_mysql_host"].GetStringValue(),
			Port:          uint32(record.Properties["options_mysql_port"].GetNumberValue()),
			DbName:        record.Properties["options_mysql_db_name"].GetStringValue(),
			DefaultSchema: record.Properties["options_mysql_default_schema"].GetStringValue(),
		}

		result.Params = options
	}

	if result.Backend == model.DataSourceBackendType_REDIS {
		options := new(model.DataSource_RedisParams)

		options.RedisParams = &model.RedisParams{
			Addr:     record.Properties["options_redis_addr"].GetStringValue(),
			Password: record.Properties["options_redis_password"].GetStringValue(),
			Db:       int32(record.Properties["options_redis_db"].GetNumberValue()),
		}

		result.Params = options
	}

	if result.Backend == model.DataSourceBackendType_MONGODB {
		options := new(model.DataSource_MongoParams)

		options.MongoParams = &model.MongoParams{
			Uri: record.Properties["options_mongo_uri"].GetStringValue(),
		}

		result.Params = options
	}

	return result
}
