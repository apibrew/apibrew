package data_source

import (
	"data-handler/test2/lib"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	lib.Setup()

	code := m.Run()
	os.Exit(code)
}
