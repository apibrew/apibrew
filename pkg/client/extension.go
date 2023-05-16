package client

import (
	"context"
	"errors"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type ExternalFunction func(ctx context.Context, req *model.Event) (*model.Event, error)

func CreateRecordTypedFunction[T abs.Entity[T]](instanceProvider func() T, fn func(ctx context.Context, entity T) (T, error)) ExternalFunction {
	return CreateRecordFunction(func(ctx context.Context, record *model.Record) (*model.Record, error) {
		instance := instanceProvider()

		instance.FromRecord(record)
		instance, err := fn(ctx, instance)

		if err != nil {
			return nil, err
		}

		return instance.ToRecord(), nil
	})
}

func CreateRecordFunction(fn func(ctx context.Context, record *model.Record) (*model.Record, error)) ExternalFunction {
	return func(ctx context.Context, event *model.Event) (*model.Event, error) {
		if event.GetAction() != model.Event_CREATE {
			return nil, errors.New("create action is expected")
		}

		var responseRecords []*model.Record

		for _, record := range event.Records {
			processedRecord, err := fn(ctx, record)
			if err != nil {
				return nil, err
			}
			responseRecords = append(responseRecords, processedRecord)
		}

		event.Records = responseRecords

		return event, nil
	}
}

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

func (d *dhClient) NewExtension(host string) Extension {
	return &extension{
		client:     d,
		host:       host,
		remoteHost: host,
		functions:  make(map[string]ExternalFunction),
	}
}
