package service

import (
	"data-handler/model"
	"google.golang.org/protobuf/proto"
)

type Response interface {
	proto.Message
	GetError() *model.Error
}
