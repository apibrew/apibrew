package dhctl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/dhctl/output"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/proto"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create - Create resource from existing table",
}

var writer = output.NewOutputWriter("describe", os.Stdout)

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
				log.Error(err)
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

var createRecordCmd = &cobra.Command{
	Use:                "record",
	Short:              "Create record <resource>",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		defineRootFlags(cmd)

		err := cmd.PersistentFlags().Parse(args)

		if err != nil {
			log.Error(err)
		}

		parseRootFlags(cmd)

		namespace := cmd.Flags().StringP("namespace", "n", "", "")

		if len(args) == 0 {
			log.Fatal("resource not specified")
		}

		resourceName := args[0]

		resource, err := GetDhClient().GetResourceServiceClient().GetByName(cmd.Context(), &stub.GetResourceByNameRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *namespace,
			Name:      resourceName,
		})

		if err != nil {
			log.Error(err)
		}

		fp := flags.NewRecordParserFlags(resource.Resource)

		fp.Declare(cmd)

		var record = &model.Record{}
		fp.Parse(record, cmd, args)

		res, err := GetDhClient().GetRecordServiceClient().Create(cmd.Context(), &stub.CreateRecordRequest{
			Token:     GetDhClient().GetToken(),
			Namespace: *namespace,
			Resource:  resourceName,
			Record:    record,
		})

		if err != nil {
			log.Error(err)
			return
		}

		writer.WriteRecords(resource.Resource, util.ArrToChan(res.Records))
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

			writer.WriteResources(resp.Resources)
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

			writer.WriteRecords(resources.DataSourceResource, util.ArrToChan(result))
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

			writer.WriteRecords(resources.NamespaceResource, util.ArrToChan(result))
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

			writer.WriteRecords(resources.UserResource, util.ArrToChan(result))
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

			writer.WriteRecords(resources.UserResource, util.ArrToChan(result))
		},
	}))

	createCmd.AddCommand(createRecordCmd)
}
