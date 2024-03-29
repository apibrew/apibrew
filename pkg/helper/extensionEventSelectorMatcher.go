package helper

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
)

type ExtensionEventSelectorMatcher struct {
}

func (b *ExtensionEventSelectorMatcher) SelectorMatches(incoming *model.Event, selector *model.EventSelector) bool {
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
			if action == incoming.Action {
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

func (b *ExtensionEventSelectorMatcher) recordSelectorMatches(incoming *model.Event, selector *model.BooleanExpression) bool {
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

		return proto.Equal(left, right)
	}

	return true
}

func (b *ExtensionEventSelectorMatcher) resolve(incoming *model.Event, left *model.Expression) proto.Message {
	if left.GetProperty() != "" {
		if len(incoming.Records) == 0 {
			return nil
		}
		return incoming.Records[0].Properties[left.GetProperty()]
	}

	if left.GetValue() != nil {
		return left.GetValue()
	}

	return nil
}
