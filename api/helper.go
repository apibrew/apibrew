package api

import (
	"context"
	"data-handler/service"
	"data-handler/stub/model"
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
	serviceCaller  func(ctx context.Context, req T) (R, error)
	payload        T
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

func (s ServiceCaller[T, R]) Payload(payload T) ServiceCaller[T, R] {
	s.payload = payload

	return s
}

func (s ServiceCaller[T, R]) ServiceCall(f func(ctx context.Context, req T) (R, error)) ServiceCaller[T, R] {
	s.serviceCaller = f

	return s
}

func (s ServiceCaller[T, R]) ResponseMapper(f func(response R) proto.Message) ServiceCaller[T, R] {
	s.responseMapper = f

	return s
}

func (s ServiceCaller[T, R]) Respond() {
	serviceResult, err := s.serviceCaller(s.request.Context(), s.payload)

	s.writer.Header().Set("Content-Type", "application/json")

	isSuccess := err == nil && serviceResult.GetError() == nil
	if !isSuccess {
		handleServiceError(s.writer, serviceResult, err)
	} else {
		s.writer.WriteHeader(200)
	}

	var msg proto.Message = serviceResult
	if isSuccess && s.responseMapper != nil {
		msg = s.responseMapper(serviceResult)
	}

	body, err := mo.Marshal(msg)

	if err != nil {
		log.Error(err)
	}

	s.writer.Write(body)
}

func handleServiceError(writer http.ResponseWriter, serviceResult service.Response, err error) {
	if err != nil {
		log.Error(err)
		writer.WriteHeader(500)
	} else if serviceResult.GetError() != nil {
		writer.WriteHeader(errorCodeHttpStatusMap[serviceResult.GetError().GetCode()])
	}
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
}

func getToken(request *http.Request) string {
	return request.Header.Get("Authorization")
}
