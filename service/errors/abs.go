package errors

import (
	"data-handler/stub/model"
	"fmt"
)

type ServiceError interface {
	Error() string
	ProtoError() *model.Error
	WithMessage(msg string) ServiceError
	WithDetails(details string) ServiceError
}

type serviceError struct {
	code    model.ErrorCode
	message string
	details string
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

func newServiceError(code model.ErrorCode, message string) ServiceError {
	return &serviceError{code: code, message: message}
}
