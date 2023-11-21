package helper

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
)

type ResourceMigrationBuilder interface {
	AddResource(resource *model.Resource) ResourceMigrationBuilder
	UpdateResource(existing, updated *model.Resource) ResourceMigrationBuilder
	DeleteResource(resource *model.Resource) ResourceMigrationBuilder
	AddProperty(prop *model.ResourceProperty) ResourceMigrationBuilder
	UpdateProperty(resource *model.Resource, existing, updated *model.ResourceProperty) ResourceMigrationBuilder
	DeleteProperty(prop *model.ResourceProperty) ResourceMigrationBuilder
	AddIndex(prop *model.ResourceIndex) ResourceMigrationBuilder
	DeleteIndex(prop *model.ResourceIndex) ResourceMigrationBuilder
	Exec() errors.ServiceError
}

type ResourceMigrationBuilderConstructor func(ctx context.Context, runner QueryRunner, schema *abs.Schema, params abs.UpgradeResourceParams, forceMigration bool) ResourceMigrationBuilder

func ResourceMigrateTableViaResourceMigrationBuilder(hp ResourceMigrationBuilder, migrationPlan *model.ResourceMigrationPlan, forceMigration bool) errors.ServiceError {
	var currentPropertyMap = util.GetNamedMap(migrationPlan.CurrentResource.Properties)
	var existingPropertyMap = util.GetNamedMap(migrationPlan.ExistingResource.Properties)

	for _, step := range migrationPlan.Steps {
		switch sk := step.Kind.(type) {
		case *model.ResourceMigrationStep_CreateResource:
			hp.AddResource(migrationPlan.CurrentResource)
		case *model.ResourceMigrationStep_DeleteResource:
			if forceMigration {
				hp.DeleteResource(migrationPlan.ExistingResource)
			}
		case *model.ResourceMigrationStep_CreateProperty:
			if sk.CreateProperty.SubType != "" {
				continue
			}

			property := currentPropertyMap[sk.CreateProperty.Property]
			if annotations.IsEnabled(property, annotations.PrimaryProperty) { // skip primary properties because they are already created as upon table creation, this logic should be reworked
				continue
			}
			if IsPropertyOmitted(property) {
				continue
			}
			hp.AddProperty(property)
		case *model.ResourceMigrationStep_UpdateProperty:
			if sk.UpdateProperty.SubType != "" {
				continue
			}

			hp.UpdateProperty(migrationPlan.CurrentResource, existingPropertyMap[sk.UpdateProperty.ExistingProperty], currentPropertyMap[sk.UpdateProperty.Property])
		case *model.ResourceMigrationStep_DeleteProperty:
			if sk.DeleteProperty.SubType != "" {
				continue
			}

			if forceMigration {
				hp.DeleteProperty(existingPropertyMap[sk.DeleteProperty.ExistingProperty])
			}
		}
	}

	for _, step := range migrationPlan.Steps {
		switch sk := step.Kind.(type) {
		case *model.ResourceMigrationStep_CreateIndex:
			hp.AddIndex(migrationPlan.CurrentResource.Indexes[sk.CreateIndex.Index])
		case *model.ResourceMigrationStep_DeleteIndex:
			hp.DeleteIndex(migrationPlan.ExistingResource.Indexes[sk.DeleteIndex.ExistingIndex])
		}
	}

	return hp.Exec()
}
