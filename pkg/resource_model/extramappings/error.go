package extramappings

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
)

func ErrorToProto(result resource_model.ExtensionError) *model.Error {
	var err = new(model.Error)

	err.Message = util.DePointer(result.Message, "")
	err.Code = model.ErrorCode(model.ErrorCode_value[string(util.DePointer(result.Code, resource_model.ExtensionCode_UNKNOWNERROR))])

	for _, field := range result.Fields {
		err.Fields = append(err.Fields, &model.ErrorField{
			Property: util.DePointer(field.Property, ""),
			Message:  util.DePointer(field.Message, ""),
		})
	}

	return err
}

func ErrorFromProto(result *model.Error) resource_model.ExtensionError {
	var err = resource_model.ExtensionError{
		Message: util.Pointer(result.Message),
		Code:    util.Pointer(resource_model.ExtensionCode(model.ErrorCode_name[int32(result.Code)])),
	}

	for _, field := range result.Fields {
		err.Fields = append(err.Fields, resource_model.ExtensionErrorField{
			Property: util.Pointer(field.Property),
			Message:  util.Pointer(field.Message),
		})
	}

	return err
}
