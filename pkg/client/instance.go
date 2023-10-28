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
	"strconv"
	"time"
)

type Params struct {
	Addr     string
	Insecure bool
	Token    string
}

type client struct {
	conn                 *grpc.ClientConn
	params               Params
	recordClient         stub.RecordClient
	authenticationClient stub.AuthenticationClient
	resourceClient       stub.ResourceClient
	dataSourceClient     stub.DataSourceClient
	watchClient          stub.WatchClient
	eventChannelClient   stub.EventChannelClient
	token                string
}

func (d *client) CreateResource(ctx context.Context, resource *model.Resource, migration bool, force bool) error {
	_, err := d.resourceClient.Create(ctx, &stub.CreateResourceRequest{
		Token:          d.token,
		Resources:      []*model.Resource{resource},
		DoMigration:    migration,
		ForceMigration: force,
	})

	return err
}

func (d *client) UpdateResource(ctx context.Context, resource *model.Resource, migration bool, force bool) error {
	_, err := d.resourceClient.Update(ctx, &stub.UpdateResourceRequest{
		Token:          d.token,
		Resources:      []*model.Resource{resource},
		DoMigration:    migration,
		ForceMigration: force,
	})

	return err
}

func (d *client) ListenRecords(ctx context.Context, namespace string, resource string, consumer func(records []*model.Record)) error {
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

func (d *client) DeleteResource(ctx context.Context, id string, doMigration bool, forceMigration bool) error {
	_, err := d.resourceClient.Delete(ctx, &stub.DeleteResourceRequest{
		Token:          d.token,
		Ids:            []string{id},
		DoMigration:    doMigration,
		ForceMigration: forceMigration,
	})

	return err
}

func (d *client) GetResourceByName(ctx context.Context, namespace string, getType string) (*model.Resource, error) {
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

func (d *client) ReadRecordStream(ctx context.Context, params service.RecordListParams, recordsChan chan *model.Record) error {
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

func (d *client) CreateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
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

func (d *client) UpdateRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
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

func (d *client) ApplyRecord(ctx context.Context, namespace string, resource string, record *model.Record) (*model.Record, error) {
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

func (d *client) GetRecord(ctx context.Context, namespace string, resource string, id string) (*model.Record, error) {
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

func (d *client) ListRecords(ctx context.Context, params service.RecordListParams) ([]*model.Record, uint32, error) {
	req := params.ToRequest()

	req.Token = d.token

	resp, err := d.recordClient.List(ctx, req)

	if err != nil {
		return nil, 0, err
	}

	return resp.Content, resp.Total, nil
}

func (d *client) ListResources(ctx context.Context) ([]*model.Resource, error) {
	resp, err := d.resourceClient.List(ctx, &stub.ListResourceRequest{
		Token: d.token,
	})

	if err != nil {
		return nil, err
	}

	return resp.Resources, nil
}

func (d *client) UpdateTokenFromContext(ctx context.Context) {
	d.params.Token = ctx.Value("token").(string)
	d.token = ctx.Value("token").(string)
}

func (d *client) AuthenticateWithUsernameAndPassword(username string, password string) error {
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

func (d *client) AuthenticateWithToken(token string) {
	d.params.Token = token
	d.token = token
}

func (d *client) DeleteRecord(ctx context.Context, namespace string, name string, record *model.Record) error {
	_, err := d.recordClient.Delete(ctx, &stub.DeleteRecordRequest{
		Token:     d.token,
		Namespace: namespace,
		Resource:  name,
		Ids:       []string{record.Properties["id"].GetStringValue()},
	})

	return err
}

func (d *client) PollEvents(ctx context.Context, channelKey string) (<-chan *model.Event, error) {
	log.Infof("Polling events for channel: %v", channelKey)
	var eventChan = make(chan *model.Event, 1000)

	go func() {
		for ctx.Err() == nil {
			resp, err := d.eventChannelClient.Poll(ctx, &stub.EventPollRequest{
				Token:      d.token,
				ChannelKey: channelKey,
			})

			if err != nil {
				log.Warn("Error while polling events: ", err)
				continue
			}

			defer func() {
				_ = resp.CloseSend()
			}()
			for {
				event, err := resp.Recv()

				if err != nil {
					break
				}

				if event.Id == "heartbeat-message" {
					continue
				}

				log.Debug("Received event: ", event.Id)

				select {
				case <-ctx.Done():
					break
				default:
				}

				eventChan <- event
			}
		}

		log.Infof("Polling events for channel: %v - done", channelKey)
	}()

	return eventChan, nil
}

func (d *client) WriteEvent(ctx context.Context, key string, event *model.Event) error {
	_, err := d.eventChannelClient.Write(ctx, &stub.EventWriteRequest{
		Token: d.token,
		Event: event,
	})

	return err
}

func NewClientWithParams(params Params) (Client, error) {
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

	return &client{
		conn:                 conn,
		params:               params,
		recordClient:         stub.NewRecordClient(conn),
		authenticationClient: stub.NewAuthenticationClient(conn),
		resourceClient:       stub.NewResourceClient(conn),
		dataSourceClient:     stub.NewDataSourceClient(conn),
		watchClient:          stub.NewWatchClient(conn),
		eventChannelClient:   stub.NewEventChannelClient(conn),
		token:                params.Token,
	}, nil
}

func NewClientWithServerName(serverName string) (Client, error) {
	configServer := LocateConfigServer(serverName)

	return NewClientWithConfigServer(configServer)
}

func NewClientWithConfigServer(configServer ServerConfig) (Client, error) {
	var params = Params{
		Addr:     configServer.Host + ":" + strconv.Itoa(int(configServer.Port)),
		Insecure: configServer.Insecure,
	}

	cl, err := NewClientWithParams(params)

	if err != nil {
		return nil, err
	}

	if configServer.Authentication.Token != "" {
		cl.AuthenticateWithToken(configServer.Authentication.Token)
	} else {
		err = cl.AuthenticateWithUsernameAndPassword(configServer.Authentication.Username, configServer.Authentication.Password)

		if err != nil {
			return nil, err
		}
	}

	return cl, nil
}
