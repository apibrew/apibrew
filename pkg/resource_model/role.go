package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type Role struct {
	Id                  *uuid.UUID                 `json:"id"`
	Version             int32                      `json:"version"`
	CreatedBy           *string                    `json:"createdBy"`
	UpdatedBy           *string                    `json:"updatedBy"`
	CreatedOn           *time.Time                 `json:"createdOn"`
	UpdatedOn           *time.Time                 `json:"updatedOn"`
	Name                string                     `json:"name"`
	SecurityConstraints []*SecurityConstraint      `json:"securityConstraints"`
	Details             *unstructured.Unstructured `json:"details"`
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
