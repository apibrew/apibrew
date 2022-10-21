package test

import (
	"data-handler/stub/model"
)

func prepareInitData() *model.InitData {
	return &model.InitData{
		Config: &model.AppConfig{
			GrpcAddr: "localhost:17981",
			HttpAddr: "localhost:17982",
		},
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
	return systemDataSource
}
