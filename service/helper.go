package service

import (
	"data-handler/service/errors"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
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

	log.Print("Internal error: ", err)

	return errors.InternalError.ProtoError()
}
