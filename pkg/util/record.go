package util

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type PropertyAccessor struct {
	Property *model.ResourceProperty
	Get      func(record *model.Record) interface{}
	Set      func(record *model.Record, val interface{})
}

func IsSameRecord(existing, updated *model.Record) bool {
	for key := range updated.Properties {
		if !proto.Equal(updated.Properties[key], existing.Properties[key]) {
			return false
		}
	}

	return true
}

func RecordIdentifierProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, error) {
	if props, ok := RecordIdentifierPrimaryProperties(resource, properties); ok {
		return props, nil
	}

	if props, ok := RecordIdentifierUniqueProperties(resource, properties); ok {
		return props, nil
	}

	return nil, fmt.Errorf("could not find identifiable properties of %s", resource.Name)
}

func RecordIdentifierPrimaryProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, bool) {
	identifierProps := make(map[string]*structpb.Value)

	for _, prop := range resource.Properties {
		if !prop.Primary {
			continue
		}

		val, ok := properties[prop.Name]

		if !ok {
			return nil, false
		}

		typ := types.ByResourcePropertyType(prop.Type)

		unpacked, err := typ.UnPack(val)

		if err != nil {
			log.Error(err)
			return nil, false
		}

		if typ.Equals(unpacked, typ.Default()) {
			return nil, false
		}

		identifierProps[prop.Name] = val
	}

	return identifierProps, true
}

func RecordIdentifierUniqueProperties(resource *model.Resource, properties map[string]*structpb.Value) (map[string]*structpb.Value, bool) {
	for _, prop := range resource.Properties {
		identifierProps := make(map[string]*structpb.Value)
		if !prop.Unique {
			continue
		}

		val, ok := properties[prop.Name]

		if !ok {
			continue
		}

		typ := types.ByResourcePropertyType(prop.Type)

		unpacked, err := typ.UnPack(val)

		if err != nil {
			log.Error(err)
			return nil, false
		}

		if typ.Equals(unpacked, typ.Default()) {
			continue
		}

		identifierProps[prop.Name] = val

		return identifierProps, true
	}

	propMap := GetNamedMap(resource.Properties)

	for _, index := range resource.Indexes {
		if index.Unique {
			var valid = true

			identifierProps := make(map[string]*structpb.Value)

			for _, indexProp := range index.Properties {
				prop := propMap[indexProp.Name]

				val, ok := properties[prop.Name]

				if !ok {
					valid = false
					break
				}

				typ := types.ByResourcePropertyType(prop.Type)

				unpacked, err := typ.UnPack(val)

				if err != nil {
					log.Error(err)

					valid = false
					break
				}

				if typ.Equals(unpacked, typ.Default()) {
					valid = false
					break
				}

				identifierProps[prop.Name] = val
			}

			if valid {
				return identifierProps, true
			}
		}
	}

	return nil, false
}

func RecordMatchIdentifiableProperties(resource *model.Resource, record *model.Record, properties map[string]*structpb.Value) (bool, error) {
	idProps, err := RecordIdentifierProperties(resource, properties)

	if err != nil {
		return false, err
	}

	for key, val := range idProps {
		if !proto.Equal(record.Properties[key], val) {
			return false, nil
		}
	}

	return true, nil
}

func RecordPropertyAccessorByPath(properties map[string]*structpb.Value, path string) (getter func() *structpb.Value, setter func(val *structpb.Value)) {
	path = strings.ReplaceAll(path, "[]", ".[]")

	if strings.HasPrefix(path, "$.") {
		path = path[2:]
	}

	parts := strings.Split(path, ".")

	if len(parts) == 0 {
		return nil, nil
	}

	if len(parts) == 1 {
		getter = func() *structpb.Value {
			return properties[parts[0]]
		}

		return getter, func(val *structpb.Value) {
			properties[parts[0]] = val
		}
	}

	left := parts[0]
	next := parts[1]
	right := strings.Join(parts[1:], ".")

	rightProperties := properties[left]

	if rightProperties == nil {
		return nil, func(val *structpb.Value) {
			properties[left] = val
		}
	}

	if next == "[]" {
		if right != "[]" {
			panic("invalid path; array accessor must be at the end")
		}
		getter = func() *structpb.Value {
			return rightProperties
		}

		return getter, func(val *structpb.Value) {
			properties[left] = val
		}
	} else {
		return RecordPropertyAccessorByPath(rightProperties.GetStructValue().Fields, right)
	}
}
