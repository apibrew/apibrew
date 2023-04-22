package mapping

import (
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

func mapSpecialColumnsToRecord(instance any, properties *map[string]*structpb.Value) {
	ref := reflect.Indirect(reflect.ValueOf(instance))
	versionCol := ref.FieldByName("Version")

	if versionCol.IsValid() {
		(*properties)["version"] = structpb.NewNumberValue(float64(versionCol.Interface().(uint32)))
	}

	auditData := ref.FieldByName("AuditData")

	if auditData.IsValid() {
		auditVal := auditData.Interface().(*model.AuditData)
		if auditVal != nil {
			(*properties)["createdBy"] = structpb.NewStringValue(auditVal.CreatedBy)

			if auditVal.UpdatedBy != "" {
				(*properties)["updatedBy"] = structpb.NewStringValue(auditVal.UpdatedBy)
			}

			if auditVal.CreatedOn != nil {
				val, err := types.TimestampType.Pack(auditVal.CreatedOn.AsTime())

				if err != nil {
					panic(err)
				}

				(*properties)["createdOn"] = val
			}

			if auditVal.UpdatedOn != nil {
				val, err := types.TimestampType.Pack(auditVal.UpdatedOn.AsTime())

				if err != nil {
					panic(err)
				}

				(*properties)["updatedOn"] = val
			}
		}
	}
}

func mapSpecialColumnsFromRecord(instance any, properties *map[string]*structpb.Value) {
	ref := reflect.Indirect(reflect.ValueOf(instance))
	versionCol := ref.FieldByName("Version")

	if versionCol.IsValid() {
		versionCol.SetUint(uint64((*properties)["version"].GetNumberValue()))
	}

	auditData := ref.FieldByName("AuditData")

	if auditData.IsValid() {
		auditVal := &model.AuditData{}

		if (*properties)["createdBy"] != nil {
			auditVal.CreatedBy = (*properties)["createdBy"].GetStringValue()
		}

		if (*properties)["createdOn"] != nil {
			val, err := types.TimestampType.UnPack((*properties)["createdOn"])

			if err != nil {
				panic(err)
			}

			auditVal.CreatedOn = timestamppb.New(val.(time.Time))
		}

		if (*properties)["updatedBy"] != nil {
			auditVal.UpdatedBy = (*properties)["updatedBy"].GetStringValue()
		}

		if (*properties)["updatedOn"] != nil {
			val, err := types.TimestampType.UnPack((*properties)["updatedOn"])

			if err != nil {
				panic(err)
			}

			auditVal.UpdatedOn = timestamppb.New(val.(time.Time))
		}

		auditData.Set(reflect.ValueOf(auditVal))
	}
}
