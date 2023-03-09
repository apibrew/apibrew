package test

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

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
