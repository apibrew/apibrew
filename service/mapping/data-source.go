package mapping

import (
	"data-handler/model"
	"data-handler/service/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func DataSourceToRecord(dataSource *model.DataSource) *model.Record {
	properties := make(map[string]interface{})

	properties["name"] = dataSource.Name
	properties["description"] = dataSource.Description
	properties["backend"] = int(dataSource.Backend)

	if options, ok := dataSource.Options.(*model.DataSource_PostgresqlParams); ok {
		properties["options_postgres_username"] = options.PostgresqlParams.Username
		properties["options_postgres_password"] = options.PostgresqlParams.Password
		properties["options_postgres_host"] = options.PostgresqlParams.Host
		properties["options_postgres_port"] = options.PostgresqlParams.Port
		properties["options_postgres_db_name"] = options.PostgresqlParams.DbName
		properties["options_postgres_default_schema"] = options.PostgresqlParams.DefaultSchema
	}

	structProperties, err := structpb.NewStruct(properties)

	if err != nil {
		panic(err)
	}

	return &model.Record{
		Id:         dataSource.Id,
		Resource:   system.DataSourceResource.Name,
		Type:       dataSource.Type,
		Properties: structProperties,
		AuditData:  dataSource.AuditData,
		Version:    dataSource.Version,
	}
}

func DataSourceFromRecord(record *model.Record) *model.DataSource {
	if record == nil {
		return nil
	}

	backendNumber := record.Properties.Fields["backend"].GetNumberValue()

	result := &model.DataSource{
		Id:          record.Id,
		Type:        record.Type,
		Backend:     model.DataSourceBackendType(backendNumber),
		Name:        record.Properties.Fields["name"].GetStringValue(),
		Description: record.Properties.Fields["description"].GetStringValue(),
		AuditData:   record.AuditData,
		Version:     record.Version,
	}

	if result.Backend == model.DataSourceBackendType_POSTGRESQL {
		options := new(model.DataSource_PostgresqlParams)

		options.PostgresqlParams = &model.PostgresqlOptions{
			Username:      record.Properties.Fields["options_postgres_username"].GetStringValue(),
			Password:      record.Properties.Fields["options_postgres_password"].GetStringValue(),
			Host:          record.Properties.Fields["options_postgres_host"].GetStringValue(),
			Port:          uint32(record.Properties.Fields["options_postgres_port"].GetNumberValue()),
			DbName:        record.Properties.Fields["options_postgres_db_name"].GetStringValue(),
			DefaultSchema: record.Properties.Fields["options_postgres_default_schema"].GetStringValue(),
		}

		result.Options = options
	}

	return result
}
