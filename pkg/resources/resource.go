package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourcePropertyProperties = []*model.ResourceProperty{
	{
		Name:     "name",
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false, // can be optional for item types
	},
	{
		Name: "type",
		Type: model.ResourceProperty_ENUM,
		EnumValues: []string{
			"BOOL",
			"STRING",
			"FLOAT32",
			"FLOAT64",
			"INT32",
			"INT64",
			"BYTES",
			"UUID",
			"DATE",
			"TIME",
			"TIMESTAMP",
			"OBJECT",
			"MAP",
			"LIST",
			"REFERENCE",
			"ENUM",
			"STRUCT",
		},
		Required: true,
	},
	{
		Name:     "typeRef",
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	{
		Name:         "primary",
		Type:         model.ResourceProperty_BOOL,
		Required:     true,
		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:         "required",
		Type:         model.ResourceProperty_BOOL,
		Required:     true,
		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:         "unique",
		Type:         model.ResourceProperty_BOOL,
		Required:     true,
		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:         "immutable",
		Type:         model.ResourceProperty_BOOL,
		Required:     true,
		DefaultValue: structpb.NewBoolValue(false),
	},
	{
		Name:         "length",
		Type:         model.ResourceProperty_INT32,
		Required:     true,
		DefaultValue: structpb.NewNumberValue(256),
	},
	{
		Name:     "item",
		Type:     model.ResourceProperty_STRUCT,
		Required: false,
		TypeRef:  util.Pointer("Property"),
	},
	{
		Name:    "reference",
		Type:    model.ResourceProperty_STRUCT,
		TypeRef: util.Pointer("Reference"),
	},
	{
		Name:     "defaultValue",
		Type:     model.ResourceProperty_OBJECT,
		Required: false,
	},
	{
		Name: "enumValues",
		Type: model.ResourceProperty_LIST,
		Item: &model.ResourceProperty{
			Type: model.ResourceProperty_STRING,
		},
		Required: false,
	},
	{
		Name:     "exampleValue",
		Type:     model.ResourceProperty_OBJECT,
		Required: false,
	},
	{
		Name:     "title",
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	{
		Name:     "description",
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	special.AnnotationsProperty,
}

var PropertyType = &model.ResourceSubType{
	Name: "Property",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: ResourcePropertyProperties,
}

var SubTypeType = &model.ResourceSubType{
	Name: "SubType",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: []*model.ResourceProperty{
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "title",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "properties",
			Type:     model.ResourceProperty_LIST,
			Required: true,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
		},
	},
}

var ReferenceType = &model.ResourceSubType{
	Name: "Reference",
	Annotations: map[string]string{
		annotations.CommonType: annotations.Enabled,
	},
	Properties: []*model.ResourceProperty{
		{
			Name: "resource",
			Type: model.ResourceProperty_REFERENCE,
			Reference: &model.Reference{
				Namespace: "system",
				Resource:  "Resource",
			},
		},
		{
			Name: "cascade",
			Type: model.ResourceProperty_BOOL,
		},
		{
			Name: "backReference",
			Type: model.ResourceProperty_STRING,
		},
	},
}

var ResourceResource = &model.Resource{
	Name:      "Resource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
	},
	Types: []*model.ResourceSubType{
		PropertyType,
		SubTypeType,
		{
			Name: "IndexProperty",
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "order",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"UNKNOWN", "ASC", "DESC",
					},
					DefaultValue: structpb.NewStringValue("ASC"),
				},
			},
		},
		{
			Name: "Index",
			Properties: []*model.ResourceProperty{
				{
					Name: "properties",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("IndexProperty"),
					},
				},
				{
					Name:     "indexType",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"BTREE", "HASH",
					},
					DefaultValue: structpb.NewStringValue("BTREE"),
				},
				{
					Name:     "unique",
					Type:     model.ResourceProperty_BOOL,
					Required: false,
				},
				special.AnnotationsProperty,
			},
		},
		ReferenceType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "namespace",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				Resource:  NamespaceResource.Name,
				Namespace: NamespaceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:         "virtual",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:     "properties",
			Type:     model.ResourceProperty_LIST,
			Required: true,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
		},
		{
			Name:     "indexes",
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Index"),
			},
		},
		{
			Name:     "types",
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("SubType"),
			},
		},
		{
			Name:         "immutable",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:         "abstract",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:         "checkReferences",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
		},
		{
			Name:     "dataSource",
			Type:     model.ResourceProperty_REFERENCE,
			Required: false,
			Reference: &model.Reference{
				Resource:  DataSourceResource.Name,
				Namespace: DataSourceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:     "entity",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "catalog",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "title",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		special.AnnotationsProperty,
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "namespace",
				},
				{
					Name: "name",
				},
			},
			Unique: true,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:     annotations.Enabled,
		annotations.RestApiDisabled: annotations.Enabled,
		annotations.OpenApiGroup:    OpenApiMeta,
	},
}
