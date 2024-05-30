package apbr

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
	"os"
	"path/filepath"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run nano script",
	Long:  `run nano script on apibrew`,
	RunE: func(cmd *cobra.Command, args []string) error {

		inputFilePath, err := cmd.Flags().GetString("file")
		if err != nil {
			return fmt.Errorf("failed to get input file path: %w", err)
		}
		if inputFilePath == "" {
			return errors.New("file must be provided")
		}

		if err != nil {
			return err
		}

		// locate nano code resource

		parseRootFlags(cmd)

		client := GetClient()

		_, err = client.GetResourceByName(cmd.Context(), "nano", "Code")

		if err != nil {
			return err
		}

		if err := runNanoCode(cmd.Context(), inputFilePath); err != nil {
			return err
		}

		return nil
	},
}

func runNanoCode(ctx context.Context, path string) error {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	var contentBytes []byte
	var language string
	var contentFormat string
	var content string
	if fileInfo.IsDir() {
		contentFormat = "TAR_GZ"
		if _, err := os.Stat(filepath.Join(path, "/index.js")); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			language = "JAVASCRIPT"
		}

		if _, err := os.Stat(filepath.Join(path, "/index.py")); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			language = "PYTHON"
		}

		if language == "" {
			return errors.New("cannot find index.js or index.py in nano code directory")
		}

		if data, err := prepareTarGz(path); err != nil {
			return err
		} else {
			contentBytes = data
		}

		content = base64.StdEncoding.EncodeToString(contentBytes)
	} else {
		contentFormat = "TEXT"
		contentBytes, err = os.ReadFile(path)

		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".js" {
			language = "JAVASCRIPT"
		} else if filepath.Ext(path) == ".py" {
			language = "PYTHON"
		} else {
			return errors.New("invalid file extension for nano code file. Only .js and .py are supported")
		}

		content = string(contentBytes)
	}

	result, err := GetClient().CreateRecord(ctx, "nano", "Script", &model.Record{
		Properties: map[string]*structpb.Value{
			"source":        structpb.NewStringValue(content),
			"language":      structpb.NewStringValue(language),
			"contentFormat": structpb.NewStringValue(contentFormat),
		},
	})

	if err != nil {
		return err
	}

	if result.GetProperties()["output"] != nil {
		var output = result.GetProperties()["output"].AsInterface()

		data, err := json.Marshal(output)

		if err != nil {
			return err
		}

		fmt.Println(string(data))
	}

	return nil
}

func init() {
	runCmd.PersistentFlags().StringP("file", "f", "", "Input file")
}
