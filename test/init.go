package test

import (
	"data-handler/stub/model"
)

func prepareInitData() *model.InitData {
	return &model.InitData{
		SystemDataSource: prepareSystemDataSource(),
		SystemWorkSpace:  prepareSystemWorkSpace(),
		InitDataSources:  prepareInitDataSources(),
		InitWorkSpaces:   prepareInitWorkSpaces(),
		InitUsers:        prepareInitUsers(),
		InitResources:    prepareInitResources(),
		InitRecords:      prepareInitRecords(),
	}
}

func prepareInitRecords() []*model.Record {
	return nil
}

func prepareInitResources() []*model.Resource {
	return nil
}

func prepareInitUsers() []*model.User {
	return []*model.User{
		{
			Type:     model.DataType_STATIC,
			Username: "admin",
			Password: "admin",
			Scopes: []string{
				"super-user",
			},
		},
	}
}

func prepareInitWorkSpaces() []*model.Workspace {
	return nil
}

func prepareInitDataSources() []*model.DataSource {
	return nil
}

func prepareSystemWorkSpace() *model.Workspace {
	return &model.Workspace{
		Name: "system",
		Type: model.DataType_SYSTEM,
	}
}

func prepareSystemDataSource() *model.DataSource {
	return &model.DataSource{
		Id:      "system",
		Backend: model.DataSourceBackend_POSTGRESQL,
		Type:    model.DataType_SYSTEM,
		Options: &model.DataSource_PostgresqlParams{
			PostgresqlParams: &model.PostgresqlOptions{
				Username:      "postgres",
				Password:      "postgres",
				Host:          "127.0.0.1",
				Port:          55432,
				DbName:        "postgres",
				DefaultSchema: "public",
			},
		},
	}
}
