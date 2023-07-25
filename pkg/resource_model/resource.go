package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Resource struct {
	Id          *uuid.UUID                 `json:"id"`
	Version     int32                      `json:"version"`
	CreatedBy   *string                    `json:"createdBy"`
	UpdatedBy   *string                    `json:"updatedBy"`
	CreatedOn   *time.Time                 `json:"createdOn"`
	UpdatedOn   *time.Time                 `json:"updatedOn"`
	Name        string                     `json:"name"`
	Namespace   *Namespace                 `json:"namespace"`
	Virtual     bool                       `json:"virtual"`
	Types       *unstructured.Unstructured `json:"types"`
	Immutable   bool                       `json:"immutable"`
	Abstract    bool                       `json:"abstract"`
	DataSource  *DataSource                `json:"dataSource"`
	Entity      *string                    `json:"entity"`
	Catalog     *string                    `json:"catalog"`
	Annotations map[string]string          `json:"annotations"`
	Indexes     *unstructured.Unstructured `json:"indexes"`
	Title       *string                    `json:"title"`
	Description *string                    `json:"description"`
}

func (s *Resource) GetId() *uuid.UUID {
	return s.Id
}
func (s *Resource) GetVersion() int32 {
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
func (s *Resource) GetTitle() *string {
	return s.Title
}
func (s *Resource) GetDescription() *string {
	return s.Description
}
