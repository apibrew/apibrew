package yamlformat

import (
	"github.com/tislib/apibrew/pkg/formats"
	"io"
)

type reader struct {
	input io.Reader
}

func NewReader(input io.Reader) formats.Reader {
	return &reader{input: input}
}
