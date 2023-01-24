package test

import (
	"context"
	"data-handler/model"
	"data-handler/server/stub"
	"data-handler/server/util"
	"data-handler/service/types"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"testing"
)

type TestRecordCreationValidationSubCase struct {
	resource   *model.Resource
	recordType model.ResourcePropertyType
}

func TestRecordCreationValidationBasedOnTypes(t *testing.T) {
	ctx := prepareTextContext()

	subCases := prepareTestRecordCreationValidationSubCase()
	var resourceIdsForRemoval []string
	var newResources []*model.Resource

	defer func() {
		if len(resourceIdsForRemoval) > 0 {
			_, err := resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
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

	resp, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
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
			testRecordCreationValidationValidCase(ctx, t, subCase)
		})
		t.Run(subCase.recordType.String()+" - Create - Invalid", func(t *testing.T) {
			testRecordCreationValidationInvalidCase(ctx, t, subCase)
		})
		// update
		t.Run(subCase.recordType.String()+" - Update - Valid", func(t *testing.T) {
			testRecordUpdateValidationValidCase(ctx, t, subCase)
		})
		//t.Run(subCase.recordType.String()+" - Invalid", func(t *testing.T) {
		//	testRecordUpdateValidationInvalidCase(ctx, t, subCase)
		//})
	}
}

func testRecordCreationValidationValidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	var records []*model.Record
	for i := 0; i < 30; i += 3 {
		var propertiesMap = make(map[string]interface{}, 3)

		propertiesMap[subCase.resource.Properties[0].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[1].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[2].Name] = fakeValidValue(subCase.recordType)

		properties, err := structpb.NewStruct(propertiesMap)

		if err != nil {
			t.Error(err)
			return
		}

		validRecord := &model.Record{
			Resource:   subCase.resource.Name,
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	resp, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Records: records,
	})

	if err != nil {
		t.Error(err)
		return
	}

	propertyType := types.ByResourcePropertyType(subCase.recordType)

	for i := 0; i < len(resp.Records); i++ {
		createdRecord := resp.Records[i]
		record := records[i]

		createdRecordValue0, _ := propertyType.UnPack(createdRecord.Properties.GetFields()[subCase.resource.Properties[0].Name])
		createdRecordValue1, _ := propertyType.UnPack(createdRecord.Properties.GetFields()[subCase.resource.Properties[1].Name])
		createdRecordValue2, _ := propertyType.UnPack(createdRecord.Properties.GetFields()[subCase.resource.Properties[2].Name])

		recordValue0, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[0].Name])
		recordValue1, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[1].Name])
		recordValue2, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[2].Name])

		if !propertyType.Equals(createdRecordValue0, recordValue0) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue0, recordValue0))
		}

		if !propertyType.Equals(createdRecordValue1, recordValue1) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue1, recordValue1))
		}

		if !propertyType.Equals(createdRecordValue2, recordValue2) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue2, recordValue2))
		}
	}
}

func testRecordUpdateValidationValidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	var records []*model.Record
	for i := 0; i < 30; i += 3 {
		var propertiesMap = make(map[string]interface{}, 3)

		propertiesMap[subCase.resource.Properties[0].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[1].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[2].Name] = fakeValidValue(subCase.recordType)

		properties, err := structpb.NewStruct(propertiesMap)

		if err != nil {
			t.Error(err)
			return
		}

		validRecord := &model.Record{
			Resource:   subCase.resource.Name,
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	resp, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Records: records,
	})

	if err != nil {
		t.Error(err)
		return
	}

	propertyType := types.ByResourcePropertyType(subCase.recordType)

	for i := 0; i < len(resp.Records); i++ {
		createdRecord := resp.Records[i]
		record := records[i]
		record.Id = createdRecord.Id
	}

	for _, record := range records {
		var propertiesMap = record.Properties.AsMap()

		propertiesMap[subCase.resource.Properties[0].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[1].Name] = fakeValidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[2].Name] = fakeValidValue(subCase.recordType)

		properties, err := structpb.NewStruct(propertiesMap)

		if err != nil {
			t.Error(err)
			return
		}

		record.Properties = properties
	}

	updateResp, err := recordServiceClient.Update(ctx, &stub.UpdateRecordRequest{
		Records: records,
	})

	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < len(updateResp.Records); i++ {
		updatedRecord := updateResp.Records[i]
		record := records[i]

		createdRecordValue0, _ := propertyType.UnPack(updatedRecord.Properties.GetFields()[subCase.resource.Properties[0].Name])
		createdRecordValue1, _ := propertyType.UnPack(updatedRecord.Properties.GetFields()[subCase.resource.Properties[1].Name])
		createdRecordValue2, _ := propertyType.UnPack(updatedRecord.Properties.GetFields()[subCase.resource.Properties[2].Name])

		recordValue0, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[0].Name])
		recordValue1, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[1].Name])
		recordValue2, _ := propertyType.UnPack(record.Properties.GetFields()[subCase.resource.Properties[2].Name])

		if !propertyType.Equals(createdRecordValue0, recordValue0) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue0, recordValue0))
		}

		if !propertyType.Equals(createdRecordValue1, recordValue1) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue1, recordValue1))
		}

		if !propertyType.Equals(createdRecordValue2, recordValue2) {
			t.Error(fmt.Sprintf("values are different: %s <=> %s", createdRecordValue2, recordValue2))
		}
	}
}

func testRecordCreationValidationInvalidCase(ctx context.Context, t *testing.T, subCase TestRecordCreationValidationSubCase) {
	if fakeInvalidValue(subCase.recordType) == nil {
		return
	}

	var records []*model.Record
	for i := 0; i < 30; i += 3 {
		var propertiesMap = make(map[string]interface{}, 3)

		propertiesMap[subCase.resource.Properties[0].Name] = fakeInvalidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[1].Name] = fakeInvalidValue(subCase.recordType)
		propertiesMap[subCase.resource.Properties[2].Name] = fakeInvalidValue(subCase.recordType)

		properties, err := structpb.NewStruct(propertiesMap)

		if err != nil {
			t.Error(err)
			return
		}

		validRecord := &model.Record{
			Resource:   subCase.resource.Name,
			Properties: properties,
		}

		records = append(records, validRecord)
	}

	_, err := recordServiceClient.Create(ctx, &stub.CreateRecordRequest{
		Records: records,
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

	if len(errorFields) != len(records)*3 {
		t.Error("Wrong error count: " + strconv.Itoa(len(errorFields)))
		return
	}
}

func prepareTestRecordCreationValidationSubCase() []TestRecordCreationValidationSubCase {
	typs := types.GetAllResourcePropertyTypes()

	var cases []TestRecordCreationValidationSubCase

	for _, typ := range typs {
		propNames := []string{
			fakePropertyName(),
			fakePropertyName(),
			fakePropertyName(),
			fakePropertyName(),
		}
		length := uint32(32)
		resource := fakeResource(
			&model.ResourceProperty{
				Name: propNames[0],
				Type: typ,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: propNames[0],
					},
				},
				Required: false,
				Primary:  false,
				Length:   length,
				Unique:   false,
			},
			&model.ResourceProperty{
				Name: propNames[1],
				Type: typ,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: propNames[1],
					},
				},
				Required: true,
				Primary:  false,
				Length:   length,
				Unique:   false,
			},
			&model.ResourceProperty{
				Name: propNames[2],
				Type: typ,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: propNames[2],
					},
				},
				Required: false,
				Primary:  false,
				Length:   length,
				Unique:   typ != model.ResourcePropertyType_TYPE_BOOL,
			},
		)

		cases = append(cases, TestRecordCreationValidationSubCase{
			recordType: typ,
			resource:   resource,
		})
	}

	return cases
}
