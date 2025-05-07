package test

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
)

var recordClient stub.RecordClient

var container service.Container

func init() {
	recordClient = setup.RecordClient
	container = setup.GetContainer()
}
