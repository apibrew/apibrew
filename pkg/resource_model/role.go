package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
import "encoding/json"

type Role struct {
	Id                  *uuid.UUID
	Version             *int32
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Name                string
	SecurityConstraints []*SecurityConstraint
	Details             *unstructured.Unstructured
}

func (s *Role) GetId() *uuid.UUID {
	return s.Id
}
func (s *Role) GetVersion() *int32 {
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
