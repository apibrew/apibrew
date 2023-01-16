package data

import (
	"data-handler/model"
)

func PreparePersonResource() *model.Resource {
	personResource := &model.Resource{
		Name:      "person",
		Namespace: "default",
		DataType:  2,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
			Mapping:    "person",
		},
		Properties: []*model.ResourceProperty{
			{
				Name: "name",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "name",
					},
				},
				Length:   255,
				Required: true,
			}, {
				Name: "lastName",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "last_name",
					},
				},
				Length:   255,
				Required: true,
			}, {
				Name: "age",
				Type: model.ResourcePropertyType_TYPE_INT32,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "age",
					},
				},
				Length:   255,
				Required: false,
			}, {
				Name: "gender",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "gender",
					},
				},
				Length:   10,
				Required: true,
			},
		},
		References: []*model.ResourceReference{
			{
				PropertyName:       "country",
				ReferencedResource: "rf-2-country",
				Cascade:            true,
			},
		},
	}

	return personResource
}
