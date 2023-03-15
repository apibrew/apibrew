package test

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/test/setup"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

type testExtension struct {
	abs.Extension
	extensionConfig *model.ExtensionConfig
}

var simpleVirtualResourceRecords = []*model.Record{
	{
		Id: "5429846c-a309-11ed-a8fc-0242ac120002",
		Properties: map[string]*structpb.Value{
			"name":        structpb.NewStringValue("rec-1"),
			"description": structpb.NewStringValue("rec-1-desc"),
		},
	},
	{
		Id: "54298994-a309-11ed-a8fc-0242ac120002",
		Properties: map[string]*structpb.Value{
			"name":        structpb.NewStringValue("rec-2"),
			"description": structpb.NewStringValue("rec-2-desc"),
		},
	},
}

func (t *testExtension) BeforeList(ctx context.Context, in *ext.BeforeListRecordRequest) (*ext.BeforeListRecordResponse, error) {
	return &ext.BeforeListRecordResponse{}, nil
}

func (t *testExtension) List(ctx context.Context, in *ext.ListRecordRequest) (*ext.ListRecordResponse, error) {
	return &ext.ListRecordResponse{
		Records: simpleVirtualResourceRecords,
		Total:   2,
	}, nil
}

func (t *testExtension) Create(ctx context.Context, in *ext.CreateRecordRequest) (*ext.CreateRecordResponse, error) {
	return nil, nil
}

func (t *testExtension) Update(ctx context.Context, in *ext.UpdateRecordRequest) (*ext.UpdateRecordResponse, error) {
	return nil, nil
}

func (t *testExtension) Delete(ctx context.Context, in *ext.DeleteRecordRequest) (*ext.DeleteRecordResponse, error) {
	return nil, nil
}

func (t *testExtension) GetExtensionConfig() *model.ExtensionConfig {
	return t.extensionConfig
}

func TestListResourceWithExtension(t *testing.T) {
	var te abs.Extension = &testExtension{
		extensionConfig: &model.ExtensionConfig{
			Namespace: setup.SimpleVirtualResource1.Namespace,
			Resource:  setup.SimpleVirtualResource1.Name,
			Operations: []*model.ExtensionOperation{
				{
					OperationType: model.ExtensionOperationType_ExtensionOperationTypeList,
					Step:          model.ExtensionOperationStep_ExtensionOperationStepInstead,
					Sync:          true,
				},
			},
		},
	}

	container.GetExtensionService().RegisterExtension(te)
	defer container.GetExtensionService().UnRegisterExtension(te)

	resp, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
		Namespace: setup.SimpleVirtualResource1.Namespace,
		Resource:  setup.SimpleVirtualResource1.Name,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if resp.Total != 2 {
		t.Error("resp.Total should be 2")
		return
	}

	if resp.Content[0].Id != simpleVirtualResourceRecords[0].Id {
		t.Error("record[0].id does not match")
		return
	}

	if resp.Content[1].Id != simpleVirtualResourceRecords[1].Id {
		t.Error("record[1].id does not match")
		return
	}

	if resp.Content[0].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["name"].GetStringValue() {
		t.Error("record[0].name does not match")
		return
	}

	if resp.Content[1].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["name"].GetStringValue() {
		t.Error("record[1].name does not match")
		return
	}

	if resp.Content[0].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["description"].GetStringValue() {
		t.Error("record[0].description does not match")
		return
	}

	if resp.Content[1].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["description"].GetStringValue() {
		t.Error("record[1].description does not match")
		return
	}

}
