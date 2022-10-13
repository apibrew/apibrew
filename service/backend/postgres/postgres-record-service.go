package postgres

import (
	"data-handler/service/backend"
	"data-handler/stub/model"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) AddRecords(params backend.AddRecordsParams) ([]*model.Record, error) {
	err := p.withBackend(params.Backend, func(tx *sql.Tx) error {
		return recordInsert(tx, params.Resource, params.Records)
	})

	if err != nil {
		log.Error("Unable to insert records", err)
		return nil, err
	}

	return params.Records, nil
}

func (p *postgresResourceServiceBackend) GetRecord(bck backend.DataSourceBackend, resource *model.Resource, id string) (*model.Record, error) {
	var record *model.Record = nil
	err := p.withBackend(bck, func(tx *sql.Tx) error {
		var err error
		record, err = readRecord(tx, resource, id)

		return err
	})

	return record, err
}
