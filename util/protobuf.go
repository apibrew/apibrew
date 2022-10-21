package util

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
)

var mo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
}

var umo = protojson.UnmarshalOptions{}

func Write(fileName string, msg proto.Message) error {
	out, err := proto.Marshal(msg)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, out, os.ModePerm)
}

func WriteJson(fileName string, msg proto.Message) error {
	out, err := mo.Marshal(msg)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, out, os.ModePerm)
}

func Read(fileName string, msg proto.Message) error {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	return proto.Unmarshal(bytes, msg)
}

func ReadJson(fileName string, msg proto.Message) error {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	return umo.Unmarshal(bytes, msg)
}
