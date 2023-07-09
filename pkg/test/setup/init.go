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
		InitRoles:        prepareInitRoles(),
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

func prepareInitRoles() []*model.Role {
	return []*model.Role{
		{
			Name: "test_user",
			SecurityConstraints: []*model.SecurityConstraint{
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
	}
}

func strPointer(s string) *string {
	return &s
}

func prepareInitUsers() []*model.User {
	return []*model.User{
		{
			Username: "admin",
			Password: strPointer("admin"),
			Roles:    []string{"root"},
		},
		{
			Username: "dh_test",
			Password: strPointer("dh_test"),
			Roles:    []string{"test_user"},
		},
	}
}

func prepareInitNamespaces() []*model.Namespace {
	return nil
}

func prepareInitDataSources() []*modelnew.DataSource {
	return []*modelnew.DataSource{
		DefaultDataSource,
	}
}

func prepareSystemNamespace() *model.Namespace {
	return &model.Namespace{
		Name: "system",
	}
}

func prepareSystemDataSource() *modelnew.DataSource {
	return SystemDataSource
}
