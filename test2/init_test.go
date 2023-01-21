package test2

import (
	"data-handler/logging"
	"data-handler/model"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.TraceLevel)
	//log.SetReportCaller(true)

	logging.SetupGrayLog("tiswork.tisserv.net:12201", "test")
}

func prepareInitData() *model.InitData {
	return &model.InitData{
		Config: &model.AppConfig{
			Host:                  "localhost",
			Port:                  17981,
			JwtPrivateKey:         "../data/jwt.key",
			JwtPublicKey:          "../data/jwt.key.pub",
			DisableAuthentication: true,
			DisableCache:          true,
		},
		SystemDataSource: prepareSystemDataSource(),
		SystemNamespace:  prepareSystemNamespace(),
		InitDataSources:  prepareInitDataSources(),
		InitNamespaces:   prepareInitNamespaces(),
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

func prepareInitNamespaces() []*model.Namespace {
	return nil
}

func prepareInitDataSources() []*model.DataSource {
	return nil
}

func prepareSystemNamespace() *model.Namespace {
	return &model.Namespace{
		Name: "system",
		Type: model.DataType_SYSTEM,
	}
}

func prepareSystemDataSource() *model.DataSource {
	return systemDataSource
}
