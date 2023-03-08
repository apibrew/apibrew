package common

import "github.com/tislib/data-handler/pkg/model"

type resourceMigrationPlanner struct {
}

type ResourceMigrationPlanner interface {
}

type ResourceMigrationPlannerParams struct {
	ExistingResource           *model.Resource
	CurrentResource            *model.Resource
	NewResourcePropertyHandler func()
}

func NewResourceMigrationPlanner() ResourceMigrationPlanner {
	return &resourceMigrationPlanner{}
}
