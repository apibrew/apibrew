package apbr

import (
	"errors"
	"github.com/apibrew/apibrew/pkg/apbr/output"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var toolsCmd = &cobra.Command{
	Use:    "tools",
	Hidden: true,
	Short:  "tools - tools resource/record docs: https://apibrew.io/docs/cli#tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		parseRootFlags(cmd)

		if len(args) == 0 {
			return errors.New("tool name is expected in args")
		}

		toolName := args[0]

		switch toolName {
		case "predict-resource-from-json":
			return predictResourceFromJson(cmd, args[1:])
		default:
			return errors.New("unknown tool name")
		}
	},
}

func predictResourceFromJson(cmd *cobra.Command, args []string) error {
	writer := output.NewOutputWriter("yaml", os.Stdout, map[string]string{
		"for-apply": "true",
	})

	resource := new(resource_model.Resource)
	resource.Properties = make(map[string]resource_model.Property)

	for _, arg := range args {
		// read file
		file, err := os.Open(arg)

		if err != nil {
			return err
		}

		// parse file

		decoder := yaml.NewDecoder(file)

		var data = new(interface{})
		err = decoder.Decode(data)

		if err != nil {
			return err
		}

		if arr, ok := (*data).([]interface{}); ok {
			for _, item := range arr {
				predictNext(resource, item.(map[string]interface{}))
			}
		} else {
			predictNext(resource, (*data).(map[string]interface{}))
		}
	}

	return writer.WriteResource(resource)
}

func predictNext(resource *resource_model.Resource, properties map[string]interface{}) {
	for key, value := range properties {
		if _, ok := resource.Properties[key]; !ok {
			resource.Properties[key] = resource_model.Property{}
		}

		prop := resource.Properties[key]

		switch value.(type) {
		case string:
			prop.Type = resource_model.ResourceType_STRING
		}

		resource.Properties[key] = prop
	}
}
