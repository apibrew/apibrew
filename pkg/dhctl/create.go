package dhctl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/dhctl/output"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/proto"
	"os"
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
	Use:   "record",
	Short: "Create record",
	Run: func(cmd *cobra.Command, args []string) {

		defineRootFlags(cmd)
		cmd.Flags().String("name", "", "")

		err := cmd.Flags().Parse(args)
		check(err)

		parseRootFlags(cmd)

		log.Print(cmd.Flags().Args())
	},
}

func initCreateCmd() {
	writer := output.NewOutputWriter("describe", os.Stdout)

	var migrate *bool
	var force *bool

	createCmd.AddCommand(protoMessageCreateCmd[*model.Resource](protoMessageCreateCmdParams[*model.Resource]{
		msg: &model.Resource{},
		use: "resource",
		before: func(cmd *cobra.Command) {
			migrate = cmd.PersistentFlags().BoolP("migrate", "m", true, "")
			force = cmd.PersistentFlags().BoolP("force", "f", true, "")
		},
		handle: func(resource *model.Resource) {
			log.Println(resource)
			resp, err := resourceServiceClient.Create(context.TODO(), &stub.CreateResourceRequest{
				Token:          authToken,
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
	}))
}
