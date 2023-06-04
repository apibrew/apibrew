package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/experimental"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
	"os"
)

func main() {
	var descriptor = &descriptorpb.FileDescriptorSet{}

	data, _ := os.ReadFile("/Users/taleh/Projects/data-handler/proto/image.pb")

	_ = proto.Unmarshal(data, descriptor)

	//protoHelper := helper.NewProtoHelper(helper.ProtoHelperParams{})

	var cd protoreflect.MessageDescriptor

	for _, file := range descriptor.File {
		if *file.Name != "experimental/country.proto" {
			continue
		}

		//result := protoHelper.ParseFileDescriptorProto(file)

		cd = file.MessageType[0].ProtoReflect().Descriptor()
	}

	country := &experimental.Country{
		Description: "asdasdsadasdsadasd",
	}

	var item1 = new(anypb.Any)

	err := item1.MarshalFrom(country)

	log.Println(err)

	dpb := dynamicpb.NewMessage(cd)

	err2 := proto.Unmarshal(item1.Value, dpb)

	log.Println(err, dpb, err2)
}
