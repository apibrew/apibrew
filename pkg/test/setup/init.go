package setup

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
)

func prepareInitData() *model.AppConfig {
	return &model.AppConfig{
		Host:                  "localhost",
		Port:                  17981,
		DisableAuthentication: false,
		DisableCache:          true,
		SystemDataSource:      resource_model.DataSourceMapperInstance.ToRecord(prepareSystemDataSource()),
		InitResources:         prepareInitResources(),
		InitRecords:           prepareInitRecords(),
	}
}

func prepareInitRecords() []*model.InitRecord {
	return []*model.InitRecord{
		{
			Namespace: resources.RoleResource.Namespace,
			Resource:  resources.RoleResource.Name,
			Record: resource_model.RoleMapperInstance.ToRecord(&resource_model.Role{
				Name: "root",
				SecurityConstraints: []*resource_model.SecurityConstraint{
					{
						Namespace: "*",
						Resource:  "*",
						Property:  "*",
						Operation: resource_model.SecurityConstraintOperation_FULL,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
				},
			}),
		},
		{
			Namespace: resources.RoleResource.Namespace,
			Resource:  resources.RoleResource.Name,
			Record: resource_model.RoleMapperInstance.ToRecord(&resource_model.Role{
				Name: "test_user",
				SecurityConstraints: []*resource_model.SecurityConstraint{
					{
						Resource:  "user",
						Operation: resource_model.SecurityConstraintOperation_READ,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
					{
						Resource:  "namespace",
						Operation: resource_model.SecurityConstraintOperation_CREATE,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
					{
						Resource:  "namespace",
						Operation: resource_model.SecurityConstraintOperation_READ,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
				},
			}),
		},
		{
			Namespace: resources.UserResource.Namespace,
			Resource:  resources.UserResource.Name,
			Record: resource_model.UserMapperInstance.ToRecord(&resource_model.User{
				Username: "admin",
				Password: strPointer("admin"),
				Roles:    []string{"root"},
			}),
		},
		{
			Namespace: resources.UserResource.Namespace,
			Resource:  resources.UserResource.Name,
			Record: resource_model.UserMapperInstance.ToRecord(&resource_model.User{
				Username: "dh_test",
				Password: strPointer("dh_test"),
				Roles:    []string{"test_user"},
			}),
		},
	}
}

func prepareInitResources() []*model.Resource {
	return nil
}

func strPointer(s string) *string {
	return &s
}

func prepareSystemDataSource() *resource_model.DataSource {
	return SystemDataSource
}
