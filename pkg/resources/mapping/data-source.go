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

	if options, ok := dataSource.Options.(*model.DataSource_PostgresqlParams); ok {
		properties["options_postgres_username"] = structpb.NewStringValue(options.PostgresqlParams.Username)
		properties["options_postgres_password"] = structpb.NewStringValue(options.PostgresqlParams.Password)
		properties["options_postgres_host"] = structpb.NewStringValue(options.PostgresqlParams.Host)
		properties["options_postgres_port"] = structpb.NewNumberValue(float64(options.PostgresqlParams.Port))
		properties["options_postgres_db_name"] = structpb.NewStringValue(options.PostgresqlParams.DbName)
		properties["options_postgres_default_schema"] = structpb.NewStringValue(options.PostgresqlParams.DefaultSchema)
	}

	if options, ok := dataSource.Options.(*model.DataSource_MysqlParams); ok {
		properties["options_mysql_username"] = structpb.NewStringValue(options.MysqlParams.Username)
		properties["options_mysql_password"] = structpb.NewStringValue(options.MysqlParams.Password)
		properties["options_mysql_host"] = structpb.NewStringValue(options.MysqlParams.Host)
		properties["options_mysql_port"] = structpb.NewNumberValue(float64(options.MysqlParams.Port))
		properties["options_mysql_db_name"] = structpb.NewStringValue(options.MysqlParams.DbName)
		properties["options_mysql_default_schema"] = structpb.NewStringValue(options.MysqlParams.DefaultSchema)
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

		options.PostgresqlParams = &model.PostgresqlOptions{
			Username:      record.Properties["options_postgres_username"].GetStringValue(),
			Password:      record.Properties["options_postgres_password"].GetStringValue(),
			Host:          record.Properties["options_postgres_host"].GetStringValue(),
			Port:          uint32(record.Properties["options_postgres_port"].GetNumberValue()),
			DbName:        record.Properties["options_postgres_db_name"].GetStringValue(),
			DefaultSchema: record.Properties["options_postgres_default_schema"].GetStringValue(),
		}

		result.Options = options
	}

	if result.Backend == model.DataSourceBackendType_MYSQL {
		options := new(model.DataSource_MysqlParams)

		options.MysqlParams = &model.MysqlOptions{
			Username:      record.Properties["options_mysql_username"].GetStringValue(),
			Password:      record.Properties["options_mysql_password"].GetStringValue(),
			Host:          record.Properties["options_mysql_host"].GetStringValue(),
			Port:          uint32(record.Properties["options_mysql_port"].GetNumberValue()),
			DbName:        record.Properties["options_mysql_db_name"].GetStringValue(),
			DefaultSchema: record.Properties["options_mysql_default_schema"].GetStringValue(),
		}

		result.Options = options
	}

	return result
}
