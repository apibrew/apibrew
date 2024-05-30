package apbr

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SyncConfigItem struct {
	Type               string   `json:"type"`
	PathByProperty     string   `json:"pathByProperty"`
	FileNameByProperty string   `json:"fileNameByProperty"`
	Path               string   `json:"path"`
	ResolveReferences  []string `json:"resolveReferences"`
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync apibrew resources and data",
	Long:  `sync apibrew resources and data`,
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		var path = cmd.Flag("path").Value.String()
		var mode = cmd.Flag("mode").Value.String()
		var config = cmd.Flag("config").Value.String()

		var configItems []SyncConfigItem
		var configAbs = path + "/" + config

		if strings.HasPrefix("/", config) {
			configAbs = config
		}

		if err := readSyncConfig(configAbs, &configItems); err != nil {
			return err
		}

		var api = client.NewInterface(dhClient)

		if mode == "pull" {
			if err := syncPull(cmd.Context(), api, configItems, path); err != nil {
				return err
			}
		} else if mode == "push-plan" {
			if err := syncPush(cmd.Context(), api, configItems, path, false); err != nil {
				return err
			}
		} else if mode == "push" {
			if err := syncPush(cmd.Context(), api, configItems, path, true); err != nil {
				return err
			}
		} else {
			return errors.New("Invalid mode")
		}

		return nil
	},
}

func readSyncConfig(configFile string, i *[]SyncConfigItem) error {
	// read configFile as json into i
	fp, err := os.Open(configFile)

	if err != nil {
		return err
	}

	defer fp.Close()

	if err := json.NewDecoder(fp).Decode(i); err != nil {
		return err
	}

	return nil
}

func syncPull(ctx context.Context, api api.Interface, configItems []SyncConfigItem, path string) error {
	for _, item := range configItems {
		if err := syncPullItem(ctx, api, item, path); err != nil {
			return err
		}
	}

	return nil
}

func syncPullItem(ctx context.Context, apiInterface api.Interface, config SyncConfigItem, basePath string) error {
	var path = basePath + "/" + config.Path

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	var resource, err = apiInterface.GetResourceByType(ctx, config.Type)

	if err != nil {
		return err
	}

	sourceData, err := apiInterface.List(ctx, api.ListParams{
		Type:              config.Type,
		Limit:             10000,
		ResolveReferences: config.ResolveReferences,
	})

	if err != nil {
		return err
	}

	for _, data := range sourceData.Content {
		var resourcePath = path
		data["type"] = config.Type

		if config.Type == "system/Resource" {
			var namespace = data["namespace"].(map[string]interface{})["name"]
			if namespace == "system" || namespace == "studio" || namespace == "apps" || namespace == "nano" {
				continue
			}
		}

		if config.PathByProperty != "" {
			if value, ok := data[config.PathByProperty]; ok {
				resourcePath = path + "/" + value.(string)
			}
		}

		var fileNameProperty = "id"

		if config.FileNameByProperty != "" {
			fileNameProperty = config.FileNameByProperty
		}

		if err := runSyncApplyModifiers(ctx, resource, data, config); err != nil {
			return err
		}

		if config.Type == "nano/Code" {
			if err := storeNanoCode(resourcePath, data, fileNameProperty); err != nil {
				return err
			}
		} else {
			if err := storeFile(resourcePath, data, fileNameProperty); err != nil {
				return err
			}
		}

	}

	return nil
}

func runSyncApplyModifiers(ctx context.Context, resource *resource_model.Resource, data unstructured.Unstructured, config SyncConfigItem) error {
	var typ = data["type"].(string)

	if typ == "system/Resource" {
		return runSyncApplyModifiersForResource(ctx, data, resource, config)
	}

	return nil
}

func runSyncApplyModifiersForResource(ctx context.Context, data unstructured.Unstructured, resource *resource_model.Resource, config SyncConfigItem) error {
	return nil
}

func storeFile(resourcePath string, data unstructured.Unstructured, fileNameProperty string) error {
	fileName := resourcePath + "/" + data[fileNameProperty].(string) + ".yaml"

	folderPath := filepath.Dir(fileName)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	fp, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer fp.Close()

	var yamlWriter = yamlformat.NewWriter(fp, make(map[string]string))

	if err := yamlWriter.WriteRecord(data["type"].(string), data); err != nil {
		return err
	}

	log.Println("Wrote file: " + fileName)

	return nil
}

func storeNanoCode(resourcePath string, data unstructured.Unstructured, fileNameProperty string) error {
	fileName := resourcePath + "/" + data[fileNameProperty].(string) + ".js"

	folderPath := filepath.Dir(fileName)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	fp, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer fp.Close()

	if _, err := fp.WriteString(data["content"].(string)); err != nil {
		return err
	}

	log.Println("Wrote file: " + fileName)

	return nil
}

func syncPush(ctx context.Context, apiInterface api.Interface, configItems []SyncConfigItem, path string, execute bool) error {
	for _, item := range configItems {
		if err := syncPushItem(ctx, apiInterface, item, path, execute); err != nil {
			return err
		}
	}

	return nil
}

func syncPushItem(ctx context.Context, apiInterface api.Interface, item SyncConfigItem, path string, execute bool) error {
	return nil
}

func init() {
	syncCmd.PersistentFlags().StringP("path", "p", ".", "Path")
	syncCmd.PersistentFlags().StringP("config", "c", "apbr-sync-config.json", "Config")
	syncCmd.PersistentFlags().StringP("mode", "m", "pull", "Mode")

}
