package main

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"os"
)

func main() {
	var descriptor = &descriptorpb.FileDescriptorSet{}

	data, _ := os.ReadFile("proto/image.pb")

	_ = proto.Unmarshal(data, descriptor)

	//for _, file := range descriptor.File {
	//	for _, service := range file.Service {
	//		for _, method := range service.Method {
	//			ext1 := proto.GetExtension(method.GetOptions(), stub.E_Resources).(*stub.Logic)
	//			if *method.Name == "Create" && *service.Name == "RecordService" {
	//
	//				log.Print(ext1)
	//			}
	//			log.Print(ext1 == (*stub.Logic)(nil))
	//			log.Print(ext1)
	//		}
	//	}
	//}

}
