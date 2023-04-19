package test

import (
	dhctl "github.com/tislib/data-handler/pkg/dhctl"
	"testing"
)

// testing if other parts of application is building

func TestBuild(t *testing.T) {
	// this is just to check if other parts of application is building

	// dhctl
	_ = dhctl.PrepareRootCmd()
}
