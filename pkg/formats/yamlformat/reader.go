package yamlformat

import (
	"github.com/apibrew/apibrew/pkg/formats"
	"io"
)

type reader struct {
	input io.Reader
}

func NewReader(input io.Reader) formats.Reader {
	return &reader{input: input}
}
