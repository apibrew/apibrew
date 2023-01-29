package params

import (
	"github.com/tislib/data-handler/pkg/model"
)

type RecordListParams struct {
	Query             *model.BooleanExpression
	Namespace         string
	Resource          string
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences bool
}

type RecordCreateParams struct {
	Namespace      string
	Records        []*model.Record
	IgnoreIfExists bool
}

type RecordUpdateParams struct {
	Namespace    string
	Records      []*model.Record
	CheckVersion bool
}

type RecordGetParams struct {
	Namespace string
	Resource  string
	Id        string
}

type RecordDeleteParams struct {
	Namespace string
	Resource  string
	Ids       []string
}
