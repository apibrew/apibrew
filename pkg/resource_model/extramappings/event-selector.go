package extramappings

import (
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
)

func EventSelectorToProto(result resource_model.EventSelector) *core.EventSelector {
	var eventSelector = new(core.EventSelector)

	eventSelector.Ids = result.Ids
	eventSelector.Actions = util.ArrayMap(result.Actions, func(t resource_model.EventAction) core.Event_Action {
		return core.Event_Action(model.Event_Action_value[string(t)])
	})

	eventSelector.Annotations = result.Annotations
	eventSelector.Namespaces = result.Namespaces
	eventSelector.Resources = result.Resources
	if result.RecordSelector != nil {
		eventSelector.RecordSelector = BooleanExpressionToProto(*result.RecordSelector)
	}

	return eventSelector
}
