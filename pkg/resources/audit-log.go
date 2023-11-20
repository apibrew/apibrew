package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var AuditLogResource = &model.Resource{
	Name:      "AuditLog",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "audit_log",
	},
	Immutable: true,
	Properties: map[string]*model.ResourceProperty{
		"id":      special.IdProperty,
		"version": special.VersionProperty,
		"namespace": {
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		"resource": {
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		"recordId": {
			Length:   256,
			Type:     model.ResourceProperty_STRING,
			Required: true,
			Annotations: map[string]string{
				annotations.SourceDef: "record_id",
			},
		},
		"time": {
			Type:     model.ResourceProperty_TIMESTAMP,
			Required: true,
		},
		"username": {
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		"operation": {
			Type:     model.ResourceProperty_ENUM,
			Required: true,
			EnumValues: []string{
				"CREATE",
				"UPDATE",
				"DELETE",
			},
		},
		"properties": {
			Type:     model.ResourceProperty_OBJECT,
			Required: false,
		},
		"annotations": special.AnnotationsProperty,
	},
	Annotations: map[string]string{
		annotations.OpenApiGroup:     OpenApiInternal,
		annotations.BypassExtensions: annotations.Enabled,
	},
}
