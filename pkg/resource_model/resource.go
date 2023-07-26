package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Resource struct {
	Id          *uuid.UUID         `json:"id,omitempty"`
	Version     int32              `json:"version,omitempty"`
	CreatedBy   *string            `json:"createdBy,omitempty"`
	UpdatedBy   *string            `json:"updatedBy,omitempty"`
	CreatedOn   *time.Time         `json:"createdOn,omitempty"`
	UpdatedOn   *time.Time         `json:"updatedOn,omitempty"`
	Name        string             `json:"name,omitempty"`
	Namespace   *Namespace         `json:"namespace,omitempty"`
	Virtual     bool               `json:"virtual,omitempty"`
	Properties  []ResourceProperty `json:"properties,omitempty"`
	Indexes     []ResourceIndex    `json:"indexes,omitempty"`
	Types       []ResourceSubType  `json:"types,omitempty"`
	Immutable   bool               `json:"immutable,omitempty"`
	Abstract    bool               `json:"abstract,omitempty"`
	DataSource  *DataSource        `json:"dataSource,omitempty"`
	Entity      *string            `json:"entity,omitempty"`
	Catalog     *string            `json:"catalog,omitempty"`
	Title       *string            `json:"title,omitempty"`
	Description *string            `json:"description,omitempty"`
	Annotations map[string]string  `json:"annotations,omitempty"`
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
func (s *Resource) GetIndexes() []ResourceIndex {
	return s.Indexes
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
func (s *Resource) GetTitle() *string {
	return s.Title
}
func (s *Resource) GetDescription() *string {
	return s.Description
}
func (s *Resource) GetAnnotations() map[string]string {
	return s.Annotations
}

type ResourceProperty struct {
	Name         string                     `json:"name,omitempty"`
	Type         int32                      `json:"type,omitempty"`
	TypeRef      *string                    `json:"typeRef,omitempty"`
	Mapping      string                     `json:"mapping,omitempty"`
	Primary      bool                       `json:"primary,omitempty"`
	Required     bool                       `json:"required,omitempty"`
	Unique       bool                       `json:"unique,omitempty"`
	Immutable    bool                       `json:"immutable,omitempty"`
	Length       int32                      `json:"length,omitempty"`
	Item         *ResourceProperty          `json:"item,omitempty"`
	Properties   []ResourceProperty         `json:"properties,omitempty"`
	Reference    *ResourceReference         `json:"reference,omitempty"`
	DefaultValue *unstructured.Unstructured `json:"defaultValue,omitempty"`
	EnumValues   []string                   `json:"enumValues,omitempty"`
	ExampleValue *unstructured.Unstructured `json:"exampleValue,omitempty"`
	Title        *string                    `json:"title,omitempty"`
	Description  *string                    `json:"description,omitempty"`
	Annotations  map[string]string          `json:"annotations,omitempty"`
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
func (s *ResourceProperty) GetItem() *ResourceProperty {
	return s.Item
}
func (s *ResourceProperty) GetProperties() []ResourceProperty {
	return s.Properties
}
func (s *ResourceProperty) GetReference() *ResourceReference {
	return s.Reference
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
	Name       string             `json:"name,omitempty"`
	Properties []ResourceProperty `json:"properties,omitempty"`
}

func (s *ResourceSubType) GetName() string {
	return s.Name
}
func (s *ResourceSubType) GetProperties() []ResourceProperty {
	return s.Properties
}

type ResourceIndexProperty struct {
	Name  string         `json:"name,omitempty"`
	Order *ResourceOrder `json:"order,omitempty"`
}

func (s *ResourceIndexProperty) GetName() string {
	return s.Name
}
func (s *ResourceIndexProperty) GetOrder() *ResourceOrder {
	return s.Order
}

type ResourceIndex struct {
	Properties  []ResourceIndexProperty `json:"properties,omitempty"`
	IndexType   *ResourceIndexType      `json:"indexType,omitempty"`
	Unique      *bool                   `json:"unique,omitempty"`
	Annotations map[string]string       `json:"annotations,omitempty"`
}

func (s *ResourceIndex) GetProperties() []ResourceIndexProperty {
	return s.Properties
}
func (s *ResourceIndex) GetIndexType() *ResourceIndexType {
	return s.IndexType
}
func (s *ResourceIndex) GetUnique() *bool {
	return s.Unique
}
func (s *ResourceIndex) GetAnnotations() map[string]string {
	return s.Annotations
}

type ResourceReference struct {
	Resource      *Resource `json:"resource,omitempty"`
	Cascade       *bool     `json:"cascade,omitempty"`
	BackReference *string   `json:"backReference,omitempty"`
}

func (s *ResourceReference) GetResource() *Resource {
	return s.Resource
}
func (s *ResourceReference) GetCascade() *bool {
	return s.Cascade
}
func (s *ResourceReference) GetBackReference() *string {
	return s.BackReference
}

type ResourceOrder string

const (
	ResourceOrder_UNKNOWN ResourceOrder = "UNKNOWN"
	ResourceOrder_ASC     ResourceOrder = "ASC"
	ResourceOrder_DESC    ResourceOrder = "DESC"
)

type ResourceIndexType string

const (
	ResourceIndexType_BTREE ResourceIndexType = "BTREE"
	ResourceIndexType_HASH  ResourceIndexType = "HASH"
)
