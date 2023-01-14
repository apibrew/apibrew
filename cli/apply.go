package main

import (
	"bytes"
	"context"
	"data-handler/grpc/stub"
	"data-handler/model"
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
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

		if strings.HasSuffix(file, "yml") || strings.HasSuffix(file, "yaml") {
			var body map[string]interface{}

			decoder := yaml.NewDecoder(bytes.NewReader(fileData))

			for {
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
							Workspace: resource.Workspace,
							Name:      resource.Name,
						})

						check(err)

						if resp.Error != nil {
							if resp.Error.Code != model.ErrorCode_RECORD_NOT_FOUND {
								check(errors.New(resp.Error.Message))
							}
						}

						checkError(resp.Error)

						if resp.Resource != nil {
							resource.Id = resp.Resource.Id
						}
					}

					if resource.Id != "" {
						resp, err := resourceServiceClient.Update(context.TODO(), &stub.UpdateResourceRequest{
							Token:          authToken,
							Resources:      []*model.Resource{resource},
							DoMigration:    migrate,
							ForceMigration: force,
						})

						check(err)

						checkError(resp.Error)

						log.Println("Resource updated: " + resource.Name)
					} else {
						resp, err := resourceServiceClient.Create(context.TODO(), &stub.CreateResourceRequest{
							Token:          authToken,
							Resources:      []*model.Resource{resource},
							DoMigration:    migrate,
							ForceMigration: force,
						})

						check(err)

						checkError(resp.Error)

						log.Println("Resource created: " + resource.Name)
					}
				}
			}
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
	applyCmd.PersistentFlags().BoolP("migrate", "m", false, "Migrate")
	applyCmd.PersistentFlags().Bool("force", false, "Force")
}