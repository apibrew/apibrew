package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type SecurityConstraint struct {
	Id            *uuid.UUID
	Version       *int32
	CreatedBy     *string
	UpdatedBy     *string
	CreatedOn     *time.Time
	UpdatedOn     *time.Time
	Namespace     string
	Resource      string
	Property      string
	PropertyValue *string
	PropertyMode  *SecurityConstraintPropertyMode
	Operation     SecurityConstraintOperation
	RecordIds     []string
	Before        *time.Time
	After         *time.Time
	Username      *string
	Role          *string
	Permit        SecurityConstraintPermit
	LocalFlags    *unstructured.Unstructured
}

func (s *SecurityConstraint) GetId() *uuid.UUID {
	return s.Id
}
func (s *SecurityConstraint) GetVersion() *int32 {
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
func (s *SecurityConstraint) GetNamespace() string {
	return s.Namespace
}
func (s *SecurityConstraint) GetResource() string {
	return s.Resource
}
func (s *SecurityConstraint) GetProperty() string {
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
func (s *SecurityConstraint) GetUsername() *string {
	return s.Username
}
func (s *SecurityConstraint) GetRole() *string {
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
