package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

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
	Types               *unstructured.Unstructured
	Immutable           bool
	Abstract            bool
	DataSource          *DataSource
	Entity              *string
	Catalog             *string
	Annotations         map[string]string
	Indexes             *unstructured.Unstructured
	SecurityConstraints []*SecurityConstraint
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
func (s *Resource) GetTypes() *unstructured.Unstructured {
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
func (s *Resource) GetIndexes() *unstructured.Unstructured {
	return s.Indexes
}
func (s *Resource) GetSecurityConstraints() []*SecurityConstraint {
	return s.SecurityConstraints
}
func (s *Resource) GetTitle() *string {
	return s.Title
}
func (s *Resource) GetDescription() *string {
	return s.Description
}
