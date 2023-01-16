package params

import "data-handler/model"

type WatchParams struct {
	Namespace  string
	Resource   string
	Query      *model.BooleanExpression
	BufferSize int
}
