package test

import (
	"context"
	"data-handler/stub"
	"data-handler/stub/model"
	"data-handler/util"
	"testing"
)

func TestResourceMigration_CrunchbaseMigration(t *testing.T) {
	withClient(func(container *SimpleAppGrpcContainer) {
		withAutoLoadedResource(t, container, dataSourceDhTest, "public.organization", func(resource1 *model.Resource) {
			withAutoLoadedResource(t, container, dataSourceDhTest, "public.organization_copy", func(resource2 *model.Resource) {
				list, err := container.recordService.List(context.TODO(), &stub.ListRecordRequest{
					Token:    "test-token",
					Resource: resource1.Name,
				})

				if err != nil {
					t.Error(err)
					return
				}

				var records = util.ArrayMap(list.Content, func(record *model.Record) *model.Record {
					record.Resource = resource2.Name

					return record
				})

				_, err = container.recordService.Create(context.TODO(), &stub.CreateRecordRequest{
					Token:   "test-token",
					Records: records,
				})

				if err != nil {
					t.Error(err)
				}

				_, err = container.recordService.Delete(context.TODO(), &stub.DeleteRecordRequest{
					Token:    "test-token",
					Resource: resource2.Name,
					Ids: util.ArrayMap(list.Content, func(record *model.Record) string {
						return record.Id
					}),
				})

				//log.Print(list, resp)
			})
		})
	})
}
