package test

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

func randDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func fakeResourceName() string {
	return RandStringRunes(16)
}

func fakePropertyName() string {
	return RandStringRunes(16)
}

func fakeResource(properties ...*model.ResourceProperty) *model.Resource {
	name := fakeResourceName()
	return &model.Resource{
		Name:      name,
		Namespace: "default",
		SourceConfig: &model.ResourceSourceConfig{
			DataSource: setup.DhTest.Name,
			Catalog:    "public",
			Entity:     name,
		},
		Properties: properties,
	}
}

func fakeValidValues(propertyType model.ResourceProperty_Type) []interface{} {
	switch propertyType {
	case model.ResourceProperty_INT32:
		return []interface{}{r.Int31()}
	case model.ResourceProperty_INT64:
		return []interface{}{r.Int63()}
	case model.ResourceProperty_STRING:
		return []interface{}{
			RandStringRunes(16),
		}
	case model.ResourceProperty_UUID:
		randUUid, _ := uuid.NewUUID()
		return []interface{}{
			randUUid.String(),
		}
	case model.ResourceProperty_BYTES:
		return []interface{}{
			[]byte(RandStringRunes(16)),
		}
	case model.ResourceProperty_BOOL:
		return []interface{}{
			false, true,
		}
	case model.ResourceProperty_DATE:
		return []interface{}{
			randDate().Format("2006-01-02"),
		}
	case model.ResourceProperty_FLOAT32:
		return []interface{}{
			r.Float32(),
		}
	case model.ResourceProperty_FLOAT64:
		return []interface{}{
			r.Float64(),
		}
	case model.ResourceProperty_TIME:
		return []interface{}{
			randDate().Format("15:04:05"),
		}
	case model.ResourceProperty_TIMESTAMP:
		return []interface{}{
			randDate().Format(time.RFC3339),
		}
	case model.ResourceProperty_REFERENCE:
		randUUid, _ := uuid.NewUUID()
		return []interface{}{
			randUUid.String(),
		}
	case model.ResourceProperty_ENUM:
		return []interface{}{
			"enum1",
			"enum2",
			"enum3",
		}
	case model.ResourceProperty_MAP:
		return []interface{}{
			map[string]interface{}{
				"val1": RandStringRunes(16),
			},
		}
	case model.ResourceProperty_LIST:
		return []interface{}{
			[]interface{}{
				RandStringRunes(16),
				RandStringRunes(16),
			},
		}
	case model.ResourceProperty_OBJECT:
		return []interface{}{
			randDate().Format(time.RFC3339),
			map[string]interface{}{
				"a123": r.Int63(),
				"a124": r.Float64(),
				"a125": map[string]interface{}{
					"a123": RandStringRunes(32),
					"a124": 124,
					"a125": 124,
				},
			},
		}
	case model.ResourceProperty_STRUCT:
		return []interface{}{
			map[string]interface{}{
				"field-1": "abc",
				"field-2": 123,
			},
			map[string]interface{}{
				"field-1": "xyz",
			},
		}
	default:
		panic("Unknown type: " + propertyType.String())
	}
}

func fakeValidValue(propertyType model.ResourceProperty_Type) interface{} {
	values := fakeValidValues(propertyType)

	return values[r.Intn(len(values))]
}

func fakeInvalidValues(propertyType model.ResourceProperty_Type) []interface{} {
	switch propertyType {
	case model.ResourceProperty_INT32:
		return []interface{}{
			r.Int63(),
			"1234",
			false,
			true,
			r.Float64(),
		}
	case model.ResourceProperty_INT64:
		return []interface{}{
			"1234",
			false,
			true,
			r.Float64(),
		}
	case model.ResourceProperty_STRING:
		return []interface{}{
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
		}
	case model.ResourceProperty_UUID:
		return []interface{}{
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_BYTES:
		return []interface{}{
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
			RandStringRunes(13),
			RandStringRunes(14),
			RandStringRunes(15),
		}
	case model.ResourceProperty_BOOL:
		return []interface{}{
			randDate().Format(time.RFC3339),
			r.Int63(),
			r.Int31(),
			r.Float32(),
			r.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_DATE:
		return []interface{}{
			randDate().Format(time.RFC3339),
			r.Int63(),
			r.Int31(),
			r.Float32(),
			r.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_FLOAT32:
		return []interface{}{
			randDate().Format(time.RFC3339),
			RandStringRunes(32),
		}
	case model.ResourceProperty_FLOAT64:
		return []interface{}{
			randDate().Format(time.RFC3339),
			RandStringRunes(32),
		}
	case model.ResourceProperty_TIME:
		return []interface{}{
			randDate().Format(time.RFC3339),
			r.Int63(),
			r.Int31(),
			r.Float32(),
			r.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_TIMESTAMP:
		return []interface{}{
			r.Int63(),
			r.Int31(),
			r.Float32(),
			r.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_OBJECT:
		return []interface{}{}
	case model.ResourceProperty_STRUCT:
		return []interface{}{
			map[string]interface{}{
				"field-2": 123,
			},
			map[string]interface{}{
				"field-1": "asd",
				"field-2": 123,
				"field-3": "asd",
			},
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_REFERENCE:
		return []interface{}{
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_ENUM:
		return []interface{}{
			r.Int31(),
			r.Int63(),
			false,
			true,
			r.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_MAP:
		return []interface{}{}
	case model.ResourceProperty_LIST:
		return []interface{}{}
	default:
		panic("Unknown type: " + propertyType.String())
	}
}

func fakeInvalidValue(propertyType model.ResourceProperty_Type) interface{} {
	values := fakeInvalidValues(propertyType)

	if len(values) == 0 {
		return nil
	}

	return values[r.Intn(len(values))]
}
