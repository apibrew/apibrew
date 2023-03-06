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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create - Create resource from existing table",
}

type protoMessageCreateCmdParams[T proto.Message] struct {
	msg    T
	use    string
	handle func(elem T)
	before func(cmd *cobra.Command)
}

func protoMessageCreateCmd[T proto.Message](params protoMessageCreateCmdParams[T]) *cobra.Command {
	fp := flags.NewProtoMessageParserFlags[T](params.msg.ProtoReflect())

	res := &cobra.Command{
		Use: params.use,
		Run: func(cmd *cobra.Command, args []string) {
			defineRootFlags(cmd)

			err := cmd.PersistentFlags().Parse(args)

			if err != nil {
				log.Fatal(err)
			}

			parseRootFlags(cmd)

			if params.before != nil {
				params.before(cmd)
			}

			fp.Parse(params.msg, cmd, args)

			params.handle(params.msg)
		},
		DisableFlagParsing: true,
	}

	fp.Declare(res)

	return res
}

var createRecordNamespace *string

var createRecordCmd = &cobra.Command{
	Use:                "record",
	Short:              "Create record <resource>",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		defineRootFlags(cmd)

		_ = cmd.PersistentFlags().Parse(args)

		parseRootFlags(cmd)

		if len(args) == 0 {
			log.Fatal("resource not specified")
		}

		resourceName := args[0]

		resource := check2(GetDhClient().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *createRecordNamespace,
			Name:      resourceName,
		})).Resource

		fp := flags.NewRecordParserFlags(resource)

		fp.Declare(cmd)

		err := cmd.PersistentFlags().Parse(args)

		if err != nil {
			log.Fatal(err)
		}

		var record = &model.Record{}
		fp.Parse(record, cmd, args)

		records := check2(GetDhClient().GetRecordServiceClient().Create(cmd.Context(), &stub.CreateRecordRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *createRecordNamespace,
			Resource:  resourceName,
			Record:    record,
		})).Records

		describeWriter.WriteRecords(resource, 0, util.ArrToChan(records))
	},
}

func initCreateCmd() {
	var migrate *bool
	var force *bool

	createCmd.AddCommand(protoMessageCreateCmd[*model.Resource](protoMessageCreateCmdParams[*model.Resource]{
		msg: &model.Resource{},
		use: "resource",
		before: func(cmd *cobra.Command) {
			migrate = cmd.PersistentFlags().BoolP("migrate", "m", true, "")
			force = cmd.PersistentFlags().BoolP("force", "f", false, "")
		},
		handle: func(resource *model.Resource) {
			resp, err := GetDhClient().GetResourceServiceClient().Create(context.TODO(), &stub.CreateResourceRequest{
				Token:          GetDhClient().GetToken(),
				Resources:      []*model.Resource{resource},
				DoMigration:    *migrate,
				ForceMigration: *force,
			})

			if err != nil {
				log.Fatal(err)
			}

			describeWriter.WriteResources(resp.Resources)
		},
	}))
	createCmd.AddCommand(protoMessageCreateCmd[*model.DataSource](protoMessageCreateCmdParams[*model.DataSource]{
		msg: &model.DataSource{},
		use: "data-source",
		handle: func(dataSource *model.DataSource) {
			resp, err := GetDhClient().GetDataSourceServiceClient().Create(context.TODO(), &stub.CreateDataSourceRequest{
				Token:       GetDhClient().GetToken(),
				DataSources: []*model.DataSource{dataSource},
			})

			if err != nil {
				log.Fatal(err)
			}

			result := util.ArrayMap[*model.DataSource, *model.Record](resp.DataSources, mapping.DataSourceToRecord)

			describeWriter.WriteRecords(resources.DataSourceResource, 0, util.ArrToChan(result))
		},
	}))
	createCmd.AddCommand(protoMessageCreateCmd[*model.Namespace](protoMessageCreateCmdParams[*model.Namespace]{
		msg: &model.Namespace{},
		use: "namespace",
		handle: func(namespace *model.Namespace) {
			resp, err := GetDhClient().GetNamespaceServiceClient().Create(context.TODO(), &stub.CreateNamespaceRequest{
				Token:      GetDhClient().GetToken(),
				Namespaces: []*model.Namespace{namespace},
			})

			if err != nil {
				log.Fatal(err)
			}

			result := util.ArrayMap[*model.Namespace, *model.Record](resp.Namespaces, mapping.NamespaceToRecord)

			describeWriter.WriteRecords(resources.NamespaceResource, 0, util.ArrToChan(result))
		},
	}))
	createCmd.AddCommand(protoMessageCreateCmd[*model.User](protoMessageCreateCmdParams[*model.User]{
		msg: &model.User{},
		use: "user",
		handle: func(user *model.User) {
			resp, err := GetDhClient().GetUserServiceClient().Create(context.TODO(), &stub.CreateUserRequest{
				Token: GetDhClient().GetToken(),
				Users: []*model.User{user},
			})

			if err != nil {
				log.Fatal(err)
			}

			result := util.ArrayMap[*model.User, *model.Record](resp.Users, mapping.UserToRecord)

			describeWriter.WriteRecords(resources.UserResource, 0, util.ArrToChan(result))
		},
	}))
	createCmd.AddCommand(protoMessageCreateCmd[*model.RemoteExtension](protoMessageCreateCmdParams[*model.RemoteExtension]{
		msg: &model.RemoteExtension{},
		use: "extension",
		handle: func(extension *model.RemoteExtension) {
			resp, err := GetDhClient().GetExtensionServiceClient().Create(context.TODO(), &stub.CreateExtensionRequest{
				Token:      GetDhClient().GetToken(),
				Extensions: []*model.RemoteExtension{extension},
			})

			if err != nil {
				log.Fatal(err)
			}

			result := util.ArrayMap[*model.RemoteExtension, *model.Record](resp.Extensions, mapping.ExtensionToRecord)

			describeWriter.WriteRecords(resources.UserResource, 0, util.ArrToChan(result))
		},
	}))

	createRecordNamespace = createRecordCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")

	createCmd.AddCommand(createRecordCmd)
}
