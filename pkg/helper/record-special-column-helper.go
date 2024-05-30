package helper

import (
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type RecordSpecialColumnHelper struct {
	Resource *model.Resource
	Record   abs.RecordLike
}

func (h RecordSpecialColumnHelper) IsAuditEnabled() bool {
	return annotations.IsEnabled(h.Resource, annotations.EnableAudit)
}

func (h RecordSpecialColumnHelper) IsVersionEnabled() bool {
	return !annotations.IsEnabled(h.Resource, annotations.DisableVersion)
}

func (h RecordSpecialColumnHelper) IncreaseVersion() {
	h.Record.GetProperties()["version"] = structpb.NewNumberValue(h.Record.GetProperties()["version"].GetNumberValue() + 1)
}

func (h RecordSpecialColumnHelper) InitVersion() {
	h.Record.GetProperties()["version"] = structpb.NewNumberValue(1)
}

func (h RecordSpecialColumnHelper) GetCreatedOn() *timestamppb.Timestamp {
	if h.Record.GetProperties()["auditData"] == nil {
		return nil
	}

	val, err := types.TimestampType.UnPack(h.Record.GetProperties()["auditData"].GetStructValue().GetFields()["createdOn"])

	if err != nil {
		panic(err)
	}

	return timestamppb.New(val.(time.Time))
}

func (h RecordSpecialColumnHelper) SetCreatedOn(createdOn *timestamppb.Timestamp) {
	if createdOn == nil {
		delete(h.Record.GetProperties(), "auditData")
	}

	val, err := types.TimestampType.Pack(createdOn.AsTime())

	if err != nil {
		panic(err)
	}

	h.ensureAuditData()

	h.Record.GetProperties()["auditData"].GetStructValue().Fields["createdOn"] = val
}

func (h RecordSpecialColumnHelper) ensureAuditData() {
	if h.Record.GetProperties()["auditData"] == nil || h.Record.GetProperties()["auditData"].AsInterface() == nil {
		h.Record.GetProperties()["auditData"] = structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{},
		})
	}
}

func (h RecordSpecialColumnHelper) GetCreatedBy() *string {
	if h.Record.GetProperties()["auditData"] == nil || h.Record.GetProperties()["auditData"].AsInterface() == nil {
		return nil
	}

	val := h.Record.GetProperties()["auditData"].GetStructValue().Fields["createdBy"].GetStringValue()
	return &val
}

func (h RecordSpecialColumnHelper) SetCreatedBy(createdBy string) {
	h.ensureAuditData()

	h.Record.GetProperties()["auditData"].GetStructValue().Fields["createdBy"] = structpb.NewStringValue(createdBy)
}

func (h RecordSpecialColumnHelper) SetUpdatedOn(updatedOn *timestamppb.Timestamp) {
	if updatedOn == nil {
		delete(h.Record.GetProperties()["auditData"].GetStructValue().Fields, "updatedOn")
	}

	val, err := types.TimestampType.Pack(updatedOn.AsTime())

	if err != nil {
		panic(err)
	}

	h.ensureAuditData()

	h.Record.GetProperties()["auditData"].GetStructValue().Fields["updatedOn"] = val
}

func (h RecordSpecialColumnHelper) SetId(id string) {
	h.Record.GetProperties()["id"] = structpb.NewStringValue(id)
}

func (h RecordSpecialColumnHelper) SetUpdatedBy(updatedBy string) {

	h.ensureAuditData()

	h.Record.GetProperties()["auditData"].GetStructValue().Fields["updatedBy"] = structpb.NewStringValue(updatedBy)
}

func (h RecordSpecialColumnHelper) GetVersion() uint32 {
	if h.Record.GetProperties()["version"] == nil {
		return 0
	}

	return uint32(h.Record.GetProperties()["version"].GetNumberValue())
}
