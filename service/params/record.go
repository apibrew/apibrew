package params

import "data-handler/model"

type RecordListParams struct {
	Query             *model.BooleanExpression
	Workspace         string
	Resource          string
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences bool
}

type RecordCreateParams struct {
	Workspace      string
	Resource       string
	Records        []*model.Record
	IgnoreIfExists bool
}

type RecordUpdateParams struct {
	Workspace    string
	Records      []*model.Record
	CheckVersion bool
}

type RecordGetParams struct {
	Workspace string
	Resource  string
	Id        string
}

type RecordDeleteParams struct {
	Workspace string
	Resource  string
	Ids       []string
}
