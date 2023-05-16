package setup

import (
	"github.com/apibrew/apibrew/pkg/model"
)

func prepareInitData() *model.InitData {
	return &model.InitData{
		Config: &model.AppConfig{
			Host:                  "localhost",
			Port:                  17981,
			DisableAuthentication: false,
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
	return []*model.User{
		{
			Username: "admin",
			Password: "admin",
			SecurityContext: &model.SecurityContext{
				Constraints: []*model.SecurityConstraint{
					{
						Operation: model.OperationType_FULL,
						Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
					},
				},
			},
		},
		{
			Username: "dh_test",
			Password: "dh_test",
			SecurityContext: &model.SecurityContext{
				Constraints: []*model.SecurityConstraint{
					{
						Resource:  "user",
						Operation: model.OperationType_OPERATION_TYPE_READ,
						Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
					},
					{
						Resource:  "namespace",
						Operation: model.OperationType_OPERATION_TYPE_CREATE,
						Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
					},
					{
						Resource:  "namespace",
						Operation: model.OperationType_OPERATION_TYPE_READ,
						Permit:    model.PermitType_PERMIT_TYPE_ALLOW,
					},
				},
			},
		},
	}
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
	}
}

func prepareSystemDataSource() *model.DataSource {
	return SystemDataSource
}
