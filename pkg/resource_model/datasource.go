package resource_model

import "github.com/google/uuid"
import "time"

type DataSource struct {
	Id          *uuid.UUID        `json:"id"`
	Version     int32             `json:"version"`
	CreatedBy   *string           `json:"createdBy"`
	UpdatedBy   *string           `json:"updatedBy"`
	CreatedOn   *time.Time        `json:"createdOn"`
	UpdatedOn   *time.Time        `json:"updatedOn"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Backend     DataSourceBackend `json:"backend"`
	Options     map[string]string `json:"options"`
}

func (s *DataSource) GetId() *uuid.UUID {
	return s.Id
}
func (s *DataSource) GetVersion() int32 {
	return s.Version
}
func (s *DataSource) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *DataSource) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *DataSource) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *DataSource) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
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

type DataSourceBackend string

const (
	DataSourceBackend_POSTGRESQL DataSourceBackend = "POSTGRESQL"
	DataSourceBackend_MYSQL      DataSourceBackend = "MYSQL"
	DataSourceBackend_MONGODB    DataSourceBackend = "MONGODB"
	DataSourceBackend_REDIS      DataSourceBackend = "REDIS"
)
