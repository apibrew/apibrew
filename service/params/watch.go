package params

import "github.com/tislib/data-handler/model"

type WatchParams struct {
	Namespace  string
	Resource   string
	Query      *model.BooleanExpression
	BufferSize int
}
