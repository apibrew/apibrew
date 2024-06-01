package core

import (
	"google.golang.org/protobuf/types/known/structpb"
)

type Record struct {
	Properties       map[string]*structpb.Value `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PropertiesPacked []*structpb.Value          `protobuf:"bytes,5,rep,name=propertiesPacked,proto3" json:"propertiesPacked,omitempty"`
}
