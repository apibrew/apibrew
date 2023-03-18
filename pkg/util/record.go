package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/security"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func ComputeRecordIdFromProperties(resource *model.Resource, record *model.Record) error {
	var idParts []string
	for _, prop := range resource.Properties {
		val := record.Properties[prop.Name]
		if val != nil && prop.Primary {
			typ := types.ByResourcePropertyType(prop.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return err
			}
			if unpacked == nil {
				continue
			}
			idParts = append(idParts, typ.String(unpacked))
		}
	}
	record.Id = strings.Join(idParts, "-")

	return nil
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

func PrepareUpdateForRecord(ctx context.Context, record *model.Record) {
	if record.AuditData == nil {
		record.AuditData = &model.AuditData{}
	}

	now := time.Now()
	record.AuditData.UpdatedOn = timestamppb.New(now)
	record.AuditData.UpdatedBy = security.GetUserPrincipalFromContext(ctx)
	record.Version++
}
