package special

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

var IdProperty = &model.ResourceProperty{
	Name:        "id",
	Type:        model.ResourceProperty_UUID,
	Primary:     true,
	Required:    true,
	Immutable:   true,
	Description: Pointer("The unique identifier of the resource. It is randomly generated and immutable."),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
	ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),
}

var VersionProperty = &model.ResourceProperty{
	Name:         "version",
	Title:        Pointer("Version"),
	Description:  Pointer("The version of the resource/record. It is incremented on every update."),
	Type:         model.ResourceProperty_INT32,
	Required:     true,
	DefaultValue: structpb.NewNumberValue(1),
	ExampleValue: structpb.NewNumberValue(1),
	Annotations: map[string]string{
		annotations.SpecialProperty:     annotations.Enabled,
		annotations.AllowEmptyPrimitive: annotations.Enabled,
	},
}

var AuditPropertyCreatedBy = &model.ResourceProperty{
	Name:         "createdBy",
	Title:        Pointer("Created By"),
	Description:  Pointer("The user who created the resource/record."),
	Type:         model.ResourceProperty_STRING,
	Length:       256,
	Immutable:    true,
	ExampleValue: structpb.NewStringValue("admin"),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedBy = &model.ResourceProperty{
	Name:         "updatedBy",
	Title:        Pointer("Updated By"),
	Description:  Pointer("The user who last updated the resource/record."),
	Type:         model.ResourceProperty_STRING,
	Length:       256,
	Required:     false,
	ExampleValue: structpb.NewStringValue("admin"),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyCreatedOn = &model.ResourceProperty{
	Name:         "createdOn",
	Title:        Pointer("Created On"),
	Description:  Pointer("The timestamp when the resource/record was created."),
	Type:         model.ResourceProperty_TIMESTAMP,
	Immutable:    true,
	ExampleValue: structpb.NewStringValue(time.Now().Format(time.RFC3339)),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditPropertyUpdatedOn = &model.ResourceProperty{
	Name:         "updatedOn",
	Title:        Pointer("Updated On"),
	Description:  Pointer("The timestamp when the resource/record was last updated."),
	Type:         model.ResourceProperty_TIMESTAMP,
	Required:     false,
	ExampleValue: structpb.NewStringValue(time.Now().Format(time.RFC3339)),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AuditDataSubType = &model.ResourceSubType{
	Name:        "AuditData",
	Title:       "Audit Data",
	Description: "Audit Data is a type that represents the audit data of a resource/record. ",
	Properties: []*model.ResourceProperty{
		AuditPropertyCreatedBy,
		AuditPropertyUpdatedBy,
		AuditPropertyCreatedOn,
		AuditPropertyUpdatedOn,
	},
}

var AuditProperty = &model.ResourceProperty{
	Name:        "auditData",
	Title:       Pointer("Audit Data"),
	Description: Pointer("The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated."),
	Type:        model.ResourceProperty_STRUCT,
	ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
		"createdBy": structpb.NewStringValue("admin"),
		"updatedBy": structpb.NewStringValue("admin"),
		"createdOn": structpb.NewStringValue(time.Now().Format(time.RFC3339)),
		"updatedOn": structpb.NewStringValue(time.Now().Format(time.RFC3339)),
	}}),
	TypeRef: Pointer("AuditData"),
	Annotations: map[string]string{
		annotations.SpecialProperty: "true",
	},
}

var AnnotationsProperty = &model.ResourceProperty{
	Name:        "annotations",
	Title:       Pointer("Annotations"),
	Description: Pointer("The annotations of the resource/record. It contains information about the resource/record. For example, it can contain information about the UI representation of the resource/record."),
	Type:        model.ResourceProperty_MAP,
	Required:    false,
	Immutable:   false,
	Item: &model.ResourceProperty{
		Type: model.ResourceProperty_STRING,
	},
	ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
		annotations.CheckVersion:   structpb.NewStringValue(annotations.Enabled),
		annotations.IgnoreIfExists: structpb.NewStringValue(annotations.Enabled),
		annotations.CommonType:     structpb.NewStringValue("testType"),
	}}),
	Annotations: map[string]string{
		annotations.SpecialProperty: annotations.Enabled,
	},
}

func Pointer[T interface{}](val T) *T {
	var pointer = new(T)

	*pointer = val

	return pointer
}
