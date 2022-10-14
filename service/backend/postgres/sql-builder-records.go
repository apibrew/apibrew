package postgres

import (
	"data-handler/service/backend"
	"data-handler/stub/model"
	"data-handler/util"
	"errors"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func recordInsert(runner QueryRunner, resource *model.Resource, records []*model.Record) error {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	insertBuilder := sqlbuilder.InsertInto(resource.SourceConfig.Mapping)
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)

	cols := prepareResourceRecordCols(resource)

	insertBuilder.Cols(cols...)

	for _, record := range records {
		recordNewId, _ := uuid.NewUUID()
		record.Id = recordNewId.String()
		now := time.Now()
		record.AuditData = &model.AuditData{
			CreatedOn: timestamppb.New(now),
			UpdatedOn: timestamppb.New(now),
			CreatedBy: "test-user",
			UpdatedBy: "",
		}
		record.Version = 1

		var row []interface{}

		row = append(row, record.Id)

		for _, property := range resource.Properties {
			if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
				val := record.Properties.AsMap()[property.Name]
				row = append(row, val)
			}
		}

		row = append(row, record.AuditData.CreatedOn.AsTime())
		row = append(row, record.AuditData.UpdatedOn.AsTime())
		row = append(row, record.AuditData.CreatedBy)
		row = append(row, record.AuditData.UpdatedBy)
		row = append(row, record.Version)

		insertBuilder.Values(row...)
	}

	sqlQuery, args := insertBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return err
	}

	return err
}

func recordUpdate(runner QueryRunner, resource *model.Resource, record *model.Record) (err error) {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	updateBuilder := sqlbuilder.Update(resource.SourceConfig.Mapping)
	updateBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	updateBuilder.Where(updateBuilder.Equal("id", record.Id), updateBuilder.Equal("version", record.Version))

	now := time.Now()

	record.AuditData.UpdatedOn = timestamppb.New(now)
	record.AuditData.UpdatedBy = "test-user"

	record.Version++

	for _, property := range resource.Properties {
		if source, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			val := record.Properties.AsMap()[property.Name]
			updateBuilder.SetMore(updateBuilder.Equal(source.Mapping, val))
		}
	}

	updateBuilder.SetMore(updateBuilder.Equal("updated_on", record.AuditData.UpdatedOn.AsTime()))
	updateBuilder.SetMore(updateBuilder.Equal("updated_by", record.AuditData.UpdatedBy))
	updateBuilder.SetMore(updateBuilder.Equal("version", record.Version))

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
		return errors.New("record not found or version is wrong")
	}

	return
}

func recordList(runner QueryRunner, params backend.ListRecordParams) (result []*model.Record, total uint32, err error) {
	// find count
	countBuilder := sqlbuilder.Select("count(*)")
	countBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	countBuilder.From(params.Resource.SourceConfig.Mapping)
	applyCondition(params.Query, countBuilder)
	countQuery, _ := countBuilder.Build()
	countRow := runner.QueryRow(countQuery)
	err = countRow.Scan(&total)

	if err != nil {
		return
	}

	if total == 0 {
		return
	}

	selectBuilder := sqlbuilder.Select(prepareResourceRecordCols(params.Resource)...)
	selectBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	selectBuilder.From(params.Resource.SourceConfig.Mapping)
	applyCondition(params.Query, selectBuilder)
	sqlQuery, _ := selectBuilder.Build()
	rows, err := runner.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return
	}

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

func applyCondition(query *model.BooleanExpression, builder *sqlbuilder.SelectBuilder) {

}

func scanRecord(record *model.Record, resource *model.Resource, scanner QueryResultScanner) error {
	var rowScanFields []any

	rowScanFields = append(rowScanFields, &record.Id)

	var propertyPointers = make(map[string]interface{})
	var properties = make(map[string]interface{})
	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			val := getPropertyPointer(property.Type, property.Required)
			rowScanFields = append(rowScanFields, val)
			propertyPointers[property.Name] = val
		}
	}

	record.AuditData = &model.AuditData{}

	var createdOn = new(time.Time)
	var updatedOn = new(*time.Time)
	var updatedBy = new(*string)

	rowScanFields = append(rowScanFields, createdOn)
	rowScanFields = append(rowScanFields, updatedOn)
	rowScanFields = append(rowScanFields, &record.AuditData.CreatedBy)
	rowScanFields = append(rowScanFields, updatedBy)
	rowScanFields = append(rowScanFields, &record.Version)

	err := scanner.Scan(rowScanFields...)

	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			properties[property.Name] = dereferenceProperty(propertyPointers[property.Name], property.Type, property.Required)
		}
	}

	propStruct, err := structpb.NewStruct(properties)

	record.Properties = propStruct

	if err != nil {
		return err
	}

	record.AuditData.CreatedOn = timestamppb.New(*createdOn)
	if *updatedOn != nil {
		record.AuditData.UpdatedOn = timestamppb.New(**updatedOn)
	}
	if *updatedBy != nil {
		record.AuditData.UpdatedBy = **updatedBy
	}
	return nil
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

	deleteBuilder.Where(deleteBuilder.In("id", util.ArrayMapToInterface(ids)...))

	sqlQuery, args := deleteBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	if err != nil {
		return err
	}

	return nil
}

func prepareResourceRecordCols(resource *model.Resource) []string {
	var cols []string

	cols = append(cols, "id")

	for _, property := range resource.Properties {
		if source, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			cols = append(cols, source.Mapping)
		}
	}

	cols = append(cols, "created_on")
	cols = append(cols, "updated_on")
	cols = append(cols, "created_by")
	cols = append(cols, "updated_by")
	cols = append(cols, "version")
	return cols
}
