package test

import (
	"testing"
)

func TestResourceMigration_CrunchbaseMigration(t *testing.T) {
	t.SkipNow()

	//withAutoLoadedResource(setup.Ctx, t, setup.DataSourceDhTest, "public", "organization", func(resource1 *model.Resource) {
	//	withAutoLoadedResource(setup.Ctx, t, setup.DataSourceDhTest, "public", "organization_copy", func(resource2 *model.Resource) {
	//		list, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
	//			Resource: resource1.Name,
	//		})
	//
	//		if err != nil {
	//			t.Error(err)
	//			return
	//		}
	//
	//		var records = util.ArrayMap(list.Content, func(record *model.Record) *model.Record {
	//
	//			return record
	//		})
	//
	//		_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{
	//			Resource: resource2.Name,
	//			Records:  records,
	//		})
	//
	//		if err != nil {
	//			t.Error(err)
	//		}
	//
	//		_, err = recordClient.Delete_(setup.Ctx, &stub.DeleteRecordRequest{
	//			Resource: resource2.Name,
	//			Ids: util.ArrayMap(list.Content, func(record *model.Record) string {
	//				return record.Id
	//			}),
	//		})
	//
	//		if err != nil {
	//			t.Error(err)
	//		}
	//	})
	//})
}

// @todo fix me
//func TestResourceMigration_CrunchbaseMigrationWithResourceCreation(t *testing.T) {
//
//
//	withAutoLoadedResource(ctx, t, container, dataSourceDhTest, "public", "organization", func(resource1 *model.resource) {
//		resource2 := proto.Clone(resource1).(*model.resource)
//
//		resource2.Name = "organization_copy_new"
//		resource2.SourceConfig.Catalog = "public"
//		resource2.SourceConfig.Entity = "organization_copy_new"
//
//		defer container.resourceService.Delete_(ctx, &stub.DeleteResourceRequest{
//			Token:          "test-token",
//			Ids:            []string{resource2.Id},
//			DoMigration:    true,
//			ForceMigration: false,
//		})
//
//		createRes, err := container.resourceService.Create(ctx, &stub.CreateResourceRequest{
//			Token:          "test-token",
//			Resources:      []*model.resource{resource2},
//			DoMigration:    true,
//			ForceMigration: false,
//		})
//
//		if err != nil {
//			if errors.GetErrorCode(err) == model.ErrorCode_ALREADY_EXISTS {
//				res2, _ := container.resourceService.GetByName(ctx, &stub.GetResourceByNameRequest{
//					Token:     "test-token",
//					Namespace: resource2.Namespace,
//					Name:      resource2.Name,
//				})
//				resource2.Id = res2.resource.Id
//
//			} else {
//				t.Error(err)
//				return
//			}
//		} else {
//			resource2.Id = createRes.Resources[0].Id
//		}
//
//		list, err := container.recordService.List(ctx, &stub.ListRecordRequest{
//			Token:    "test-token",
//			resource: resource1.Name,
//		})
//
//		if err != nil {
//			t.Error(err)
//			return
//		}
//
//		if err != nil {
//			t.Error(err)
//			return
//		}
//
//		var records = util.ArrayMap(list.Content, func(record *model.Record) *model.Record {
//			record.resource = resource2.Name
//
//			return record
//		})
//
//		_, err = container.recordService.Create(ctx, &stub.CreateRecordRequest{
//			Token:   "test-token",
//			Records: records,
//		})
//
//		if err != nil {
//			t.Error(err)
//		}
//
//		if err != nil {
//			t.Error(err)
//		}
//
//		_, err = container.recordService.Delete_(ctx, &stub.DeleteRecordRequest{
//			Token:    "test-token",
//			resource: resource2.Name,
//			Ids: util.ArrayMap(list.Content, func(record *model.Record) string {
//				return record.Id
//			}),
//		})
//
//		//log.Print(list, resp)
//	})
//}
