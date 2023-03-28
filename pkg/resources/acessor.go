package resources

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
)

type PropertyAccessor struct {
	Property *model.ResourceProperty
	Get      func(record *model.Record) interface{}
	Set      func(record *model.Record, val interface{})
}

func GetResourceSpecialProperties(resource *model.Resource) []PropertyAccessor {
	var specialProps []PropertyAccessor

	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		specialProps = append(specialProps, PropertyAccessor{
			Property: IdProperty,
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
