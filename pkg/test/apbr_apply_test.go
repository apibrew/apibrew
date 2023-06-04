package test

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"os"
	"testing"
)

func apbrApply(inputFilePath string) error {
	in, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to open YAML file: %w", err)
	}
	defer in.Close()

	dhClient := setup.GetTestDhClient()
	token := dhClient.GetToken()

	executor, err := yamlformat.NewExecutor(yamlformat.ExecutorParams{
		Input:          in,
		Token:          token,
		DhClient:       dhClient,
		DoMigration:    true,
		ForceMigration: true,
	})
	if err != nil {
		return fmt.Errorf("failed to create YAML executor: %w", err)
	}

	return executor.Restore(setup.Ctx)
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
