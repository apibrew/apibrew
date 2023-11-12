package main

import (
	"github.com/apibrew/apibrew/pkg/generator/templates/golang"
	"github.com/apibrew/apibrew/pkg/resources"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := golang.GenerateGoResourceCode("github.com/apibrew/apibrew/resource_model", resources.GetAllSystemResources(), nil, "pkg/resource_model", "resource_model")

	if err != nil {
		log.Fatal(err)
	}
}
