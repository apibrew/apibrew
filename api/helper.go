package api

import (
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

var mo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
}

var umo = protojson.UnmarshalOptions{}

func parseRequestMessage[T proto.Message](request *http.Request, msg T) error {
	data, err := io.ReadAll(request.Body)

	if err != nil {
		return err
	}

	err = umo.Unmarshal(data, msg)

	return err
}

func ServiceResponder[T proto.Message, R service.Response]() ServiceCaller[T, R] {
	return ServiceCaller[T, R]{}
}

type ServiceCaller[T proto.Message, R service.Response] struct {
	request        *http.Request
	writer         http.ResponseWriter
	responseMapper func(response R) proto.Message
}

func (s ServiceCaller[T, R]) Request(request *http.Request) ServiceCaller[T, R] {
	s.request = request

	return s
}

func (s ServiceCaller[T, R]) Writer(writer http.ResponseWriter) ServiceCaller[T, R] {
	s.writer = writer

	return s
}

func (s ServiceCaller[T, R]) Respond(serviceResult proto.Message, serviceError errors.ServiceError) {
	s.writer.Header().Set("Content-Type", "application/json")

	isSuccess := serviceError == nil
	if !isSuccess {
		s.writer.WriteHeader(400)
		handleServiceError(s.writer, serviceError)
	} else {
		s.writer.WriteHeader(200)
	}

	body, err := mo.Marshal(serviceResult)

	if err != nil {
		log.Error(err)
	}

	s.writer.Write(body)
}

func handleServiceError(writer http.ResponseWriter, err errors.ServiceError) {
	writer.WriteHeader(errorCodeHttpStatusMap[err.ProtoError().GetCode()])
}

func handleClientError(writer http.ResponseWriter, err error) {
	if err != nil {
		log.Error(err)
		writer.WriteHeader(500)
		writer.Write([]byte("Invalid Request Data: " + err.Error()))
	}
}

var errorCodeHttpStatusMap = map[model.ErrorCode]int{
	model.ErrorCode_RECORD_NOT_FOUND:             404,
	model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY: 400,
	model.ErrorCode_INTERNAL_ERROR:               500,
	model.ErrorCode_PROPERTY_NOT_FOUND:           400,
	model.ErrorCode_RECORD_VALIDATION_ERROR:      400,
}

func getToken(request *http.Request) string {
	return request.Header.Get("Authorization")
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
