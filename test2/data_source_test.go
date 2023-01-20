package test2

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Setup()

	code := m.Run()
	os.Exit(code)
}
