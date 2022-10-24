package test

import (
	"context"
	"data-handler/stub"
	"data-handler/stub/model"
	"data-handler/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestResourceMigration_CrunchbaseMigration(t *testing.T) {
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
}

func TestResourceMigration_CrunchbaseMigrationWithResourceCreation(t *testing.T) {
	withAutoLoadedResource(t, container, dataSourceDhTest, "public.organization", func(resource1 *model.Resource) {
		resource2 := proto.Clone(resource1).(*model.Resource)

		resource2.Name = "organization_copy_new"
		resource2.SourceConfig.Mapping = "public.organization_copy_new"

		defer container.resourceService.Delete(context.TODO(), &stub.DeleteResourceRequest{
			Token:          "test-token",
			Ids:            []string{resource2.Name},
			DoMigration:    true,
			ForceMigration: false,
		})

		res, err := container.resourceService.Create(context.TODO(), &stub.CreateResourceRequest{
			Token:          "test-token",
			Resources:      []*model.Resource{resource2},
			DoMigration:    true,
			ForceMigration: false,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if res.Error != nil {
			t.Error(res.Error)
		}

		list, err := container.recordService.List(context.TODO(), &stub.ListRecordRequest{
			Token:    "test-token",
			Resource: resource1.Name,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if list.Error != nil {
			log.Print("WEEEE ARREEE FFAAAIILLEEDDD")
			t.Error(list.Error.Message)
			return
		}

		var records = util.ArrayMap(list.Content, func(record *model.Record) *model.Record {
			record.Resource = resource2.Name

			return record
		})

		createRes, err := container.recordService.Create(context.TODO(), &stub.CreateRecordRequest{
			Token:   "test-token",
			Records: records,
		})

		if err != nil {
			t.Error(err)
		}

		if createRes.Error != nil {
			t.Error(createRes.Error.Message)
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
}
