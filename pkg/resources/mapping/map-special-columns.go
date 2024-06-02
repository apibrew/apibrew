package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

func MapSpecialColumnsToRecord(instance any, properties *map[string]interface{}) {
	ref := reflect.Indirect(reflect.ValueOf(instance))
	versionCol := ref.FieldByName("Version")

	if versionCol.IsValid() {
		(*properties)["version"] = structpb.NewNumberValue(float64(versionCol.Interface().(uint32)))
	}

	auditData := ref.FieldByName("AuditData")

	if auditData.IsValid() {
		auditVal := auditData.Interface().(*model.AuditData)
		if auditVal != nil {
			if (*properties)["auditData"] == nil {
				(*properties)["auditData"] = structpb.NewStructValue(&structpb.Struct{Fields: make(map[string]*structpb.Value)})
			}

			(*properties)["auditData"].GetStructValue().Fields["createdBy"] = structpb.NewStringValue(auditVal.CreatedBy)

			if auditVal.UpdatedBy != "" {
				(*properties)["auditData"].GetStructValue().Fields["updatedBy"] = structpb.NewStringValue(auditVal.UpdatedBy)
			}

			if auditVal.CreatedOn != nil {
				val, err := types.TimestampType.Pack(auditVal.CreatedOn.AsTime())

				if err != nil {
					panic(err)
				}

				(*properties)["auditData"].GetStructValue().Fields["createdOn"] = val
			}

			if auditVal.UpdatedOn != nil {
				val, err := types.TimestampType.Pack(auditVal.UpdatedOn.AsTime())

				if err != nil {
					panic(err)
				}

				(*properties)["auditData"].GetStructValue().Fields["updatedOn"] = val
			}
		}
	}
}

func MapSpecialColumnsFromRecord(instance any, properties *map[string]*structpb.Value) {
	ref := reflect.Indirect(reflect.ValueOf(instance))
	versionCol := ref.FieldByName("Version")

	if versionCol.IsValid() {
		versionCol.SetUint(uint64((*properties)["version"].GetNumberValue()))
	}

	auditData := ref.FieldByName("AuditData")

	if auditData.IsValid() && (*properties)["auditData"] != nil {
		auditVal := &model.AuditData{}
		var auditProperties = (*properties)["auditData"].GetStructValue().Fields

		if auditProperties["createdBy"] != nil {
			auditVal.CreatedBy = auditProperties["createdBy"].GetStringValue()
		}

		if auditProperties["createdOn"] != nil {
			val, err := types.TimestampType.UnPack(auditProperties["createdOn"])

			if err != nil {
				panic(err)
			}

			auditVal.CreatedOn = timestamppb.New(val.(time.Time))
		}

		if auditProperties["updatedBy"] != nil {
			auditVal.UpdatedBy = auditProperties["updatedBy"].GetStringValue()
		}

		if auditProperties["updatedOn"] != nil {
			val, err := types.TimestampType.UnPack(auditProperties["updatedOn"])

			if err != nil {
				panic(err)
			}

			auditVal.UpdatedOn = timestamppb.New(val.(time.Time))
		}

		auditData.Set(reflect.ValueOf(auditVal))
	}
}
