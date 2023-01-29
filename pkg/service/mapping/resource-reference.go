package mapping

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/system"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourceReferenceToRecord(property *model.ResourceReference, resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["propertyName"] = structpb.NewStringValue(property.PropertyName)
	properties["referencedResource"] = structpb.NewStringValue(property.ReferencedResource)
	properties["cascade"] = structpb.NewBoolValue(property.Cascade)

	return &model.Record{
		Resource:   system.ResourceReferenceResource.Name,
		DataType:   model.DataType_SYSTEM,
		Properties: properties,
	}
}

func ResourceReferenceFromRecord(record *model.Record) *model.ResourceReference {
	if record == nil {
		return nil
	}

	var resource = &model.ResourceReference{
		PropertyName:       record.Properties["propertyName"].GetStringValue(),
		ReferencedResource: record.Properties["referencedResource"].GetStringValue(),
		Cascade:            record.Properties["cascade"].GetBoolValue(),
	}

	return resource
}
