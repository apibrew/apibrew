package service

import (
	"data-handler/service/errors"
	"data-handler/stub/model"
	"google.golang.org/protobuf/proto"
)

type Response interface {
	proto.Message
	GetError() *model.Error
}

func toProtoError(err error) *model.Error {
	if serviceError, ok := err.(errors.ServiceError); ok {
		return serviceError.ProtoError()
	}

	return errors.InternalError.ProtoError()
}
