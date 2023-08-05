package setup

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/system"
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
			Namespace: resources.NamespaceResource.Namespace,
			Resource:  resources.NamespaceResource.Name,
			Record: resource_model.NamespaceMapperInstance.ToRecord(&resource_model.Namespace{
				Name: "default",
			}),
		},
		{
			Namespace: resources.RoleResource.Namespace,
			Resource:  resources.RoleResource.Name,
			Record: resource_model.RoleMapperInstance.ToRecord(&resource_model.Role{
				Name: "root",
				SecurityConstraints: []*resource_model.SecurityConstraint{
					{
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
						Namespace: &system.UserResourceModel.Namespace.Name,
						Resource:  &system.UserResourceModel.Name,
						Operation: resource_model.SecurityConstraintOperation_READ,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
					{
						Namespace: &system.NamespaceResourceModel.Namespace.Name,
						Resource:  &system.NamespaceResourceModel.Name,
						Operation: resource_model.SecurityConstraintOperation_CREATE,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
					{
						Namespace: &system.NamespaceResourceModel.Namespace.Name,
						Resource:  &system.NamespaceResourceModel.Name,
						Operation: resource_model.SecurityConstraintOperation_READ,
						Permit:    resource_model.SecurityConstraintPermit_ALLOW,
					},
				},
			}),
		},
		{
			Namespace: resources.DataSourceResource.Namespace,
			Resource:  resources.DataSourceResource.Name,
			Record:    resource_model.DataSourceMapperInstance.ToRecord(DefaultDataSource),
		},
		{
			Namespace: resources.UserResource.Namespace,
			Resource:  resources.UserResource.Name,
			Record: resource_model.UserMapperInstance.ToRecord(&resource_model.User{
				Username: "admin",
				Password: strPointer("admin"),
				Roles: []*resource_model.Role{
					{
						Name: "root",
					},
				},
			}),
		},
		{
			Namespace: resources.UserResource.Namespace,
			Resource:  resources.UserResource.Name,
			Record: resource_model.UserMapperInstance.ToRecord(&resource_model.User{
				Username: "dh_test",
				Password: strPointer("dh_test"),
				Roles: []*resource_model.Role{
					{
						Name: "test_user",
					},
				},
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
