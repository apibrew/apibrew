package main

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/proto"
	"testing"
)

var msg = &model.Event{
	Id:                "esdaskdsajkdsajkdjsa",
	Action:            model.Event_UPDATE,
	ActionSummary:     "sadaskdjsahdjsahdjsahjda",
	ActionDescription: "sdsajdksajdksajdjaskd",
	Resource: &model.Resource{
		Id:   "dsajkdsajkdjsakdjsakdjsa",
		Name: "dsajkdsajkdjsakdjsakdjsa",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: "dsadas",
			Catalog:    "asdsadsa",
			Entity:     "sadsadsadsa",
		},
	},
}

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(msg)
	}
}

func BenchmarkProtoc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(msg)
	}
}
