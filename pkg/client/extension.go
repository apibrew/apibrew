package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type ExternalFunction func(ctx context.Context, req *model.Event) (*model.Event, error)

type Extension interface {
	Run(ctx context.Context) error
	RegisterFunction(s string, f ExternalFunction)
	GetRemoteHost() string
	RegisterExtension(newExtension *resource_model.Extension)
	getServiceKey() string
	WithServiceKey(serviceKey string) Extension
}

type extension struct {
	serviceKey string
	host       string
	remoteHost string
	client     *dhClient
	ext.FunctionServer
	functions            map[string]ExternalFunction
	registeredExtensions []*resource_model.Extension
}

func (e *extension) getServiceKey() string {
	return e.serviceKey
}

func (e *extension) RegisterExtension(newExtension *resource_model.Extension) {
	e.registeredExtensions = append(e.registeredExtensions, newExtension)
}

func (e *extension) GetRemoteHost() string {
	return e.remoteHost
}

func (e *extension) RegisterFunction(name string, handler ExternalFunction) {
	e.functions[name] = handler
}

// WithServiceKey
func (e *extension) WithServiceKey(serviceKey string) Extension {
	e.serviceKey = serviceKey
	return e
}

func (e *extension) FunctionCall(ctx context.Context, req *ext.FunctionCallRequest) (*ext.FunctionCallResponse, error) {
	if e.functions[req.Name] == nil {
		return nil, status.Error(codes.NotFound, "External function not found: "+req.Name)
	}

	result, err := e.functions[req.Name](ctx, req.Event)

	if err != nil {
		return nil, err
	}

	return &ext.FunctionCallResponse{Event: result}, nil
}

func (e *extension) Run(ctx context.Context) error {
	server := grpc.NewServer()

	ext.RegisterFunctionServer(server, e)

	l, err := net.Listen("tcp", e.host)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()

		server.Stop()
	}()

	if e.serviceKey != "" {
		extensions, _, err := e.client.ListRecords(ctx, service.RecordListParams{
			Namespace: resources.ExtensionResource.Namespace,
			Resource:  resources.ExtensionResource.Name,
		})

		if err != nil {
			return err
		}

		for _, ex := range extensions {
			exr := resource_model.ExtensionMapperInstance.FromRecord(ex)
			if annotations.Get(exr, annotations.ServiceKey) == e.serviceKey {
				// check if it is in registration list
				var found = false

				for _, re := range e.registeredExtensions {
					if re.Id.String() == exr.Id.String() {
						found = true
						break
					}
				}

				if !found {
					log.Warn("Removing orphaned extension: ", exr.Id)
					err = e.client.DeleteRecord(ctx, resources.ExtensionResource.Namespace, resources.ExtensionResource.Name, exr.Id.String())

					if err != nil {
						return err
					}
				}
			}
		}
	}

	return server.Serve(l)
}

func (d *dhClient) NewExtension(host string, remoteHost string) Extension {
	return &extension{
		client:     d,
		host:       host,
		remoteHost: remoteHost,
		functions:  make(map[string]ExternalFunction),
	}
}
