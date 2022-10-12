package postgres

import (
	"data-handler/stub/model"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"time"
)

func recordInsert(runner QueryRunner, resource *model.Resource, records []*model.Record) error {
	if resource.Flags == nil {
		resource.Flags = &model.ResourceFlags{}
	}

	insertBuilder := sqlbuilder.InsertInto(resource.SourceConfig.Mapping)
	insertBuilder.SetFlavor(sqlbuilder.PostgreSQL)

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

	insertBuilder.Cols(cols...)

	for _, record := range records {
		recordNewId, _ := uuid.NewUUID()

		var row []interface{}

		row = append(row, recordNewId)

		for _, property := range resource.Properties {
			if _, ok := property.SourceConfig.(*model.ResourceProperty_Mapping); ok {
				val := record.Properties.AsMap()[property.Name]
				row = append(row, val)
			}
		}

		row = append(row, time.Now())
		row = append(row, nil)
		row = append(row, "test-usre")
		row = append(row, nil)
		row = append(row, 1)

		insertBuilder.Values(row...)
	}

	sqlQuery, args := insertBuilder.Build()

	_, err := runner.Exec(sqlQuery, args...)

	return err
}
