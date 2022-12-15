package test

import (
	"data-handler/model"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.TraceLevel)
	//log.SetReportCaller(true)
}

func prepareInitData() *model.InitData {
	return &model.InitData{
		Config: &model.AppConfig{
			GrpcAddr:              "localhost:17981",
			HttpAddr:              "localhost:17982",
			JwtPrivateKey:         "../data/jwt.key",
			JwtPublicKey:          "../data/jwt.key.pub",
			DisableAuthentication: true,
			DisableCache:          true,
		},
		SystemDataSource: prepareSystemDataSource(),
		SystemWorkspace:  prepareSystemWorkspace(),
		InitDataSources:  prepareInitDataSources(),
		InitWorkspaces:   prepareInitWorkspaces(),
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
	return nil
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
	return systemDataSource
}
