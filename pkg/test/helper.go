package test

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"reflect"
	"testing"
	"unicode"
)

func CheckTwoRecordEquals(t *testing.T, resource *model.Resource, a *model.Record, b *model.Record) {
	for _, prop := range resource.Properties {
		if annotations.IsEnabled(prop, annotations.SpecialProperty) {
			continue
		}

		if (a.Properties[prop.Name] != nil) != (b.Properties[prop.Name] != nil) {
			t.Errorf("[%s]; different: %v <=> %v", prop.Name, a.Properties[prop.Name], b.Properties[prop.Name])
		}

		if (a.Properties[prop.Name] == nil) && (b.Properties[prop.Name] == nil) {
			continue
		}

		typeHelper := types.ByResourcePropertyType(prop.Type)

		va, err := typeHelper.UnPack(a.Properties[prop.Name])

		if err != nil {
			t.Error(err)
			return
		}
		vb, err := typeHelper.UnPack(b.Properties[prop.Name])

		if err != nil {
			t.Error(err)
			return
		}
		if !typeHelper.Equals(va, vb) {
			t.Errorf("[%s]; different: %v <=> %v", prop.Name, va, vb)
		}
	}
}

func DeepEqual(t *testing.T, a interface{}, b interface{}, prefix string) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Type() != vb.Type() {
		t.Errorf("%sType: %s != %s\n", prefix, va.Type(), vb.Type())
		return
	}

	switch va.Kind() {
	case reflect.Ptr:
		if va.IsNil() != vb.IsNil() {
			t.Errorf("%sPointer[Nil]: %v != %v\n", prefix, va.IsNil(), vb.IsNil())
			return
		}

		if va.IsNil() {
			return
		}

		DeepEqual(t, va.Elem().Interface(), vb.Elem().Interface(), prefix)
	case reflect.Struct:
		for i := 0; i < va.NumField(); i++ {
			fa := va.Field(i)
			fb := vb.Field(i)

			if !unicode.IsUpper(rune(va.Type().Field(i).Name[0])) {
				continue // skip unexported fields
			}

			DeepEqual(t, fa.Interface(), fb.Interface(), prefix)
		}
	case reflect.Array, reflect.Slice:
		if va.Len() != vb.Len() {
			t.Errorf("%sLength: %d != %d\n", prefix, va.Len(), vb.Len())
			return
		}

		for i := 0; i < va.Len(); i++ {
			DeepEqual(t, va.Index(i).Interface(), vb.Index(i).Interface(), fmt.Sprintf("%s[%d]", prefix, i))
		}
	case reflect.Map:
		if va.Len() != vb.Len() {
			t.Errorf("%sLength: %d != %d\n", prefix, va.Len(), vb.Len())
			return
		}

		for _, key := range va.MapKeys() {
			DeepEqual(t, va.MapIndex(key).Interface(), vb.MapIndex(key).Interface(), fmt.Sprintf("%s[%v]", prefix, key.Interface()))
		}
	default:
		if va.Interface() != vb.Interface() {
			t.Errorf("%sValue: %v != %v\n", prefix, va.Interface(), vb.Interface())
		}
	}
}
