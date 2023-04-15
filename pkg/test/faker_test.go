package test

import (
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/test/setup"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
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
		return []interface{}{rand.Int31()}
	case model.ResourceProperty_INT64:
		return []interface{}{rand.Int63()}
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
			rand.Float32(),
		}
	case model.ResourceProperty_FLOAT64:
		return []interface{}{
			rand.Float64(),
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
				"a123": rand.Int63(),
				"a124": rand.Float64(),
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

	return values[rand.Intn(len(values))]
}

func fakeInvalidValues(propertyType model.ResourceProperty_Type) []interface{} {
	switch propertyType {
	case model.ResourceProperty_INT32:
		return []interface{}{
			rand.Int63(),
			"1234",
			false,
			true,
			rand.Float64(),
		}
	case model.ResourceProperty_INT64:
		return []interface{}{
			"1234",
			false,
			true,
			rand.Float64(),
		}
	case model.ResourceProperty_STRING:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
		}
	case model.ResourceProperty_UUID:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_BYTES:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
			RandStringRunes(13),
			RandStringRunes(14),
			RandStringRunes(15),
		}
	case model.ResourceProperty_BOOL:
		return []interface{}{
			randDate().Format(time.RFC3339),
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_DATE:
		return []interface{}{
			randDate().Format(time.RFC3339),
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
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
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourceProperty_TIMESTAMP:
		return []interface{}{
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
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
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_REFERENCE:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
			RandStringRunes(12),
		}
	case model.ResourceProperty_ENUM:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
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

	return values[rand.Intn(len(values))]
}
