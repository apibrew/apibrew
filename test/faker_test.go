package test

import (
	"data-handler/model"
	"github.com/google/uuid"
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
			DataSource: dhTest.Id,
			Catalog:    "public",
			Entity:     name,
		},
		Properties: properties,
		Flags:      &model.ResourceFlags{},
	}
}

func fakeValidValues(propertyType model.ResourcePropertyType) []interface{} {
	switch propertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		return []interface{}{rand.Int31()}
	case model.ResourcePropertyType_TYPE_INT64:
		return []interface{}{rand.Int63()}
	case model.ResourcePropertyType_TYPE_STRING:
		return []interface{}{
			RandStringRunes(16),
		}
	case model.ResourcePropertyType_TYPE_UUID:
		randUUid, _ := uuid.NewUUID()
		return []interface{}{
			randUUid.String(),
		}
	case model.ResourcePropertyType_TYPE_BYTES:
		return []interface{}{
			[]byte(RandStringRunes(16)),
		}
	case model.ResourcePropertyType_TYPE_BOOL:
		return []interface{}{
			false, true,
		}
	case model.ResourcePropertyType_TYPE_DATE:
		return []interface{}{
			randDate().Format("2006-01-02"),
		}
	case model.ResourcePropertyType_TYPE_FLOAT32:
		return []interface{}{
			rand.Float32(),
		}
	case model.ResourcePropertyType_TYPE_FLOAT64:
		return []interface{}{
			rand.Float64(),
		}
	case model.ResourcePropertyType_TYPE_TIME:
		return []interface{}{
			randDate().Format("15:04:05"),
		}
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return []interface{}{
			randDate().Format(time.RFC3339),
		}
	case model.ResourcePropertyType_TYPE_OBJECT:
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
	default:
		panic("Unknown type: " + propertyType.String())
	}
}

func fakeValidValue(propertyType model.ResourcePropertyType) interface{} {
	values := fakeValidValues(propertyType)

	return values[rand.Intn(len(values))]
}

func fakeInvalidValues(propertyType model.ResourcePropertyType) []interface{} {
	switch propertyType {
	case model.ResourcePropertyType_TYPE_INT32:
		return []interface{}{
			rand.Int63(),
			"1234",
			false,
			true,
			rand.Float64(),
		}
	case model.ResourcePropertyType_TYPE_INT64:
		return []interface{}{
			"1234",
			false,
			true,
			rand.Float64(),
		}
	case model.ResourcePropertyType_TYPE_STRING:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
		}
	case model.ResourcePropertyType_TYPE_UUID:
		return []interface{}{
			rand.Int31(),
			rand.Int63(),
			false,
			true,
			rand.Float64(),
			RandStringRunes(12),
		}
	case model.ResourcePropertyType_TYPE_BYTES:
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
	case model.ResourcePropertyType_TYPE_BOOL:
		return []interface{}{
			randDate().Format(time.RFC3339),
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_DATE:
		return []interface{}{
			randDate().Format(time.RFC3339),
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_FLOAT32:
		return []interface{}{
			randDate().Format(time.RFC3339),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_FLOAT64:
		return []interface{}{
			randDate().Format(time.RFC3339),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_TIME:
		return []interface{}{
			randDate().Format(time.RFC3339),
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_TIMESTAMP:
		return []interface{}{
			rand.Int63(),
			rand.Int31(),
			rand.Float32(),
			rand.Float64(),
			RandStringRunes(32),
		}
	case model.ResourcePropertyType_TYPE_OBJECT:
		return []interface{}{}
	default:
		panic("Unknown type: " + propertyType.String())
	}
}

func fakeInvalidValue(propertyType model.ResourcePropertyType) interface{} {
	values := fakeInvalidValues(propertyType)

	if len(values) == 0 {
		return nil
	}

	return values[rand.Intn(len(values))]
}
