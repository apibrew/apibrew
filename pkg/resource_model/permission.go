// AUTOGENERATED FILE

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Permission struct {
	Id            *uuid.UUID                 `json:"id,omitempty"`
	Version       int32                      `json:"version,omitempty"`
	CreatedBy     *string                    `json:"createdBy,omitempty"`
	UpdatedBy     *string                    `json:"updatedBy,omitempty"`
	CreatedOn     *time.Time                 `json:"createdOn,omitempty"`
	UpdatedOn     *time.Time                 `json:"updatedOn,omitempty"`
	Namespace     *string                    `json:"namespace,omitempty"`
	Resource      *string                    `json:"resource,omitempty"`
	Property      *string                    `json:"property,omitempty"`
	PropertyValue *string                    `json:"propertyValue,omitempty"`
	PropertyMode  *PermissionPropertyMode    `json:"propertyMode,omitempty"`
	Operation     PermissionOperation        `json:"operation,omitempty"`
	RecordIds     []string                   `json:"recordIds,omitempty"`
	Before        *time.Time                 `json:"before,omitempty"`
	After         *time.Time                 `json:"after,omitempty"`
	User          *User                      `json:"user,omitempty"`
	Role          *Role                      `json:"role,omitempty"`
	Permit        PermissionPermit           `json:"permit,omitempty"`
	LocalFlags    *unstructured.Unstructured `json:"localFlags,omitempty"`
}

func (s *Permission) GetId() *uuid.UUID {
	return s.Id
}
func (s *Permission) GetVersion() int32 {
	return s.Version
}
func (s *Permission) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Permission) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Permission) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Permission) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Permission) GetNamespace() *string {
	return s.Namespace
}
func (s *Permission) GetResource() *string {
	return s.Resource
}
func (s *Permission) GetProperty() *string {
	return s.Property
}
func (s *Permission) GetPropertyValue() *string {
	return s.PropertyValue
}
func (s *Permission) GetPropertyMode() *PermissionPropertyMode {
	return s.PropertyMode
}
func (s *Permission) GetOperation() PermissionOperation {
	return s.Operation
}
func (s *Permission) GetRecordIds() []string {
	return s.RecordIds
}
func (s *Permission) GetBefore() *time.Time {
	return s.Before
}
func (s *Permission) GetAfter() *time.Time {
	return s.After
}
func (s *Permission) GetUser() *User {
	return s.User
}
func (s *Permission) GetRole() *Role {
	return s.Role
}
func (s *Permission) GetPermit() PermissionPermit {
	return s.Permit
}
func (s *Permission) GetLocalFlags() *unstructured.Unstructured {
	return s.LocalFlags
}

type PermissionPropertyMode string

const (
	PermissionPropertyMode_PROPERTYMATCHONLY PermissionPropertyMode = "PROPERTY_MATCH_ONLY"
	PermissionPropertyMode_PROPERTYMATCHANY  PermissionPropertyMode = "PROPERTY_MATCH_ANY"
)

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
