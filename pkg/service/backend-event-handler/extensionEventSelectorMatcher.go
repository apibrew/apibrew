package backend_event_handler

import (
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type ExtensionEventSelectorMatcher struct {
}

func (b *ExtensionEventSelectorMatcher) SelectorMatches(incoming *core.Event, selector *core.EventSelector) bool {
	if incoming.Shallow && !selector.Shallow {
		log.Tracef("Event is shallow, but selector is not")
		return false
	}

	if selector == nil {
		return true
	}

	if len(selector.Resources) > 0 {
		var found = false
		for _, resource := range selector.Resources {
			if resource == incoming.Resource.Name {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if len(selector.Actions) > 0 {
		var found = false
		for _, action := range selector.Actions {
			if int(action) == int(incoming.Action) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if len(selector.Ids) > 0 {
		var found = false
		for _, id := range selector.Ids {
			if id == incoming.Id {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if len(selector.Namespaces) > 0 {
		var found = false
		for _, namespace := range selector.Namespaces {
			if namespace == incoming.Resource.Namespace {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if len(selector.Annotations) > 0 {
		for key, value := range selector.Annotations {
			if incoming.Resource.Annotations[key] == value {
				break
			}

			if value == "*" {
				if _, ok := incoming.Resource.Annotations[key]; ok {
					break
				}
			}

			return false
		}
	}

	if selector.RecordSelector != nil {
		return b.recordSelectorMatches(incoming, selector.RecordSelector)
	}

	return true
}

func (b *ExtensionEventSelectorMatcher) recordSelectorMatches(incoming *core.Event, selector *model.BooleanExpression) bool {
	if selector == nil {
		return true
	}

	if selector.GetAnd() != nil {
		for _, child := range selector.GetAnd().Expressions {
			if !b.recordSelectorMatches(incoming, child) {
				return false
			}
		}

		return true
	}

	if selector.GetOr() != nil {
		for _, child := range selector.GetOr().Expressions {
			if b.recordSelectorMatches(incoming, child) {
				return true
			}
		}

		return false
	}

	if selector.GetNot() != nil {
		return !b.recordSelectorMatches(incoming, selector.GetNot())
	}

	if selector.GetEqual() != nil {
		left := b.resolve(incoming, selector.GetEqual().Left)
		right := b.resolve(incoming, selector.GetEqual().Right)

		if left != nil && right != nil {
			leftO := left.AsInterface()
			rightO := right.AsInterface()

			leftM, ok1 := leftO.(unstructured.Unstructured)
			rightS, ok2 := rightO.(string)

			if ok1 && ok2 {
				return leftM["id"] == rightS
			}

			return leftO == rightO
		}

		return proto.Equal(left, right)
	}

	return true
}

func (b *ExtensionEventSelectorMatcher) resolve(incoming *core.Event, left *model.Expression) *structpb.Value {
	if left.GetProperty() != "" {
		if len(incoming.Records) == 0 {
			return nil
		}
		return incoming.Records[0].GetStructProperty(left.GetProperty())
	}

	if left.GetValue() != nil {
		return left.GetValue()
	}

	return nil
}
