package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"

	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func recordInsert(ctx context.Context, runner QueryRunner, resource *model.Resource, records []*model.Record, ignoreIfExists bool, schema *abs.Schema, history bool) (bool, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	query := fmt.Sprintf("INSERT INTO %s", getTableName(resource.SourceConfig, history))

	cols := prepareResourceRecordCols(resource)

	cols = util.ArrayMap(cols, func(t string) string {
		return fmt.Sprintf("\"%s\"", t)
	})

	query = query + fmt.Sprintf(" (%s)", strings.Join(cols, ","))
	var args []interface{}

	argPlaceHolder := func(val interface{}) string {
		idx := len(args) + 1
		args = append(args, val)
		return fmt.Sprintf("$%d", idx)
	}

	var values []string

	for _, record := range records {
		if !history {
			recordNewId, _ := uuid.NewUUID()
			record.Id = recordNewId.String()
		}
		now := time.Now()
		if !history {
			record.AuditData = &model.AuditData{
				CreatedOn: timestamppb.New(now),
				UpdatedOn: timestamppb.New(now),
				CreatedBy: "test-user",
				UpdatedBy: "",
			}
			record.Version = 1
		}

		var row []string

		if checkHasOwnId(resource) {
			row = append(row, argPlaceHolder(record.Id))
		}

		for _, property := range resource.Properties {
			packedVal := record.Properties[property.Name]
			if packedVal == nil {
				row = append(row, argPlaceHolder(nil))
				continue
			}

			val, serviceError := DbEncode(property, packedVal)
			if serviceError != nil {
				return false, serviceError
			}

			if property.Type == model.ResourcePropertyType_TYPE_REFERENCE {
				row = append(row, resolveReference(val, argPlaceHolder, schema, resource, property))

				continue
			}

			row = append(row, argPlaceHolder(val))

		}

		if !annotations.IsEnabled(resource, annotations.DisableAudit) {
			row = append(row, argPlaceHolder(record.AuditData.CreatedOn.AsTime()))
			row = append(row, argPlaceHolder(record.AuditData.UpdatedOn.AsTime()))
			row = append(row, argPlaceHolder(record.AuditData.CreatedBy))
			row = append(row, argPlaceHolder(record.AuditData.UpdatedBy))
			row = append(row, argPlaceHolder(record.Version))
		}

		values = append(values, fmt.Sprintf("(%s)", strings.Join(row, ",")))
	}

	query = query + " VALUES " + strings.Join(values, ",")

	if ignoreIfExists {
		query = query + " ON CONFLICT DO NOTHING"
	}

	logger.Tracef("SQL: %s", query)

	_, err := runner.ExecContext(ctx, query, args...)

	if len(records) == 1 && records[0].Resource == "author" {
		ra := runner.QueryRow("select count(*) from author")
		var i = new(int32)
		err3 := ra.Scan(&i)

		log.Print(i, err3)
	}

	if err != nil {
		logger.Error("SQL ERROR: ", err)

		ra := runner.QueryRow("select count(*) from author")
		var i = new(int32)
		err3 := ra.Scan(&i)

		log.Print(i, err3)

		return false, handleDbError(ctx, err)
	}

	return true, handleDbError(ctx, err)
}

func resolveReference(val interface{}, argPlaceHolder func(val interface{}) string, schema *abs.Schema, resource *model.Resource, property *model.ResourceProperty) string {
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

func recordUpdate(ctx context.Context, runner QueryRunner, resource *model.Resource, record *model.Record, checkVersion bool, schema *abs.Schema) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	if record.AuditData == nil {
		record.AuditData = &model.AuditData{}
	}

	updateBuilder := sqlbuilder.Update(getTableName(resource.SourceConfig, false))
	updateBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	if checkVersion {
		updateBuilder.Where(updateBuilder.Equal("id", record.Id), updateBuilder.Equal("version", record.Version))
	} else {
		updateBuilder.Where(updateBuilder.Equal("id", record.Id))
	}

	now := time.Now()

	record.AuditData.UpdatedOn = timestamppb.New(now)
	record.AuditData.UpdatedBy = "test-user"

	record.Version++

	for _, property := range resource.Properties {
		packedVal, exists := record.Properties[property.Name]

		if !exists {
			continue
		}

		val, serviceError := DbEncode(property, packedVal)

		if serviceError != nil {
			return serviceError
		}

		if property.Type == model.ResourcePropertyType_TYPE_REFERENCE {
			updateBuilder.SetMore(fmt.Sprintf("\"%s\"=%s", property.Mapping, resolveReference(val, updateBuilder.Var, schema, resource, property)))
		} else {
			updateBuilder.SetMore(updateBuilder.Equal(fmt.Sprintf("\"%s\"", property.Mapping), val))
		}
	}

	updateBuilder.SetMore(updateBuilder.Equal("updated_on", record.AuditData.UpdatedOn.AsTime()))
	updateBuilder.SetMore(updateBuilder.Equal("updated_by", record.AuditData.UpdatedBy))
	updateBuilder.SetMore("version = version + 1")

	sqlQuery, args := updateBuilder.Build()

	logger.Tracef("SQL: %s", sqlQuery)

	result, err := runner.ExecContext(ctx, sqlQuery, args...)

	if err != nil {
		return handleDbError(ctx, err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return handleDbError(ctx, err)
	}

	if affected == 0 {
		return errors.RecordNotFoundError
	}

	return nil
}

func readRecord(ctx context.Context, runner QueryRunner, resource *model.Resource, schema *abs.Schema, id string) (*model.Record, errors.ServiceError) {
	list, total, err := recordList(ctx, runner, abs.ListRecordParams{
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
		UseHistory:        false,
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

func deleteRecords(ctx context.Context, runner QueryRunner, resource *model.Resource, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	deleteBuilder := sqlbuilder.DeleteFrom(getTableName(resource.SourceConfig, false) + " as t")
	deleteBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	if checkHasOwnId(resource) {
		deleteBuilder.Where(deleteBuilder.In("t.id", util.ArrayMapToInterface(ids)...))
	} else {
		idField, err := locatePrimaryKey(resource)

		if err != nil {
			return err
		}

		deleteBuilder.Where(deleteBuilder.In(idField, util.ArrayMapToInterface(ids)...))
	}

	sqlQuery, args := deleteBuilder.Build()

	logger.Tracef("SQL: %s", sqlQuery)

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return handleDbError(ctx, err)
	}

	return nil
}
