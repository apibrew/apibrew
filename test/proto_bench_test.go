package test

import (
	"data-handler/model"
	"data-handler/test/data"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func BenchmarkSerializeProtobufSmall(b *testing.B) {
	newInitData := &model.InitData{}

	for i := 0; i < 1000; i++ {
		newInitData.InitResources = append(newInitData.InitResources, data.PreparePersonResource())
	}

	data, err := proto.Marshal(newInitData)

	log.Print(len(data), err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			proto.Marshal(newInitData)
		}
	})
}
