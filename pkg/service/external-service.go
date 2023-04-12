package service

import (
	"bytes"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"net/http"
)

type externalService struct {
	functionClientMap map[string]ext.FunctionClient
}

func (e *externalService) Call(ctx context.Context, call *model.ExternalCall, in map[string]proto.Message, out map[string]proto.Message) errors.ServiceError {
	if call.GetFunctionCall() != nil {
		return e.CallFunction(ctx, call.GetFunctionCall(), in, out)
	} else if call.GetHttpCall() != nil {
		return e.CallHttp(ctx, call.GetHttpCall(), in, out)
	} else {
		return errors.LogicalError.WithMessage("Both function call and http call is empty")
	}
}

func (e *externalService) CallFunction(ctx context.Context, call *model.FunctionCall, in map[string]proto.Message, out map[string]proto.Message) errors.ServiceError {
	if e.functionClientMap[call.Host+"/"+call.FunctionName] == nil {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial(call.Host, opts...)

		if err != nil {
			log.Error(err)
			return errors.ExternalBackendCommunicationError.WithDetails(err.Error())
		}

		e.functionClientMap[call.Host+"/"+call.FunctionName] = ext.NewFunctionClient(conn)
	}

	functionService := e.functionClientMap[call.Host+"/"+call.FunctionName]

	var request = make(map[string]*anypb.Any)

	for key, item := range in {
		var anyItem = new(anypb.Any)
		err := anyItem.MarshalFrom(item)

		if err != nil {
			log.Error(err)
			return errors.InternalError.WithDetails(err.Error())
		}

		request[key] = anyItem
	}

	result, err := functionService.FunctionCall(ctx, &ext.FunctionCallRequest{
		Name:    call.FunctionName,
		Request: request,
	})

	if err != nil {
		log.Warn(err.Error())

		if sterr, ok := status.FromError(err); ok {
			return errors.ExternalBackendError.WithDetails(sterr.Message())
		} else {
			return errors.ExternalBackendError.WithDetails(err.Error())
		}
	}

	for key, item := range result.Response {
		msg := out[key]

		if msg == nil {
			continue
		}

		err = item.UnmarshalTo(msg)

		if err != nil {
			log.Error(err)
			return errors.InternalError.WithDetails(err.Error())
		}
	}

	return nil
}

func (e *externalService) CallHttp(ctx context.Context, call *model.HttpCall, in map[string]proto.Message, out map[string]proto.Message) errors.ServiceError {
	var request = make(map[string]*anypb.Any)

	for key, item := range in {
		var anyItem = new(anypb.Any)
		err := anyItem.MarshalFrom(item)

		if err != nil {
			log.Error(err)
			return errors.InternalError.WithDetails(err.Error())
		}

		request[key] = anyItem
	}

	requestWrap := &model.MapAnyWrap{Content: request}

	body, err := protojson.Marshal(requestWrap)

	if err != nil {
		log.Error(err)
		return errors.InternalError.WithDetails(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, call.Method, call.Uri, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error(err)
		return errors.InternalError.WithDetails(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error(err)
		return errors.ExternalBackendCommunicationError.WithDetails(err.Error())
	}

	responseWrap := &model.MapAnyWrap{Content: request}
	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return errors.ExternalBackendCommunicationError.WithDetails(err.Error())
	}

	err = protojson.Unmarshal(responseData, responseWrap)

	if err != nil {
		log.Error(err)
		return errors.ExternalBackendCommunicationError.WithDetails(err.Error())
	}

	for key, item := range responseWrap.Content {
		msg := out[key]

		if msg == nil {
			continue
		}

		err = item.UnmarshalTo(msg)

		if err != nil {
			log.Error(err)
			return errors.InternalError.WithDetails(err.Error())
		}
	}

	return nil
}

func NewExternalService() abs.ExternalService {
	return &externalService{
		functionClientMap: make(map[string]ext.FunctionClient),
	}
}
