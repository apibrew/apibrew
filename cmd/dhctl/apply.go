package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/model"
	"github.com/tislib/data-handler/server/stub"
	"github.com/tislib/data-handler/server/util"
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
		initClient(cmd.Context())

		file, err := cmd.Flags().GetString("file")
		check(err)

		migrate, err := cmd.Flags().GetBool("migrate")
		check(err)

		namespace, err := cmd.Flags().GetString("namespace")
		check(err)

		force, err := cmd.Flags().GetBool("force")
		check(err)

		if file == "" {
			log.Fatal("file should provided")
		}

		fileData, err := os.ReadFile(file)

		check(err)

		var jsonUMo = protojson.UnmarshalOptions{
			AllowPartial:   false,
			DiscardUnknown: false,
			Resolver:       nil,
		}

		var createRecords []*model.Record
		var updateRecords []*model.Record

		if strings.HasSuffix(file, "yml") || strings.HasSuffix(file, "yaml") {
			decoder := yaml.NewDecoder(bytes.NewReader(fileData))

			for {
				var body map[string]interface{}
				err = decoder.Decode(&body)

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
						resp, err := resourceServiceClient.GetByName(context.TODO(), &stub.GetResourceByNameRequest{
							Token:     authToken,
							Namespace: resource.Namespace,
							Name:      resource.Name,
						})

						if err != nil && util.GetErrorCode(err) != model.ErrorCode_RECORD_NOT_FOUND {
							panic(err)
						}

						if resp.Resource != nil {
							resource.Id = resp.Resource.Id
						}
					}

					if resource.Id != "" {
						_, err := resourceServiceClient.Update(context.TODO(), &stub.UpdateResourceRequest{
							Token:          authToken,
							Resources:      []*model.Resource{resource},
							DoMigration:    migrate,
							ForceMigration: force,
						})

						check(err)

						log.Println("Resource updated: " + resource.Name)
					} else {
						_, err := resourceServiceClient.Create(context.TODO(), &stub.CreateResourceRequest{
							Token:          authToken,
							Resources:      []*model.Resource{resource},
							DoMigration:    migrate,
							ForceMigration: force,
						})

						check(err)

						log.Println("Resource created: " + resource.Name)
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
		}

		if len(updateRecords) > 0 {
			_, err := recordServiceClient.Update(context.TODO(), &stub.UpdateRecordRequest{
				Token:        authToken,
				Namespace:    namespace,
				Records:      updateRecords,
				CheckVersion: false,
			})

			check(err)

			log.Println("Record updated: " + strconv.Itoa(len(updateRecords)))
		}

		if len(createRecords) > 0 {
			_, err := recordServiceClient.Create(context.TODO(), &stub.CreateRecordRequest{
				Token:          authToken,
				Namespace:      namespace,
				Records:        createRecords,
				IgnoreIfExists: true,
			})

			check(err)

			log.Println("Record created: " + strconv.Itoa(len(createRecords)))
		}

		log.Println(migrate)

	},
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
}
