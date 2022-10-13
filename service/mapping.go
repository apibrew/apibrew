package service

import (
	"data-handler/stub/model"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func mapToRecord[T proto.Message](list []T, mapper func(T) *model.Record) []*model.Record {
	var result []*model.Record

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}

func mapFromRecord[T proto.Message](list []*model.Record, mapper func(*model.Record) T) []T {
	var result []T

	for _, item := range list {
		result = append(result, mapper(item))
	}

	return result
}

func dataSourceToRecord(department *model.DataSource) *model.Record {
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
		Resource:   dataSourceResource.Name,
		Type:       model.DataType_SYSTEM,
		Properties: structProperties,
		AuditData:  department.AuditData,
		Version:    department.Version,
	}
}

func dataSourceFromRecord(record *model.Record) *model.DataSource {
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
