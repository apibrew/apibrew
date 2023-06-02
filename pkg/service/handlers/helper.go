package handlers

import (
	"github.com/apibrew/apibrew/pkg/model"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
)

func prepareStdHandler(order int, action model.Event_Action, handlerFunc backend_event_handler.HandlerFunc, resource *model.Resource) backend_event_handler.Handler {
	handlerId := "std-handler-" + resource.Namespace + "-" + resource.Name + "-" + util.RandomHex(6)
	return backend_event_handler.Handler{
		Id:   handlerId,
		Name: handlerId,
		Fn:   handlerFunc,
		Selector: &model.EventSelector{
			Actions:    []model.Event_Action{action},
			Namespaces: []string{resource.Namespace},
			Resources:  []string{resource.Name},
		},
		Order:     order,
		Finalizes: false,
		Sync:      true,
		Responds:  true,
	}
}
