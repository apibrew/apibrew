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
	Properties  []ResourceProperty         `json:"properties"`
	Types       []ResourceSubType          `json:"types"`
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
func (s *Resource) GetProperties() []ResourceProperty {
	return s.Properties
}
func (s *Resource) GetTypes() []ResourceSubType {
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

type ResourceProperty struct {
	Name                  string                     `json:"name"`
	Type                  int32                      `json:"type"`
	TypeRef               *string                    `json:"typeRef"`
	Mapping               string                     `json:"mapping"`
	Primary               bool                       `json:"primary"`
	Required              bool                       `json:"required"`
	Unique                bool                       `json:"unique"`
	Immutable             bool                       `json:"immutable"`
	Length                int32                      `json:"length"`
	Resource              *Resource                  `json:"resource"`
	Item                  *ResourceProperty          `json:"item"`
	Properties            []ResourceProperty         `json:"properties"`
	ReferenceResource     *Resource                  `json:"referenceResource"`
	ReferenceCascade      *bool                      `json:"referenceCascade"`
	BackReferenceProperty *bool                      `json:"backReferenceProperty"`
	DefaultValue          *unstructured.Unstructured `json:"defaultValue"`
	EnumValues            []string                   `json:"enumValues"`
	ExampleValue          *unstructured.Unstructured `json:"exampleValue"`
	Title                 *string                    `json:"title"`
	Description           *string                    `json:"description"`
	Annotations           map[string]string          `json:"annotations"`
}

func (s *ResourceProperty) GetName() string {
	return s.Name
}
func (s *ResourceProperty) GetType() int32 {
	return s.Type
}
func (s *ResourceProperty) GetTypeRef() *string {
	return s.TypeRef
}
func (s *ResourceProperty) GetMapping() string {
	return s.Mapping
}
func (s *ResourceProperty) GetPrimary() bool {
	return s.Primary
}
func (s *ResourceProperty) GetRequired() bool {
	return s.Required
}
func (s *ResourceProperty) GetUnique() bool {
	return s.Unique
}
func (s *ResourceProperty) GetImmutable() bool {
	return s.Immutable
}
func (s *ResourceProperty) GetLength() int32 {
	return s.Length
}
func (s *ResourceProperty) GetResource() *Resource {
	return s.Resource
}
func (s *ResourceProperty) GetItem() *ResourceProperty {
	return s.Item
}
func (s *ResourceProperty) GetProperties() []ResourceProperty {
	return s.Properties
}
func (s *ResourceProperty) GetReferenceResource() *Resource {
	return s.ReferenceResource
}
func (s *ResourceProperty) GetReferenceCascade() *bool {
	return s.ReferenceCascade
}
func (s *ResourceProperty) GetBackReferenceProperty() *bool {
	return s.BackReferenceProperty
}
func (s *ResourceProperty) GetDefaultValue() *unstructured.Unstructured {
	return s.DefaultValue
}
func (s *ResourceProperty) GetEnumValues() []string {
	return s.EnumValues
}
func (s *ResourceProperty) GetExampleValue() *unstructured.Unstructured {
	return s.ExampleValue
}
func (s *ResourceProperty) GetTitle() *string {
	return s.Title
}
func (s *ResourceProperty) GetDescription() *string {
	return s.Description
}
func (s *ResourceProperty) GetAnnotations() map[string]string {
	return s.Annotations
}

type ResourceSubType struct {
	Name       string             `json:"name"`
	Properties []ResourceProperty `json:"properties"`
}

func (s *ResourceSubType) GetName() string {
	return s.Name
}
func (s *ResourceSubType) GetProperties() []ResourceProperty {
	return s.Properties
}
