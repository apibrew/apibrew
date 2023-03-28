package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func ComputeRecordIdFromProperties(resource *model.Resource, record *model.Record) {
	var idParts []string
	for _, prop := range resource.Properties {
		val := record.Properties[prop.Name]
		if val != nil && prop.Primary {
			typ := types.ByResourcePropertyType(prop.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				panic(err)
			}
			if unpacked == nil {
				continue
			}
			idParts = append(idParts, typ.String(unpacked))
		}
	}
	record.Id = strings.Join(idParts, "-")
}

func InitRecord(ctx context.Context, resource *model.Resource, record *model.Record) {
	now := time.Now()
	recordNewId, _ := uuid.NewUUID()
	record.Id = recordNewId.String()
	if record.Properties == nil {
		record.Properties = make(map[string]*structpb.Value)
	}

	ah := helper.RecordSpecialColumnHelper{
		Resource: resource,
		Record:   record,
	}

	if ah.IsAuditEnabled() {
		ah.SetCreatedOn(timestamppb.New(now))
		ah.SetCreatedBy(security.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.InitVersion()
	}
}

func NormalizeRecord(resource *model.Resource, record *model.Record) {
	if record.Properties == nil {
		record.Properties = make(map[string]*structpb.Value)
	}

	specialProps := resources.GetResourceSpecialProperties(resource)

	for _, prop := range specialProps {
		var err error
		val := prop.Get(record)

		if val != nil {
			record.Properties[prop.Property.Name], err = types.ByResourcePropertyType(prop.Property.Type).Pack(val)
		}

		if err != nil {
			panic(err)
		}
	}
}

func DeNormalizeRecord(resource *model.Resource, record *model.Record) {
	if record.Properties == nil {
		return
	}

	specialProps := resources.GetResourceSpecialProperties(resource)

	for _, prop := range specialProps {
		if record.Properties[prop.Property.Name] == nil {
			continue
		}

		val, err := types.ByResourcePropertyType(prop.Property.Type).UnPack(record.Properties[prop.Property.Name])

		if err != nil {
			panic(err)
		}

		prop.Set(record, val)
	}

	if record.Id == "" {
		ComputeRecordIdFromProperties(resource, record)
	}
}

func PrepareUpdateForRecord(ctx context.Context, resource *model.Resource, record *model.Record) {
	ah := &helper.RecordSpecialColumnHelper{
		Resource: resource,
		Record:   record,
	}

	if ah.IsAuditEnabled() {
		if ah.GetCreatedOn() == nil {
			ah.SetCreatedOn(timestamppb.New(time.Now()))
		}

		if ah.GetCreatedBy() == nil {
			ah.SetCreatedBy(security.GetUserPrincipalFromContext(ctx))
		}

		ah.SetUpdatedOn(timestamppb.New(time.Now()))
		ah.SetUpdatedBy(security.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.IncreaseVersion()
	}
}

func IsSameRecord(existing, updated *model.Record) bool {
	for key := range updated.Properties {
		if !proto.Equal(updated.Properties[key], existing.Properties[key]) {
			return false
		}
	}

	return true
}
