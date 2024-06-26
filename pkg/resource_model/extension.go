// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type Extension struct {
	Id          *uuid.UUID          `json:"id,omitempty"`
	Version     int32               `json:"version,omitempty"`
	AuditData   *ExtensionAuditData `json:"auditData,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description *string             `json:"description,omitempty"`
	Selector    *EventSelector      `json:"selector,omitempty"`
	Order       int32               `json:"order,omitempty"`
	Finalizes   bool                `json:"finalizes,omitempty"`
	Sync        bool                `json:"sync,omitempty"`
	Responds    bool                `json:"responds,omitempty"`
	Call        ExternalCall        `json:"call,omitempty"`
	Annotations map[string]string   `json:"annotations,omitempty"`
}

func (s Extension) GetId() *uuid.UUID {
	return s.Id
}
func (s Extension) GetVersion() int32 {
	return s.Version
}
func (s Extension) GetAuditData() *ExtensionAuditData {
	return s.AuditData
}
func (s Extension) GetName() string {
	return s.Name
}
func (s Extension) GetDescription() *string {
	return s.Description
}
func (s Extension) GetSelector() *EventSelector {
	return s.Selector
}
func (s Extension) GetOrder() int32 {
	return s.Order
}
func (s Extension) GetFinalizes() bool {
	return s.Finalizes
}
func (s Extension) GetSync() bool {
	return s.Sync
}
func (s Extension) GetResponds() bool {
	return s.Responds
}
func (s Extension) GetCall() ExternalCall {
	return s.Call
}
func (s Extension) GetAnnotations() map[string]string {
	return s.Annotations
}

type BooleanExpression struct {
	And                []BooleanExpression    `json:"and,omitempty"`
	Or                 []BooleanExpression    `json:"or,omitempty"`
	Not                *BooleanExpression     `json:"not,omitempty"`
	Equal              *PairExpression        `json:"equal,omitempty"`
	LessThan           *PairExpression        `json:"lessThan,omitempty"`
	GreaterThan        *PairExpression        `json:"greaterThan,omitempty"`
	LessThanOrEqual    *PairExpression        `json:"lessThanOrEqual,omitempty"`
	GreaterThanOrEqual *PairExpression        `json:"greaterThanOrEqual,omitempty"`
	In                 *PairExpression        `json:"in,omitempty"`
	Like               *PairExpression        `json:"like,omitempty"`
	Ilike              *PairExpression        `json:"ilike,omitempty"`
	Regex              *PairExpression        `json:"regex,omitempty"`
	IsNull             *Expression            `json:"isNull,omitempty"`
	Filters            map[string]interface{} `json:"filters,omitempty"`
}

func (s BooleanExpression) GetAnd() []BooleanExpression {
	return s.And
}
func (s BooleanExpression) GetOr() []BooleanExpression {
	return s.Or
}
func (s BooleanExpression) GetNot() *BooleanExpression {
	return s.Not
}
func (s BooleanExpression) GetEqual() *PairExpression {
	return s.Equal
}
func (s BooleanExpression) GetLessThan() *PairExpression {
	return s.LessThan
}
func (s BooleanExpression) GetGreaterThan() *PairExpression {
	return s.GreaterThan
}
func (s BooleanExpression) GetLessThanOrEqual() *PairExpression {
	return s.LessThanOrEqual
}
func (s BooleanExpression) GetGreaterThanOrEqual() *PairExpression {
	return s.GreaterThanOrEqual
}
func (s BooleanExpression) GetIn() *PairExpression {
	return s.In
}
func (s BooleanExpression) GetLike() *PairExpression {
	return s.Like
}
func (s BooleanExpression) GetIlike() *PairExpression {
	return s.Ilike
}
func (s BooleanExpression) GetRegex() *PairExpression {
	return s.Regex
}
func (s BooleanExpression) GetIsNull() *Expression {
	return s.IsNull
}
func (s BooleanExpression) GetFilters() map[string]interface{} {
	return s.Filters
}

type PairExpression struct {
	Left  *Expression `json:"left,omitempty"`
	Right *Expression `json:"right,omitempty"`
}

func (s PairExpression) GetLeft() *Expression {
	return s.Left
}
func (s PairExpression) GetRight() *Expression {
	return s.Right
}

type Expression struct {
	Property *string     `json:"property,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

func (s Expression) GetProperty() *string {
	return s.Property
}
func (s Expression) GetValue() interface{} {
	return s.Value
}

type ExtensionAuditData struct {
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (s ExtensionAuditData) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s ExtensionAuditData) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s ExtensionAuditData) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s ExtensionAuditData) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}

type FunctionCall struct {
	Host         string `json:"host,omitempty"`
	FunctionName string `json:"functionName,omitempty"`
}

func (s FunctionCall) GetHost() string {
	return s.Host
}
func (s FunctionCall) GetFunctionName() string {
	return s.FunctionName
}

type HttpCall struct {
	Uri    string `json:"uri,omitempty"`
	Method string `json:"method,omitempty"`
}

func (s HttpCall) GetUri() string {
	return s.Uri
}
func (s HttpCall) GetMethod() string {
	return s.Method
}

type ChannelCall struct {
	ChannelKey string `json:"channelKey,omitempty"`
}

func (s ChannelCall) GetChannelKey() string {
	return s.ChannelKey
}

type ExternalCall struct {
	FunctionCall *FunctionCall `json:"functionCall,omitempty"`
	HttpCall     *HttpCall     `json:"httpCall,omitempty"`
	ChannelCall  *ChannelCall  `json:"channelCall,omitempty"`
}

func (s ExternalCall) GetFunctionCall() *FunctionCall {
	return s.FunctionCall
}
func (s ExternalCall) GetHttpCall() *HttpCall {
	return s.HttpCall
}
func (s ExternalCall) GetChannelCall() *ChannelCall {
	return s.ChannelCall
}

type EventSelector struct {
	Actions        []EventAction      `json:"actions,omitempty"`
	RecordSelector *BooleanExpression `json:"recordSelector,omitempty"`
	Namespaces     []string           `json:"namespaces,omitempty"`
	Resources      []string           `json:"resources,omitempty"`
	Ids            []string           `json:"ids,omitempty"`
	Annotations    map[string]string  `json:"annotations,omitempty"`
}

func (s EventSelector) GetActions() []EventAction {
	return s.Actions
}
func (s EventSelector) GetRecordSelector() *BooleanExpression {
	return s.RecordSelector
}
func (s EventSelector) GetNamespaces() []string {
	return s.Namespaces
}
func (s EventSelector) GetResources() []string {
	return s.Resources
}
func (s EventSelector) GetIds() []string {
	return s.Ids
}
func (s EventSelector) GetAnnotations() map[string]string {
	return s.Annotations
}

type RecordSearchParams struct {
	Query             *BooleanExpression `json:"query,omitempty"`
	Limit             *int32             `json:"limit,omitempty"`
	Offset            *int32             `json:"offset,omitempty"`
	ResolveReferences []string           `json:"resolveReferences,omitempty"`
}

func (s RecordSearchParams) GetQuery() *BooleanExpression {
	return s.Query
}
func (s RecordSearchParams) GetLimit() *int32 {
	return s.Limit
}
func (s RecordSearchParams) GetOffset() *int32 {
	return s.Offset
}
func (s RecordSearchParams) GetResolveReferences() []string {
	return s.ResolveReferences
}

type Event struct {
	Id                 string              `json:"id,omitempty"`
	Action             ExtensionAction     `json:"action,omitempty"`
	RecordSearchParams *RecordSearchParams `json:"recordSearchParams,omitempty"`
	Resource           *Resource           `json:"resource,omitempty"`
	Records            []*Record           `json:"records,omitempty"`
	Finalizes          *bool               `json:"finalizes,omitempty"`
	Sync               *bool               `json:"sync,omitempty"`
	Time               *time.Time          `json:"time,omitempty"`
	Total              *int64              `json:"total,omitempty"`
	Annotations        map[string]string   `json:"annotations,omitempty"`
	Error              *Error              `json:"error,omitempty"`
}

func (s Event) GetId() string {
	return s.Id
}
func (s Event) GetAction() ExtensionAction {
	return s.Action
}
func (s Event) GetRecordSearchParams() *RecordSearchParams {
	return s.RecordSearchParams
}
func (s Event) GetResource() *Resource {
	return s.Resource
}
func (s Event) GetRecords() []*Record {
	return s.Records
}
func (s Event) GetFinalizes() *bool {
	return s.Finalizes
}
func (s Event) GetSync() *bool {
	return s.Sync
}
func (s Event) GetTime() *time.Time {
	return s.Time
}
func (s Event) GetTotal() *int64 {
	return s.Total
}
func (s Event) GetAnnotations() map[string]string {
	return s.Annotations
}
func (s Event) GetError() *Error {
	return s.Error
}

type ErrorField struct {
	RecordId *string     `json:"recordId,omitempty"`
	Property *string     `json:"property,omitempty"`
	Message  *string     `json:"message,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

func (s ErrorField) GetRecordId() *string {
	return s.RecordId
}
func (s ErrorField) GetProperty() *string {
	return s.Property
}
func (s ErrorField) GetMessage() *string {
	return s.Message
}
func (s ErrorField) GetValue() interface{} {
	return s.Value
}

type Error struct {
	Code    *ExtensionCode `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
	Fields  []ErrorField   `json:"fields,omitempty"`
}

func (s Error) GetCode() *ExtensionCode {
	return s.Code
}
func (s Error) GetMessage() *string {
	return s.Message
}
func (s Error) GetFields() []ErrorField {
	return s.Fields
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

type ExtensionAction string

const (
	ExtensionAction_CREATE  ExtensionAction = "CREATE"
	ExtensionAction_UPDATE  ExtensionAction = "UPDATE"
	ExtensionAction_DELETE  ExtensionAction = "DELETE"
	ExtensionAction_GET     ExtensionAction = "GET"
	ExtensionAction_LIST    ExtensionAction = "LIST"
	ExtensionAction_OPERATE ExtensionAction = "OPERATE"
)

type ExtensionCode string

const (
	ExtensionCode_UNKNOWNERROR                      ExtensionCode = "UNKNOWN_ERROR"
	ExtensionCode_RECORDNOTFOUND                    ExtensionCode = "RECORD_NOT_FOUND"
	ExtensionCode_UNABLETOLOCATEPRIMARYKEY          ExtensionCode = "UNABLE_TO_LOCATE_PRIMARY_KEY"
	ExtensionCode_INTERNALERROR                     ExtensionCode = "INTERNAL_ERROR"
	ExtensionCode_PROPERTYNOTFOUND                  ExtensionCode = "PROPERTY_NOT_FOUND"
	ExtensionCode_RECORDVALIDATIONERROR             ExtensionCode = "RECORD_VALIDATION_ERROR"
	ExtensionCode_RESOURCEVALIDATIONERROR           ExtensionCode = "RESOURCE_VALIDATION_ERROR"
	ExtensionCode_AUTHENTICATIONFAILED              ExtensionCode = "AUTHENTICATION_FAILED"
	ExtensionCode_ALREADYEXISTS                     ExtensionCode = "ALREADY_EXISTS"
	ExtensionCode_ACCESSDENIED                      ExtensionCode = "ACCESS_DENIED"
	ExtensionCode_BACKENDERROR                      ExtensionCode = "BACKEND_ERROR"
	ExtensionCode_UNIQUEVIOLATION                   ExtensionCode = "UNIQUE_VIOLATION"
	ExtensionCode_REFERENCEVIOLATION                ExtensionCode = "REFERENCE_VIOLATION"
	ExtensionCode_RESOURCENOTFOUND                  ExtensionCode = "RESOURCE_NOT_FOUND"
	ExtensionCode_UNSUPPORTEDOPERATION              ExtensionCode = "UNSUPPORTED_OPERATION"
	ExtensionCode_EXTERNALBACKENDCOMMUNICATIONERROR ExtensionCode = "EXTERNAL_BACKEND_COMMUNICATION_ERROR"
	ExtensionCode_EXTERNALBACKENDERROR              ExtensionCode = "EXTERNAL_BACKEND_ERROR"
	ExtensionCode_RATELIMITERROR                    ExtensionCode = "RATE_LIMIT_ERROR"
)
