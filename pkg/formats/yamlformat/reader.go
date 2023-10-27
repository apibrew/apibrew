package yamlformat

import (
	"context"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"gopkg.in/yaml.v3"
	"io"
)

func Parse(r io.Reader, ctx context.Context, handler func(unstructured.Unstructured) error) error {
	decoder := yaml.NewDecoder(r)

	for {
		if ctx.Err() != nil {
			break
		}

		var body unstructured.Unstructured
		var err = decoder.Decode(&body)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		err = handler(body)

		if err != nil {
			return err
		}

	}

	return nil
}
