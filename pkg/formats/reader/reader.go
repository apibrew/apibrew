package reader

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats/jsonformat"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
	"io"
	"os"
	"strings"
)

type Reader struct {
	RecordReceiver func(record interface{}) error
}

func (r *Reader) Read(ctx context.Context, inputFilePath string, format string, handler func(unstructured.Unstructured) error) error {
	if strings.HasSuffix(inputFilePath, "json") {
		format = "json"
	} else if strings.HasSuffix(inputFilePath, ".yaml") || strings.HasSuffix(inputFilePath, ".yml") {
		format = "yaml"
	}

	if format == "yml" {
		format = "yaml"
	}

	var in io.Reader
	if strings.HasPrefix(inputFilePath, "https://") {
		in = strings.NewReader("$include: " + inputFilePath)
	} else {
		file, err := os.Open(inputFilePath)
		if err != nil {
			return fmt.Errorf("failed to open YAML file: %w", err)
		}
		defer file.Close()
		in = file
	}

	var parseFunc func(r io.Reader, ctx context.Context, handler func(unstructured.Unstructured) error) error

	switch {
	case format == "yaml":
		parseFunc = yamlformat.Parse
	case format == "json":
		parseFunc = jsonformat.Parse
	default:
		return fmt.Errorf("unsupported file format: %s", inputFilePath)
	}

	return parseFunc(in, ctx, handler)
}

func (r *Reader) ReadWithPattern(ctx context.Context, inputFilePath string, format string, handler func(unstructured.Unstructured) error) error {
	log.Info("Apply pattern: ", inputFilePath, " ...")

	if strings.Contains(inputFilePath, "*") {
		filenames, err := filepathx.Glob(inputFilePath)

		if err != nil {
			return err
		}

		for _, filename := range filenames {
			log.Info("Apply file: ", filename)

			err := r.Read(ctx, filename, format, handler)
			if err != nil {
				return err
			}
		}
	} else {
		return r.Read(ctx, inputFilePath, format, handler)
	}

	return nil
}
