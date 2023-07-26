package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Role struct {
	Id                  *uuid.UUID                 `json:"id,omitempty"`
	Version             int32                      `json:"version,omitempty"`
	CreatedBy           *string                    `json:"createdBy,omitempty"`
	UpdatedBy           *string                    `json:"updatedBy,omitempty"`
	CreatedOn           *time.Time                 `json:"createdOn,omitempty"`
	UpdatedOn           *time.Time                 `json:"updatedOn,omitempty"`
	Name                string                     `json:"name,omitempty"`
	SecurityConstraints []*SecurityConstraint      `json:"securityConstraints,omitempty"`
	Details             *unstructured.Unstructured `json:"details,omitempty"`
}

func (s *Role) GetId() *uuid.UUID {
	return s.Id
}
func (s *Role) GetVersion() int32 {
	return s.Version
}
func (s *Role) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Role) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Role) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Role) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Role) GetName() string {
	return s.Name
}
func (s *Role) GetSecurityConstraints() []*SecurityConstraint {
	return s.SecurityConstraints
}
func (s *Role) GetDetails() *unstructured.Unstructured {
	return s.Details
}
