package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type SecurityConstraint struct {
	Id            *uuid.UUID                      `json:"id"`
	Version       int32                           `json:"version"`
	CreatedBy     *string                         `json:"createdBy"`
	UpdatedBy     *string                         `json:"updatedBy"`
	CreatedOn     *time.Time                      `json:"createdOn"`
	UpdatedOn     *time.Time                      `json:"updatedOn"`
	Namespace     *Namespace                      `json:"namespace"`
	Resource      *Resource                       `json:"resource"`
	Property      *string                         `json:"property"`
	PropertyValue *string                         `json:"propertyValue"`
	PropertyMode  *SecurityConstraintPropertyMode `json:"propertyMode"`
	Operation     SecurityConstraintOperation     `json:"operation"`
	RecordIds     []string                        `json:"recordIds"`
	Before        *time.Time                      `json:"before"`
	After         *time.Time                      `json:"after"`
	User          *User                           `json:"user"`
	Role          *Role                           `json:"role"`
	Permit        SecurityConstraintPermit        `json:"permit"`
	LocalFlags    *unstructured.Unstructured      `json:"localFlags"`
}

func (s *SecurityConstraint) GetId() *uuid.UUID {
	return s.Id
}
func (s *SecurityConstraint) GetVersion() int32 {
	return s.Version
}
func (s *SecurityConstraint) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *SecurityConstraint) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *SecurityConstraint) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *SecurityConstraint) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *SecurityConstraint) GetNamespace() *Namespace {
	return s.Namespace
}
func (s *SecurityConstraint) GetResource() *Resource {
	return s.Resource
}
func (s *SecurityConstraint) GetProperty() *string {
	return s.Property
}
func (s *SecurityConstraint) GetPropertyValue() *string {
	return s.PropertyValue
}
func (s *SecurityConstraint) GetPropertyMode() *SecurityConstraintPropertyMode {
	return s.PropertyMode
}
func (s *SecurityConstraint) GetOperation() SecurityConstraintOperation {
	return s.Operation
}
func (s *SecurityConstraint) GetRecordIds() []string {
	return s.RecordIds
}
func (s *SecurityConstraint) GetBefore() *time.Time {
	return s.Before
}
func (s *SecurityConstraint) GetAfter() *time.Time {
	return s.After
}
func (s *SecurityConstraint) GetUser() *User {
	return s.User
}
func (s *SecurityConstraint) GetRole() *Role {
	return s.Role
}
func (s *SecurityConstraint) GetPermit() SecurityConstraintPermit {
	return s.Permit
}
func (s *SecurityConstraint) GetLocalFlags() *unstructured.Unstructured {
	return s.LocalFlags
}

type SecurityConstraintPropertyMode string

const (
	SecurityConstraintPropertyMode_PROPERTYMATCHONLY SecurityConstraintPropertyMode = "PROPERTY_MATCH_ONLY"
	SecurityConstraintPropertyMode_PROPERTYMATCHANY  SecurityConstraintPropertyMode = "PROPERTY_MATCH_ANY"
)

type SecurityConstraintOperation string

const (
	SecurityConstraintOperation_READ   SecurityConstraintOperation = "READ"
	SecurityConstraintOperation_CREATE SecurityConstraintOperation = "CREATE"
	SecurityConstraintOperation_UPDATE SecurityConstraintOperation = "UPDATE"
	SecurityConstraintOperation_DELETE SecurityConstraintOperation = "DELETE"
	SecurityConstraintOperation_FULL   SecurityConstraintOperation = "FULL"
)

type SecurityConstraintPermit string

const (
	SecurityConstraintPermit_ALLOW  SecurityConstraintPermit = "ALLOW"
	SecurityConstraintPermit_REJECT SecurityConstraintPermit = "REJECT"
)
