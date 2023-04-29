package server

import (
	. "github.com/tislib/apibrew/pkg/apbradm/model"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/resources"
	"github.com/tislib/apibrew/pkg/resources/mapping"
	"github.com/tislib/apibrew/pkg/service/annotations"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var NodeResource = &model.Resource{
	Name: "node",
	Properties: []*model.ResourceProperty{
		resources.IdProperty,
		resources.VersionProperty,
		resources.AuditProperties[0],
		resources.AuditProperties[1],
		resources.AuditProperties[2],
		resources.AuditProperties[3],
		{
			Name:      "name",
			Mapping:   "name",
			Primary:   false,
			Type:      model.ResourceProperty_STRING,
			Length:    256,
			Required:  true,
			Unique:    true,
			Immutable: true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name: "description",

			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "url",

			Mapping:  "url",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name: "authenticationToken",

			Mapping:  "authentication_token",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   2048,
			Required: true,
		},
	},
}

func NodeToRecord(node *Node) *model.Record {
	if node.AuditData == nil {
		node.AuditData = &model.AuditData{
			CreatedBy: "system",
			UpdatedBy: "system",
			CreatedOn: timestamppb.New(time.Now()),
			UpdatedOn: timestamppb.New(time.Now()),
		}
	}
	record := &model.Record{
		Id: node.Id,
		Properties: map[string]*structpb.Value{
			"id":                  structpb.NewStringValue(node.Id),
			"name":                structpb.NewStringValue(node.Name),
			"description":         structpb.NewStringValue(node.Description),
			"url":                 structpb.NewStringValue(node.Url),
			"authenticationToken": structpb.NewStringValue(node.AuthenticationToken),
		},
	}

	mapping.MapSpecialColumnsToRecord(node, &record.Properties)

	return record
}

func RecordToNode(record *model.Record) *Node {
	node := &Node{
		Id:                  record.Id,
		Name:                record.Properties["name"].GetStringValue(),
		Description:         record.Properties["description"].GetStringValue(),
		Url:                 record.Properties["url"].GetStringValue(),
		AuthenticationToken: record.Properties["authenticationToken"].GetStringValue(),
	}

	mapping.MapSpecialColumnsFromRecord(node, &record.Properties)

	return node
}
