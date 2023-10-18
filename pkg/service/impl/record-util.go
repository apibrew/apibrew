package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func GetResourceSpecialProperties(resource *model.Resource) []util.PropertyAccessor {
	var specialProps []util.PropertyAccessor

	idProp := util.GetResourceSinglePrimaryProp(resource)

	if idProp != nil && idProp.Name == "id" && idProp.Type == model.ResourceProperty_UUID {
		specialProps = append(specialProps, util.PropertyAccessor{
			Property: special.IdProperty,
			Get: func(record *model.Record) interface{} {
				val, err := uuid.Parse(record.Id)

				if err != nil {
					log.Warn(err)
				}

				return val
			},
			Set: func(record *model.Record, val interface{}) {
				record.Id = val.(uuid.UUID).String()
			},
		})
	}

	return specialProps
}

func ComputeRecordIdFromProperties(resource *model.Resource, record *model.Record) {
	var idParts []string
	for _, prop := range resource.Properties {
		val := record.Properties[prop.Name]
		if val != nil && annotations.IsEnabled(prop, annotations.PrimaryProperty) {
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
	recordNewId := uuid.Must(uuid.NewRandom())
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
		ah.SetCreatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.InitVersion()
	}

	if ah.HasIdSpecialProperty() {
		ah.SetId(record.Id)
	}
}

func NormalizeRecord(resource *model.Resource, record *model.Record) {
	if record.Properties == nil {
		record.Properties = make(map[string]*structpb.Value)
	}

	specialProps := GetResourceSpecialProperties(resource)

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

	specialProps := GetResourceSpecialProperties(resource)

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
			ah.SetCreatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
		}
		ah.SetUpdatedOn(timestamppb.New(time.Now()))
		ah.SetUpdatedBy(jwt_model.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.IncreaseVersion()
	}

	if record.Id != "" && ah.HasIdSpecialProperty() {
		ah.SetId(record.Id)
	}
}
