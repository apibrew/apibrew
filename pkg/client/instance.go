package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"time"
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
	watchClient          stub.WatchClient
	token                string
}

func (d *dhClient) ListenRecords(ctx context.Context, namespace string, resource string, consumer func(records []*model.Record)) error {
	resp, err := d.watchClient.Watch(ctx, &stub.WatchRequest{
		Token: d.token,
		Selector: &model.EventSelector{
			Actions:    []model.Event_Action{model.Event_CREATE, model.Event_DELETE, model.Event_UPDATE},
			Namespaces: []string{namespace},
			Resources:  []string{resource},
		},
	})

	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = resp.CloseSend()

			time.Sleep(1 * time.Second)

			log.Println("Reconnecting to watch stream")

			err := d.ListenRecords(ctx, namespace, resource, consumer)

			if err != nil {
				panic(err) // need to improve
			}
		}()
		for {
			_, err := resp.Recv()

			if err != nil {
				break
			}

			list, _, err := d.ListRecords(ctx, service.RecordListParams{
				Namespace: namespace,
				Resource:  resource,
				Limit:     10000,
			})

			if err != nil {
				log.Error(err)
				break
			}

			consumer(list)

			select {
			case <-ctx.Done():
				break
			default:
				continue
			}
		}
	}()

	return nil
}

func (d *dhClient) DeleteResource(ctx context.Context, id string, doMigration bool, forceMigration bool) error {
	_, err := d.resourceClient.Delete(ctx, &stub.DeleteResourceRequest{
		Token:          d.token,
		Ids:            []string{id},
		DoMigration:    doMigration,
		ForceMigration: forceMigration,
	})

	return err
}

func (d *dhClient) GetResourceByName(ctx context.Context, namespace string, getType string) (*model.Resource, error) {
	resp, err := d.resourceClient.GetByName(ctx, &stub.GetResourceByNameRequest{
		Token:     d.token,
		Namespace: namespace,
		Name:      getType,
	})

	if err != nil {
		return nil, err
	}

	return resp.Resource, nil
}

func (d *dhClient) ReadRecordStream(ctx context.Context, params service.RecordListParams, recordsChan chan *model.Record) error {
	resp, err := d.recordClient.ReadStream(ctx, &stub.ReadStreamRequest{})

	if err != nil {
		return err
	}

	go func() {
		defer func() {
			close(recordsChan)
			_ = resp.CloseSend()
		}()
		for {
			record, err := resp.Recv()

			if err != nil {
				break
			}

			recordsChan <- record

			select {
			case <-ctx.Done():
				break
			default:
				continue
			}
		}
	}()

	return nil
}

func (d *dhClient) CreateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := d.recordClient.Create(ctx, &stub.CreateRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp.Records[0], nil
}

func (d *dhClient) UpdateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := d.recordClient.Update(ctx, &stub.UpdateRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp.Records[0], nil
}

func (d *dhClient) ApplyRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
	resp, err := d.recordClient.Apply(ctx, &stub.ApplyRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  resource,
		Records:   []*model.Record{record},
	})

	if err != nil {
		return nil, err
	}

	return resp.Records[0], nil
}

func (d *dhClient) GetRecord(ctx context.Context, namespace string, resource string, id string) (*model.Record, error) {
	resp, err := d.recordClient.Get(ctx, &stub.GetRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  resource,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return resp.Record, nil
}

func (d *dhClient) ListRecords(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	req := params.ToRequest()

	req.Token = d.token

	resp, err := d.recordClient.List(ctx, req)

	if err != nil {
		return nil, 0, err
	}

	return resp.Content, resp.Total, nil
}

func (d *dhClient) ListResources(ctx context.Context) ([]*model.Resource, error) {
	resp, err := d.resourceClient.List(ctx, &stub.ListResourceRequest{
		Token: d.token,
	})

	if err != nil {
		return nil, err
	}

	return resp.Resources, nil
}

func (d *dhClient) UpdateTokenFromContext(ctx context.Context) {
	d.params.Token = ctx.Value("token").(string)
	d.token = ctx.Value("token").(string)
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
	d.token = authResp.Token.Content

	return nil
}

func (d *dhClient) AuthenticateWithToken(token string) {
	d.params.Token = token
	d.token = token
}

func (d *dhClient) DeleteRecord(ctx context.Context, namespace string, name string, id string) error {
	_, err := d.recordClient.Delete(ctx, &stub.DeleteRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  name,
		Ids:       []string{id},
	})

	return err
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
		token:                params.Token,
	}, nil
}

func NewDhClientLocal(serverName string) (DhClient, error) {
	configServer := locateConfigServer(serverName)

	var params = DhClientParams{
		Addr:     configServer.Host,
		Insecure: configServer.Insecure,
	}
	var opts []grpc.DialOption
	if params.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		creds := credentials.NewClientTLSFromCert(nil, "")

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(params.Addr, opts...)

	if err != nil {
		return nil, err
	}

	var dhc = &dhClient{
		conn:                 conn,
		params:               params,
		recordClient:         stub.NewRecordClient(conn),
		authenticationClient: stub.NewAuthenticationClient(conn),
		resourceClient:       stub.NewResourceClient(conn),
		dataSourceClient:     stub.NewDataSourceClient(conn),
		watchClient:          stub.NewWatchClient(conn),
		token:                params.Token,
	}

	if configServer.Authentication.Token != "" {
		dhc.AuthenticateWithToken(configServer.Authentication.Token)
	} else {
		err = dhc.AuthenticateWithUsernameAndPassword(configServer.Authentication.Username, configServer.Authentication.Password)

		if err != nil {
			return nil, err
		}
	}

	return dhc, nil
}
