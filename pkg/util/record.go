package util

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type PropertyAccessor struct {
	Property *model.ResourceProperty
	Get      func(record *model.Record) interface{}
	Set      func(record *model.Record, val interface{})
}

func GetResourceSpecialProperties(resource *model.Resource) []PropertyAccessor {
	var specialProps []PropertyAccessor

	idProp := GetResourceSinglePrimaryProp(resource)

	if idProp != nil && idProp.Name == "id" && idProp.Type == model.ResourceProperty_UUID {
		specialProps = append(specialProps, PropertyAccessor{
			Property: resources.IdProperty,
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
			ah.SetCreatedBy(security.GetUserPrincipalFromContext(ctx))
		}
		ah.SetUpdatedOn(timestamppb.New(time.Now()))
		ah.SetUpdatedBy(security.GetUserPrincipalFromContext(ctx))
	}

	if ah.IsVersionEnabled() {
		ah.IncreaseVersion()
	}

	if record.Id != "" && ah.HasIdSpecialProperty() {
		ah.SetId(record.Id)
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

func RecordIdentifierProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, error) {
	if props, ok := RecordIdentifierPrimaryProperties(resource, properties); ok {
		return props, nil
	}

	if props, ok := RecordIdentifierUniqueProperties(resource, properties); ok {
		return props, nil
	}

	return nil, fmt.Errorf("could not find identifiable properties of %s", resource.Name)
}

func RecordIdentifierPrimaryProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, bool) {
	identifierProps := make(map[string]*structpb.Value)

	for _, prop := range resource.Properties {
		if !prop.Primary {
			continue
		}

		val, ok := properties[prop.Name]

		if !ok {
			return nil, false
		}

		typ := types.ByResourcePropertyType(prop.Type)

		unpacked, err := typ.UnPack(val)

		if err != nil {
			log.Error(err)
			return nil, false
		}

		if typ.Equals(unpacked, typ.Default()) {
			return nil, false
		}

		identifierProps[prop.Name] = val
	}

	return identifierProps, true
}

func RecordIdentifierUniqueProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, bool) {
	for _, prop := range resource.Properties {
		identifierProps := make(map[string]*structpb.Value)
		if !prop.Unique {
			continue
		}

		val, ok := properties[prop.Name]

		if !ok {
			continue
		}

		typ := types.ByResourcePropertyType(prop.Type)

		unpacked, err := typ.UnPack(val)

		if err != nil {
			log.Error(err)
			return nil, false
		}

		if typ.Equals(unpacked, typ.Default()) {
			continue
		}

		identifierProps[prop.Name] = val

		return identifierProps, true
	}

	propMap := GetNamedMap(resource.Properties)

	for _, index := range resource.Indexes {
		if index.Unique {
			var valid = true

			identifierProps := make(map[string]*structpb.Value)

			for _, indexProp := range index.Properties {
				prop := propMap[indexProp.Name]

				val, ok := properties[prop.Name]

				if !ok {
					valid = false
					break
				}

				typ := types.ByResourcePropertyType(prop.Type)

				unpacked, err := typ.UnPack(val)

				if err != nil {
					log.Error(err)

					valid = false
					break
				}

				if typ.Equals(unpacked, typ.Default()) {
					valid = false
					break
				}

				identifierProps[prop.Name] = val
			}

			if valid {
				return identifierProps, true
			}
		}
	}

	return nil, false
}
