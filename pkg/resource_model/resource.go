package resource_model

import "github.com/google/uuid"
import "time"

type Resource struct {
	Id                  *uuid.UUID
	Version             *int32
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Name                string
	Namespace           *Namespace
	Virtual             bool
	Types               *interface{}
	Immutable           bool
	Abstract            bool
	DataSource          *DataSource
	Entity              *string
	Catalog             *string
	Annotations         map[string]string
	Indexes             *interface{}
	SecurityConstraints []ResourceSecurityConstraint
	Title               *string
	Description         *string
}

func (s *Resource) GetId() *uuid.UUID {
	return s.Id
}
func (s *Resource) GetVersion() *int32 {
	return s.Version
}
func (s *Resource) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Resource) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Resource) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Resource) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Resource) GetName() string {
	return s.Name
}
func (s *Resource) GetNamespace() *Namespace {
	return s.Namespace
}
func (s *Resource) GetVirtual() bool {
	return s.Virtual
}
func (s *Resource) GetTypes() *interface{} {
	return s.Types
}
func (s *Resource) GetImmutable() bool {
	return s.Immutable
}
func (s *Resource) GetAbstract() bool {
	return s.Abstract
}
func (s *Resource) GetDataSource() *DataSource {
	return s.DataSource
}
func (s *Resource) GetEntity() *string {
	return s.Entity
}
func (s *Resource) GetCatalog() *string {
	return s.Catalog
}
func (s *Resource) GetAnnotations() map[string]string {
	return s.Annotations
}
func (s *Resource) GetIndexes() *interface{} {
	return s.Indexes
}
func (s *Resource) GetSecurityConstraints() []ResourceSecurityConstraint {
	return s.SecurityConstraints
}
func (s *Resource) GetTitle() *string {
	return s.Title
}
func (s *Resource) GetDescription() *string {
	return s.Description
}

type ResourceSecurityConstraint struct {
	Namespace     string
	Resource      string
	Property      string
	PropertyValue *string
	PropertyMode  *ResourcePropertyMode
	Operation     *ResourceOperation
	RecordIds     []string
	Before        *time.Time
	After         *time.Time
	Username      *string
	Role          *string
	Permit        ResourcePermit
	LocalFlags    *interface{}
}

func (s *ResourceSecurityConstraint) GetNamespace() string {
	return s.Namespace
}
func (s *ResourceSecurityConstraint) GetResource() string {
	return s.Resource
}
func (s *ResourceSecurityConstraint) GetProperty() string {
	return s.Property
}
func (s *ResourceSecurityConstraint) GetPropertyValue() *string {
	return s.PropertyValue
}
func (s *ResourceSecurityConstraint) GetPropertyMode() *ResourcePropertyMode {
	return s.PropertyMode
}
func (s *ResourceSecurityConstraint) GetOperation() *ResourceOperation {
	return s.Operation
}
func (s *ResourceSecurityConstraint) GetRecordIds() []string {
	return s.RecordIds
}
func (s *ResourceSecurityConstraint) GetBefore() *time.Time {
	return s.Before
}
func (s *ResourceSecurityConstraint) GetAfter() *time.Time {
	return s.After
}
func (s *ResourceSecurityConstraint) GetUsername() *string {
	return s.Username
}
func (s *ResourceSecurityConstraint) GetRole() *string {
	return s.Role
}
func (s *ResourceSecurityConstraint) GetPermit() ResourcePermit {
	return s.Permit
}
func (s *ResourceSecurityConstraint) GetLocalFlags() *interface{} {
	return s.LocalFlags
}

type ResourcePropertyMode string

const (
	ResourcePropertyMode_PROPERTYMATCHONLY ResourcePropertyMode = "PROPERTY_MATCH_ONLY"
	ResourcePropertyMode_PROPERTYMATCHANY  ResourcePropertyMode = "PROPERTY_MATCH_ANY"
)

type ResourceOperation string

const (
	ResourceOperation_OPERATIONTYPEREAD   ResourceOperation = "OPERATION_TYPE_READ"
	ResourceOperation_OPERATIONTYPECREATE ResourceOperation = "OPERATION_TYPE_CREATE"
	ResourceOperation_OPERATIONTYPEUPDATE ResourceOperation = "OPERATION_TYPE_UPDATE"
	ResourceOperation_OPERATIONTYPEDELETE ResourceOperation = "OPERATION_TYPE_DELETE"
	ResourceOperation_FULL                ResourceOperation = "FULL"
)

type ResourcePermit string

const (
	ResourcePermit_PERMITTYPEALLOW  ResourcePermit = "PERMIT_TYPE_ALLOW"
	ResourcePermit_PERMITTYPEREJECT ResourcePermit = "PERMIT_TYPE_REJECT"
)
