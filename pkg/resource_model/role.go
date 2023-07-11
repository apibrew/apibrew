package resource_model

import "github.com/google/uuid"
import "time"

type Role struct {
	Id                  *uuid.UUID
	Name                string
	SecurityConstraints []interface{}
	Details             *interface{}
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Version             *int32
}

func (s *Role) GetId() *uuid.UUID {
	return s.Id
}
func (s *Role) GetName() string {
	return s.Name
}
func (s *Role) GetSecurityConstraints() []interface{} {
	return s.SecurityConstraints
}
func (s *Role) GetDetails() *interface{} {
	return s.Details
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
func (s *Role) GetVersion() *int32 {
	return s.Version
}
