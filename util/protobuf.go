package util

import (
	"google.golang.org/protobuf/proto"
	"os"
)

func Write(fileName string, msg proto.Message) error {
	out, err := proto.Marshal(msg)

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
