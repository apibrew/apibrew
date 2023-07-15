package test

import (
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
)

var authenticationClient stub.AuthenticationClient

var resourceClient stub.ResourceClient
var recordClient stub.RecordClient
var dataSourceClient stub.DataSourceClient

var container service.Container

func init() {
	dhClient := setup.GetTestDhClient()
	recordClient = dhClient.GetRecordClient()
	authenticationClient = dhClient.GetAuthenticationClient()
	resourceClient = dhClient.GetResourceClient()
	dataSourceClient = dhClient.GetDataSourceClient()
	container = setup.GetContainer()
}
