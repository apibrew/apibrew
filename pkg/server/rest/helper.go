package rest

import (
	"encoding/json"
	errors "github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

func parseRequestMessage[T interface{}](request *http.Request, msg T) error {
	data, err := io.ReadAll(request.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, msg)

	return err
}

func ServiceResponder[T proto.Message]() ServiceCaller[T] {
	return ServiceCaller[T]{}
}

type ServiceCaller[T proto.Message] struct {
	request        *http.Request
	writer         http.ResponseWriter
	responseMapper func(response proto.Message) proto.Message
}

func (s ServiceCaller[T]) Request(request *http.Request) ServiceCaller[T] {
	s.request = request

	return s
}

func (s ServiceCaller[T]) Writer(writer http.ResponseWriter) ServiceCaller[T] {
	s.writer = writer

	return s
}

func (s ServiceCaller[T]) Respond(result interface{}, serviceError errors.ServiceError) {
	s.writer.Header().Set("Content-Type", "application/json")

	isSuccess := serviceError == nil
	if !isSuccess {
		handleServiceError(s.writer, serviceError)
		return
	} else {
		s.writer.WriteHeader(200)
	}

	body, err := json.Marshal(result)

	if err != nil {
		log.Error(err)
	}

	s.writer.Write(body)
}

func handleServiceError(writer http.ResponseWriter, err errors.ServiceError) {
	writer.WriteHeader(errorCodeHttpStatusMap[err.ProtoError().GetCode()])

	body, xerr := protojson.Marshal(err.ProtoError())

	if xerr != nil {
		log.Error(xerr)
		return
	}

	writer.Write(body)
}

func handleClientError(writer http.ResponseWriter, err error) {
	if err != nil {
		log.Error(err)
		writer.WriteHeader(400)
		writer.Write([]byte(err.Error()))
	}
}

func handleClientErrorText(writer http.ResponseWriter, err string) {
	if err != "" {
		log.Error(err)
		writer.WriteHeader(400)
		writer.Write([]byte(err))
	}
}

var errorCodeHttpStatusMap = map[model.ErrorCode]int{
	model.ErrorCode_RECORD_NOT_FOUND:                     404,
	model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY:         400,
	model.ErrorCode_INTERNAL_ERROR:                       500,
	model.ErrorCode_PROPERTY_NOT_FOUND:                   400,
	model.ErrorCode_RECORD_VALIDATION_ERROR:              400,
	model.ErrorCode_RESOURCE_VALIDATION_ERROR:            400,
	model.ErrorCode_AUTHENTICATION_FAILED:                401,
	model.ErrorCode_ALREADY_EXISTS:                       409,
	model.ErrorCode_ACCESS_DENIED:                        403,
	model.ErrorCode_BACKEND_ERROR:                        500,
	model.ErrorCode_UNIQUE_VIOLATION:                     409,
	model.ErrorCode_REFERENCE_VIOLATION:                  409,
	model.ErrorCode_RESOURCE_NOT_FOUND:                   404,
	model.ErrorCode_UNSUPPORTED_OPERATION:                400,
	model.ErrorCode_EXTERNAL_BACKEND_COMMUNICATION_ERROR: 500,
	model.ErrorCode_EXTERNAL_BACKEND_ERROR:               500,
}

func getRequestBoolFlag(request *http.Request, s string) bool {
	return request.URL.Query().Has(s)
}

func toProtoError(err errors.ServiceError) *model.Error {
	if err == nil {
		return nil
	}

	return err.ProtoError()
}