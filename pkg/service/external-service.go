package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
)

type externalService struct {
	functionClientMap map[string]ext.FunctionClient
}

func (e *externalService) Call(ctx context.Context, call *model.ExternalCall, event *model.Event) (*model.Event, errors.ServiceError) {
	if call.GetFunctionCall() != nil {
		return e.CallFunction(ctx, call.GetFunctionCall(), event)
	} else if call.GetHttpCall() != nil {
		return e.CallHttp(ctx, call.GetHttpCall(), event)
	} else {
		return nil, errors.LogicalError.WithMessage("Both function call and http call is empty")
	}
}

func (e *externalService) CallFunction(ctx context.Context, call *model.FunctionCall, event *model.Event) (*model.Event, errors.ServiceError) {
	if e.functionClientMap[call.Host+"/"+call.FunctionName] == nil {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial(call.Host, opts...)

		if err != nil {
			log.Error(err)
			return nil, errors.ExternalBackendCommunicationError.WithMessage(err.Error())
		}

		e.functionClientMap[call.Host+"/"+call.FunctionName] = ext.NewFunctionClient(conn)
	}

	functionService := e.functionClientMap[call.Host+"/"+call.FunctionName]

	result, err := functionService.FunctionCall(ctx, &ext.FunctionCallRequest{
		Name:  call.FunctionName,
		Event: event,
	})

	if err != nil {
		log.Warn(err.Error())

		if sterr, ok := status.FromError(err); ok {
			return nil, errors.ExternalBackendError.WithMessage(sterr.Message())
		} else {
			return nil, errors.ExternalBackendError.WithMessage(err.Error())
		}
	}

	return result.Event, nil
}

func (e *externalService) CallHttp(ctx context.Context, call *model.HttpCall, event *model.Event) (*model.Event, errors.ServiceError) {
	body, err := protojson.Marshal(event)

	if err != nil {
		log.Error(err)
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, call.Method, call.Uri, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error(err)
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error(err)
		return nil, errors.ExternalBackendCommunicationError.WithMessage(err.Error())
	}

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return nil, errors.ExternalBackendCommunicationError.WithMessage(err.Error())
	}

	if len(responseData) == 0 {
		return nil, nil
	}

	var result = new(model.Event)

	err = protojson.Unmarshal(responseData, result)

	if err != nil {
		var responseError = &model.Error{}

		err = protojson.Unmarshal(responseData, responseError)

		if err != nil {
			return nil, errors.ExternalBackendCommunicationError.WithMessage(fmt.Sprintf("Error: %s; content: %s", err.Error(), string(responseData)))
		}

		return nil, errors.RecordValidationError.WithDetails(responseError.Message)
	}

	return result, nil
}

func NewExternalService() abs.ExternalService {
	return &externalService{
		functionClientMap: make(map[string]ext.FunctionClient),
	}
}
