package resources

import (
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PropertyAccessor struct {
	Property *model.ResourceProperty
	Get      func(record *model.Record) interface{}
	Set      func(record *model.Record, val interface{})
}

func GetResourceSpecialProperties(resource *model.Resource) []PropertyAccessor {
	var specialProps []PropertyAccessor

	if !annotations.IsEnabled(resource, annotations.DisableAudit) {
		specialProps = append(specialProps, PropertyAccessor{
			Property: AuditPropertyCreatedBy,
			Get: func(record *model.Record) interface{} {
				return record.AuditData.CreatedBy
			},
			Set: func(record *model.Record, val interface{}) {
				record.AuditData.CreatedBy = val.(string)
			},
		})
		specialProps = append(specialProps, PropertyAccessor{
			Property: AuditPropertyUpdatedBy,
			Get: func(record *model.Record) interface{} {
				return record.AuditData.UpdatedBy
			},
			Set: func(record *model.Record, val interface{}) {
				record.AuditData.UpdatedBy = val.(string)
			},
		})
		specialProps = append(specialProps, PropertyAccessor{
			Property: AuditPropertyCreatedOn,
			Get: func(record *model.Record) interface{} {
				return record.AuditData.CreatedOn.AsTime()
			},
			Set: func(record *model.Record, val interface{}) {
				record.AuditData.CreatedOn = timestamppb.New(val.(time.Time))
			},
		})
		specialProps = append(specialProps, PropertyAccessor{
			Property: AuditPropertyUpdatedOn,
			Get: func(record *model.Record) interface{} {
				return record.AuditData.UpdatedOn.AsTime()
			},
			Set: func(record *model.Record, val interface{}) {
				record.AuditData.UpdatedOn = timestamppb.New(val.(time.Time))
			},
		})
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		specialProps = append(specialProps, PropertyAccessor{
			Property: VersionProperty,
			Get: func(record *model.Record) interface{} {
				return record.Version
			},
			Set: func(record *model.Record, val interface{}) {
				record.Version = uint32(val.(int32))
			},
		})
	}

	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		specialProps = append(specialProps, PropertyAccessor{
			Property: IdProperty,
			Get: func(record *model.Record) interface{} {
				return uuid.MustParse(record.Id)
			},
			Set: func(record *model.Record, val interface{}) {
				record.Id = val.(uuid.UUID).String()
			},
		})
	}

	return specialProps
}
