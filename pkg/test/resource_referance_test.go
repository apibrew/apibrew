package test

import (
	"github.com/tislib/data-handler/pkg/model"
	util2 "github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func prepareTestResourceReferenceResources() []*model.Resource {
	return []*model.Resource{
		{
			Name: "author",
			SourceConfig: &model.ResourceSourceConfig{
				DataSource: dhTest.Name,
				Catalog:    "",
				Entity:     "author",
			},
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourcePropertyType_TYPE_STRING,
					Mapping:  "name",
					Required: true,
					Length:   255,
				},
				{
					Name:         "description",
					Type:         model.ResourcePropertyType_TYPE_STRING,
					Mapping:      "description",
					Required:     true,
					Length:       255,
					DefaultValue: structpb.NewStringValue("no-description"),
				},
			},
		},
		{
			Name: "book",
			SourceConfig: &model.ResourceSourceConfig{
				DataSource: dhTest.Name,
				Catalog:    "",
				Entity:     "book",
			},
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourcePropertyType_TYPE_STRING,
					Mapping:  "name",
					Required: true,
					Length:   255,
				},
				{
					Name:         "description",
					Type:         model.ResourcePropertyType_TYPE_STRING,
					Mapping:      "description",
					Required:     true,
					Length:       255,
					DefaultValue: structpb.NewStringValue("no-description"),
				},
				{
					Name:     "author",
					Type:     model.ResourcePropertyType_TYPE_REFERENCE,
					Mapping:  "author",
					Required: true,
					Reference: &model.Reference{
						ReferencedResource: "author",
						Cascade:            false,
					},
				},
			},
		},
	}
}

func TestResourceReferenceViolation(t *testing.T) {
	resources := prepareTestResourceReferenceResources()

	resp, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:   resources,
		DoMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		_, err = resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
			Ids:            util.ArrayMapToId(resp.Resources),
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			t.Error(err)
			return
		}
	}()

	_, err = recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: "book",
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name":        structpb.NewStringValue("test-book"),
					"description": structpb.NewStringValue("descp-1"),
					"author": util.MapStructValue(map[string]interface{}{
						"id": "11c3135a-a4e3-11ed-b9df-0242ac120003",
					}),
				},
			},
		},
	})

	if err == nil {
		t.Error("It should not create records")
		return
	}

	if util2.GetErrorCode(err) != model.ErrorCode_REFERENCE_VIOLATION {
		t.Error("Error should be model.ErrorCode_REFERENCE_VIOLATION but: " + util2.GetErrorCode(err).String())
		return
	}
}

func TestResourceReferenceSuccess(t *testing.T) {
	resources := prepareTestResourceReferenceResources()

	resp, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:   resources,
		DoMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		_, err = resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
			Ids:            util.ArrayMapToId(resp.Resources),
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			t.Error(err)
			return
		}
	}()

	_, err = recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: "author",
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name":        structpb.NewStringValue("test-author"),
					"description": structpb.NewStringValue("descp-1"),
				},
			},
		},
	})

	if err != nil {
		t.Error("It should create records")
		return
	}

	_, err = recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: "book",
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name":        structpb.NewStringValue("test-book"),
					"description": structpb.NewStringValue("descp-1"),
					"author": util.MapStructValue(map[string]interface{}{
						"name": "test-author",
					}),
				},
			},
		},
	})

	if err != nil {
		t.Error("It should create records")
		return
	}
}
