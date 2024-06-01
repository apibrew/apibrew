package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/core"
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
	"google.golang.org/protobuf/types/known/structpb"
	"net"
)

type remoteExtension struct {
	serviceKey string
	host       string
	remoteHost string
	client     *client
	ext.FunctionServer
	functions            map[string]ExternalFunction
	registeredExtensions []*resource_model.Extension
}

func (e *remoteExtension) PrepareCall(ext *resource_model.Extension) resource_model.ExternalCall {
	return resource_model.ExternalCall{
		FunctionCall: &resource_model.FunctionCall{
			Host:         e.remoteHost,
			FunctionName: ext.Name,
		},
	}
}

func (e *remoteExtension) getServiceKey() string {
	return e.serviceKey
}

func (e *remoteExtension) RegisterExtension(newExtension *resource_model.Extension) {
	e.registeredExtensions = append(e.registeredExtensions, newExtension)
}

func (e *remoteExtension) RegisterFunction(name string, handler ExternalFunction) {
	e.functions[name] = handler
}

// WithServiceKey
func (e *remoteExtension) WithServiceKey(serviceKey string) Extension {
	e.serviceKey = serviceKey
	return e
}

func (e *remoteExtension) FunctionCall(ctx context.Context, req *ext.FunctionCallRequest) (*ext.FunctionCallResponse, error) {
	if e.functions[req.Name] == nil {
		return nil, status.Error(codes.NotFound, "External function not found: "+req.Name)
	}

	result, err := e.functions[req.Name](ctx, core.FromProtoEvent(req.Event))

	if err != nil {
		return nil, err
	}

	return &ext.FunctionCallResponse{Event: result.ToProtoEvent()}, nil
}

func (e *remoteExtension) Run(ctx context.Context) error {
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
					log.Warn("Removing orphaned remoteExtension: ", exr.Id)
					err = e.client.DeleteRecord(ctx, resources.ExtensionResource.Namespace, resources.ExtensionResource.Name, &model.Record{
						Properties: map[string]*structpb.Value{
							"id": structpb.NewStringValue(exr.Id.String()),
						},
					})

					if err != nil {
						return err
					}
				}
			}
		}
	}

	return server.Serve(l)
}

func (d *client) NewRemoteExtension(host string, remoteHost string) Extension {
	return &remoteExtension{
		client:     d,
		host:       host,
		remoteHost: remoteHost,
		functions:  make(map[string]ExternalFunction),
	}
}
