package common

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
)

type CustomType struct {
	CustomEquals  func(a, b interface{}) bool
	CustomPack    func(value interface{}) (*structpb.Value, error)
	CustomUnPack  func(val *structpb.Value) (interface{}, error)
	CustomDefault func() any
	CustomPointer func(required bool) any
	CustomString  func(val any) string
	CustomIsEmpty func(value any) bool
}

func (u CustomType) Equals(a, b interface{}) bool {
	if u.CustomEquals == nil {
		return a == b
	} else {
		return u.CustomEquals(a, b)
	}
}

func (u CustomType) Pack(val interface{}) (*structpb.Value, error) {
	if u.CustomPack == nil {
		return structpb.NewValue(val)
	} else {
		return u.CustomPack(val)
	}
}

func (u CustomType) UnPack(val *structpb.Value) (interface{}, error) {
	if u.CustomPack == nil {
		return structpb.NewValue(val)
	} else {
		return u.CustomUnPack(val)
	}
}

func (u CustomType) Default() any {
	if u.CustomDefault == nil {
		return new(interface{})
	} else {
		return u.CustomDefault()
	}
}

func (u CustomType) Pointer(required bool) any {
	if u.CustomPointer == nil {
		if required {
			return new(interface{})
		} else {
			return new(*interface{})
		}
	} else {
		return u.CustomPointer(required)
	}
}

func (u CustomType) String(val any) string {
	if u.CustomString == nil {
		return fmt.Sprintf("%v", val)
	} else {
		return u.CustomString(val)
	}
}

func (u CustomType) IsEmpty(value any) bool {
	if u.CustomIsEmpty == nil {
		return value == nil
	} else {
		return u.CustomIsEmpty(value)
	}
}

func CustomTypeFromType(typ types.PropertyType, override CustomType) CustomType {
	customType := CustomType{
		CustomEquals:  typ.Equals,
		CustomPack:    typ.Pack,
		CustomUnPack:  typ.UnPack,
		CustomDefault: typ.Default,
		CustomPointer: typ.Pointer,
		CustomString:  typ.String,
		CustomIsEmpty: typ.IsEmpty,
	}

	if override.CustomEquals != nil {
		customType.CustomEquals = override.CustomEquals
	}

	if override.CustomPack != nil {
		customType.CustomPack = override.CustomPack
	}

	if override.CustomUnPack != nil {
		customType.CustomUnPack = override.CustomUnPack
	}

	if override.CustomDefault != nil {
		customType.CustomDefault = override.CustomDefault
	}

	if override.CustomPointer != nil {
		customType.CustomPointer = override.CustomPointer
	}

	if override.CustomString != nil {
		customType.CustomString = override.CustomString
	}

	if override.CustomIsEmpty != nil {
		customType.CustomIsEmpty = override.CustomIsEmpty
	}

	return customType
}
