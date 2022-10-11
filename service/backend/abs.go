package backend

import (
	"data-handler/stub/model"
)

type DataSourceBackend interface {
	getDataSourceId() string
}

type AddResourceParams struct {
	Backend              DataSourceBackend
	Resource             *model.Resource
	AllowSystemAndStatic bool
	Migrate              bool
	ForceMigrate         bool
}

type ResourceServiceBackend interface {
	AddResource(params AddResourceParams) (*model.Resource, error)
	Init(backend DataSourceBackend)
}
