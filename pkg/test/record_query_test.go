package test

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	util2 "github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestListRecord1(t *testing.T) {

	withAutoLoadedResource(ctx, t, dataSource1, "public", "organization", func(resource *model.Resource) {
		val1, err := structpb.NewValue("month")

		if err != nil {
			t.Error(err)
			return
		}

		val2, err := structpb.NewValue("c_00001_00010")

		if err != nil {
			t.Error(err)
			return
		}

		res, err := recordServiceClient.Search(ctx, &stub.SearchRecordRequest{
			Resource: resource.Name,
			Query: &model.BooleanExpression{
				Expression: &model.BooleanExpression_And{
					And: &model.CompoundBooleanExpression{
						Expressions: []*model.BooleanExpression{
							{
								Expression: &model.BooleanExpression_Equal{
									Equal: &model.PairExpression{
										Left: &model.Expression{
											Expression: &model.Expression_Property{Property: "founded_on_precision"},
										},
										Right: &model.Expression{
											Expression: &model.Expression_Value{Value: val1},
										},
									},
								},
							},
							{
								Expression: &model.BooleanExpression_Equal{
									Equal: &model.PairExpression{
										Left: &model.Expression{
											Expression: &model.Expression_Property{Property: "num_employees_enum"},
										},
										Right: &model.Expression{
											Expression: &model.Expression_Value{Value: val2},
										},
									},
								},
							},
						},
					},
				},
			},
		})

		log.Print(res, err)

		if err != nil {
			t.Error(err)
			return
		}

		if res.Total != 0 { // fix by inserting real data
			t.Error("Unknown record count")
		}
	})
}

func withAutoLoadedResource(ctx context.Context, t testing.TB, dataSource *model.DataSource, catalog, entity string, exec func(resource *model.Resource)) {
	log.Print("begin PrepareResourceFromEntity", catalog, entity, dataSource.Id)
	res, err := dataSourceServiceClient.PrepareResourceFromEntity(ctx, &stub.PrepareResourceFromEntityRequest{
		Id:      dataSource.Id,
		Catalog: catalog,
		Entity:  entity,
	})
	log.Print("end PrepareResourceFromEntity", catalog, entity, dataSource.Id)

	if err != nil {
		t.Error(err)
		return
	}

	var resourceId string

	defer func() {
		if resourceId == "" {
			return
		}

		log.Print("begin delete resource without migration", res.Resource.Namespace, res.Resource.Name)
		_, err := resourceServiceClient.Delete(ctx, &stub.DeleteResourceRequest{
			Ids:            []string{resourceId},
			DoMigration:    false,
			ForceMigration: false,
		})

		if err != nil {
			t.Error(err)
			return
		}

		log.Info("resource deleted: " + res.Resource.Name)
	}()

	log.Print("finish PrepareResourceFromEntity", catalog, entity, dataSource.Id)

	log.Print("begin create resource without migration", res.Resource.Namespace, res.Resource.Name)
	createRes, err := resourceServiceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{res.Resource},
		DoMigration:    false,
		ForceMigration: false,
	})

	if err != nil {
		if util2.GetErrorCode(err) == model.ErrorCode_ALREADY_EXISTS {
			res2, _ := resourceServiceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
				Namespace: res.Resource.Namespace,
				Name:      res.Resource.Name,
			})
			resourceId = res2.Resource.Id

		} else {
			t.Error(err)
			return
		}
	} else {
		resourceId = createRes.Resources[0].Id
	}

	log.Print("finish create resource without migration", res.Resource.Namespace, res.Resource.Name)

	log.Print("Calling exec: ", res.Resource.Namespace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)
	exec(res.Resource)
	log.Print("Finished exec: ", res.Resource.Namespace, " ", res.Resource.Name, " ", res.Resource.SourceConfig.DataSource)
}
