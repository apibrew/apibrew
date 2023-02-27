package client

import (
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DhClientParams struct {
	Addr     string
	Insecure bool
	Token    string
}

type dhClient struct {
	conn                        *grpc.ClientConn
	params                      DhClientParams
	recordServiceClient         stub.RecordServiceClient
	authenticationServiceClient stub.AuthenticationServiceClient
	resourceServiceClient       stub.ResourceServiceClient
	dataSourceServiceClient     stub.DataSourceServiceClient
	userServiceClient           stub.UserServiceClient
	extensionServiceClient      stub.ExtensionServiceClient
	genericServiceClient        stub.GenericServiceClient
}

func (d *dhClient) GetToken() string {
	return d.params.Token
}

func (d *dhClient) GetAuthenticationServiceClient() stub.AuthenticationServiceClient {
	return d.authenticationServiceClient
}

func (d *dhClient) GetDataSourceServiceClient() stub.DataSourceServiceClient {
	return d.dataSourceServiceClient
}

func (d *dhClient) GetResourceServiceClient() stub.ResourceServiceClient {
	return d.resourceServiceClient
}

func (d *dhClient) GetRecordServiceClient() stub.RecordServiceClient {
	return d.recordServiceClient
}

func (d *dhClient) GetGenericServiceClient() stub.GenericServiceClient {
	return d.genericServiceClient
}

func (d *dhClient) GetExtensionServiceClient() stub.ExtensionServiceClient {
	return d.extensionServiceClient
}

func (d *dhClient) GetUserServiceClient() stub.UserServiceClient {
	return d.userServiceClient
}

func NewDhClient(params DhClientParams) (DhClient, error) {
	var opts []grpc.DialOption
	if params.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(params.Addr, opts...)

	if err != nil {
		return nil, err
	}

	return &dhClient{
		conn:                        conn,
		params:                      params,
		recordServiceClient:         stub.NewRecordServiceClient(conn),
		authenticationServiceClient: stub.NewAuthenticationServiceClient(conn),
		resourceServiceClient:       stub.NewResourceServiceClient(conn),
		dataSourceServiceClient:     stub.NewDataSourceServiceClient(conn),
		userServiceClient:           stub.NewUserServiceClient(conn),
		extensionServiceClient:      stub.NewExtensionServiceClient(conn),
		genericServiceClient:        stub.NewGenericServiceClient(conn),
	}, nil
}