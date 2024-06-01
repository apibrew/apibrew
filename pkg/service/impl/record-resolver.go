package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type recordResolver struct {
	recordService   service.RecordService
	resourceService service.ResourceService
	resource        *model.Resource
	records         []abs.RecordLike
	paths           []string
}

func (r *recordResolver) resolveReferences(ctx context.Context) error {
	var pathMap = make(map[string]bool)

	for _, path := range r.paths {
		pathMap[path] = true
	}

	var recordValues = util.ArrayToMap(r.records, func(record abs.RecordLike) string {
		return util.GetRecordId(record)
	}, func(t abs.RecordLike) abs.RecordLike {
		return t
	})

	err := r._recordListWalkOperator(ctx, "$", r.resource.Properties, recordValues, pathMap)

	if err != nil {
		return err
	}

	return nil
}

func (r *recordResolver) _recordListWalkOperator(ctx context.Context, path string, properties []*model.ResourceProperty, recordValueMap map[string]abs.RecordLike, pathsToOperate map[string]bool) error {
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
			subValues[recordId] = value.GetStructProperty(prop.Name)
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

		var referenceRecords []abs.RecordLike

		if prop.Type == model.ResourceProperty_REFERENCE {
			if prop.BackReference != nil {
				var ids []string
				for _, record := range r.records {
					recordId := util.GetRecordId(record)
					if recordId == "" {
						continue
					}
					ids = append(ids, recordId)
				}

				if len(ids) == 0 {
					log.Debug("No records to resolve references")
					continue
				}

				// get referenced records
				list, _, err := r.recordService.List(ctx, service.RecordListParams{
					Limit:     10000,
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
					recordValueMap[id].SetStructProperty(prop.Name, subValues[id])
				}

				for _, record := range list {
					actualReference := record.GetStructProperty(prop.BackReference.Property).GetStructValue()
					var id = actualReference.Fields["id"].GetStringValue()

					subValues[id].GetListValue().Values = append(subValues[id].GetListValue().Values, structpb.NewStructValue(record.ToStruct()))
					referenceRecords = append(referenceRecords, record)
				}
			} else {
				var referencedResource = r.resourceService.LocateResourceByReference(r.resource, prop.Reference)
				var query *model.BooleanExpression

				if referencedResource == nil {
					return errors.LogicalError.WithDetails(fmt.Sprintf("Could not locate resource %s.%s", prop.Reference.Namespace, prop.Reference.Resource))
				}

				for _, referenceValue := range subValues {
					if referenceValue == nil {
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
					} else if referenceValue.GetStructValue() != nil {
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
					Limit:             10000,
					Namespace:         prop.Reference.Namespace,
					Resource:          prop.Reference.Resource,
					Query:             query,
					ResolveReferences: []string{},
				})

				if err != nil {
					return err
				}

				for recordId, referenceValue := range subValues {
					if referenceValue == nil {
						continue
					}

					if referenceValue.GetListValue() != nil {
						subValues[recordId] = structpb.NewListValue(&structpb.ListValue{})
						recordValueMap[recordId].SetStructProperty(prop.Name, subValues[recordId])

						for _, subRefValue := range referenceValue.GetListValue().Values {
							for _, item := range list {
								matches, err := util.RecordMatchIdentifiableProperties(referencedResource, item, subRefValue.GetStructValue().Fields)

								if err != nil {
									return errors.LogicalError.WithDetails(err.Error())
								}

								if matches {
									subValues[recordId].GetListValue().Values = append(subValues[recordId].GetListValue().Values, structpb.NewStructValue(item.ToStruct()))
									referenceRecords = append(referenceRecords, item)
								}
							}
						}
					} else if referenceValue.GetStructValue() != nil {
						for _, item := range list {
							matches, err := util.RecordMatchIdentifiableProperties(referencedResource, item, referenceValue.GetStructValue().Fields)

							if err != nil {
								return errors.LogicalError.WithDetails(err.Error())
							}

							if matches {
								subValues[recordId] = structpb.NewStructValue(item.ToStruct())
								recordValueMap[recordId].SetStructProperty(prop.Name, subValues[recordId])
								referenceRecords = append(referenceRecords, item)
								break
							}
						}
					}
				}
			}
		}

		if len(pathToOperateNextReference) > 0 {
			if prop.Type == model.ResourceProperty_STRUCT {
				err := r._recordListWalkCheckOperator(ctx, newPath, r.getTypeProperties(*prop.TypeRef), subValues, pathToOperateNextReferenceMap)

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

func (r *recordResolver) getTypeProperties(typeRef string) []*model.ResourceProperty {
	var properties []*model.ResourceProperty

	for _, typ := range r.resource.Types {
		if typ.Name == typeRef {
			properties = typ.Properties
			break
		}
	}
	return properties
}

func (r *recordResolver) _recordListWalkCheckOperator(ctx context.Context, path string, properties []*model.ResourceProperty, recordValueMap map[string]*structpb.Value, pathsToOperate map[string]bool) error {
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
			if valueSt.Fields[prop.Name] == nil {
				continue
			}
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

		var referenceRecords []abs.RecordLike

		if prop.Type == model.ResourceProperty_REFERENCE {
			if prop.BackReference != nil {
				continue
			}

			var referencedResource = r.resourceService.LocateResourceByReference(r.resource, prop.Reference)
			var query *model.BooleanExpression

			for _, referenceValue := range subValues {
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
				} else if referenceValue.GetStructValue() != nil {
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

			_, count, err := r.recordService.List(ctx, service.RecordListParams{
				Limit:             10000,
				Namespace:         prop.Reference.Namespace,
				Resource:          prop.Reference.Resource,
				Query:             query,
				ResolveReferences: []string{},
			})

			if err != nil {
				return err
			}

			if count < uint32(len(subValues)) {
				return errors.ReferenceViolation.WithDetails(
					fmt.Sprintf("Resolved reference count does not match: required %d reference but found %d reference", count, len(subValues)))
			}
		}

		if len(pathToOperateNextReference) > 0 {
			if prop.Type == model.ResourceProperty_STRUCT {
				log.Println("Not implemented")
				//err := r._recordListWalkOperator(ctx, newPath, r.getTypeProperties(*prop.TypeRef), subValues, pathToOperateNextReferenceMap)
				//err := nil
				//
				//if err != nil {
				//	return err
				//}
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

func (r *recordResolver) checkReferences(ctx context.Context) error {
	var pathMap = make(map[string]bool)

	for _, path := range r.paths {
		pathMap[path] = true
	}

	var recordValues = util.ArrayToMap(r.records, func(record abs.RecordLike) string {
		return util.GetRecordId(record)
	}, func(t abs.RecordLike) *structpb.Value {
		return structpb.NewStructValue(t.ToStruct())
	})

	return r._recordListWalkCheckOperator(ctx, "$", r.resource.Properties, recordValues, pathMap)
}
