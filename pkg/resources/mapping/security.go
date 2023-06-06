package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func SecurityContextToValue(securityConstraints []*model.SecurityConstraint) *structpb.Value {
	var list []interface{}

	for _, item := range securityConstraints {
		var properties = make(map[string]interface{})

		properties["namespace"] = item.Namespace
		properties["resource"] = item.Resource
		properties["property"] = item.Property
		properties["property"] = item.Property
		properties["before"] = item.Before.AsTime().UnixMilli()
		properties["after"] = item.After.AsTime().UnixMilli()
		properties["username"] = item.Username
		properties["operation"] = int32(item.Operation.Number())
		properties["permit"] = int32(item.Permit.Number())
		properties["recordIds"] = util.ArrayMap(item.GetRecordIds(), func(t string) interface{} {
			return t
		})

		list = append(list, properties)
	}

	listVal, err := structpb.NewList(list)

	if err != nil {
		panic(err)
	}

	return structpb.NewListValue(listVal)
}

func SecurityContextFromValue(value *structpb.Value) []*model.SecurityConstraint {
	if value == nil {
		return nil
	}

	securityContext := new(model.SecurityContext)

	if value.GetListValue() == nil {
		return nil
	}

	for _, value := range value.GetListValue().Values {
		obj := value.GetStructValue()

		securityConstraint := new(model.SecurityConstraint)

		securityConstraint.Namespace = obj.Fields["namespace"].GetStringValue()
		securityConstraint.Resource = obj.Fields["resource"].GetStringValue()
		securityConstraint.Property = obj.Fields["property"].GetStringValue()
		securityConstraint.Before = timestamppb.New(time.UnixMilli(int64(obj.Fields["before"].GetNumberValue())))
		securityConstraint.After = timestamppb.New(time.UnixMilli(int64(obj.Fields["after"].GetNumberValue())))
		securityConstraint.Username = obj.Fields["username"].GetStringValue()
		securityConstraint.Operation = model.OperationType(obj.Fields["operation"].GetNumberValue())
		securityConstraint.Permit = model.PermitType(obj.Fields["permit"].GetNumberValue())

		if obj.Fields["recordIds"] != nil {
			securityConstraint.RecordIds = util.ArrayMap(obj.Fields["recordIds"].GetListValue().Values, func(t *structpb.Value) string {
				return t.GetStringValue()
			})
		}

		securityContext.Constraints = append(securityContext.Constraints, securityConstraint)
	}

	return securityContext.Constraints
}
