package common

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	helper2 "github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

func (p *sqlBackend) recordInsert(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, records []*model.Record, ignoreIfExists bool, schema *abs.Schema) errors.ServiceError {
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

			if helper.IsPropertyOmitted(property) {
				continue
			}

			if exists || !annotations.IsEnabled(property, annotations.Identity) {
				if packedVal == nil {
					row = append(row, args.Add(nil))
					continue
				}

				val, serviceError := p.options.DbEncode(property, packedVal)
				if serviceError != nil {
					return serviceError
				}

				if property.Type == model.ResourceProperty_REFERENCE {
					if property.BackReference != nil {
						continue
					}
					referenceNamespace := property.Reference.Namespace
					if referenceNamespace == "" {
						referenceNamespace = resource.Namespace
					}
					referencedResource := schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]
					item, err := p.resolveReference(packedVal.GetStructValue().Fields, args.Add, referencedResource)

					if err != nil {
						return err
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

	return p.handleDbError(ctx, err)
}

func (p *sqlBackend) resolveReference(properties map[string]*structpb.Value, argPlaceHolder func(val interface{}) string, referencedResource *model.Resource) (string, errors.ServiceError) {
	identifierProps, err := util.RecordIdentifierProperties(referencedResource, properties)

	if err != nil {
		return "", errors.LogicalError.WithDetails(err.Error())
	}

	namedProps := util.GetNamedMap(referencedResource.Properties)
	if util.HasResourceSinglePrimaryProp(referencedResource) {
		idProp := util.GetResourceSinglePrimaryProp(referencedResource)

		if val, ok := identifierProps[idProp.Name]; ok {
			typ := types.ByResourcePropertyType(idProp.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return "", errors.LogicalError.WithDetails(err.Error())
			}

			return argPlaceHolder(unpacked), nil
		}
	}

	var where []string
	for k, v := range identifierProps {
		typ := types.ByResourcePropertyType(namedProps[k].Type)

		if typ == types.ReferenceType { // skip reference checking for now, it is not implemented yet
			continue
		}

		unpacked, err := typ.UnPack(v)

		if err != nil {
			return "", errors.LogicalError.WithDetails(err.Error())
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

func (p *sqlBackend) createRecordIdMatchQuery(resource *model.Resource, record *model.Record, argPlaceHolder func(value interface{}) string) (string, errors.ServiceError) {
	identifierProps, err := util.RecordIdentifierProperties(resource, record.Properties)
	namedProps := util.GetNamedMap(resource.Properties)
	if util.HasResourceSinglePrimaryProp(resource) {
		idProp := util.GetResourceSinglePrimaryProp(resource)

		if val, ok := identifierProps[idProp.Name]; ok {
			typ := types.ByResourcePropertyType(idProp.Type)
			unpacked, err := typ.UnPack(val)
			if err != nil {
				return "", errors.LogicalError.WithDetails(err.Error())
			}

			return fmt.Sprintf("%s=%s", idProp.Mapping, argPlaceHolder(unpacked)), nil
		}
	}

	if err != nil {
		return "", errors.LogicalError.WithDetails(err.Error())
	}

	var where []string
	for k, v := range identifierProps {
		typ := types.ByResourcePropertyType(namedProps[k].Type)

		if typ == types.ReferenceType { // skip reference checking for now, it is not implemented yet
			continue
		}

		unpacked, err := typ.UnPack(v)
		if err != nil {
			return "", errors.LogicalError.WithDetails(err.Error())
		}
		where = append(where, fmt.Sprintf("%s=%s", k, argPlaceHolder(unpacked)))
	}

	if len(where) == 0 {
		return argPlaceHolder(nil), nil
	} else {
		return strings.Join(where, " AND "), nil
	}
}

func (p *sqlBackend) recordUpdate(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, record *model.Record, checkVersion bool, schema *abs.Schema) errors.ServiceError {
	updateBuilder := sqlbuilder.Update(p.getFullTableName(resource.SourceConfig))
	updateBuilder.SetFlavor(p.options.GetFlavor())

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		checkVersion = false
	}

	sqlPart, err := p.createRecordIdMatchQuery(resource, record, updateBuilder.Var)
	if err != nil {
		return err
	}
	if checkVersion {
		ah := helper2.RecordSpecialColumnHelper{
			Resource: resource,
			Record:   record,
		}
		sqlPart += " AND " + updateBuilder.Equal("version", ah.GetVersion())
	}
	updateBuilder.Where(sqlPart)

	for _, property := range resource.Properties {
		if helper.IsPropertyOmitted(property) {
			continue
		}

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
			referenceNamespace := property.Reference.Namespace
			if referenceNamespace == "" {
				referenceNamespace = resource.Namespace
			}
			referencedResource := schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]
			item, err := p.resolveReference(packedVal.GetStructValue().Fields, updateBuilder.Var, referencedResource)

			if err != nil {
				return err
			}
			updateBuilder.SetMore(fmt.Sprintf("%s=%s", p.options.Quote(property.Mapping), item))
		} else {
			updateBuilder.SetMore(updateBuilder.Equal(p.options.Quote(property.Mapping), val))
		}
	}

	sqlQuery, args := updateBuilder.Build()

	result, sqlErr := runner.ExecContext(ctx, sqlQuery, args...)

	if sqlErr != nil {
		return p.handleDbError(ctx, err)
	}

	affected, sqlErr := result.RowsAffected()

	if sqlErr != nil {
		return p.handleDbError(ctx, err)
	}

	if affected == 0 {
		return errors.RecordNotFoundError.WithDetails("No records are affected by update")
	}

	return nil
}

func (p *sqlBackend) readRecord(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	list, total, err := p.recordList(ctx, runner, resource, abs.ListRecordParams{
		Query: &model.BooleanExpression{
			Expression: &model.BooleanExpression_Equal{
				Equal: &model.PairExpression{
					Left:  &model.Expression{Expression: &model.Expression_Property{Property: "id"}},
					Right: &model.Expression{Expression: &model.Expression_Value{Value: structpb.NewStringValue(id)}},
				},
			},
		},
		Limit:  1,
		Offset: 0,
	}, nil)

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
