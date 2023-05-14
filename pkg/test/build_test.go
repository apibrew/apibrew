package test

import (
	apbr "github.com/apibrew/apibrew/pkg/apbr"
	"testing"
)

// testing if other parts of application is building

func TestBuild(t *testing.T) {
	// this is just to check if other parts of application is building

	// apbr
	_ = apbr.PrepareRootCmd()
}
