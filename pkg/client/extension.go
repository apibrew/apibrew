package client

import (
	"context"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
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
}

type extension struct {
	host       string
	remoteHost string
	client     *dhClient
	ext.FunctionServer
	functions map[string]ExternalFunction
}

func (e *extension) GetRemoteHost() string {
	return e.remoteHost
}

func (e *extension) RegisterFunction(name string, handler ExternalFunction) {
	e.functions[name] = handler
}

func (e *extension) FunctionCall(ctx context.Context, req *ext.FunctionCallRequest) (*ext.FunctionCallResponse, error) {
	if e.functions[req.Name] == nil {
		return nil, status.Error(codes.NotFound, "External function not found")
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
