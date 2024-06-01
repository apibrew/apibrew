package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type resourceMigrationService struct {
}

func (r *resourceMigrationService) PreparePlan(ctx context.Context, existingResource *model.Resource, resource *model.Resource) (*model.ResourceMigrationPlan, error) {
	if existingResource == nil && resource == nil {
		return nil, errors.LogicalError.WithDetails("Both existing resource and resource cannot be nil at the same time")
	}

	if resource != nil && existingResource != nil {
		resource.AuditData = existingResource.AuditData
		resource.Version = existingResource.Version
	}

	var plan = &model.ResourceMigrationPlan{
		ExistingResource: existingResource,
		CurrentResource:  resource,
		Steps:            nil,
	}

	// create new resource case
	if existingResource == nil {
		plan.ExistingResource = &model.Resource{} // for safety
		plan.Steps = r.preparePlanStepsForNewResource(resource)

		return plan, nil
	}

	// delete existing resource case
	if resource == nil {
		plan.Steps = r.preparePlanStepsForDeleteResource()
		plan.CurrentResource = existingResource // for consistency

		return plan, nil
	}

	// check resource fields updated
	plan.Steps = append(plan.Steps, r.preparePlanStepsForUpdateResource(resource, existingResource)...)

	// check properties
	_ = util.ArrayDiffer(existingResource.Properties,
		resource.Properties,
		util.IsSameIdentifiedResourceProperty,
		util.IsSameResourceProperty,
		func(prop *model.ResourceProperty) error { // new
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_CreateProperty{CreateProperty: &model.ResourceMigrationCreateProperty{
				Property: prop.Name,
			}}})

			return nil
		}, func(e, u *model.ResourceProperty) error { // update
			plan.Steps = append(plan.Steps, r.preparePlanStepsForUpdateResourceProperty(resource, existingResource, u, e, "")...)

			return nil
		}, func(prop *model.ResourceProperty) error { // delete
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_DeleteProperty{DeleteProperty: &model.ResourceMigrationDeleteProperty{
				ExistingProperty: prop.Name,
			}}})

			return nil
		})

	// types
	_ = util.ArrayDiffer(existingResource.Types,
		resource.Types,
		func(a, b *model.ResourceSubType) bool {
			return a.Name == b.Name
		},
		func(a, b *model.ResourceSubType) bool {
			return false // to check properties always
		},
		func(subType *model.ResourceSubType) error { // new
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_CreateSubType{CreateSubType: &model.ResourceMigrationCreateSubType{
				Name: subType.Name,
			}}})

			return nil
		}, func(e, u *model.ResourceSubType) error { // update
			// check properties
			_ = util.ArrayDiffer(e.Properties,
				u.Properties,
				util.IsSameIdentifiedResourceProperty,
				util.IsSameResourceProperty,
				func(prop *model.ResourceProperty) error { // new
					plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_CreateProperty{CreateProperty: &model.ResourceMigrationCreateProperty{
						Property: prop.Name,
						SubType:  u.Name,
					}}})

					return nil
				}, func(ep, up *model.ResourceProperty) error { // update
					plan.Steps = append(plan.Steps, r.preparePlanStepsForUpdateResourceProperty(resource, existingResource, up, ep, u.Name)...)

					return nil
				}, func(prop *model.ResourceProperty) error { // delete
					plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_DeleteProperty{DeleteProperty: &model.ResourceMigrationDeleteProperty{
						ExistingProperty: prop.Name,
						SubType:          u.Name,
					}}})

					return nil
				})

			return nil
		}, func(subType *model.ResourceSubType) error { // delete
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_DeleteSubType{DeleteSubType: &model.ResourceMigrationDeleteSubType{
				Name: subType.Name,
			}}})

			return nil
		})

	// check indexes
	_ = util.ArrayDiffer(existingResource.Indexes,
		resource.Indexes,
		util.IsSameIdentifiedResourceIndex,
		util.IsSameResourceIndex,
		func(prop *model.ResourceIndex) error { // new
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_CreateIndex{CreateIndex: &model.ResourceMigrationCreateIndex{
				Index: uint32(util.GetArrayIndex(resource.Indexes, prop, util.IsSameResourceIndex)),
			}}})

			return nil
		}, func(e, u *model.ResourceIndex) error { // update
			log.Fatal("Not implemented, not possible")
			return nil
		}, func(prop *model.ResourceIndex) error { // delete
			plan.Steps = append(plan.Steps, &model.ResourceMigrationStep{Kind: &model.ResourceMigrationStep_DeleteIndex{DeleteIndex: &model.ResourceMigrationDeleteIndex{
				ExistingIndex: uint32(util.GetArrayIndex(existingResource.Indexes, prop, util.IsSameResourceIndex)),
			}}})

			return nil
		})

	return plan, nil
}

func (r *resourceMigrationService) preparePlanStepsForNewResource(resource *model.Resource) []*model.ResourceMigrationStep {
	var steps []*model.ResourceMigrationStep

	steps = append(steps, &model.ResourceMigrationStep{
		Kind: &model.ResourceMigrationStep_CreateResource{CreateResource: &model.ResourceMigrationCreateResource{}},
	})

	for _, prop := range resource.Properties {
		steps = append(steps, &model.ResourceMigrationStep{
			Kind: &model.ResourceMigrationStep_CreateProperty{CreateProperty: &model.ResourceMigrationCreateProperty{
				Property: prop.Name,
			}},
		})
	}

	for index := range resource.Indexes {
		steps = append(steps, &model.ResourceMigrationStep{
			Kind: &model.ResourceMigrationStep_CreateIndex{CreateIndex: &model.ResourceMigrationCreateIndex{
				Index: uint32(index),
			}},
		})
	}

	return steps
}

func (r *resourceMigrationService) preparePlanStepsForDeleteResource() []*model.ResourceMigrationStep {
	var steps []*model.ResourceMigrationStep

	steps = append(steps, &model.ResourceMigrationStep{
		Kind: &model.ResourceMigrationStep_DeleteResource{DeleteResource: &model.ResourceMigrationDeleteResource{}},
	})

	return steps
}

func (r *resourceMigrationService) preparePlanStepsForUpdateResource(resource, existingResource *model.Resource) []*model.ResourceMigrationStep {
	var steps []*model.ResourceMigrationStep

	resourceRecord := mapping.ResourceToRecord(resource)
	existingResourceRecord := mapping.ResourceToRecord(existingResource)

	var changedFields []string
	for _, prop := range resources.ResourceResource.Properties {
		if prop.Name == "properties" {
			continue
		}
		if !proto.Equal(resourceRecord.GetStructProperty(prop.Name), existingResourceRecord.GetStructProperty(prop.Name)) {
			changedFields = append(changedFields, prop.Name)
		}
	}

	if len(changedFields) > 0 {
		steps = append(steps, &model.ResourceMigrationStep{
			Kind: &model.ResourceMigrationStep_UpdateResource{UpdateResource: &model.ResourceMigrationUpdateResource{
				ChangedFields: changedFields,
			}},
		})
	}

	return steps
}

func (r *resourceMigrationService) preparePlanStepsForUpdateResourceProperty(resource, existingResource *model.Resource, resourceProperty, existingResourceProperty *model.ResourceProperty, subType string) []*model.ResourceMigrationStep {
	var steps []*model.ResourceMigrationStep

	var changed bool

	changed = changed || resourceProperty.Type != existingResourceProperty.Type
	changed = changed || resourceProperty.Required != existingResourceProperty.Required

	changed = changed || resourceProperty.Unique != existingResourceProperty.Unique
	changed = changed || resourceProperty.DefaultValue != existingResourceProperty.DefaultValue
	changed = changed || resourceProperty.Virtual != existingResourceProperty.Virtual
	changed = changed || resourceProperty.Primary != existingResourceProperty.Primary
	if len(resourceProperty.Annotations) != len(existingResourceProperty.Annotations) {
		changed = true
	}

	for key, value := range resourceProperty.Annotations {
		if existingResourceProperty.Annotations[key] != value {
			changed = true
			break
		}
	}

	if resourceProperty.Type != model.ResourceProperty_REFERENCE {
		if (resourceProperty.Reference != nil) != (existingResourceProperty.Reference != nil) {
			changed = true
		} else if (resourceProperty.Reference != nil) && (existingResourceProperty.Reference != nil) {
			changed = changed || resourceProperty.Reference.Resource != existingResourceProperty.Reference.Resource
			changed = changed || resourceProperty.Reference.Namespace != existingResourceProperty.Reference.Namespace
		}
	}

	if resourceProperty.Type == model.ResourceProperty_STRING {
		changed = changed || resourceProperty.Length != existingResourceProperty.Length
	}

	if resourceProperty.Type == model.ResourceProperty_STRUCT {
		changed = changed || resourceProperty.TypeRef != existingResourceProperty.TypeRef
	}

	if changed {
		steps = append(steps, &model.ResourceMigrationStep{
			Kind: &model.ResourceMigrationStep_UpdateProperty{UpdateProperty: &model.ResourceMigrationUpdateProperty{
				ExistingProperty: existingResourceProperty.Name,
				Property:         resourceProperty.Name,
				SubType:          subType,
			}},
		})
	}

	return steps
}

func NewResourceMigrationService() service.ResourceMigrationService {
	return &resourceMigrationService{}
}
