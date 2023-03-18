package common

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/helper"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
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
	var args []interface{}

	argPlaceHolder := func(val interface{}) string {
		idx := len(args) + 1
		args = append(args, val)
		return fmt.Sprintf("$%d", idx)
	}

	var values []string

	for _, record := range records {
		if annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
			err := util.ComputeRecordIdFromProperties(resource, record)

			if err != nil {
				return false, p.handleDbError(ctx, err)
			}
		}

		var row []string

		if p.checkHasOwnId(resource) {
			row = append(row, argPlaceHolder(record.Id))
		}

		for _, property := range resource.Properties {
			packedVal, exists := record.Properties[property.Name]

			if exists {
				if packedVal == nil {
					row = append(row, argPlaceHolder(nil))
					continue
				}

				val, serviceError := p.DbEncode(property, packedVal)
				if serviceError != nil {
					return false, serviceError
				}

				if property.Type == model.ResourceProperty_REFERENCE {
					row = append(row, p.resolveReference(val, argPlaceHolder, schema, resource, property))

					continue
				}

				row = append(row, argPlaceHolder(val))
			} else {
				row = append(row, "DEFAULT")
			}
		}

		if !annotations.IsEnabled(resource, annotations.DisableAudit) {
			row = append(row, argPlaceHolder(record.AuditData.CreatedOn.AsTime()))
			row = append(row, argPlaceHolder(record.AuditData.UpdatedOn.AsTime()))
			row = append(row, argPlaceHolder(record.AuditData.CreatedBy))
			row = append(row, argPlaceHolder(record.AuditData.UpdatedBy))
		}

		if !annotations.IsEnabled(resource, annotations.DisableVersion) {
			row = append(row, argPlaceHolder(record.Version))
		}

		values = append(values, fmt.Sprintf("(%s)", strings.Join(row, ",")))
	}

	query = query + " VALUES " + strings.Join(values, ",")

	if ignoreIfExists {
		query = query + " ON CONFLICT DO NOTHING"
	}

	_, err := runner.ExecContext(ctx, query, args...)

	if err != nil {
		logger.Error(err)
	}

	return true, p.handleDbError(ctx, err)
}

func (p *sqlBackend) resolveReference(val interface{}, argPlaceHolder func(val interface{}) string, schema *abs.Schema, resource *model.Resource, property *model.ResourceProperty) string {
	refType := val.(types.ReferenceType)

	if refType["id"] != nil {
		return argPlaceHolder(refType["id"])
	} else {
		var where []string
		for k, v := range refType {
			where = append(where, fmt.Sprintf("%s=%s", k, argPlaceHolder(v)))
		}

		if len(where) == 0 {
			return argPlaceHolder(nil)
		} else {
			referencedResource := schema.ResourceByNamespaceSlashName[resource.Namespace+"/"+property.Reference.ReferencedResource]
			innerSql := fmt.Sprintf("select id from %s where %s", referencedResource.SourceConfig.Entity, strings.Join(where, " AND "))

			return fmt.Sprintf("(%s)", innerSql)
		}
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
			updateBuilder.Where(updateBuilder.Equal("id", record.Id), updateBuilder.Equal("version", record.Version))
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

		val, serviceError := p.DbEncode(property, packedVal)

		if serviceError != nil {
			return serviceError
		}

		if property.Type == model.ResourceProperty_REFERENCE {
			updateBuilder.SetMore(fmt.Sprintf("%s=%s", p.options.Quote(property.Mapping), p.resolveReference(val, updateBuilder.Var, schema, resource, property)))
		} else {
			updateBuilder.SetMore(updateBuilder.Equal(fmt.Sprintf("%s", p.options.Quote(property.Mapping)), val))
		}
	}

	if !annotations.IsEnabled(resource, annotations.DisableAudit) {
		updateBuilder.SetMore(updateBuilder.Equal("updated_on", record.AuditData.UpdatedOn.AsTime()))
		updateBuilder.SetMore(updateBuilder.Equal("updated_by", record.AuditData.UpdatedBy))
	}

	if !annotations.IsEnabled(resource, annotations.DisableVersion) {
		updateBuilder.SetMore("version = version + 1")
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

	if annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		err := util.ComputeRecordIdFromProperties(resource, record)

		if err != nil {
			return p.handleDbError(ctx, err)
		}
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

	if p.checkHasOwnId(resource) {
		deleteBuilder.Where(deleteBuilder.In("t.id", util.ArrayMapToInterface(ids)...))
	} else {
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
			typ := types.ByResourcePropertyType(prop.Type)
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
