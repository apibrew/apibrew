package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"

type User struct {
	Id                  *uuid.UUID                 `json:"id,omitempty"`
	Version             int32                      `json:"version,omitempty"`
	CreatedBy           *string                    `json:"createdBy,omitempty"`
	UpdatedBy           *string                    `json:"updatedBy,omitempty"`
	CreatedOn           *time.Time                 `json:"createdOn,omitempty"`
	UpdatedOn           *time.Time                 `json:"updatedOn,omitempty"`
	Username            string                     `json:"username,omitempty"`
	Password            *string                    `json:"password,omitempty"`
	Roles               []*Role                    `json:"roles,omitempty"`
	SecurityConstraints []*SecurityConstraint      `json:"securityConstraints,omitempty"`
	Details             *unstructured.Unstructured `json:"details,omitempty"`
}

func (s *User) GetId() *uuid.UUID {
	return s.Id
}
func (s *User) GetVersion() int32 {
	return s.Version
}
func (s *User) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *User) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *User) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *User) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *User) GetUsername() string {
	return s.Username
}
func (s *User) GetPassword() *string {
	return s.Password
}
func (s *User) GetRoles() []*Role {
	return s.Roles
}
func (s *User) GetSecurityConstraints() []*SecurityConstraint {
	return s.SecurityConstraints
}
func (s *User) GetDetails() *unstructured.Unstructured {
	return s.Details
}
