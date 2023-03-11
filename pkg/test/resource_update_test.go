package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestPrepareResourceMigrationPlan(t *testing.T) {
	resource1 := &model.Resource{
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			}, {
				Name:     "prop-2",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-2",
			}, {
				Name:     "prop-3",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-3",
			},
		},
	}

	resourceCreateRes, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if resourceCreateRes != nil {
			_, _ = resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
				Ids:            []string{resourceCreateRes.Resources[0].Id},
				DoMigration:    true,
				ForceMigration: true,
			})
		}
	}()

	resource1 = &model.Resource{
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			},
			{
				Name:     "prop-2a",
				Type:     model.ResourcePropertyType_TYPE_FLOAT32,
				Length:   128,
				Required: false,
				Mapping:  "prop-2",
			},
			{
				Name:     "prop-5",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   127,
				Required: false,
				Mapping:  "prop-5",
			},
		},
	}

	resource1.Id = resourceCreateRes.Resources[0].Id

	res, err := resourceServiceClient.PrepareResourceMigrationPlan(ctx, &stub.PrepareResourceMigrationPlanRequest{
		Resources: []*model.Resource{resource1},
	})

	if err != nil {
		assert.Error(t, err)
	}

	assert.Len(t, res.Plans, 1)
	assert.Len(t, res.Plans[0].Steps, 4)

	if t.Failed() {
		return
	}

	steps := res.Plans[0].Steps

	assert.IsType(t, steps[0].Kind, &model.ResourceMigrationStep_UpdateProperty{})
	assert.IsType(t, steps[1].Kind, &model.ResourceMigrationStep_DeleteProperty{})
	assert.IsType(t, steps[2].Kind, &model.ResourceMigrationStep_UpdateProperty{})
	assert.IsType(t, steps[3].Kind, &model.ResourceMigrationStep_CreateProperty{})

	if t.Failed() {
		return
	}

	assert.Equal(t, steps[0].Kind.(*model.ResourceMigrationStep_UpdateProperty).UpdateProperty.ChangedFields, []string{"name", "type", "required", "subType"})
	assert.Equal(t, steps[2].Kind.(*model.ResourceMigrationStep_UpdateProperty).UpdateProperty.ChangedFields, []string{"name", "type", "required", "subType"})

}

func TestResourceUpdateCreateNewPropertyAndMarkAsRequired(t *testing.T) {
	resource1 := &model.Resource{
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			},
		},
	}

	resourceCreateRes, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if resourceCreateRes != nil {
			_, _ = resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
				Ids:            []string{resourceCreateRes.Resources[0].Id},
				DoMigration:    true,
				ForceMigration: true,
			})
		}
	}()

	recordCreateResult1, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []*model.Record{
		{
			Properties: map[string]*structpb.Value{
				"prop-1": structpb.NewStringValue("test-123321"),
			},
		},
	}})

	if err != nil {
		t.Error(err)
		return
	}

	resource1 = &model.Resource{
		Id:        resourceCreateRes.Resources[0].Id,
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			},
			{
				Name:     "prop-2",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: false,
				Mapping:  "prop-2",
			},
		},
	}

	_, err = resourceServiceClient.Update(ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = recordServiceClient.Create(ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []*model.Record{
		{
			Properties: map[string]*structpb.Value{
				"prop-1": structpb.NewStringValue("test-123321"),
				"prop-2": structpb.NewStringValue("test-12332144"),
			},
		},
	}})

	if err != nil {
		t.Error(err)
		return
	}

	resource1 = &model.Resource{
		Id:        resourceCreateRes.Resources[0].Id,
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			},
			{
				Name:     "prop-2",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-2",
			},
		},
	}

	_, err = resourceServiceClient.Update(ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})

	if err == nil {
		t.Error("marking property prop-2 should be failed, because it is containing null values")
		return
	}

	_, err = recordServiceClient.Update(ctx, &stub.UpdateRecordRequest{Resource: resource1.Name, Records: []*model.Record{
		{
			Id: recordCreateResult1.Records[0].Id,
			Properties: map[string]*structpb.Value{
				"prop-1": structpb.NewStringValue("test-123321"),
				"prop-2": structpb.NewStringValue("test-12332144"),
			},
		},
	}})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = resourceServiceClient.Update(ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}})

	if err != nil {
		t.Error(err)
		return
	}

	resource1 = &model.Resource{
		Id:        resourceCreateRes.Resources[0].Id,
		Name:      "test-resource-for-update-1",
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: dhTest.Name,
			Entity:     "test-resource-for-update-1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name:     "prop-1",
				Type:     model.ResourcePropertyType_TYPE_STRING,
				Length:   128,
				Required: true,
				Mapping:  "prop-1",
			},
		},
	}

	_, err = resourceServiceClient.Update(ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = recordServiceClient.Create(ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []*model.Record{
		{
			Properties: map[string]*structpb.Value{
				"prop-1": structpb.NewStringValue("test-123321"),
				"prop-2": structpb.NewStringValue("test-12332144"),
			},
		},
	}})

	if err == nil {
		t.Error("prop-2 should complaint about property not exists")
		return
	}
}
