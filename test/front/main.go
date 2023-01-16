package main

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkYXRhLWhhbmRsZXIiLCJzdWIiOiJhZG1pbiIsImF1ZCI6WyJkYXRhLWhhbmRsZXIiXSwiZXhwIjoxNzI5ODY3NDM0LCJuYmYiOjE2NjY3OTU0MzQsImlhdCI6MTY2Njc5NTQzNCwianRpIjoiZDYwNjg5OTQtZTkxMy00M2NlLTg1MmQtMzdhMjAzNjBjMDY4Iiwic2NvcGVzIjpbInN1cGVyLXVzZXIiXSwidXNlcm5hbWUiOiJhZG1pbiJ9.ajMXk2TNfGwZzaSqtgQLGxofJ7Fddz2ZzsYZOOG9mWiSdDJoGC-VKyAl5zhfVBKYubqxwI2aL340nHrFB-DQ9r1yyc-6glxPpdRYxxDArXZpIIUZA4f6oKUnoYfsmCkgMdzrKdQfoetV3ABhekPwgHk1hSlMm6BDBxPCr5voClU"

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("tw:9009", opts...)

	if err != nil {
		panic(err)
	}

	resourceService := stub.NewResourceServiceClient(conn)
	//recordService := stub.NewRecordServiceClient(conn)

	//prepareRichResource(resourceService)
	//prepareReferencedResource(resourceService)
	//prepareCityResource(resourceService)
	prepareCategoryResource(resourceService)
}

func prepareCategoryResource(service stub.ResourceServiceClient) {
	personResource := &model.Resource{
		Name:      "product",
		Namespace: "default",
		Type:      2,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
			Mapping:    "product",
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
				Name: "description",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "description",
					},
				},
				Length:   255,
				Required: true,
			}, {
				Name: "category",
				Type: model.ResourcePropertyType_TYPE_UUID,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "category",
					},
				},
				Required: true,
			}},
		References: []*model.ResourceReference{
			{
				PropertyName:       "category",
				ReferencedResource: "category",
				Cascade:            true,
			},
		},
	}

	res2, err := service.Create(context.TODO(), &stub.CreateResourceRequest{
		Token:       token,
		Resources:   []*model.Resource{personResource},
		DoMigration: true,
	})

	log.Print(res2, err)
}

func prepareCityResource(service stub.ResourceServiceClient) {
	personResource := &model.Resource{
		Name:      "city",
		Namespace: "default",
		Type:      2,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
			Mapping:    "city",
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
				Name: "description",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "description",
					},
				},
				Length:   255,
				Required: true,
			}, {
				Name: "country",
				Type: model.ResourcePropertyType_TYPE_UUID,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "country",
					},
				},
				Required: false,
			}},
		References: []*model.ResourceReference{
			{
				PropertyName:       "country",
				ReferencedResource: "country",
				Cascade:            true,
			},
		},
	}

	res2, err := service.Create(context.TODO(), &stub.CreateResourceRequest{
		Token:       token,
		Resources:   []*model.Resource{personResource},
		DoMigration: true,
	})

	log.Print(res2, err)
}
func prepareReferencedResource(service stub.ResourceServiceClient) {
	//countryResource := &model.Resource{
	//	Name:      "rf--country",
	//	Namespace: "default",
	//	Type:      2,
	//	SourceConfig: &model.ResourceSourceConfig{
	//		DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
	//		Mapping:    "rf_2_country",
	//	},
	//	Properties: []*model.ResourceProperty{
	//		{
	//			Name: "name",
	//			Type: model.ResourcePropertyType_TYPE_STRING,
	//			SourceConfig: &model.ResourceProperty_Mapping{
	//				Mapping: &model.ResourcePropertyMappingConfig{
	//					Mapping: "name",
	//				},
	//			},
	//			Length:   255,
	//			Required: true,
	//			Unique:   true,
	//		},
	//	},
	//}
	//
	//res, err := service.Create(ctx, &stub.CreateResourceRequest{
	//	Token:       token,
	//	Resources:   []*model.Resource{countryResource},
	//	DoMigration: true,
	//})
	//
	//log.Print(res, err)

	cityResource := &model.Resource{
		Name:      "rf-3-city",
		Namespace: "default",
		Type:      2,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
			Mapping:    "rf_3_city",
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
			},
			{
				Name: "country",
				Type: model.ResourcePropertyType_TYPE_UUID,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "country",
					},
				},
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

	res2, err := service.Create(context.TODO(), &stub.CreateResourceRequest{
		Token:       token,
		Resources:   []*model.Resource{cityResource},
		DoMigration: true,
	})

	log.Print(res2, err)
}

func prepareRichResource(service stub.ResourceServiceClient) {
	richResource := &model.Resource{
		Name:      "rich-test-2",
		Namespace: "default",
		Type:      2,
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "0f96d8ca-4d48-11ed-a348-b29c4ac91271",
			Mapping:    "rich_test_1",
		},
		Properties: []*model.ResourceProperty{
			{
				Name: "int32_o",
				Type: model.ResourcePropertyType_TYPE_INT32,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "int32_o",
					},
				},
				Required: false,
			},

			{
				Name: "int32",
				Type: model.ResourcePropertyType_TYPE_INT32,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "int32",
					},
				},
				Required: true,
			},

			{
				Name: "int64",
				Type: model.ResourcePropertyType_TYPE_INT64,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "int64",
					},
				},
				Required: true,
			},

			{
				Name: "float",
				Type: model.ResourcePropertyType_TYPE_FLOAT,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "float",
					},
				},
				Required: true,
			},

			{
				Name: "double",
				Type: model.ResourcePropertyType_TYPE_DOUBLE,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "double",
					},
				},
				Required: true,
			},

			{
				Name: "numeric",
				Type: model.ResourcePropertyType_TYPE_NUMERIC,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "numeric",
					},
				},
				Required: true,
			},

			{
				Name: "text",
				Type: model.ResourcePropertyType_TYPE_TEXT,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "text",
					},
				},
				Required: true,
			},

			{
				Name: "string",
				Type: model.ResourcePropertyType_TYPE_STRING,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "string",
					},
				},
				Required: true,
				Length:   255,
			},
			{
				Name: "uuid",
				Type: model.ResourcePropertyType_TYPE_UUID,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "uuid",
					},
				},
				Required: true,
			},

			{
				Name: "date",
				Type: model.ResourcePropertyType_TYPE_DATE,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "date",
					},
				},
				Required: true,
			},

			{
				Name: "time",
				Type: model.ResourcePropertyType_TYPE_TIME,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "time",
					},
				},
				Required: true,
			},

			{
				Name: "timestamp",
				Type: model.ResourcePropertyType_TYPE_TIMESTAMP,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "timestamp",
					},
				},
				Required: true,
			},

			{
				Name: "bool",
				Type: model.ResourcePropertyType_TYPE_BOOL,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "bool",
					},
				},
				Required: true,
			},

			{
				Name: "object",
				Type: model.ResourcePropertyType_TYPE_OBJECT,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "object",
					},
				},
				Required: true,
			},

			{
				Name: "bytes",
				Type: model.ResourcePropertyType_TYPE_BYTES,
				SourceConfig: &model.ResourceProperty_Mapping{
					Mapping: &model.ResourcePropertyMappingConfig{
						Mapping: "bytes",
					},
				},
				Required: true,
			},
		},
	}

	res, err := service.Create(context.TODO(), &stub.CreateResourceRequest{
		Token:       "",
		Resources:   []*model.Resource{richResource},
		DoMigration: true,
	})

	log.Print(res, err)
}
