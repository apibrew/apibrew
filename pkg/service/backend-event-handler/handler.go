package backend_event_handler

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
)

const NaturalOrder = 100

type HandlerFunc func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError)

type Handler struct {
	Id        string
	Name      string
	Fn        HandlerFunc
	Selector  *model.EventSelector
	Order     int
	Finalizes bool
	Sync      bool
	Responds  bool
	Internal  bool
}

type ByOrder []Handler

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }
