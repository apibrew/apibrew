package postgres

import (
	"data-handler/stub/model"
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

	//row.Scan()

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

	err := row.Scan(rowScanFields...)

	for _, property := range resource.Properties {
		if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
			properties[property.Name] = dereferenceProperty(propertyPointers[property.Name], property.Type, property.Required)
		}
	}

	propStruct, err := structpb.NewStruct(properties)

	record.Properties = propStruct

	if err != nil {
		return nil, err
	}

	record.AuditData.CreatedOn = timestamppb.New(*createdOn)
	if *updatedOn != nil {
		record.AuditData.UpdatedOn = timestamppb.New(**updatedOn)
	}
	if *updatedBy != nil {
		record.AuditData.UpdatedBy = **updatedBy
	}

	return record, nil
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
