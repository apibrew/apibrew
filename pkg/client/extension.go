package client

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
)

type ExternalFunctionData map[string]*anypb.Any

func (d ExternalFunctionData) GetAction() string {
	return d.GetString("action")
}

func (d ExternalFunctionData) GetString(key string) string {
	if d[key] == nil {
		return ""
	}

	strWrapper := wrapperspb.String("")

	err := d[key].UnmarshalTo(strWrapper)

	if err != nil {
		log.Error(err)
		return ""
	}

	return strWrapper.Value
}

type ExternalFunction func(ctx context.Context, req ExternalFunctionData) (ExternalFunctionData, error)

func CreateRecordTypedFunction[T Entity[T]](instanceProvider func() T, fn func(ctx context.Context, entity T) (T, error)) ExternalFunction {
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
	return func(ctx context.Context, req ExternalFunctionData) (ExternalFunctionData, error) {
		if req.GetAction() != "Create" {
			return nil, errors.New("create action is expected")
		}

		var request = &stub.CreateRecordRequest{}
		err := req["request"].UnmarshalTo(request)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		var responseRecords []*model.Record

		for _, record := range request.Records {
			processedRecord, err := fn(ctx, record)
			if err != nil {
				return nil, err
			}
			responseRecords = append(responseRecords, processedRecord)
		}

		var response = &stub.CreateRecordResponse{
			Records: responseRecords,
		}

		return map[string]*anypb.Any{
			"response": util.ToAny(response),
		}, nil
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

	result, err := e.functions[req.Name](ctx, req.Request)

	if err != nil {
		return nil, err
	}

	return &ext.FunctionCallResponse{Response: result}, nil
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
