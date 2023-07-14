package resource_model

import "github.com/google/uuid"

type Record struct {
	Id               *uuid.UUID
	Properties       map[string]interface{}
	PackedProperties []interface{}
}

func (s *Record) GetId() *uuid.UUID {
	return s.Id
}
func (s *Record) GetProperties() map[string]interface{} {
	return s.Properties
}
func (s *Record) GetPackedProperties() []interface{} {
	return s.PackedProperties
}
