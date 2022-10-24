package postgres

import (
	"data-handler/service/backend"
	"data-handler/service/errors"
	"data-handler/stub/model"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (p *postgresResourceServiceBackend) ListRecords(params backend.ListRecordParams) (result []*model.Record, total uint32, err error) {
	log.Tracef("Begin listing: %v", params)
	err = p.withBackend(params.Resource.SourceConfig.DataSource, true, func(tx *sql.Tx) error {
		result, total, err = recordList(tx, params)

		return err
	})
	log.Tracef("Begin listed: %v", params)

	return
}

func (p *postgresResourceServiceBackend) AddRecords(params backend.BulkRecordsParams) ([]*model.Record, bool, error) {
	var inserted bool
	var err error

	log.Tracef("Begin creating: %v; %v", params.Records)

	err = p.withBackend(params.Resource.SourceConfig.DataSource, false, func(tx *sql.Tx) error {
		inserted, err = recordInsert(tx, params.Resource, params.Records, params.IgnoreIfExists, false)

		if err != nil {
			return err
		}

		if inserted && params.Resource.Flags.KeepHistory {
			_, err = recordInsert(tx, params.Resource, params.Records, false, true)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, inserted, err
	}

	log.Tracef("Records created: %v; %v", params.Records, inserted)

	return params.Records, inserted, nil
}

func (p *postgresResourceServiceBackend) UpdateRecords(params backend.BulkRecordsParams) ([]*model.Record, error) {
	err := p.withBackend(params.Resource.SourceConfig.DataSource, false, func(tx *sql.Tx) error {
		for _, record := range params.Records {
			err := recordUpdate(tx, params.Resource, record, params.CheckVersion)

			if err != nil {
				return err
			}
		}

		if params.Resource.Flags.KeepHistory {
			_, err := recordInsert(tx, params.Resource, params.Records, false, false)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return params.Records, nil
}

func (p *postgresResourceServiceBackend) GetRecord(resource *model.Resource, id string) (*model.Record, error) {
	var record *model.Record = nil
	err := p.withBackend(resource.SourceConfig.DataSource, true, func(tx *sql.Tx) error {
		var err error
		record, err = readRecord(tx, resource, id)

		if err == sql.ErrNoRows {
			return errors.NotFoundError.WithDetails(fmt.Sprintf("workspace %s; resource %s; id %v", resource.Workspace, resource.Name, id))
		}

		if err != nil {
			return err
		}

		return err
	})

	return record, err
}

func (p *postgresResourceServiceBackend) DeleteRecords(resource *model.Resource, ids []string) error {
	log.Tracef("Begin deleting records: %v / %v / %v", resource.Workspace, resource.Name, ids)
	err := p.withBackend(resource.SourceConfig.DataSource, false, func(tx *sql.Tx) error {
		return deleteRecords(tx, resource, ids)
	})
	if err != nil {
		log.Print(err)
	} else {
		log.Tracef("records deleted: %v / %v / %v", resource.Workspace, resource.Name, ids)
	}

	return err
}
