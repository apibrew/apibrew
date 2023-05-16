package helper

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type RecordSpecialColumnHelper struct {
	Resource *model.Resource
	Record   *model.Record
}

func (h RecordSpecialColumnHelper) IsAuditEnabled() bool {
	if h.Resource.Namespace != "system" {
		log.Print("found")
	}
	return annotations.IsEnabled(h.Resource, annotations.EnableAudit)
}

func (h RecordSpecialColumnHelper) IsVersionEnabled() bool {
	return !annotations.IsEnabled(h.Resource, annotations.DisableVersion)
}

func (h RecordSpecialColumnHelper) IncreaseVersion() {
	h.Record.Properties["version"] = structpb.NewNumberValue(h.Record.Properties["version"].GetNumberValue() + 1)
}

func (h RecordSpecialColumnHelper) InitVersion() {
	h.Record.Properties["version"] = structpb.NewNumberValue(1)
}

func (h RecordSpecialColumnHelper) GetCreatedOn() *timestamppb.Timestamp {
	if h.Record.Properties["createdOn"] == nil {
		return nil
	}

	val, err := types.TimestampType.UnPack(h.Record.Properties["createdOn"])

	if err != nil {
		panic(err)
	}

	return timestamppb.New(val.(time.Time))
}

func (h RecordSpecialColumnHelper) SetCreatedOn(createdOn *timestamppb.Timestamp) {
	if createdOn == nil {
		delete(h.Record.Properties, "createdOn")
	}

	val, err := types.TimestampType.Pack(createdOn.AsTime())

	if err != nil {
		panic(err)
	}

	h.Record.Properties["createdOn"] = val
}

func (h RecordSpecialColumnHelper) GetCreatedBy() *string {
	if h.Record.Properties["createdBy"] == nil {
		return nil
	}

	val := h.Record.Properties["createdBy"].GetStringValue()
	return &val
}

func (h RecordSpecialColumnHelper) SetCreatedBy(createdBy string) {
	h.Record.Properties["createdBy"] = structpb.NewStringValue(createdBy)
}

func (h RecordSpecialColumnHelper) SetUpdatedOn(updatedOn *timestamppb.Timestamp) {
	if updatedOn == nil {
		delete(h.Record.Properties, "updatedOn")
	}

	val, err := types.TimestampType.Pack(updatedOn.AsTime())

	if err != nil {
		panic(err)
	}

	h.Record.Properties["updatedOn"] = val
}

func (h RecordSpecialColumnHelper) HasIdSpecialProperty() bool {
	for _, prop := range h.Resource.Properties {
		if prop.Name == "id" && prop.Type == model.ResourceProperty_UUID && prop.Primary {
			return true
		}
	}

	return false
}

func (h RecordSpecialColumnHelper) SetId(id string) {
	h.Record.Properties["id"] = structpb.NewStringValue(id)
}

func (h RecordSpecialColumnHelper) SetUpdatedBy(updatedBy string) {
	h.Record.Properties["updatedBy"] = structpb.NewStringValue(updatedBy)
}

func (h RecordSpecialColumnHelper) GetVersion() uint32 {
	if h.Record.Properties["version"] == nil {
		return 0
	}

	return uint32(h.Record.Properties["version"].GetNumberValue())
}
