package dhctl

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/batch"
	"github.com/tislib/data-handler/pkg/dhctl/flags"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/server/util"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply - apply resources",
	Run: func(cmd *cobra.Command, args []string) {
		parseRootFlags(cmd)

		file, err := cmd.Flags().GetString("file")
		check(err)

		migrate, err := cmd.Flags().GetBool("migrate")
		check(err)

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		force, err := cmd.Flags().GetBool("force")
		check(err)

		var overrideConfig = new(flags.OverrideConfig)
		overrideFlags.Parse(overrideConfig, cmd, args)

		if file == "" {
			log.Fatal("file should provided")
		}

		if strings.HasSuffix(file, ".pbe") {
			in, err := os.Open(file)

			check(err)

			batchExecutor := batch.NewExecutor(batch.ExecutorParams{
				Input:                 in,
				Token:                 GetDhClient().GetToken(),
				ResourceServiceClient: GetDhClient().GetResourceServiceClient(),
				RecordServiceClient:   GetDhClient().GetRecordServiceClient(),
				OverrideConfig: batch.OverrideConfig{
					Namespace:  overrideConfig.Namespace,
					DataSource: overrideConfig.DataSource,
				},
			})

			err = batchExecutor.Restore(context.TODO(), in)

			check(err)

			return
		} else if strings.HasSuffix(file, "yml") || strings.HasSuffix(file, "yaml") {
			fileData, err := os.ReadFile(file)

			check(err)
			applyYaml(fileData, migrate, namespace, force, overrideConfig)
		}

		log.Println(migrate)

	},
}

func applyYaml(fileData []byte, migrate bool, namespace string, force bool, overrideConfig *flags.OverrideConfig) {
	var jsonUMo = protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: false,
		Resolver:       nil,
	}

	var createRecords []*model.Record
	var updateRecords []*model.Record
	decoder := yaml.NewDecoder(bytes.NewReader(fileData))

	for {
		var body map[string]interface{}
		var err = decoder.Decode(&body)

		if err == io.EOF {
			break
		}

		check(err)

		body = convert(body).(map[string]interface{})

		switch body["type"].(string) {
		case "resource":
			delete(body, "type")

			jsonData, err := json.Marshal(body)

			check(err)

			var resource = new(model.Resource)
			err = jsonUMo.Unmarshal(jsonData, resource)

			check(err)

			// locating resource
			if resource.Id == "" {
				resp, err := GetDhClient().GetResourceServiceClient().GetByName(context.TODO(), &stub.GetResourceByNameRequest{
					Token:     GetDhClient().GetToken(),
					Namespace: resource.Namespace,
					Name:      resource.Name,
				})

				if err != nil && util.GetErrorCode(err) != model.ErrorCode_RESOURCE_NOT_FOUND {
					panic(err)
				}

				if resp != nil && resp.Resource != nil {
					resource.Id = resp.Resource.Id
				}
			}

			if resource.Id != "" {
				_, err := GetDhClient().GetResourceServiceClient().Update(context.TODO(), &stub.UpdateResourceRequest{
					Token:          GetDhClient().GetToken(),
					Resources:      []*model.Resource{resource},
					DoMigration:    migrate,
					ForceMigration: force,
				})

				check(err)

				log.Println("resource updated: " + resource.Name)
			} else {
				_, err := GetDhClient().GetResourceServiceClient().Create(context.TODO(), &stub.CreateResourceRequest{
					Token:          GetDhClient().GetToken(),
					Resources:      []*model.Resource{resource},
					DoMigration:    migrate,
					ForceMigration: force,
				})

				check(err)

				log.Println("resource created: " + resource.Name)
			}
		case "record":
			delete(body, "type")

			jsonData, err := json.Marshal(body)

			check(err)

			var record = new(model.Record)
			err = jsonUMo.Unmarshal(jsonData, record)

			check(err)

			if record.Id != "" {
				updateRecords = append(updateRecords, record)
			} else {
				createRecords = append(createRecords, record)
			}

		}
	}

	if len(updateRecords) > 0 {
		_, err := GetDhClient().GetRecordServiceClient().Update(context.TODO(), &stub.UpdateRecordRequest{
			Token:        GetDhClient().GetToken(),
			Namespace:    namespace,
			Records:      updateRecords,
			CheckVersion: false,
		})

		check(err)

		log.Println("Record updated: " + strconv.Itoa(len(updateRecords)))
	}

	if len(createRecords) > 0 {
		_, err := GetDhClient().GetRecordServiceClient().Create(context.TODO(), &stub.CreateRecordRequest{
			Token:          GetDhClient().GetToken(),
			Namespace:      namespace,
			Records:        createRecords,
			IgnoreIfExists: true,
		})

		check(err)

		log.Println("Record created: " + strconv.Itoa(len(createRecords)))
	}
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func init() {
	applyCmd.PersistentFlags().StringP("file", "f", "", "Output file")
	applyCmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace")
	applyCmd.PersistentFlags().BoolP("migrate", "m", false, "Migrate")
	applyCmd.PersistentFlags().Bool("force", false, "Force")

	overrideFlags.Declare(applyCmd)
}
