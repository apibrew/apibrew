package errors

import (
	"fmt"
	"github.com/tislib/data-handler/pkg/model"
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
}

type serviceError struct {
	code        model.ErrorCode
	message     string
	details     string
	errorFields []*model.ErrorField
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

func (s serviceError) Is(err error) bool {
	if err == nil {
		return false
	}

	if se, ok := err.(serviceError); ok {
		return s.Code() == se.Code()
	}

	return false
}

func newServiceError(code model.ErrorCode, message string) ServiceError {
	return &serviceError{code: code, message: message}
}
