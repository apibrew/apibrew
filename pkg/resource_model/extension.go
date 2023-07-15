package resource_model

import "github.com/google/uuid"
import "time"

type Extension struct {
	Id          *uuid.UUID
	Version     int32
	CreatedBy   *string
	UpdatedBy   *string
	CreatedOn   *time.Time
	UpdatedOn   *time.Time
	Name        string
	Description *string
	Selector    *ExtensionEventSelector
	Order       int32
	Finalizes   bool
	Sync        bool
	Responds    bool
	Call        ExtensionExternalCall
	Annotations map[string]string
}

func (s *Extension) GetId() *uuid.UUID {
	return s.Id
}
func (s *Extension) GetVersion() int32 {
	return s.Version
}
func (s *Extension) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Extension) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Extension) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Extension) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Extension) GetName() string {
	return s.Name
}
func (s *Extension) GetDescription() *string {
	return s.Description
}
func (s *Extension) GetSelector() *ExtensionEventSelector {
	return s.Selector
}
func (s *Extension) GetOrder() int32 {
	return s.Order
}
func (s *Extension) GetFinalizes() bool {
	return s.Finalizes
}
func (s *Extension) GetSync() bool {
	return s.Sync
}
func (s *Extension) GetResponds() bool {
	return s.Responds
}
func (s *Extension) GetCall() ExtensionExternalCall {
	return s.Call
}
func (s *Extension) GetAnnotations() map[string]string {
	return s.Annotations
}

type ExtensionBooleanExpression struct {
}

type ExtensionFunctionCall struct {
	Host         string
	FunctionName string
}

func (s *ExtensionFunctionCall) GetHost() string {
	return s.Host
}
func (s *ExtensionFunctionCall) GetFunctionName() string {
	return s.FunctionName
}

type ExtensionHttpCall struct {
	Uri    string
	Method string
}

func (s *ExtensionHttpCall) GetUri() string {
	return s.Uri
}
func (s *ExtensionHttpCall) GetMethod() string {
	return s.Method
}

type ExtensionExternalCall struct {
	FunctionCall *ExtensionFunctionCall
	HttpCall     *ExtensionHttpCall
}

func (s *ExtensionExternalCall) GetFunctionCall() *ExtensionFunctionCall {
	return s.FunctionCall
}
func (s *ExtensionExternalCall) GetHttpCall() *ExtensionHttpCall {
	return s.HttpCall
}

type ExtensionEventSelector struct {
	Actions        []EventAction
	RecordSelector *ExtensionBooleanExpression
	Namespaces     []string
	Resources      []string
	Ids            []string
	Annotations    map[string]string
}

func (s *ExtensionEventSelector) GetActions() []EventAction {
	return s.Actions
}
func (s *ExtensionEventSelector) GetRecordSelector() *ExtensionBooleanExpression {
	return s.RecordSelector
}
func (s *ExtensionEventSelector) GetNamespaces() []string {
	return s.Namespaces
}
func (s *ExtensionEventSelector) GetResources() []string {
	return s.Resources
}
func (s *ExtensionEventSelector) GetIds() []string {
	return s.Ids
}
func (s *ExtensionEventSelector) GetAnnotations() map[string]string {
	return s.Annotations
}

type ExtensionRecordSearchParams struct {
	Query             *ExtensionBooleanExpression
	Limit             *int32
	Offset            *int32
	ResolveReferences []string
}

func (s *ExtensionRecordSearchParams) GetQuery() *ExtensionBooleanExpression {
	return s.Query
}
func (s *ExtensionRecordSearchParams) GetLimit() *int32 {
	return s.Limit
}
func (s *ExtensionRecordSearchParams) GetOffset() *int32 {
	return s.Offset
}
func (s *ExtensionRecordSearchParams) GetResolveReferences() []string {
	return s.ResolveReferences
}

type ExtensionEvent struct {
	Id                 *uuid.UUID
	Action             EventAction
	RecordSearchParams *ExtensionRecordSearchParams
	ActionSummary      *string
	ActionDescription  *string
	Resource           *Resource
	Records            []*Record
	Ids                []string
	Finalizes          *bool
	Sync               *bool
	Time               *time.Time
	Annotations        map[string]string
}

func (s *ExtensionEvent) GetId() *uuid.UUID {
	return s.Id
}
func (s *ExtensionEvent) GetAction() EventAction {
	return s.Action
}
func (s *ExtensionEvent) GetRecordSearchParams() *ExtensionRecordSearchParams {
	return s.RecordSearchParams
}
func (s *ExtensionEvent) GetActionSummary() *string {
	return s.ActionSummary
}
func (s *ExtensionEvent) GetActionDescription() *string {
	return s.ActionDescription
}
func (s *ExtensionEvent) GetResource() *Resource {
	return s.Resource
}
func (s *ExtensionEvent) GetRecords() []*Record {
	return s.Records
}
func (s *ExtensionEvent) GetIds() []string {
	return s.Ids
}
func (s *ExtensionEvent) GetFinalizes() *bool {
	return s.Finalizes
}
func (s *ExtensionEvent) GetSync() *bool {
	return s.Sync
}
func (s *ExtensionEvent) GetTime() *time.Time {
	return s.Time
}
func (s *ExtensionEvent) GetAnnotations() map[string]string {
	return s.Annotations
}

type EventAction string

const (
	EventAction_CREATE  EventAction = "CREATE"
	EventAction_UPDATE  EventAction = "UPDATE"
	EventAction_DELETE  EventAction = "DELETE"
	EventAction_GET     EventAction = "GET"
	EventAction_LIST    EventAction = "LIST"
	EventAction_OPERATE EventAction = "OPERATE"
)
