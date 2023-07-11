package resource_model

import "github.com/google/uuid"
import "time"

type Extension struct {
	Id          *uuid.UUID
	Version     *int32
	CreatedBy   *string
	UpdatedBy   *string
	CreatedOn   *time.Time
	UpdatedOn   *time.Time
	Name        string
	Description *string
	Selector    *ExtensionSelector
	Order       int32
	Finalizes   bool
	Sync        bool
	Responds    bool
	Call        ExtensionCall
	Annotations map[string]string
}

func (s *Extension) GetId() *uuid.UUID {
	return s.Id
}
func (s *Extension) GetVersion() *int32 {
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
func (s *Extension) GetSelector() *ExtensionSelector {
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
func (s *Extension) GetCall() ExtensionCall {
	return s.Call
}
func (s *Extension) GetAnnotations() map[string]string {
	return s.Annotations
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

type ExtensionCall struct {
	FunctionCall *ExtensionFunctionCall
	HttpCall     *ExtensionHttpCall
}

func (s *ExtensionCall) GetFunctionCall() *ExtensionFunctionCall {
	return s.FunctionCall
}
func (s *ExtensionCall) GetHttpCall() *ExtensionHttpCall {
	return s.HttpCall
}

type ExtensionBooleanExpression struct {
}

type ExtensionSelector struct {
	Actions        []ExtensionActions
	RecordSelector *ExtensionBooleanExpression
	Namespaces     []string
	Resources      []string
	Ids            []string
	Annotations    map[string]string
}

func (s *ExtensionSelector) GetActions() []ExtensionActions {
	return s.Actions
}
func (s *ExtensionSelector) GetRecordSelector() *ExtensionBooleanExpression {
	return s.RecordSelector
}
func (s *ExtensionSelector) GetNamespaces() []string {
	return s.Namespaces
}
func (s *ExtensionSelector) GetResources() []string {
	return s.Resources
}
func (s *ExtensionSelector) GetIds() []string {
	return s.Ids
}
func (s *ExtensionSelector) GetAnnotations() map[string]string {
	return s.Annotations
}

type ExtensionActions string

const (
	ExtensionActions_CREATE  ExtensionActions = "CREATE"
	ExtensionActions_UPDATE  ExtensionActions = "UPDATE"
	ExtensionActions_DELETE  ExtensionActions = "DELETE"
	ExtensionActions_GET     ExtensionActions = "GET"
	ExtensionActions_LIST    ExtensionActions = "LIST"
	ExtensionActions_OPERATE ExtensionActions = "OPERATE"
)
