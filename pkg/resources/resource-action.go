package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourceActionResource = &model.Resource{
	Name:      "ResourceAction",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_action",
	},
	Types: []*model.ResourceSubType{
		SubTypeType,
		PropertyType,
		special.AuditDataSubType,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperty,
		{
			Name:     "resource",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				Resource:  ResourceResource.Name,
				Namespace: ResourceResource.Namespace,
				Cascade:   false,
			},
		},
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
			Name:     "title",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "description",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:         "internal",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
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
			Name: "input",
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
		},
		{
			Name:    "output",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("Property"),
		},
		special.AnnotationsProperty,
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "resource",
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
