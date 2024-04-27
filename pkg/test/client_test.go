package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
)

var authenticationClient stub.AuthenticationClient

var resourceClient stub.ResourceClient
var recordClient stub.RecordClient
var dataSourceClient stub.DataSourceClient

var container service.Container
var apiInterface api.Interface
var apiDirectInterface api.Interface

func init() {
	recordClient = setup.RecordClient
	authenticationClient = setup.AuthenticationClient
	resourceClient = setup.ResourceClient
	dataSourceClient = setup.DataSourceClient
	container = setup.GetContainer()
	apiInterface = client.NewInterface(setup.GetTestDhClient())
	apiDirectInterface = api.NewInterface(container)
}
