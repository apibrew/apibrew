package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func InitRecord(ctx context.Context, resource *model.Resource, record *model.Record) {
	now := time.Now()
	recordNewId := uuid.Must(uuid.NewRandom())
	record.Id = recordNewId.String()
	if record.Properties == nil {
		record.Properties = make(map[string]*structpb.Value)
	}

	record.Properties["id"] = structpb.NewStringValue(recordNewId.String())

	ah := helper.RecordSpecialColumnHelper{
		Resource: resource,
		Record:   record,
	}

	if ah.IsAuditEnabled() {
		ah.SetCreatedOn(timestamppb.New(now))
		ah.SetCreatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.InitVersion()
	}

	if ah.HasIdSpecialProperty() {
		ah.SetId(record.Id)
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
			ah.SetCreatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
		}
		ah.SetUpdatedOn(timestamppb.New(time.Now()))
		ah.SetUpdatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.IncreaseVersion()
	}
}
