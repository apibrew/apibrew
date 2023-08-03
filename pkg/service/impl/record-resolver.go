package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type recordResolver struct {
	recordService   service.RecordService
	resourceService service.ResourceService
	resource        *model.Resource
	records         []*model.Record
	paths           []string
}

func (r *recordResolver) resolveReferences(ctx context.Context) errors.ServiceError {
	var pathMap = make(map[string]bool)

	for _, path := range r.paths {
		pathMap[path] = true
	}

	var recordValues = util.ArrayToMap(r.records, func(t *model.Record) string {
		return t.GetId()
	}, func(t *model.Record) *structpb.Value {
		return structpb.NewStructValue(&structpb.Struct{Fields: t.Properties})
	})

	return r._recordListWalkOperator(ctx, "$", r.resource.Properties, recordValues, pathMap)

}

func (r *recordResolver) _recordListWalkOperator(ctx context.Context, path string, properties []*model.ResourceProperty, recordValueMap map[string]*structpb.Value, pathsToOperate map[string]bool) errors.ServiceError {
	for _, prop := range properties {
		var newPath = path + "." + prop.Name

		if prop.Type == model.ResourceProperty_LIST || prop.Type == model.ResourceProperty_MAP {
			newPath = newPath + "[]"
			theProp := prop
			prop = prop.Item
			prop.Name = theProp.Name
		}

		if !pathsToOperate[newPath] {
			continue
		}

		var subValues = make(map[string]*structpb.Value, len(recordValueMap))

		for recordId, value := range recordValueMap {
			valueSt := value.GetStructValue()
			subValues[recordId] = valueSt.Fields[prop.Name]
		}

		if len(recordValueMap) == 0 {
			continue
		}

		var pathToOperateNextReference []string
		var pathToOperateNextReferenceMap = make(map[string]bool)

		for pathToOperate := range pathsToOperate {
			if !strings.HasPrefix(pathToOperate, newPath) {
				continue
			}
			rightPath := pathToOperate[len(newPath):]
			if rightPath == "" {
				continue
			}
			pathToOperateNextReference = append(pathToOperateNextReference, "$"+rightPath)
			pathToOperateNextReferenceMap["$"+rightPath] = true
		}

		var referenceRecords []*model.Record

		if prop.Type == model.ResourceProperty_REFERENCE {
			if prop.BackReference != nil {
				var ids []string
				for _, record := range r.records {
					ids = append(ids, record.Id)
				}

				// get referenced records
				list, _, err := r.recordService.List(ctx, service.RecordListParams{
					Namespace: prop.Reference.Namespace,
					Resource:  prop.Reference.Resource,
					Query: util.QueryInExpression(prop.BackReference.Property, structpb.NewListValue(&structpb.ListValue{
						Values: util.ArrayMap(ids, func(t string) *structpb.Value {
							return structpb.NewStringValue(t)
						}),
					})),
					ResolveReferences: []string{},
				})

				if err != nil {
					return err
				}

				for id := range subValues {
					subValues[id] = structpb.NewListValue(&structpb.ListValue{})
					recordValueMap[id].GetStructValue().Fields[prop.Name] = subValues[id]
				}

				for _, record := range list {
					actualReference := record.Properties[prop.BackReference.Property].GetStructValue()
					var id = actualReference.Fields["id"].GetStringValue()

					subValues[id].GetListValue().Values = append(subValues[id].GetListValue().Values, structpb.NewStructValue(&structpb.Struct{Fields: record.Properties}))
					referenceRecords = append(referenceRecords, record)
				}
			} else {
				var referencedResource = r.resourceService.LocateResourceByReference(r.resource, prop.Reference)
				var query *model.BooleanExpression

				for _, referenceValue := range subValues {
					if referenceValue == nil || referenceValue.GetStructValue().Fields == nil {
						continue
					}
					if referenceValue.GetListValue() != nil {
						for _, value := range referenceValue.GetListValue().Values {
							subQuery, err := util.RecordIdentifierQuery(referencedResource, value.GetStructValue().Fields)

							if err != nil {
								return errors.LogicalError.WithDetails(err.Error())
							}

							if query == nil {
								query = subQuery
							} else {
								query = util.QueryOrExpression(query, subQuery)
							}
						}
					} else {
						subQuery, err := util.RecordIdentifierQuery(referencedResource, referenceValue.GetStructValue().Fields)

						if err != nil {
							return errors.LogicalError.WithDetails(err.Error())
						}

						if query == nil {
							query = subQuery
						} else {
							query = util.QueryOrExpression(query, subQuery)
						}
					}
				}

				list, _, err := r.recordService.List(ctx, service.RecordListParams{
					Namespace:         prop.Reference.Namespace,
					Resource:          prop.Reference.Resource,
					Query:             query,
					ResolveReferences: []string{},
				})

				if err != nil {
					return err
				}

				for recordId, referenceValue := range subValues {
					if referenceValue == nil || referenceValue.GetStructValue().Fields == nil {
						continue
					}

					if referenceValue.GetListValue() != nil {
						subValues[recordId] = structpb.NewListValue(&structpb.ListValue{})
						recordValueMap[recordId].GetStructValue().Fields[prop.Name] = subValues[recordId]

						for _, item := range list {
							matches, err := util.RecordMatchIdentifiableProperties(referencedResource, item, referenceValue.GetStructValue().Fields)

							if err != nil {
								return errors.LogicalError.WithDetails(err.Error())
							}

							if matches {
								subValues[recordId].GetListValue().Values = append(subValues[recordId].GetListValue().Values, structpb.NewStructValue(&structpb.Struct{Fields: item.Properties}))
							}
						}
					} else {
						for _, item := range list {
							matches, err := util.RecordMatchIdentifiableProperties(referencedResource, item, referenceValue.GetStructValue().Fields)

							if err != nil {
								return errors.LogicalError.WithDetails(err.Error())
							}

							if matches {
								subValues[recordId] = structpb.NewStructValue(&structpb.Struct{Fields: item.Properties})
								recordValueMap[recordId].GetStructValue().Fields[prop.Name] = subValues[recordId]
								break
							}
						}
					}
				}
			}
		}

		if len(pathToOperateNextReference) > 0 {
			if prop.Type == model.ResourceProperty_STRUCT {
				err := r._recordListWalkOperator(ctx, newPath, prop.Properties, subValues, pathToOperateNextReferenceMap)

				if err != nil {
					return err
				}
			}

			if prop.Type == model.ResourceProperty_REFERENCE {
				var referencedResource = r.resourceService.LocateResourceByReference(r.resource, prop.Reference)

				subRefRecordResolver := &recordResolver{
					recordService:   r.recordService,
					resourceService: r.resourceService,
					resource:        referencedResource,
					records:         referenceRecords,
					paths:           pathToOperateNextReference,
				}

				err := subRefRecordResolver.resolveReferences(ctx)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
