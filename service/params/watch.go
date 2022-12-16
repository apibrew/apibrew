package params

import "data-handler/model"

type WatchParams struct {
	Workspace  string
	Resource   string
	Query      *model.BooleanExpression
	BufferSize int
}
