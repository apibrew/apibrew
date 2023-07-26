package resource_model

import "github.com/google/uuid"
import "github.com/apibrew/apibrew/pkg/formats/unstructured"
import "encoding/json"

type Record struct {
	Id               *uuid.UUID                  `json:"id,omitempty"`
	Properties       unstructured.Unstructured   `json:"properties,omitempty"`
	PackedProperties []unstructured.Unstructured `json:"packedProperties,omitempty"`
}

func (s *Record) GetId() *uuid.UUID {
	return s.Id
}
func (s *Record) GetProperties() unstructured.Unstructured {
	return s.Properties
}
func (s *Record) GetPackedProperties() []unstructured.Unstructured {
	return s.PackedProperties
}

func (s *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Properties)
}

func (s *Record) UnmarshalJSON(data []byte) error {
	s.Properties = make(unstructured.Unstructured)

	return json.Unmarshal(data, &s.Properties)
}
