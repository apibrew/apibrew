package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func prepareTestResourceReferenceResources() []*model.Resource {
	return []*model.Resource{
		{
			Name: "author",
			SourceConfig: &model.ResourceSourceConfig{
				DataSource: setup.DhTest.Name,
				Catalog:    "",
				Entity:     "author",
			},
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Mapping:  "name",
					Required: true,
					Unique:   true,
					Length:   255,
				},
				{
					Name:         "description",
					Type:         model.ResourceProperty_STRING,
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
				DataSource: setup.DhTest.Name,
				Catalog:    "",
				Entity:     "book",
			},
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Mapping:  "name",
					Required: true,
					Length:   255,
				},
				{
					Name:         "description",
					Type:         model.ResourceProperty_STRING,
					Mapping:      "description",
					Required:     true,
					Length:       255,
					DefaultValue: structpb.NewStringValue("no-description"),
				},
				{
					Name:     "author",
					Type:     model.ResourceProperty_REFERENCE,
					Mapping:  "author",
					Required: true,
					Reference: &model.Reference{
						Resource: "author",
						Cascade:  false,
					},
				},
			},
		},
	}
}

func TestResourceReferenceViolation(t *testing.T) {
	resources := prepareTestResourceReferenceResources()

	resp, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources:   resources,
		DoMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		_, err = resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
			Ids:            util.ArrayMapToId(resp.Resources),
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			t.Error(err)
			return
		}
	}()

	_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
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

	if util.GetErrorCode(err) != model.ErrorCode_REFERENCE_VIOLATION {
		t.Error("Error should be model.ErrorCode_REFERENCE_VIOLATION but: " + util.GetErrorCode(err).String())
		return
	}
}

func TestResourceReferenceSuccess(t *testing.T) {
	resources := prepareTestResourceReferenceResources()

	resp, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Resources:   resources,
		DoMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		_, err = resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
			Ids:            util.ArrayMapToId(resp.Resources),
			DoMigration:    true,
			ForceMigration: true,
		})

		if err != nil {
			t.Error(err)
			return
		}
	}()

	_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
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

	_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
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
		log.Print(err)
		t.Error("It should create records")
		return
	}
}
