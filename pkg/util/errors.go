package util

import (
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/grpc/status"
)

func ToStatusError(err errors.ServiceError) error {
	if err == nil {
		return nil
	}

	st := status.New(err.GetGrpcErrorCode(), err.Error())

	st, _ = st.WithDetails(err.ProtoError())

	return st.Err()
}

func GetErrorCode(err error) model.ErrorCode {
	st, found := status.FromError(err)

	if !found {
		return model.ErrorCode_UNKNOWN_ERROR
	}

	a := st.Details()[0].(*model.Error)

	return a.GetCode()
}

func GetErrorFields(err error) []*model.ErrorField {
	st, found := status.FromError(err)

	if !found {
		return nil
	}

	a := st.Details()[0].(*model.Error)

	return a.Fields
}

func GetErrorMessage(err error) string {
	st, found := status.FromError(err)

	if !found {
		return ""
	}

	a := st.Details()[0].(*model.Error)

	return a.Message
}
