package test

import (
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/service/annotations"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/test/setup"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func prepareTestResourceHistoryResources() []*model.Resource {
	return []*model.Resource{
		{
			Name: "paper",
			SourceConfig: &model.ResourceSourceConfig{
				DataSource: setup.DhTest.Name,
				Catalog:    "",
				Entity:     "paper",
			},
			Annotations: map[string]string{
				annotations.KeepHistory: "true",
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
			},
		},
	}
}

func TestResourceCreateRecordWithHistory(t *testing.T) {
	resources := prepareTestResourceHistoryResources()

	fn := setup.PrepareResourcesForTest(t, resources)

	defer fn()

	createResp, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
		Resource: "paper",
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name":        structpb.NewStringValue("test-paper"),
					"description": structpb.NewStringValue("descp-1"),
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	_, err = recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{
		Resource:     "paper",
		Records:      createResp.Records,
		CheckVersion: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	listResp, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
		Resource:   "paper",
		UseHistory: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if listResp.Total != 2 {
		t.Error("It should have record history")
		return
	}
}
