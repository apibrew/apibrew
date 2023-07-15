package test

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	util2 "github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestListRecord1(t *testing.T) {

	withAutoLoadedResource(setup.Ctx, t, setup.DataSource1, "public", "organization", func(resource *model.Resource) {
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

		res, err := recordClient.Search(setup.Ctx, &stub.SearchRecordRequest{
			Resource: resource.Name,
			Query: &model.BooleanExpression{
				Expression: &model.BooleanExpression_And{
					And: &model.CompoundBooleanExpression{
						Expressions: []*model.BooleanExpression{
							{
								Expression: &model.BooleanExpression_Equal{
									Equal: &model.PairExpression{
										Left: &model.Expression{
											Expression: &model.Expression_Property{Property: "foundedOnPrecision"},
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
											Expression: &model.Expression_Property{Property: "numEmployeesEnum"},
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

func withAutoLoadedResource(ctx context.Context, t testing.TB, dataSource *resource_model.DataSource, catalog, entity string, exec func(resource *model.Resource)) {
	log.Print("begin PrepareResourceFromEntity", catalog, entity, dataSource.Id)
	res, err := dataSourceClient.PrepareResourceFromEntity(ctx, &stub.PrepareResourceFromEntityRequest{
		Id:      dataSource.Id.String(),
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
		_, err := resourceClient.Delete(ctx, &stub.DeleteResourceRequest{
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
	createRes, err := resourceClient.Create(ctx, &stub.CreateResourceRequest{
		Resources:      []*model.Resource{res.Resource},
		DoMigration:    false,
		ForceMigration: false,
	})

	if err != nil {
		if util2.GetErrorCode(err) == model.ErrorCode_ALREADY_EXISTS {
			res2, _ := resourceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
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
