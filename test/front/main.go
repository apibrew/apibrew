package main

import (
	"context"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("tw:9009", opts...)

	if err != nil {
		panic(err)
	}

	resourceService := stub.NewResourceServiceClient(conn)
	//recordService := stub.NewRecordServiceClient(conn)

	prepareRichResource(resourceService)
}

func prepareRichResource(service stub.ResourceServiceClient) {
	richResource := &model.Resource{
		Name:      "rich-test-2",
		Workspace: "default",
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
