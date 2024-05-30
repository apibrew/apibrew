package test

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

type TestRecordCreationValidationSubCase struct {
	resource   *model.Resource
	recordType model.ResourceProperty_Type
}

func TestRecordCreationValidationBasedOnTypes(t *testing.T) {
	subCases := prepareTestRecordCreationValidationSubCase()
	var resourceIdsForRemoval []string
	var newResources []*model.Resource

	defer func() {
		if len(resourceIdsForRemoval) > 0 {
			_, err := resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
				Ids:            resourceIdsForRemoval,
				DoMigration:    true,
				ForceMigration: true,
			})

			if err != nil {
				t.Error("Could not delete resources", err)
				return
			}
		}
	}()

	for _, subCase := range subCases {
		// creation of resource
		newResources = append(newResources, subCase.resource)
	}

	resp, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{
		Token:          "",
		Resources:      newResources,
		DoMigration:    true,
		ForceMigration: true,
	})

	if err != nil {
		t.Error(err)
		return
	}

	for index, newResource := range resp.Resources {
		newResources[index].Id = newResource.Id
		resourceIdsForRemoval = append(resourceIdsForRemoval, newResource.Id)
	}

	for _, subCase := range subCases {
		// create
		t.Run(subCase.recordType.String()+" - Create - Valid", func(t *testing.T) {
			testRecordCreationValidationValidCase(setup.Ctx, t, subCase)
		})
		// create
		t.Run(subCase.recordType.String()+" - Create - Default - Valid", func(t *testing.T) {
			testRecordCreationValidationDefaultValidCase(setup.Ctx, t, subCase)
		})
		t.Run(subCase.recordType.String()+" - Create - Invalid", func(t *testing.T) {
			testRecordCreationValidationInvalidCase(setup.Ctx, t, subCase)
		})
		// update
		t.Run(subCase.recordType.String()+" - Update - Valid", func(t *testing.T) {
			testRecordUpdateValidationValidCase(setup.Ctx, t, subCase)
		})
		//t.Run(subCase.recordType.String()+" - Invalid", func(t *testing.T) {
		//	testRecordUpdateValidationInvalidCase(ctx, t, subCase)
		//})
	}
}

func testRecordCreationValidationValidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	var records []abs.RecordLike
	for i := 0; i < 30; i += 3 {
		var properties = make(map[string]*structpb.Value, 3)

		properties[subCase.resource.Properties[0].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		properties[subCase.resource.Properties[1].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		properties[subCase.resource.Properties[2].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))

		validRecord := &model.Record{
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	resp, err := recordClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: subCase.resource.Name,
		Records:  abs.RecordLikeAsRecords(records),
	})

	if err != nil {
		t.Error(err)
		return
	}

	propertyType := types.ByResourcePropertyType(subCase.recordType)

	for i := 0; i < len(resp.Records); i++ {
		createdRecord := resp.Records[i]
		record := records[i]

		createdRecordValue0, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[0].Name])
		createdRecordValue1, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[1].Name])
		createdRecordValue2, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[2].Name])

		recordValue0, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[0].Name])
		recordValue1, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[1].Name])
		recordValue2, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[2].Name])

		if !propertyType.Equals(createdRecordValue0, recordValue0) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue0, recordValue0)
		}

		if !propertyType.Equals(createdRecordValue1, recordValue1) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue1, recordValue1)
		}

		if !propertyType.Equals(createdRecordValue2, recordValue2) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue2, recordValue2)
		}
	}
}

func testRecordCreationValidationDefaultValidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	var records []abs.RecordLike
	for i := 0; i < 30; i += 3 {
		var properties = make(map[string]*structpb.Value, 3)
		typ := types.ByResourcePropertyType(subCase.recordType)

		properties[subCase.resource.Properties[0].Name], _ = typ.Pack(typ.Default())
		properties[subCase.resource.Properties[1].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		properties[subCase.resource.Properties[2].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))

		if subCase.recordType == model.ResourceProperty_ENUM || subCase.recordType == model.ResourceProperty_STRUCT {
			properties[subCase.resource.Properties[0].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		}

		validRecord := &model.Record{
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	resp, err := recordClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: subCase.resource.Name,
		Records:  abs.RecordLikeAsRecords(records),
	})

	if err != nil {
		t.Error(err)
		return
	}

	propertyType := types.ByResourcePropertyType(subCase.recordType)

	for i := 0; i < len(resp.Records); i++ {
		createdRecord := resp.Records[i]
		record := records[i]

		createdRecordValue0, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[0].Name])
		createdRecordValue1, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[1].Name])
		createdRecordValue2, _ := propertyType.UnPack(createdRecord.Properties[subCase.resource.Properties[2].Name])

		recordValue0, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[0].Name])
		recordValue1, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[1].Name])
		recordValue2, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[2].Name])

		if !propertyType.Equals(createdRecordValue0, recordValue0) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue0, recordValue0)
		}

		if !propertyType.Equals(createdRecordValue1, recordValue1) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue1, recordValue1)
		}

		if !propertyType.Equals(createdRecordValue2, recordValue2) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue2, recordValue2)
		}
	}
}

func testRecordUpdateValidationValidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	var records []abs.RecordLike
	for i := 0; i < 30; i += 3 {
		var properties = make(map[string]*structpb.Value, 3)

		properties[subCase.resource.Properties[0].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		properties[subCase.resource.Properties[1].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		properties[subCase.resource.Properties[2].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))

		validRecord := &model.Record{
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	resp, err := recordClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: subCase.resource.Name,
		Records:  abs.RecordLikeAsRecords(records),
	})

	if err != nil {
		t.Error(err)
		return
	}

	propertyType := types.ByResourcePropertyType(subCase.recordType)

	for i := 0; i < len(resp.Records); i++ {
		createdRecord := resp.Records[i]
		record := records[i]
		record.GetProperties()["id"] = createdRecord.Properties["id"]
	}

	for _, record := range records {
		record.GetProperties()[subCase.resource.Properties[0].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		record.GetProperties()[subCase.resource.Properties[1].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
		record.GetProperties()[subCase.resource.Properties[2].Name], _ = structpb.NewValue(fakeValidValue(subCase.recordType))
	}

	updateResp, err := recordClient.Update(ctx, &stub.UpdateRecordRequest{
		Resource: subCase.resource.Name,
		Records:  abs.RecordLikeAsRecords(records),
	})

	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < len(updateResp.Records); i++ {
		updatedRecord := updateResp.Records[i]
		record := records[i]

		createdRecordValue0, _ := propertyType.UnPack(updatedRecord.Properties[subCase.resource.Properties[0].Name])
		createdRecordValue1, _ := propertyType.UnPack(updatedRecord.Properties[subCase.resource.Properties[1].Name])
		createdRecordValue2, _ := propertyType.UnPack(updatedRecord.Properties[subCase.resource.Properties[2].Name])

		recordValue0, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[0].Name])
		recordValue1, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[1].Name])
		recordValue2, _ := propertyType.UnPack(record.GetProperties()[subCase.resource.Properties[2].Name])

		if !propertyType.Equals(createdRecordValue0, recordValue0) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue0, recordValue0)
		}

		if !propertyType.Equals(createdRecordValue1, recordValue1) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue1, recordValue1)
		}

		if !propertyType.Equals(createdRecordValue2, recordValue2) {
			t.Errorf("values are different: %s <=> %s", createdRecordValue2, recordValue2)
		}
	}
}

func testRecordCreationValidationInvalidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	if fakeInvalidValue(subCase.recordType) == nil {
		return
	}

	var records []abs.RecordLike
	for i := 0; i < 30; i += 3 {
		var properties = make(map[string]*structpb.Value, 3)

		properties[subCase.resource.Properties[0].Name], _ = structpb.NewValue(fakeInvalidValue(subCase.recordType))
		properties[subCase.resource.Properties[1].Name], _ = structpb.NewValue(fakeInvalidValue(subCase.recordType))
		properties[subCase.resource.Properties[2].Name], _ = structpb.NewValue(fakeInvalidValue(subCase.recordType))

		validRecord := &model.Record{
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	_, err := recordClient.Create(ctx, &stub.CreateRecordRequest{
		Resource: subCase.resource.Name,
		Records:  abs.RecordLikeAsRecords(records),
	})

	if err == nil {
		t.Error("Validation should failed but not failed for: " + subCase.recordType.String())
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_RECORD_VALIDATION_ERROR {
		t.Error("Wrong error code: " + util.GetErrorCode(err).String())
		return
	}

	errorFields := util.GetErrorFields(err)

	if len(errorFields) == 0 {
		t.Error("Errors are expected")
		return
	}
}

func prepareTestRecordCreationValidationSubCase() []TestRecordCreationValidationSubCase {
	typs := types.GetAllResourcePropertyTypes()

	var cases []TestRecordCreationValidationSubCase

	for _, typ := range typs {
		if typ == model.ResourceProperty_REFERENCE {
			continue
		}

		checkUnique := typ != model.ResourceProperty_BOOL && typ != model.ResourceProperty_ENUM && typ != model.ResourceProperty_MAP && typ != model.ResourceProperty_LIST && typ != model.ResourceProperty_STRUCT

		propNames := []string{
			fakePropertyName(),
			fakePropertyName(),
			fakePropertyName(),
			fakePropertyName(),
		}
		length := uint32(32)

		resource := fakeResource(
			&model.ResourceProperty{
				Name:     propNames[0],
				Type:     typ,
				Required: false,
				Length:   length,
				Unique:   false,
			},
			&model.ResourceProperty{
				Name:     propNames[1],
				Type:     typ,
				Required: true,
				Length:   length,
				Unique:   false,
			},
			&model.ResourceProperty{
				Name:     propNames[2],
				Type:     typ,
				Required: false,
				Length:   length,
				Unique:   checkUnique,
			},
		)

		resource.Types = []*model.ResourceSubType{
			{
				Name: "SampleResource",
				Properties: []*model.ResourceProperty{
					{
						Name:     "field-1",
						Required: true,
						Type:     model.ResourceProperty_STRING,
					},
					{
						Name: "field-2",
						Type: model.ResourceProperty_INT32,
					},
				},
			},
		}

		for _, prop := range resource.Properties {
			switch typ {
			case model.ResourceProperty_ENUM:
				prop.EnumValues = []string{"ENUM1", "ENUM2", "ENUM3", "ENUM4"}
			case model.ResourceProperty_MAP:
				prop.Item = &model.ResourceProperty{
					Type: model.ResourceProperty_STRING,
				}
			case model.ResourceProperty_LIST:
				prop.Item = &model.ResourceProperty{
					Type: model.ResourceProperty_STRING,
				}
			case model.ResourceProperty_STRUCT:
				prop.TypeRef = util.Pointer("SampleResource")
			}
		}

		cases = append(cases, TestRecordCreationValidationSubCase{
			recordType: typ,
			resource:   resource,
		})
	}

	return cases
}
