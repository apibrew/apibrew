// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type DataSource struct {
	Id          *uuid.UUID           `json:"id,omitempty"`
	Version     int32                `json:"version,omitempty"`
	AuditData   *DataSourceAuditData `json:"auditData,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Backend     DataSourceBackend    `json:"backend,omitempty"`
	Options     map[string]string    `json:"options,omitempty"`
}

func (s *DataSource) GetId() *uuid.UUID {
	return s.Id
}
func (s *DataSource) GetVersion() int32 {
	return s.Version
}
func (s *DataSource) GetAuditData() *DataSourceAuditData {
	return s.AuditData
}
func (s *DataSource) GetName() string {
	return s.Name
}
func (s *DataSource) GetDescription() string {
	return s.Description
}
func (s *DataSource) GetBackend() DataSourceBackend {
	return s.Backend
}
func (s *DataSource) GetOptions() map[string]string {
	return s.Options
}

type DataSourceAuditData struct {
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (s *DataSourceAuditData) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *DataSourceAuditData) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *DataSourceAuditData) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *DataSourceAuditData) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}

type DataSourceBackend string

const (
	DataSourceBackend_POSTGRESQL DataSourceBackend = "POSTGRESQL"
	DataSourceBackend_MYSQL      DataSourceBackend = "MYSQL"
	DataSourceBackend_MONGODB    DataSourceBackend = "MONGODB"
	DataSourceBackend_REDIS      DataSourceBackend = "REDIS"
)
