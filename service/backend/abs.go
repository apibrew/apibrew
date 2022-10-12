package backend

import (
	"data-handler/stub/model"
)

type DataSourceBackend interface {
	GetDataSourceId() string
}

type AddResourceParams struct {
	Backend              DataSourceBackend
	Resource             *model.Resource
	AllowSystemAndStatic bool
	IgnoreIfExists       bool
	Migrate              bool
	ForceMigrate         bool
}

type ResourceServiceBackend interface {
	AddResource(params AddResourceParams) (*model.Resource, error)
	Init(backend DataSourceBackend)
}
