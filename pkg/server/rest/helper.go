package rest

import (
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
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

func (s ServiceCaller[T]) Respond(serviceResult proto.Message, serviceError errors.ServiceError) {
	s.writer.Header().Set("Content-Type", "application/json")

	isSuccess := serviceError == nil
	if !isSuccess {
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
	model.ErrorCode_RECORD_NOT_FOUND:             404,
	model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY: 400,
	model.ErrorCode_INTERNAL_ERROR:               500,
	model.ErrorCode_PROPERTY_NOT_FOUND:           400,
	model.ErrorCode_RECORD_VALIDATION_ERROR:      400,
	model.ErrorCode_AUTHENTICATION_FAILED:        401,
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

var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.NewWithNamespace("rest")

	if err != nil {
		log.Fatal(err)
	}
}
