package errors

import "data-handler/stub/model"

type ServiceError interface {
	Error() string
	ProtoError() *model.Error
	WithMessage(msg string) ServiceError
}

type serviceError struct {
	code    model.ErrorCode
	message string
}

func (s serviceError) Error() string {
	return s.ProtoError().Message
}

func (s serviceError) ProtoError() *model.Error {
	return &model.Error{
		Code:    s.code,
		Message: s.message,
	}
}

func (s serviceError) WithMessage(msg string) ServiceError {
	s.message = msg
	return s
}

func newServiceError(code model.ErrorCode, message string) ServiceError {
	return &serviceError{code: code, message: message}
}
