package resource_model

import "github.com/google/uuid"
import "time"

type Namespace struct {
	Id                  *uuid.UUID
	Version             *int32
	CreatedBy           *string
	UpdatedBy           *string
	CreatedOn           *time.Time
	UpdatedOn           *time.Time
	Name                string
	Description         *string
	Details             *interface{}
	SecurityConstraints []*SecurityConstraint
}

func (s *Namespace) GetId() *uuid.UUID {
	return s.Id
}
func (s *Namespace) GetVersion() *int32 {
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
func (s *Namespace) GetDetails() *interface{} {
	return s.Details
}
func (s *Namespace) GetSecurityConstraints() []*SecurityConstraint {
	return s.SecurityConstraints
}
