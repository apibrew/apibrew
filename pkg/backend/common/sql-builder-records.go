package common

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	helper2 "github.com/tislib/data-handler/pkg/helper"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func (p *sqlBackend) recordInsert(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, records []*model.Record, ignoreIfExists bool, schema *abs.Schema) (bool, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	query := fmt.Sprintf("INSERT INTO %s", p.getFullTableName(resource.SourceConfig))

	cols := p.prepareResourceRecordCols(resource)

	query = query + fmt.Sprintf(" (%s)", strings.Join(cols, ","))
	var args = sqlbuilder.Args{Flavor: p.options.GetFlavor()}

	var values []string

	for _, record := range records {
		var row []string

		for _, property := range resource.Properties {
			packedVal, exists := record.Properties[property.Name]

			if exists || !annotations.IsEnabled(property, annotations.Identity) {
				if packedVal == nil {
					row = append(row, args.Add(nil))
					continue
				}

				val, serviceError := p.options.DbEncode(property, packedVal)
				if serviceError != nil {
					return false, serviceError
				}

				if property.Type == model.ResourceProperty_REFERENCE {
					referencedResource := schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+property.Reference.ReferencedResource]
					item, err := p.resolveReference(packedVal.GetStructValue().Fields, args.Add, referencedResource)

					if err != nil {
						return false, errors.LogicalError.WithDetails(err.Error())
					}

					row = append(row, item)

					continue
				}

				row = append(row, args.Add(val))
			} else {
				row = append(row, "DEFAULT")
			}
		}

		values = append(values, fmt.Sprintf("(%s)", strings.Join(row, ",")))
	}

	query = query + " VALUES " + strings.Join(values, ",")

	if ignoreIfExists {
		query = query + " ON CONFLICT DO NOTHING"
	}

	q, a := args.Compile(query)
	_, err := runner.ExecContext(ctx, q, a...)

	if err != nil {
		logger.Error(err)
	}

	return true, p.handleDbError(ctx, err)
}

func (p *sqlBackend) resolveReference(properties map[string]*structpb.Value, argPlaceHolder func(val interface{}) string, referencedResource *model.Resource) (string, error) {
	identifierProps, err := util.RecordIdentifierProperties(referencedResource, properties)
	namedProps := util.GetNamedMap(referencedResource.Properties)
	if util.HasResourceSinglePrimaryProp(referencedResource) {
		idProp := util.GetResourceSinglePrimaryProp(referencedResource)

		if val, ok := identifierProps[idProp.Name]; ok {
			typ := types.ByResourcePropertyType(idProp.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return "", err
			}

			return argPlaceHolder(unpacked), nil
		}
	}

	if err != nil {
		return "", err
	}

	var where []string
	for k, v := range identifierProps {
		typ := types.ByResourcePropertyType(namedProps[k].Type)

		if typ == types.ReferenceType { // skip reference checking for now, it is not implemented yet
			continue
		}

		unpacked, err := typ.UnPack(v)
		if err != nil {
			return "", err
		}
		where = append(where, fmt.Sprintf("%s=%s", k, argPlaceHolder(unpacked)))
	}

	if len(where) == 0 {
		return argPlaceHolder(nil), nil
	} else {
		innerSql := fmt.Sprintf("select id from %s where %s", referencedResource.SourceConfig.Entity, strings.Join(where, " AND "))

		return fmt.Sprintf("(%s)", innerSql), nil
	}
}

func (p *sqlBackend) recordUpdate(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, record *model.Record, checkVersion bool, schema *abs.Schema) errors.ServiceError {
	updateBuilder := sqlbuilder.Update(p.getFullTableName(resource.SourceConfig))
	updateBuilder.SetFlavor(p.options.GetFlavor())

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		checkVersion = false
	}

	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		if checkVersion {
			ah := helper2.RecordSpecialColumnHelper{
				Resource: resource,
				Record:   record,
			}
			updateBuilder.Where(updateBuilder.Equal("id", record.Id), updateBuilder.Equal("version", ah.GetVersion()))
		} else {
			updateBuilder.Where(updateBuilder.Equal("id", record.Id))
		}
	} else {
		sqlPart, err := p.createRecordIdMatchQuery(ctx, resource, record, updateBuilder.Var)
		if err != nil {
			return err
		}
		updateBuilder.Where(sqlPart)
	}

	for _, property := range resource.Properties {
		packedVal, exists := record.Properties[property.Name]

		if !exists {
			continue
		}

		if property.Immutable {
			continue
		}

		if property.Name == "version" && annotations.IsEnabled(property, annotations.SpecialProperty) && !annotations.IsEnabled(resource, annotations.DisableVersion) {
			updateBuilder.SetMore("version = version + 1")
			continue
		}

		val, serviceError := p.options.DbEncode(property, packedVal)

		if serviceError != nil {
			return serviceError
		}

		if property.Type == model.ResourceProperty_REFERENCE {
			referencedResource := schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+property.Reference.ReferencedResource]
			item, err := p.resolveReference(packedVal.GetStructValue().Fields, updateBuilder.Var, referencedResource)

			if err != nil {
				return errors.LogicalError.WithDetails(err.Error())
			}
			updateBuilder.SetMore(fmt.Sprintf("%s=%s", p.options.Quote(property.Mapping), item))
		} else {
			updateBuilder.SetMore(updateBuilder.Equal(p.options.Quote(property.Mapping), val))
		}
	}

	sqlQuery, args := updateBuilder.Build()

	result, err := runner.ExecContext(ctx, sqlQuery, args...)

	if err != nil {
		return p.handleDbError(ctx, err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return p.handleDbError(ctx, err)
	}

	if affected == 0 {
		return errors.RecordNotFoundError.WithDetails("No records are affected by update")
	}

	return nil
}

func (p *sqlBackend) readRecord(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	list, total, err := p.recordList(ctx, runner, abs.ListRecordParams{
		Resource: resource,
		Query: &model.BooleanExpression{
			Expression: &model.BooleanExpression_Equal{
				Equal: &model.PairExpression{
					Left:  &model.Expression{Expression: &model.Expression_Property{Property: "id"}},
					Right: &model.Expression{Expression: &model.Expression_Value{Value: structpb.NewStringValue(id)}},
				},
			},
		},
		Limit:             1,
		Offset:            0,
		ResolveReferences: []string{"*"},
		Schema:            schema,
	})

	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.RecordNotFoundError.WithDetails("with id: " + id)
	}

	return list[0], nil
}

func (p *sqlBackend) deleteRecords(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, ids []string) errors.ServiceError {
	deleteBuilder := sqlbuilder.DeleteFrom(p.getFullTableName(resource.SourceConfig) + " as t")
	deleteBuilder.SetFlavor(p.options.GetFlavor())

	var primaryFound = false
	for _, prop := range resource.Properties {
		if prop.Primary {
			deleteBuilder.Where(deleteBuilder.In(prop.Mapping, util.ArrayMapToInterface(ids)...))
			primaryFound = true
			break
		}
	}

	if !primaryFound {
		return errors.LogicalError.WithDetails("Delete operation cannot be executed without id")
	}

	sqlQuery, args := deleteBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return p.handleDbError(ctx, err)
	}

	return nil
}

func (p *sqlBackend) createRecordIdMatchQuery(ctx context.Context, resource *model.Resource, record *model.Record, varFn func(value interface{}) string) (string, errors.ServiceError) {
	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		return fmt.Sprintf("id=%s", varFn(record.Id)), nil
	}

	for _, prop := range resource.Properties {
		val := record.Properties[prop.Name]
		if val != nil && prop.Primary {
			typ := p.options.TypeModifier(prop.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return "", p.handleDbError(ctx, err)
			}
			if unpacked == nil {
				continue
			}

			return fmt.Sprintf("%s=%s", p.options.Quote(prop.Mapping), varFn(unpacked)), nil
		}
	}

	return "", errors.LogicalError.WithDetails("No Primary key exists")
}
