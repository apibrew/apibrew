package postgres

import (
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/service/types"
	"data-handler/stub/model"
	"data-handler/util"
	"fmt"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strings"
	"time"
)

func recordInsert(runner QueryRunner, resource *model.Resource, records []*model.Record, ignoreIfExists bool, history bool) (bool, error) {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	insertBuilder := sqlbuilder.InsertInto(getTableName(resource.SourceConfig.Mapping, history))
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	cols := prepareResourceRecordCols(resource)

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
				val := record.Properties.AsMap()[property.Name]
				if val == nil {
					row = append(row, nil)
					continue
				}
				propertyType := types.ByResourcePropertyType(property.Type)
				unpackedVal, err := propertyType.UnPack(val)

				if err != nil {
					log.Print("SQL ERROR: ", err)
					return false, err
				}

				row = append(row, unpackedVal)
			}
		}

		if !resource.Flags.DisableAudit {
			row = append(row, record.AuditData.CreatedOn.AsTime())
			row = append(row, record.AuditData.UpdatedOn.AsTime())
			row = append(row, record.AuditData.CreatedBy)
			row = append(row, record.AuditData.UpdatedBy)
			row = append(row, record.Version)
		}

		insertBuilder.Values(row...)
	}

	if ignoreIfExists {
		insertBuilder.SQL("ON CONFLICT(id) DO NOTHING")
	}

	sqlQuery, args := insertBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		log.Print("SQL ERROR: ", err)
		return false, err
	}

	return true, err
}

func getTableName(mapping string, history bool) string {
	if history {
		return mapping + "_h"
	} else {
		return mapping
	}
}

func recordUpdate(runner QueryRunner, resource *model.Resource, record *model.Record, checkVersion bool) (err error) {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	updateBuilder := sqlbuilder.Update(getTableName(resource.SourceConfig.Mapping, false))
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
			val := record.Properties.AsMap()[property.Name]

			propertyType := types.ByResourcePropertyType(property.Type)
			unpackedVal, err := propertyType.UnPack(val)

			if err != nil {
				return err
			}

			updateBuilder.SetMore(updateBuilder.Equal(source.Mapping.Mapping, unpackedVal))
		}
	}

	updateBuilder.SetMore(updateBuilder.Equal("updated_on", record.AuditData.UpdatedOn.AsTime()))
	updateBuilder.SetMore(updateBuilder.Equal("updated_by", record.AuditData.UpdatedBy))
	updateBuilder.SetMore("version = version + 1")

	sqlQuery, args := updateBuilder.Build()

	result, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return
	}

	if affected == 0 {
		return errors.NotFoundError
	}

	return
}

func recordList(runner QueryRunner, params backend.ListRecordParams) (result []*model.Record, total uint32, err error) {
	// find count
	countBuilder := sqlbuilder.Select("count(*)")
	countBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	countBuilder.From(getTableName(params.Resource.SourceConfig.Mapping, params.UseHistory))
	if params.Query != nil {
		var where = ""
		where, err = applyCondition(params.Resource, params.Query, countBuilder)
		if err != nil {
			return nil, 0, err
		}
		countBuilder.Where(where)
	}
	countQuery, args := countBuilder.Build()
	countRow := runner.QueryRow(countQuery, args...)
	err = countRow.Scan(&total)

	if err != nil {
		return
	}

	if total == 0 {
		return
	}

	selectBuilder := sqlbuilder.Select(prepareResourceRecordCols(params.Resource)...)
	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	selectBuilder.From(getTableName(params.Resource.SourceConfig.Mapping, params.UseHistory))
	if params.Query != nil {
		var where = ""
		where, err = applyCondition(params.Resource, params.Query, selectBuilder)
		if err != nil {
			return nil, 0, err
		}
		selectBuilder.Where(where)
	}

	if params.Limit == 0 || params.Limit > 10000 {
		params.Limit = 100
	}

	selectBuilder.Limit(int(params.Limit))
	selectBuilder.Offset(int(params.Offset))

	sqlQuery, args := selectBuilder.Build()
	rows, err := runner.Query(sqlQuery, args...)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		record := new(model.Record)
		err = scanRecord(record, params.Resource, rows)
		if err != nil {
			return
		}

		result = append(result, record)
	}

	return
}

func applyCondition(resource *model.Resource, query *model.BooleanExpression, builder *sqlbuilder.SelectBuilder) (string, error) {
	if and, ok := query.Expression.(*model.BooleanExpression_And); ok {
		expressions, err := util.ArrayMapWithError(and.And.Expressions, func(t *model.BooleanExpression) (string, error) {
			return applyCondition(resource, t, builder)
		})
		if err != nil {
			return "", err
		}
		return builder.And(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Or); ok {
		expressions, err := util.ArrayMapWithError(and.Or.Expressions, func(t *model.BooleanExpression) (string, error) {
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

func applyExpression(resource *model.Resource, query *model.Expression, builder *sqlbuilder.SelectBuilder) (string, error) {
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

func scanRecord(record *model.Record, resource *model.Resource, scanner QueryResultScanner) error {
	var rowScanFields []any

	var hasOwnId = checkHasOwnId(resource)

	if hasOwnId {
		rowScanFields = append(rowScanFields, &record.Id)
	}

	var propertyPointers = make(map[string]interface{})
	var properties = make(map[string]interface{})
	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			propertyType := types.ByResourcePropertyType(property.Type)
			val := propertyType.Pointer(property.Required)
			rowScanFields = append(rowScanFields, val)
			propertyPointers[property.Name] = val
		}
	}

	var createdOn = new(time.Time)
	var updatedOn = new(*time.Time)
	var updatedBy = new(*string)

	if !resource.Flags.DisableAudit {
		record.AuditData = &model.AuditData{}

		rowScanFields = append(rowScanFields, createdOn)
		rowScanFields = append(rowScanFields, updatedOn)
		rowScanFields = append(rowScanFields, &record.AuditData.CreatedBy)
		rowScanFields = append(rowScanFields, updatedBy)

		rowScanFields = append(rowScanFields, &record.Version)
	}

	err := scanner.Scan(rowScanFields...)

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

			packedValue, err := propertyType.Pack(val)

			if err != nil {
				return err
			}

			properties[property.Name] = packedValue

			if property.Primary {
				ids = append(ids, propertyType.String(packedValue))
			}
		}

		if _, ok := properties[property.Name].(uuid.UUID); ok {
			properties[property.Name] = properties[property.Name].(uuid.UUID).String()
		}

		if _, ok := properties[property.Name].(time.Time); ok {
			properties[property.Name] = properties[property.Name].(time.Time).Format(time.RFC3339)
		}
	}

	if !hasOwnId {
		record.Id = strings.Join(ids, "-")
	}

	propStruct, err := structpb.NewStruct(properties)

	record.Properties = propStruct

	if err != nil {
		return err
	}

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
	record.Type = model.DataType_USER

	if record.Id == "" {
		return errors.NotFoundError
	}

	return nil
}

func checkHasOwnId(resource *model.Resource) bool {
	return !resource.Flags.AutoCreated && !resource.Flags.DoPrimaryKeyLookup
}

func readRecord(runner QueryRunner, resource *model.Resource, id string) (*model.Record, error) {
	selectBuilder := sqlbuilder.Select(prepareResourceRecordCols(resource)...)
	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	selectBuilder.From(resource.SourceConfig.Mapping)
	selectBuilder.Where(selectBuilder.Equal("id", id))

	sqlQuery, _ := selectBuilder.Build()

	row := runner.QueryRow(sqlQuery, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	record := new(model.Record)

	err := scanRecord(record, resource, row)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func deleteRecords(runner QueryRunner, resource *model.Resource, ids []string) error {
	deleteBuilder := sqlbuilder.DeleteFrom(resource.SourceConfig.Mapping)
	deleteBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	if checkHasOwnId(resource) {
		deleteBuilder.Where(deleteBuilder.In("id", util.ArrayMapToInterface(ids)...))
	} else {
		idField, err := locatePrimaryKey(resource)

		if err != nil {
			return err
		}

		deleteBuilder.Where(deleteBuilder.In(idField, util.ArrayMapToInterface(ids)...))
	}

	sqlQuery, args := deleteBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return err
	}

	return nil
}

func locatePrimaryKey(resource *model.Resource) (string, error) {
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
			cols = append(cols, source.Mapping.Mapping)
		}
	}

	if !resource.Flags.DisableAudit {
		cols = append(cols, "created_on")
		cols = append(cols, "updated_on")
		cols = append(cols, "created_by")
		cols = append(cols, "updated_by")
		cols = append(cols, "version")
	}
	return cols
}
