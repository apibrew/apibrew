package bench1

import (
	"github.com/tislib/data-handler/model"
	"github.com/tislib/data-handler/test/data"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
	"time"
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

func BenchmarkSerializeProtobufMedium(b *testing.B) {
	newInitData := &model.InitData{}

	for i := 0; i < 100000; i++ {
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

func BenchmarkSerializeProtobufLarge(b *testing.B) {
	newInitData := &model.InitData{}

	res := data.PreparePersonResource()

	for i := 0; i < 100000000; i++ {
		newInitData.InitResources = append(newInitData.InitResources, res)
	}

	data, err := proto.Marshal(newInitData)

	log.Print(len(data), err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 10; i++ {
				log.Print(i, time.Now())
				proto.Marshal(newInitData)
			}
		}
	})
}

func BenchmarkSerializeProtobufLarge2(b *testing.B) {
	chunkCount := 10
	newInitData := &model.InitData{}

	res := data.PreparePersonResource()

	for i := 0; i < 1000000/chunkCount; i++ {
		newInitData.InitResources = append(newInitData.InitResources, res)
	}

	data, err := proto.Marshal(newInitData)

	log.Print(len(data), err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for j := 0; j < chunkCount; j++ {
				proto.Marshal(newInitData)
			}
		}
	})
}

// 651428875
// 453834792
// 1791968653
