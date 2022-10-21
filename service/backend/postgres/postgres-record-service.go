package postgres

import (
	"data-handler/service/backend"
	"data-handler/stub/model"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) ListRecords(params backend.ListRecordParams) (result []*model.Record, total uint32, err error) {
	err = p.withBackend(params.Resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		result, total, err = recordList(tx, params)

		return err
	})

	if err != nil {
		log.Error("Unable to insert records", err)
		return
	}

	return
}

func (p *postgresResourceServiceBackend) AddRecords(params backend.BulkRecordsParams) ([]*model.Record, error) {
	err := p.withBackend(params.Resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		return recordInsert(tx, params.Resource, params.Records)
	})

	if err != nil {
		log.Error("Unable to insert records", err)
		return nil, err
	}

	return params.Records, nil
}

func (p *postgresResourceServiceBackend) UpdateRecords(params backend.BulkRecordsParams) ([]*model.Record, error) {
	err := p.withBackend(params.Resource.SourceConfig.DataSource, func(tx *sql.Tx) error {
		for _, record := range params.Records {
			err := recordUpdate(tx, params.Resource, record, params.CheckVersion)

			if err != nil {
				return err
			}
		}

		return nil
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

		if err != nil {
			return err
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
