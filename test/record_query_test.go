package test

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestListRecord1(t *testing.T) {
	ctx := prepareTextContext()

	withAutoLoadedResource(ctx, t, container, dataSource1, "public.organization", func(resource *model.Resource) {
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

		res, err := container.recordService.List(ctx, &stub.ListRecordRequest{
			Token:    "test-token",
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

		if res.Total != 16 {
			t.Error("Unknown record count")
		}
	})
}
