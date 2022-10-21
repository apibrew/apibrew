package test

//
//import (
//	"data-handler/stub"
//	"data-handler/stub/model"
//	"fmt"
//	"google.golang.org/protobuf/encoding/protojson"
//	"google.golang.org/protobuf/types/known/structpb"
//	"testing"
//)
//
//func TestTest(t *testing.T) {
//	val1, err := structpb.NewValue("month")
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	val2, err := structpb.NewValue("c_00001_00010")
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	var r = &stub.ListRecordRequest{
//		Query: &model.BooleanExpression{
//			Expression: &model.BooleanExpression_And{
//				And: &model.CompoundBooleanExpression{
//					Expressions: []*model.BooleanExpression{
//						{
//							Expression: &model.BooleanExpression_Equal{
//								Equal: &model.PairExpression{
//									Left: &model.Expression{
//										Expression: &model.Expression_Property{Property: "founded_on_precision"},
//									},
//									Right: &model.Expression{
//										Expression: &model.Expression_Value{Value: val1},
//									},
//								},
//							},
//						},
//						{
//							Expression: &model.BooleanExpression_Equal{
//								Equal: &model.PairExpression{
//									Left: &model.Expression{
//										Expression: &model.Expression_Property{Property: "num_employees_enum"},
//									},
//									Right: &model.Expression{
//										Expression: &model.Expression_Value{Value: val2},
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	res, err := protojson.Marshal(r)
//
//	fmt.Print(string(res), err)
//}
