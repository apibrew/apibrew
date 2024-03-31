package main

import (
	"encoding/json"
	"github.com/apibrew/apibrew/cmd/modules/gen"
	"github.com/prometheus/common/log"
	"os"
	"os/exec"
)

func main() {
	modulesJson, err := os.ReadFile("modules.json")

	if err != nil {
		panic(err)
	}

	var modules = make(map[string]string)

	err = json.Unmarshal(modulesJson, &modules)

	if err != nil {
		panic(err)
	}

	for name, version := range modules {
		importModule(name, version)
	}

	code := gen.GenerateModulesContent(modules)

	err = os.WriteFile("module/modules.go", []byte(code), 0644)

	if err != nil {
		panic(err)
	}
}

func importModule(name string, version string) {
	cmd := exec.Command("go", "get", name+"@"+version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
