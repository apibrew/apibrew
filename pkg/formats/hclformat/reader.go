package hclformat

import "io"

type reader struct {
	input io.Reader
}

func NewReader(input io.Reader) Reader {
	return &reader{input: input}
}
