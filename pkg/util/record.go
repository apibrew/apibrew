package util

import (
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"strings"
)

func ComputeRecordIdFromProperties(resource *model.Resource, record *model.Record) error {
	var idParts []string
	for _, prop := range resource.Properties {
		val := record.Properties[prop.Name]
		if val != nil && prop.Primary {
			typ := types.ByResourcePropertyType(prop.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return err
			}
			if unpacked == nil {
				continue
			}
			idParts = append(idParts, typ.String(unpacked))
		}
	}
	record.Id = strings.Join(idParts, "-")

	return nil
}
