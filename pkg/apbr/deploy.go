package apbr

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"os"
	"path/filepath"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy nano code",
	Long:  `Deploy nano code to apibrew`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var name = cmd.Flag("name").Value.String()
		var override, err = cmd.PersistentFlags().GetBool("override")

		inputFilePathArr, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			return fmt.Errorf("failed to get input file path: %w", err)
		}
		if len(inputFilePathArr) == 0 {
			return errors.New("file must be provided")
		}

		if err != nil {
			return err
		}

		if len(args) > 1 && name != "" {
			return errors.New("cannot specify name when deploying multiple nano code files")
		}

		// locate nano code resource

		parseRootFlags(cmd)

		client := GetClient()

		_, err = client.GetResourceByName(cmd.Context(), "nano", "Code")

		if err != nil {
			return err
		}

		for _, nanoCodeFile := range inputFilePathArr {
			if name == "" {
				name = nanoCodeFile
			}
			if err := deployNanoCode(cmd.Context(), nanoCodeFile, name, override); err != nil {
				return err
			}
		}

		return nil
	},
}

func deployNanoCode(ctx context.Context, path string, name string, override bool) error {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	var contentBytes []byte
	var language string
	var contentFormat string
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
	}

	content := base64.StdEncoding.EncodeToString(contentBytes)

	var record *model.Record
	if !override {
		record, err = GetClient().CreateRecord(ctx, "nano", "Code", &model.Record{
			Properties: map[string]*structpb.Value{
				"name":          structpb.NewStringValue(name),
				"content":       structpb.NewStringValue(content),
				"language":      structpb.NewStringValue(language),
				"contentFormat": structpb.NewStringValue(contentFormat),
			},
		})

		if err != nil {
			return err
		}
	} else {
		record, err = GetClient().ApplyRecord(ctx, "nano", "Code", &model.Record{
			Properties: map[string]*structpb.Value{
				"name":          structpb.NewStringValue(name),
				"content":       structpb.NewStringValue(content),
				"language":      structpb.NewStringValue(language),
				"contentFormat": structpb.NewStringValue(contentFormat),
			},
		})

		if err != nil {
			return err
		}
	}

	fmt.Printf("Deployed nano code %s with id: %s\n", path, record.Id)

	return nil
}

func prepareTarGz(path string) ([]byte, error) {
	bw := bytes.NewBuffer(nil)
	gw := gzip.NewWriter(bw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	walkFn := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}
		// Because of scoping we can reference the external root_directory variable
		new_path := path[len(path):]
		if len(new_path) == 0 {
			return nil
		}
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		if h, err := tar.FileInfoHeader(info, new_path); err != nil {
			return err
		} else {
			h.Name = new_path
			if err = tw.WriteHeader(h); err != nil {
				return err
			}
		}
		if length, err := io.Copy(tw, fr); err != nil {
			return err
		} else {
			fmt.Println(length)
		}
		return nil
	}

	if err := filepath.Walk(path, walkFn); err != nil {
		return nil, err
	}
	return bw.Bytes(), nil
}

func init() {
	deployCmd.PersistentFlags().String("name", "", "unique code name")
	deployCmd.PersistentFlags().StringArrayP("file", "f", nil, "Input file(s)")
	deployCmd.PersistentFlags().Bool("override", false, "Override if code already exists")
}
