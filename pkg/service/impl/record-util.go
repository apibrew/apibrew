package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func InitRecord(ctx context.Context, resource *model.Resource, record *model.Record) {
	now := time.Now()
	recordNewId := uuid.Must(uuid.NewRandom())
	if record.Properties == nil {
		record.Properties = make(map[string]*structpb.Value)
	}

	if util.HasResourceSinglePrimaryProp(resource) {
		idProp := util.GetResourceSinglePrimaryProp(resource)
		if idProp.Type == model.ResourceProperty_UUID {
			if record.Properties[idProp.Name] == nil || !util.IsSystemContext(ctx) {
				record.Properties[idProp.Name] = structpb.NewStringValue(recordNewId.String())
			} else {
				log.Println("FOUNDxxx")
			}
		}
	}

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
