package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/types"
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

func InitRecord(ctx context.Context, record *model.Record) {
	now := time.Now()
	recordNewId, _ := uuid.NewUUID()
	record.Id = recordNewId.String()
	record.AuditData = &model.AuditData{
		CreatedOn: timestamppb.New(now),
		CreatedBy: security.GetUserPrincipalFromContext(ctx),
	}
	record.Version = 1
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
	if record.AuditData == nil {
		record.AuditData = &model.AuditData{}
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

func PrepareUpdateForRecord(ctx context.Context, record *model.Record) {
	if record.AuditData == nil {
		record.AuditData = &model.AuditData{
			CreatedOn: timestamppb.New(time.Now()),
			CreatedBy: "unknown",
		}
	}

	now := time.Now()
	record.AuditData.UpdatedOn = timestamppb.New(now)
	record.AuditData.UpdatedBy = security.GetUserPrincipalFromContext(ctx)
	record.Version++
}
