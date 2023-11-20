// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type Role struct {
	Id          *uuid.UUID     `json:"id,omitempty"`
	Version     int32          `json:"version,omitempty"`
	AuditData   *RoleAuditData `json:"auditData,omitempty"`
	Name        string         `json:"name,omitempty"`
	Permissions []*Permission  `json:"permissions,omitempty"`
	Details     interface{}    `json:"details,omitempty"`
}

func (s *Role) GetId() *uuid.UUID {
	return s.Id
}
func (s *Role) GetVersion() int32 {
	return s.Version
}
func (s *Role) GetAuditData() *RoleAuditData {
	return s.AuditData
}
func (s *Role) GetName() string {
	return s.Name
}
func (s *Role) GetPermissions() []*Permission {
	return s.Permissions
}
func (s *Role) GetDetails() interface{} {
	return s.Details
}

type RoleAuditData struct {
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (s *RoleAuditData) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *RoleAuditData) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *RoleAuditData) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *RoleAuditData) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
