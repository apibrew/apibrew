package test

import (
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	executor2 "github.com/apibrew/apibrew/pkg/formats/executor"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"testing"
)

func apbrApply(inputFilePath string) error {
	dhClient := setup.GetTestDhClient()

	exec := executor2.NewExecutor(executor2.APPLY, dhClient, true, false, false, "", flags.OverrideConfig{})
	return exec.Apply(setup.Ctx, inputFilePath, "yaml")
}

func TestApply(t *testing.T) {
	err := apbrApply("data/test1.yml")
	if err != nil {
		t.Error(err)
	}

	err = apbrApply("data/test2.yml")
	if err != nil {
		t.Error(err)
	}
}
