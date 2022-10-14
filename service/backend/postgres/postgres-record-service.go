package postgres

import (
	"data-handler/service/backend"
	"data-handler/stub/model"
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) AddRecords(params backend.AddRecordsParams) ([]*model.Record, error) {
	err := p.withBackend(params.Resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		return recordInsert(tx, params.Resource, params.Records)
	})

	if err != nil {
		log.Error("Unable to insert records", err)
		return nil, err
	}

	return params.Records, nil
}

func (p *postgresResourceServiceBackend) GetRecord(resource *model.Resource, id string) (*model.Record, error) {
	var record *model.Record = nil
	err := p.withBackend(resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		var err error
		record, err = readRecord(tx, resource, id)

		if record.Id == "" {
			return errors.New("record does not exists")
		}

		return err
	})

	return record, err
}

func (p *postgresResourceServiceBackend) DeleteRecords(resource *model.Resource, ids []string) error {
	return p.withBackend(resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		return deleteRecords(tx, resource, ids)
	})
}
