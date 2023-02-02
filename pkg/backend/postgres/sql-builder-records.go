package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations2 "github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func recordInsert(ctx context.Context, runner QueryRunner, resource *model.Resource, records []*model.Record, ignoreIfExists bool, history bool) (bool, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	insertBuilder := sqlbuilder.InsertInto(getTableName(resource.SourceConfig, history))
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	cols := prepareResourceRecordCols(resource)

	cols = util.ArrayMap(cols, func(t string) string {
		return fmt.Sprintf("\"%s\"", t)
	})

	insertBuilder.Cols(cols...)

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

		var row []interface{}

		if checkHasOwnId(resource) {
			row = append(row, record.Id)
		}

		for _, property := range resource.Properties {
			if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
				packedVal := record.Properties[property.Name]
				if packedVal == nil {
					row = append(row, nil)
					continue
				}
				propertyType := types.ByResourcePropertyType(property.Type)

				var val interface{}

				if property.Type == model.ResourcePropertyType_TYPE_OBJECT {
					var err2 error
					val, err2 = json.Marshal(packedVal.AsInterface())

					if err2 != nil {
						return false, errors.InternalError.WithDetails(err2.Error())
					}
					val = string(val.([]byte))
				} else {
					var err error
					val, err = propertyType.UnPack(packedVal)

					if err != nil {
						return false, errors.RecordValidationError.WithDetails(err.Error())
					}
				}
				row = append(row, val)
			}
		}

		if !annotations2.IsEnabled(resource, annotations2.DisableAudit) {
			row = append(row, record.AuditData.CreatedOn.AsTime())
			row = append(row, record.AuditData.UpdatedOn.AsTime())
			row = append(row, record.AuditData.CreatedBy)
			row = append(row, record.AuditData.UpdatedBy)
			row = append(row, record.Version)
		}

		insertBuilder.Values(row...)
	}

	if ignoreIfExists {
		insertBuilder.SQL(fmt.Sprintf("ON CONFLICT DO NOTHING"))
	}

	sqlQuery, args := insertBuilder.Build()

	logger.Tracef("SQL: %s", sqlQuery)

	_, err := runner.ExecContext(ctx, sqlQuery, args...)

	if err != nil {
		logger.Error("SQL ERROR: ", err)
		return false, handleDbError(ctx, err)
	}

	return true, handleDbError(ctx, err)
}

func getTableName(sourceConfig *model.ResourceSourceConfig, history bool) string {
	def := ""

	if sourceConfig.Catalog != "" {
		def = fmt.Sprintf("\"%s\".", sourceConfig.Catalog)
	}

	def = fmt.Sprintf("\"%s\"", sourceConfig.Entity)

	if history {
		def = fmt.Sprintf("\"%s_h\"", sourceConfig.Entity)
	} else {
		def = fmt.Sprintf("\"%s\"", sourceConfig.Entity)
	}

	return def
}

func recordUpdate(ctx context.Context, runner QueryRunner, resource *model.Resource, record *model.Record, checkVersion bool) errors.ServiceError {
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
		if source, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			packedVal, exists := record.Properties[property.Name]

			if !exists {
				continue
			}

			propertyType := types.ByResourcePropertyType(property.Type)
			var val interface{}

			if property.Type == model.ResourcePropertyType_TYPE_OBJECT {
				var err2 error
				val, err2 = json.Marshal(packedVal.AsInterface())

				if err2 != nil {
					return errors.InternalError.WithDetails(err2.Error())
				}
				val = string(val.([]byte))
			} else {
				var err error
				val, err = propertyType.UnPack(packedVal)
				if err != nil {
					return errors.RecordValidationError.WithDetails(err.Error())
				}
			}

			updateBuilder.SetMore(updateBuilder.Equal(fmt.Sprintf("\"%s\"", source.Mapping.Mapping), val))
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

func recordList(ctx context.Context, runner QueryRunner, params abs.ListRecordParams) (result []*model.Record, total uint32, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	// find count
	countBuilder := sqlbuilder.Select("count(*)")
	countBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	countBuilder.From(getTableName(params.Resource.SourceConfig, params.UseHistory) + " as t")
	if params.Query != nil {
		var where = ""
		where, err = applyCondition(params.Resource, params.Query, countBuilder)
		if err != nil {
			return nil, 0, err
		}
		if where != "" {
			countBuilder.Where(where)
		}
	}
	countQuery, args := countBuilder.Build()

	logger.Tracef("SQL: %s", countQuery)

	countRow := runner.QueryRowContext(ctx, countQuery, args...)
	err = handleDbError(ctx, countRow.Scan(&total))

	if err != nil {
		return
	}

	if total == 0 {
		return
	}

	ownCols := util.ArrayMapString(prepareResourceRecordCols(params.Resource), func(s string) string {
		return fmt.Sprintf("t.\"%s\" as \"t_%s\"", s, s)
	})

	var joinCols []string
	if params.ResolveReferences {
		joinCols = recordPrepareJoinCols(runner, params.Resource)
	}

	selectBuilder := sqlbuilder.Select(append(ownCols, joinCols...)...)
	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	selectBuilder.From(getTableName(params.Resource.SourceConfig, params.UseHistory) + " as t")

	if params.ResolveReferences {
		err = recordPrepareJoins(runner, selectBuilder, params.Resource)

		if err != nil {
			return
		}
	}

	if params.Query != nil {
		var where = ""
		where, err = applyCondition(params.Resource, params.Query, selectBuilder)
		if err != nil {
			return nil, 0, err
		}
		if where != "" {
			selectBuilder.Where(where)
		}
	}

	if params.Limit == 0 || params.Limit > 10000 {
		params.Limit = 100
	}

	selectBuilder.Limit(int(params.Limit))
	selectBuilder.Offset(int(params.Offset))

	sqlQuery, args := selectBuilder.Build()

	logger.Tracef("SQL: %s", sqlQuery)

	rows, sqlErr := runner.Query(sqlQuery, args...)
	err = handleDbError(ctx, sqlErr)

	if err != nil {
		return
	}

	defer func() {
		err2 := rows.Close()

		if err2 != nil {
			logger.Print(err2)
		}
	}()

	for rows.Next() {
		record := new(model.Record)
		err = scanRecord(ctx, runner, record, params.Resource, params.ResolveReferences, rows)
		if err != nil {
			return
		}

		result = append(result, record)
	}

	return
}

func recordPrepareJoinScan(runner QueryRunner, resource *model.Resource, record *model.Record, rowScanFields *[]any) errors.ServiceError {
	//for _, reference := range resource.References {
	//	var referencedResource = new(model.Resource)
	//	err := resourceLoadDetailsByName(runner, referencedResource, resource.Namespace, reference.ReferencedResource)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	for _, property := range referencedResource.Properties {
	//		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
	//			propertyType := types.ByResourcePropertyType(property.Type)
	//			val := propertyType.Pointer(property.Required)
	//			*rowScanFields = append(*rowScanFields, val)
	//		}
	//	}
	//}

	return nil
}

func recordPrepareJoinCols(runner QueryRunner, resource *model.Resource) []string {
	var result []string
	//for _, reference := range resource.References {
	//	var referencedResource = new(model.Resource)
	//	err := resourceLoadDetailsByName(runner, referencedResource, resource.Namespace, reference.ReferencedResource)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	joinAlias := "l_" + referencedResource.SourceConfig.Mapping
	//
	//	for _, property := range referencedResource.Properties {
	//		if sourceConfig, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
	//			colName := fmt.Sprintf("%s.%s as %s_%s", joinAlias, sourceConfig.Mapping, joinAlias, sourceConfig.Mapping)
	//			result = append(result, colName)
	//		}
	//	}
	//}

	return result
}

func recordPrepareJoins(runner QueryRunner, builder *sqlbuilder.SelectBuilder, resource *model.Resource) errors.ServiceError {
	//for _, reference := range resource.References {
	//	referenceLocalDetails, err := resolveReferenceDetails(runner, resource, reference)
	//
	//	if err != nil {
	//		return nil
	//	}
	//
	//	onExpression := fmt.Sprintf("%s.%s=%s", referenceLocalDetails.joinAlias, referenceLocalDetails.referencedTableColumn, referenceLocalDetails.sourceTableColumn)
	//
	//	builder.JoinWithOption(sqlbuilder.LeftJoin, referenceLocalDetails.sourceTableColumn+" as "+referenceLocalDetails.joinAlias, onExpression)
	//}

	return nil
}

func applyCondition(resource *model.Resource, query *model.BooleanExpression, builder *sqlbuilder.SelectBuilder) (string, errors.ServiceError) {
	if and, ok := query.Expression.(*model.BooleanExpression_And); ok {
		if len(and.And.Expressions) == 0 {
			return "", nil
		}

		expressions, err := util.ArrayMapWithError(and.And.Expressions, func(t *model.BooleanExpression) (string, errors.ServiceError) {
			return applyCondition(resource, t, builder)
		})
		if err != nil {
			return "", err
		}
		return builder.And(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Or); ok {
		if len(and.Or.Expressions) == 0 {
			return "", nil
		}

		expressions, err := util.ArrayMapWithError(and.Or.Expressions, func(t *model.BooleanExpression) (string, errors.ServiceError) {
			return applyCondition(resource, t, builder)
		})
		if err != nil {
			return "", err
		}
		return builder.Or(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Not); ok {
		return applyCondition(resource, and.Not, builder)
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_Equal); ok {
		left, err := applyExpression(resource, equ.Equal.Left, builder)
		if err != nil {
			return "", err
		}
		right, err := applyExpression(resource, equ.Equal.Right, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s = %s", left, right), nil
	}

	if isn, ok := query.Expression.(*model.BooleanExpression_IsNull); ok {
		left, err := applyExpression(resource, isn.IsNull, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s is null", left), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_GreaterThan); ok {
		left, err := applyExpression(resource, equ.GreaterThan.Left, builder)
		if err != nil {
			return "", err
		}
		right, err := applyExpression(resource, equ.GreaterThan.Right, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s > %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_GreaterThanOrEqual); ok {
		left, err := applyExpression(resource, equ.GreaterThanOrEqual.Left, builder)
		if err != nil {
			return "", err
		}
		right, err := applyExpression(resource, equ.GreaterThanOrEqual.Right, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s >= %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThan); ok {
		left, err := applyExpression(resource, equ.LessThan.Left, builder)
		if err != nil {
			return "", err
		}
		right, err := applyExpression(resource, equ.LessThan.Right, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s < %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThanOrEqual); ok {
		left, err := applyExpression(resource, equ.LessThanOrEqual.Left, builder)
		if err != nil {
			return "", err
		}
		right, err := applyExpression(resource, equ.LessThanOrEqual.Right, builder)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s <= %s", left, right), nil
	}

	if _, ok := query.Expression.(*model.BooleanExpression_RegexMatch); ok {
		panic("not implemented")
	}

	panic("unknown boolean expression type: " + query.String())
}

func applyExpression(resource *model.Resource, query *model.Expression, builder *sqlbuilder.SelectBuilder) (string, errors.ServiceError) {
	if query.Expression == nil {
		return "", errors.PropertyNotFoundError.WithDetails("expression is empty")
	}

	var additionalProperties = []string{
		"id", "version",
	}

	if propEx, ok := query.Expression.(*model.Expression_Property); ok {
		for _, ap := range additionalProperties {
			if ap == propEx.Property {
				return ap, nil
			}
		}
		property := locatePropertyByName(resource, propEx.Property)

		if property == nil {
			return "", errors.PropertyNotFoundError.WithDetails(propEx.Property)
		}

		return propEx.Property, nil
	}

	if propEx, ok := query.Expression.(*model.Expression_Value); ok {
		return builder.Var(propEx.Value.AsInterface()), nil
	}

	panic("unknown expression type: " + query.String())
}

func scanRecord(ctx context.Context, runner QueryRunner, record *model.Record, resource *model.Resource, resolveReferences bool, scanner QueryResultScanner) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))

	var rowScanFields []any

	var hasOwnId = checkHasOwnId(resource)

	if hasOwnId {
		rowScanFields = append(rowScanFields, &record.Id)
	}

	var propertyPointers = make(map[string]interface{})
	var properties = make(map[string]*structpb.Value)
	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			propertyType := types.ByResourcePropertyType(property.Type)
			val := propertyType.Pointer(property.Required)
			rowScanFields = append(rowScanFields, val)
			propertyPointers[property.Name] = val
		}
	}

	if resolveReferences {
		err := recordPrepareJoinScan(runner, resource, record, &rowScanFields)

		if err != nil {
			return err
		}
	}

	var createdOn = new(time.Time)
	var updatedOn = new(*time.Time)
	var updatedBy = new(*string)

	if !annotations2.IsEnabled(resource, annotations2.DisableAudit) {
		record.AuditData = &model.AuditData{}

		rowScanFields = append(rowScanFields, createdOn)
		rowScanFields = append(rowScanFields, updatedOn)
		rowScanFields = append(rowScanFields, &record.AuditData.CreatedBy)
		rowScanFields = append(rowScanFields, updatedBy)

		rowScanFields = append(rowScanFields, &record.Version)
	}

	sqlErr := scanner.Scan(rowScanFields...)

	if sqlErr == sql.ErrNoRows {
		return errors.RecordNotFoundError.WithDetails(fmt.Sprintf("namespace: %s; resource: %s", resource.Namespace, resource.Name))
	}

	err := handleDbError(ctx, sqlErr)

	if err != nil {
		return err
	}

	var ids []string
	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			propP := propertyPointers[property.Name]

			propertyType := types.ByResourcePropertyType(property.Type)

			val := types.Dereference(propP)

			if val == nil {
				continue
			}

			if property.Type == model.ResourcePropertyType_TYPE_OBJECT {
				var data = new(interface{})
				err2 := json.Unmarshal([]byte(val.(string)), data)

				if err2 != nil {
					return errors.InternalError.WithDetails(err2.Error())
				}

				val = *data
			}

			packedValue, err := propertyType.Pack(val)

			if err != nil {
				return handleDbError(ctx, err)
			}

			properties[property.Name] = packedValue

			if property.Primary {
				ids = append(ids, propertyType.String(val))
			}
		}
	}

	if !hasOwnId {
		record.Id = strings.Join(ids, "-")
	}

	record.Properties = properties

	if record.AuditData == nil {
		record.AuditData = new(model.AuditData)
	}

	record.AuditData.CreatedOn = timestamppb.New(*createdOn)
	if *updatedOn != nil {
		record.AuditData.UpdatedOn = timestamppb.New(**updatedOn)
	}
	if *updatedBy != nil {
		record.AuditData.UpdatedBy = **updatedBy
	}

	record.Resource = resource.Name
	record.DataType = model.DataType_USER

	if record.Id == "" {
		return errors.RecordNotFoundError
	}

	logger.Tracef("Record scanned: %s" + record.Id)

	return nil
}

func checkHasOwnId(resource *model.Resource) bool {
	return !annotations2.IsEnabled(resource, annotations2.DoPrimaryKeyLookup)
}

func readRecord(ctx context.Context, runner QueryRunner, resource *model.Resource, id string) (*model.Record, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	selectBuilder := sqlbuilder.Select(prepareResourceRecordCols(resource)...)
	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	selectBuilder.From(getTableName(resource.SourceConfig, false))
	selectBuilder.Where(selectBuilder.Equal("id", id))

	sqlQuery, _ := selectBuilder.Build()

	logger.Tracef("SQL: %s", sqlQuery)

	row := runner.QueryRowContext(ctx, sqlQuery, id)

	if row.Err() != nil {
		return nil, handleDbError(ctx, row.Err())
	}

	record := new(model.Record)

	err := scanRecord(ctx, runner, record, resource, false, row)

	if err != nil {
		return nil, err
	}

	return record, nil
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

func locatePrimaryKey(resource *model.Resource) (string, errors.ServiceError) {
	for _, property := range resource.Properties {
		if property.Primary {
			return property.SourceConfig.(*model.ResourceProperty_Mapping).Mapping.Mapping, nil
		}
	}

	return "", errors.UnableToLocatePrimaryKey
}

func prepareResourceRecordCols(resource *model.Resource) []string {
	var cols []string

	if checkHasOwnId(resource) {
		cols = append(cols, "id")
	}

	for _, property := range resource.Properties {
		if source, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			col := fmt.Sprintf("%s", source.Mapping.Mapping)
			cols = append(cols, col)
		}
	}

	// referenced columns

	if !annotations2.IsEnabled(resource, annotations2.DisableAudit) {
		cols = append(cols, "created_on")
		cols = append(cols, "updated_on")
		cols = append(cols, "created_by")
		cols = append(cols, "updated_by")
		cols = append(cols, "version")
	}
	return cols
}
