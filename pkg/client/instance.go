package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DhClientParams struct {
	Addr     string
	Insecure bool
	Token    string
}

type dhClient struct {
	conn                 *grpc.ClientConn
	params               DhClientParams
	recordClient         stub.RecordClient
	authenticationClient stub.AuthenticationClient
	resourceClient       stub.ResourceClient
	dataSourceClient     stub.DataSourceClient
	genericClient        stub.GenericClient
}

func (d *dhClient) ListResources(ctx context.Context) ([]*model.Resource, error) {
	resp, err := d.resourceClient.List(ctx, &stub.ListResourceRequest{
		Token: d.GetToken(),
	})

	if err != nil {
		return nil, err
	}

	return resp.Resources, nil
}

func (d *dhClient) UpdateTokenFromContext(ctx context.Context) {
	d.params.Token = ctx.Value("token").(string)
}

func (d *dhClient) AuthenticateWithUsernameAndPassword(username string, password string) error {
	authResp, err := d.authenticationClient.Authenticate(context.TODO(), &stub.AuthenticationRequest{
		Username: username,
		Password: password,
		Term:     2,
	})

	if err != nil {
		return err
	}

	d.params.Token = authResp.Token.Content

	return nil
}

func (d *dhClient) AuthenticateWithToken(token string) {
	d.params.Token = token
}

func (d *dhClient) GetToken() string {
	return d.params.Token
}

func (d *dhClient) GetAuthenticationClient() stub.AuthenticationClient {
	return d.authenticationClient
}

func (d *dhClient) GetDataSourceClient() stub.DataSourceClient {
	return d.dataSourceClient
}

func (d *dhClient) GetResourceClient() stub.ResourceClient {
	return d.resourceClient
}

func (d *dhClient) GetRecordClient() stub.RecordClient {
	return d.recordClient
}

func (d *dhClient) GetGenericClient() stub.GenericClient {
	return d.genericClient
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
		conn:                 conn,
		params:               params,
		recordClient:         stub.NewRecordClient(conn),
		authenticationClient: stub.NewAuthenticationClient(conn),
		resourceClient:       stub.NewResourceClient(conn),
		dataSourceClient:     stub.NewDataSourceClient(conn),
		genericClient:        stub.NewGenericClient(conn),
	}, nil
}
