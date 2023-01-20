package errors

import (
	"data-handler/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceError interface {
	Code() model.ErrorCode
	Error() string
	ProtoError() *model.Error
	WithMessage(msg string) ServiceError
	WithDetails(details string) ServiceError
	WithErrorFields(errors []*model.ErrorField) ServiceError
}

type serviceError struct {
	code        model.ErrorCode
	message     string
	details     string
	errorFields []*model.ErrorField
}

func (s serviceError) Code() model.ErrorCode {
	return s.code
}

func (s serviceError) Error() string {
	return s.ProtoError().Message
}

func (s serviceError) ProtoError() *model.Error {
	message := s.message

	if s.details != "" {
		message = fmt.Sprintf("%s: %s", s.message, s.details)
	}
	return &model.Error{
		Code:    s.code,
		Message: message,
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

func newServiceError(code model.ErrorCode, message string) ServiceError {
	return &serviceError{code: code, message: message}
}

func ToStatusError(err ServiceError) error {
	if err == nil {
		return nil
	}

	st := status.New(codes.Unknown, err.Error())

	st, _ = st.WithDetails(err.ProtoError())

	return st.Err()
}

func GetErrorCode(err error) model.ErrorCode {
	st, found := status.FromError(err)

	if !found {
		return model.ErrorCode_INTERNAL_ERROR
	}

	a := st.Details()[0].(*model.Error)

	return a.GetCode()
}

func GetServiceError(err error) model.ErrorCode {
	st, found := status.FromError(err)

	if !found {
		return model.ErrorCode_INTERNAL_ERROR
	}

	a := st.Details()[0]

	log.Print(a)

	return model.ErrorCode_BACKEND_ERROR
}
