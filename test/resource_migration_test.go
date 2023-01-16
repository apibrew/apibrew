package test

import (
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestResourceMigration_CrunchbaseMigration(t *testing.T) {
	ctx := prepareTextContext()

	withAutoLoadedResource(ctx, t, container, dataSourceDhTest, "public.organization", func(resource1 *model.Resource) {
		withAutoLoadedResource(ctx, t, container, dataSourceDhTest, "public.organization_copy", func(resource2 *model.Resource) {
			list, err := container.recordService.List(ctx, &stub.ListRecordRequest{
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

			_, err = container.recordService.Create(ctx, &stub.CreateRecordRequest{
				Token:   "test-token",
				Records: records,
			})

			if err != nil {
				t.Error(err)
			}

			_, err = container.recordService.Delete(ctx, &stub.DeleteRecordRequest{
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

// @todo fix me
func TestResourceMigration_CrunchbaseMigrationWithResourceCreation(t *testing.T) {
	ctx := prepareTextContext()

	withAutoLoadedResource(ctx, t, container, dataSourceDhTest, "public.organization", func(resource1 *model.Resource) {
		resource2 := proto.Clone(resource1).(*model.Resource)

		resource2.Name = "organization_copy_new"
		resource2.SourceConfig.Mapping = "public.organization_copy_new"

		defer container.resourceService.Delete(ctx, &stub.DeleteResourceRequest{
			Token:          "test-token",
			Ids:            []string{resource2.Id},
			DoMigration:    true,
			ForceMigration: false,
		})

		createRes, err := container.resourceService.Create(ctx, &stub.CreateResourceRequest{
			Token:          "test-token",
			Resources:      []*model.Resource{resource2},
			DoMigration:    true,
			ForceMigration: false,
		})

		if err != nil {
			t.Error(err)
			return
		}

		if createRes.Error != nil {
			if createRes.Error.Code == model.ErrorCode_ALREADY_EXISTS {
				res2, _ := container.resourceService.GetByName(ctx, &stub.GetResourceByNameRequest{
					Token:     "test-token",
					Namespace: resource2.Namespace,
					Name:      resource2.Name,
				})
				resource2.Id = res2.Resource.Id

			} else {
				t.Error(createRes.Error.Message)
			}
			return
		} else {
			resource2.Id = createRes.Resources[0].Id
		}

		list, err := container.recordService.List(ctx, &stub.ListRecordRequest{
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

		createRes2, err := container.recordService.Create(ctx, &stub.CreateRecordRequest{
			Token:   "test-token",
			Records: records,
		})

		if err != nil {
			t.Error(err)
		}

		if createRes2.Error != nil {
			t.Error(createRes2.Error.Message)
		}

		_, err = container.recordService.Delete(ctx, &stub.DeleteRecordRequest{
			Token:    "test-token",
			Resource: resource2.Name,
			Ids: util.ArrayMap(list.Content, func(record *model.Record) string {
				return record.Id
			}),
		})

		//log.Print(list, resp)
	})
}
