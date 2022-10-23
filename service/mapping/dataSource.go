package mapping

import (
	"data-handler/service/system"
	"data-handler/stub/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func DataSourceToRecord(department *model.DataSource) *model.Record {
	properties := make(map[string]interface{})

	properties["backend"] = int(department.Backend)

	if options, ok := department.Options.(*model.DataSource_PostgresqlParams); ok {
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
		Id:         department.Id,
		Resource:   system.DataSourceResource.Name,
		Type:       department.Type,
		Properties: structProperties,
		AuditData:  department.AuditData,
		Version:    department.Version,
	}
}

func DataSourceFromRecord(record *model.Record) *model.DataSource {
	if record == nil {
		return nil
	}

	backendNumnber := record.Properties.Fields["backend"].GetNumberValue()

	result := &model.DataSource{
		Id:        record.Id,
		Type:      record.Type,
		Backend:   model.DataSourceBackend(backendNumnber),
		AuditData: record.AuditData,
		Version:   record.Version,
	}

	if result.Backend == model.DataSourceBackend_POSTGRESQL {
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
