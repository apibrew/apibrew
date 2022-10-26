package main

import (
	"data-handler/model"
	"data-handler/util"
)

func main() {
	initData := prepareInitData()

	util.Write("/Users/taleh/Projects/tiswork/data-handler/data/init.pb", initData)
	util.WriteJson("/Users/taleh/Projects/tiswork/data-handler/data/init.json", initData)
}

func prepareInitData() *model.InitData {
	return &model.InitData{
		Config:           prepareAppConfig(),
		SystemDataSource: prepareSystemDataSource(),
		SystemWorkspace:  prepareSystemWorkspace(),
		InitDataSources:  prepareInitDataSources(),
		InitWorkspaces:   prepareInitWorkspaces(),
		InitUsers:        prepareInitUsers(),
		InitResources:    prepareInitResources(),
		InitRecords:      prepareInitRecords(),
	}
}

func prepareAppConfig() *model.AppConfig {
	return &model.AppConfig{
		GrpcAddr:      "0.0.0.0:9009",
		HttpAddr:      "0.0.0.0:8008",
		JwtPrivateKey: "data/jwt.key",
		JwtPublicKey:  "data/jwt.key.pub",
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

func prepareInitWorkspaces() []*model.Workspace {
	return nil
}

func prepareInitDataSources() []*model.DataSource {
	return nil
}

func prepareSystemWorkspace() *model.Workspace {
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
				Username:      "root",
				Password:      "52fa536f0c5b85f9d806633937f06446",
				Host:          "tiswork.tisserv.net",
				Port:          5432,
				DbName:        "dh",
				DefaultSchema: "public",
			},
		},
	}
}
