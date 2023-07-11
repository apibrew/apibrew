package resource_model

import "github.com/google/uuid"
import "time"

type Namespace struct {
	Id                  *uuid.UUID
	Version             *int32
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Name                string
	Description         *string
	Details             *interface{}
	SecurityConstraints []NamespaceSecurityConstraint
}

func (s *Namespace) GetId() *uuid.UUID {
	return s.Id
}
func (s *Namespace) GetVersion() *int32 {
	return s.Version
}
func (s *Namespace) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Namespace) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Namespace) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Namespace) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Namespace) GetName() string {
	return s.Name
}
func (s *Namespace) GetDescription() *string {
	return s.Description
}
func (s *Namespace) GetDetails() *interface{} {
	return s.Details
}
func (s *Namespace) GetSecurityConstraints() []NamespaceSecurityConstraint {
	return s.SecurityConstraints
}

type NamespaceSecurityConstraint struct {
	Namespace     string
	Resource      string
	Property      string
	PropertyValue *string
	PropertyMode  *NamespacePropertyMode
	Operation     *NamespaceOperation
	RecordIds     []string
	Before        *time.Time
	After         *time.Time
	Username      *string
	Role          *string
	Permit        NamespacePermit
	LocalFlags    *interface{}
}

func (s *NamespaceSecurityConstraint) GetNamespace() string {
	return s.Namespace
}
func (s *NamespaceSecurityConstraint) GetResource() string {
	return s.Resource
}
func (s *NamespaceSecurityConstraint) GetProperty() string {
	return s.Property
}
func (s *NamespaceSecurityConstraint) GetPropertyValue() *string {
	return s.PropertyValue
}
func (s *NamespaceSecurityConstraint) GetPropertyMode() *NamespacePropertyMode {
	return s.PropertyMode
}
func (s *NamespaceSecurityConstraint) GetOperation() *NamespaceOperation {
	return s.Operation
}
func (s *NamespaceSecurityConstraint) GetRecordIds() []string {
	return s.RecordIds
}
func (s *NamespaceSecurityConstraint) GetBefore() *time.Time {
	return s.Before
}
func (s *NamespaceSecurityConstraint) GetAfter() *time.Time {
	return s.After
}
func (s *NamespaceSecurityConstraint) GetUsername() *string {
	return s.Username
}
func (s *NamespaceSecurityConstraint) GetRole() *string {
	return s.Role
}
func (s *NamespaceSecurityConstraint) GetPermit() NamespacePermit {
	return s.Permit
}
func (s *NamespaceSecurityConstraint) GetLocalFlags() *interface{} {
	return s.LocalFlags
}

type NamespacePropertyMode string

const (
	NamespacePropertyMode_PROPERTYMATCHONLY NamespacePropertyMode = "PROPERTY_MATCH_ONLY"
	NamespacePropertyMode_PROPERTYMATCHANY  NamespacePropertyMode = "PROPERTY_MATCH_ANY"
)

type NamespaceOperation string

const (
	NamespaceOperation_OPERATIONTYPEREAD   NamespaceOperation = "OPERATION_TYPE_READ"
	NamespaceOperation_OPERATIONTYPECREATE NamespaceOperation = "OPERATION_TYPE_CREATE"
	NamespaceOperation_OPERATIONTYPEUPDATE NamespaceOperation = "OPERATION_TYPE_UPDATE"
	NamespaceOperation_OPERATIONTYPEDELETE NamespaceOperation = "OPERATION_TYPE_DELETE"
	NamespaceOperation_FULL                NamespaceOperation = "FULL"
)

type NamespacePermit string

const (
	NamespacePermit_PERMITTYPEALLOW  NamespacePermit = "PERMIT_TYPE_ALLOW"
	NamespacePermit_PERMITTYPEREJECT NamespacePermit = "PERMIT_TYPE_REJECT"
)
