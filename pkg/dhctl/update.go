package dhctl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/proto"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update - Update resource from existing table",
}

var updateRecordCmdId *string
var updateRecordCmdNamespace *string

type protoMessageUpdateCmdParams[T proto.Message] struct {
	use    string
	msg    T
	get    func(id string) T
	handle func(elem T)
	before func(cmd *cobra.Command)
}

func protoMessageUpdateCmd[T proto.Message](params protoMessageUpdateCmdParams[T]) *cobra.Command {
	fp := flags.NewProtoMessageParserFlags[T](params.msg.ProtoReflect())

	res := &cobra.Command{
		Use: params.use,
		Run: func(cmd *cobra.Command, args []string) {
			defineRootFlags(cmd)

			if params.before != nil {
				params.before(cmd)
			}

			err := cmd.PersistentFlags().Parse(args)

			if err != nil {
				log.Fatal(err)
			}

			parseRootFlags(cmd)

			rec := params.get(check2(cmd.PersistentFlags().GetString("id")))
			fp.Parse(rec, cmd, args)

			params.handle(rec)
		},
		DisableFlagParsing: true,
	}

	fp.Declare(res)

	return res
}

var updateRecordCmd = &cobra.Command{
	Use:                "record",
	Short:              "Update record <resource>",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		defineRootFlags(cmd)

		if len(args) == 0 {
			log.Fatal("resource not specified")
		}

		resourceName := args[0]

		resource := check2(GetDhClient().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *updateRecordCmdNamespace,
			Name:      resourceName,
		}))

		fp := flags.NewRecordParserFlags(resource.Resource)

		fp.Declare(cmd)

		err := cmd.PersistentFlags().Parse(args)

		if err != nil {
			log.Fatal(err)
		}

		parseRootFlags(cmd)

		record := check2(GetDhClient().GetRecordServiceClient().Get(cmd.Context(), &stub.GetRecordRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *updateRecordCmdNamespace,
			Resource:  resourceName,
			Id:        *updateRecordCmdId,
		})).Record

		fp.Parse(record, cmd, args)

		check2(GetDhClient().GetRecordServiceClient().Update(cmd.Context(), &stub.UpdateRecordRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *updateRecordCmdNamespace,
			Resource:  resourceName,
			Record:    record,
		}))

		describeWriter.WriteRecords(resource.Resource, 0, util.ArrToChan([]*model.Record{
			check2(GetDhClient().GetRecordServiceClient().Get(cmd.Context(), &stub.GetRecordRequest{
				Token:     GetDhClient().GetToken(),
				Namespace: *updateRecordCmdNamespace,
				Resource:  resourceName,
				Id:        *updateRecordCmdId,
			})).Record,
		}))
	},
}

func initUpdateCmd() {
	var migrate *bool
	var force *bool

	updateCmd.AddCommand(protoMessageUpdateCmd[*model.Resource](protoMessageUpdateCmdParams[*model.Resource]{
		msg: &model.Resource{},
		get: func(id string) *model.Resource {
			return check2(GetDhClient().GetResourceServiceClient().Get(context.TODO(), &stub.GetResourceRequest{
				Token:       GetDhClient().GetToken(),
				Id:          id,
				Annotations: nil,
			})).Resource
		},
		use: "resource",
		before: func(cmd *cobra.Command) {
			migrate = cmd.PersistentFlags().BoolP("migrate", "m", true, "")
			force = cmd.PersistentFlags().BoolP("force", "f", false, "")
		},
		handle: func(resource *model.Resource) {
			log.Println(*migrate, *force)
			resp := check2(GetDhClient().GetResourceServiceClient().Update(context.TODO(), &stub.UpdateResourceRequest{
				Token:          GetDhClient().GetToken(),
				Resources:      []*model.Resource{resource},
				DoMigration:    *migrate,
				ForceMigration: *force,
			}))

			describeWriter.WriteResources([]*model.Resource{
				check2(GetDhClient().GetResourceServiceClient().Get(context.TODO(), &stub.GetResourceRequest{
					Token:       GetDhClient().GetToken(),
					Id:          resp.Resources[0].Id,
					Annotations: nil,
				})).Resource,
			})
		},
	}))
	updateCmd.AddCommand(protoMessageUpdateCmd[*model.DataSource](protoMessageUpdateCmdParams[*model.DataSource]{
		msg: &model.DataSource{},
		use: "data-source",
		get: func(id string) *model.DataSource {
			return check2(GetDhClient().GetDataSourceServiceClient().Get(context.TODO(), &stub.GetDataSourceRequest{
				Token: GetDhClient().GetToken(),
				Id:    id,
			})).DataSource
		},
		handle: func(dataSource *model.DataSource) {
			resp := check2(GetDhClient().GetDataSourceServiceClient().Update(context.TODO(), &stub.UpdateDataSourceRequest{
				Token:       GetDhClient().GetToken(),
				DataSources: []*model.DataSource{dataSource},
			}))

			describeWriter.WriteRecords(resources.DataSourceResource, 0, util.ArrToChan([]*model.Record{
				mapping.DataSourceToRecord(check2(GetDhClient().GetDataSourceServiceClient().Get(context.TODO(), &stub.GetDataSourceRequest{
					Token: GetDhClient().GetToken(),
					Id:    resp.DataSources[0].Id,
				})).DataSource),
			}))
		},
	}))
	updateCmd.AddCommand(protoMessageUpdateCmd[*model.Namespace](protoMessageUpdateCmdParams[*model.Namespace]{
		msg: &model.Namespace{},
		use: "namespace",
		get: func(id string) *model.Namespace {
			return check2(GetDhClient().GetNamespaceServiceClient().Get(context.TODO(), &stub.GetNamespaceRequest{
				Token: GetDhClient().GetToken(),
				Id:    id,
			})).Namespace
		},
		handle: func(namespace *model.Namespace) {
			resp := check2(GetDhClient().GetNamespaceServiceClient().Update(context.TODO(), &stub.UpdateNamespaceRequest{
				Token:      GetDhClient().GetToken(),
				Namespaces: []*model.Namespace{namespace},
			}))

			describeWriter.WriteRecords(resources.NamespaceResource, 0, util.ArrToChan([]*model.Record{
				mapping.NamespaceToRecord(check2(GetDhClient().GetNamespaceServiceClient().Get(context.TODO(), &stub.GetNamespaceRequest{
					Token: GetDhClient().GetToken(),
					Id:    resp.Namespaces[0].Id,
				})).Namespace),
			}))
		},
	}))
	updateCmd.AddCommand(protoMessageUpdateCmd[*model.User](protoMessageUpdateCmdParams[*model.User]{
		msg: &model.User{},
		use: "user",
		get: func(id string) *model.User {
			return check2(GetDhClient().GetUserServiceClient().Get(context.TODO(), &stub.GetUserRequest{
				Token: GetDhClient().GetToken(),
				Id:    id,
			})).User
		},
		handle: func(user *model.User) {
			resp := check2(GetDhClient().GetUserServiceClient().Update(context.TODO(), &stub.UpdateUserRequest{
				Token: GetDhClient().GetToken(),
				Users: []*model.User{user},
			}))

			result := util.ArrayMap[*model.User, *model.Record](resp.Users, mapping.UserToRecord)

			describeWriter.WriteRecords(resources.UserResource, 0, util.ArrToChan(result))
		},
	}))
	updateCmd.AddCommand(protoMessageUpdateCmd[*model.RemoteExtension](protoMessageUpdateCmdParams[*model.RemoteExtension]{
		msg: &model.RemoteExtension{},
		use: "extension",
		get: func(id string) *model.RemoteExtension {
			return check2(GetDhClient().GetExtensionServiceClient().Get(context.TODO(), &stub.GetExtensionRequest{
				Token: GetDhClient().GetToken(),
				Id:    id,
			})).Extension
		},
		handle: func(extension *model.RemoteExtension) {
			resp := check2(GetDhClient().GetExtensionServiceClient().Update(context.TODO(), &stub.UpdateExtensionRequest{
				Token:      GetDhClient().GetToken(),
				Extensions: []*model.RemoteExtension{extension},
			}))

			result := util.ArrayMap[*model.RemoteExtension, *model.Record](resp.Extensions, mapping.ExtensionToRecord)

			describeWriter.WriteRecords(resources.UserResource, 0, util.ArrToChan(result))
		},
	}))

	updateRecordCmdId = updateRecordCmd.PersistentFlags().String("id", "", "Id of record to update")
	updateRecordCmdNamespace = updateRecordCmd.PersistentFlags().StringP("namespace", "n", "", "")

	updateCmd.AddCommand(updateRecordCmd)
}
