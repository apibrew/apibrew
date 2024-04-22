// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type Permission struct {
	Id             *uuid.UUID           `json:"id,omitempty"`
	Version        int32                `json:"version,omitempty"`
	AuditData      *PermissionAuditData `json:"auditData,omitempty"`
	Namespace      *string              `json:"namespace,omitempty"`
	Resource       *string              `json:"resource,omitempty"`
	RecordSelector *BooleanExpression   `json:"recordSelector,omitempty"`
	Operation      PermissionOperation  `json:"operation,omitempty"`
	Before         *time.Time           `json:"before,omitempty"`
	After          *time.Time           `json:"after,omitempty"`
	User           *User                `json:"user,omitempty"`
	Role           *Role                `json:"role,omitempty"`
	Permit         PermissionPermit     `json:"permit,omitempty"`
	LocalFlags     interface{}          `json:"localFlags,omitempty"`
}

func (s Permission) GetId() *uuid.UUID {
	return s.Id
}
func (s Permission) GetVersion() int32 {
	return s.Version
}
func (s Permission) GetAuditData() *PermissionAuditData {
	return s.AuditData
}
func (s Permission) GetNamespace() *string {
	return s.Namespace
}
func (s Permission) GetResource() *string {
	return s.Resource
}
func (s Permission) GetRecordSelector() *BooleanExpression {
	return s.RecordSelector
}
func (s Permission) GetOperation() PermissionOperation {
	return s.Operation
}
func (s Permission) GetBefore() *time.Time {
	return s.Before
}
func (s Permission) GetAfter() *time.Time {
	return s.After
}
func (s Permission) GetUser() *User {
	return s.User
}
func (s Permission) GetRole() *Role {
	return s.Role
}
func (s Permission) GetPermit() PermissionPermit {
	return s.Permit
}
func (s Permission) GetLocalFlags() interface{} {
	return s.LocalFlags
}

type PermissionAuditData struct {
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (s PermissionAuditData) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s PermissionAuditData) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s PermissionAuditData) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s PermissionAuditData) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}

type PermissionOperation string

const (
	PermissionOperation_READ   PermissionOperation = "READ"
	PermissionOperation_CREATE PermissionOperation = "CREATE"
	PermissionOperation_UPDATE PermissionOperation = "UPDATE"
	PermissionOperation_DELETE PermissionOperation = "DELETE"
	PermissionOperation_FULL   PermissionOperation = "FULL"
)

type PermissionPermit string

const (
	PermissionPermit_ALLOW  PermissionPermit = "ALLOW"
	PermissionPermit_REJECT PermissionPermit = "REJECT"
)
