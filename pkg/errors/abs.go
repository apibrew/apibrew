package errors

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type ServiceError interface {
	Code() model.ErrorCode
	Error() string
	ProtoError() *model.Error
	WithMessage(msg string) ServiceError
	WithDetails(details string) ServiceError
	WithErrorFields(errors []*model.ErrorField) ServiceError
	GetDetails() string
	Is(err error) bool
	GetGrpcErrorCode() codes.Code
	GetFullMessage() string
}

type serviceError struct {
	code          model.ErrorCode
	message       string
	details       string
	errorFields   []*model.ErrorField
	grpcErrorCode codes.Code
}

func (s serviceError) GetFullMessage() string {
	message := s.message

	if s.details != "" {
		message = fmt.Sprintf("%s: %s", s.message, s.details)
	}

	if len(s.errorFields) > 0 {
		message = message + " -"
		for _, ef := range s.errorFields {
			var valueBytes []byte

			if ef.Value != nil {
				valueBytes, _ = protojson.Marshal(ef.Value)
			}

			message = fmt.Sprintf("%s (%s => %s => %s)", message, ef.Property, ef.Message, string(valueBytes))
		}
	}

	return message
}

func (s serviceError) GetGrpcErrorCode() codes.Code {
	return s.grpcErrorCode
}

func (s serviceError) GetDetails() string {
	return s.details
}

func (s serviceError) Code() model.ErrorCode {
	return s.code
}

func (s serviceError) Error() string {
	return s.ProtoError().Message
}

func (s serviceError) ProtoError() *model.Error {
	return &model.Error{
		Code:    s.code,
		Message: s.GetFullMessage(),
		Fields:  s.errorFields,
	}
}

func (s serviceError) WithMessage(msg string) ServiceError {
	s.message = msg
	return s
}

func (s serviceError) WithDetails(details string) ServiceError {
	s.details = details
	return s
}

func (s serviceError) WithErrorFields(errorFields []*model.ErrorField) ServiceError {
	s.errorFields = errorFields
	return s
}

func (s serviceError) Is(err error) bool {
	if err == nil {
		return false
	}

	if se, ok := err.(serviceError); ok {
		return s.Code() == se.Code()
	}

	if se, ok := err.(*serviceError); ok {
		return s.Code() == se.Code()
	}

	if errorCode := getErrorCode(err); errorCode != model.ErrorCode_UNKNOWN_ERROR {
		return s.Code() == errorCode
	}

	return false
}

func getErrorCode(err error) model.ErrorCode {
	st, found := status.FromError(err)

	if !found {
		return model.ErrorCode_UNKNOWN_ERROR
	}

	a := st.Details()[0].(*model.Error)

	return a.GetCode()
}

func NewServiceError(code model.ErrorCode, message string, grpcErrorCode codes.Code) ServiceError {
	return &serviceError{code: code, message: message, grpcErrorCode: grpcErrorCode}
}

func FromGrpcError(err error) ServiceError {
	st, found := status.FromError(err)

	if !found {
		return NewServiceError(model.ErrorCode_UNKNOWN_ERROR, err.Error(), codes.Unknown)
	}

	a, ok := st.Details()[0].(*model.Error)

	if !ok {
		return NewServiceError(model.ErrorCode_UNKNOWN_ERROR, err.Error(), codes.Unknown)
	}

	return FromProtoError(a)
}

func FromProtoError(err *model.Error) ServiceError {
	return &serviceError{
		code:          err.Code,
		message:       err.Message,
		details:       "",
		errorFields:   err.Fields,
		grpcErrorCode: codes.Aborted,
	}
}
