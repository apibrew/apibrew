package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Namespace struct {
	Id          *uuid.UUID                 `json:"id,omitempty"`
	Version     int32                      `json:"version,omitempty"`
	CreatedBy   *string                    `json:"createdBy,omitempty"`
	UpdatedBy   *string                    `json:"updatedBy,omitempty"`
	CreatedOn   *time.Time                 `json:"createdOn,omitempty"`
	UpdatedOn   *time.Time                 `json:"updatedOn,omitempty"`
	Name        string                     `json:"name,omitempty"`
	Description *string                    `json:"description,omitempty"`
	Details     *unstructured.Unstructured `json:"details,omitempty"`
}

func (s *Namespace) GetId() *uuid.UUID {
	return s.Id
}
func (s *Namespace) GetVersion() int32 {
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
func (s *Namespace) GetDetails() *unstructured.Unstructured {
	return s.Details
}
