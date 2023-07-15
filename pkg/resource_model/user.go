package resource_model

import "github.com/google/uuid"
import "time"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
import "encoding/json"

type User struct {
	Id                  *uuid.UUID
	Version             *int32
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Username            string
	Password            *string
	Roles               []string
	SecurityConstraints []*SecurityConstraint
	Details             *unstructured.Unstructured
}

func (s *User) GetId() *uuid.UUID {
	return s.Id
}
func (s *User) GetVersion() *int32 {
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
func (s *User) GetRoles() []string {
	return s.Roles
}
func (s *User) GetSecurityConstraints() []*SecurityConstraint {
	return s.SecurityConstraints
}
func (s *User) GetDetails() *unstructured.Unstructured {
	return s.Details
}
