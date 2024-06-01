package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resource_model/extramappings"
	"github.com/apibrew/apibrew/pkg/service"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
	"strings"
)

type externalService struct {
	functionClientMap   map[string]ext.FunctionClient
	eventChannelService service.EventChannelService
}

func (e *externalService) Call(ctx context.Context, call resource_model.ExternalCall, event *core.Event) (*core.Event, error) {
	userDetails := jwt_model.GetUserDetailsFromContext(ctx)

	if event.Annotations == nil {
		event.Annotations = make(map[string]string)
	}

	if userDetails != nil {
		event.Annotations["user"] = userDetails.Username
		event.Annotations["userId"] = userDetails.UserId
		event.Annotations["groups"] = strings.Join(userDetails.Roles, ",")
	}

	if call.GetFunctionCall() != nil {
		return e.CallFunction(ctx, call.GetFunctionCall(), event)
	} else if call.GetHttpCall() != nil {
		return e.CallHttp(ctx, call.GetHttpCall(), event)
	} else if call.GetChannelCall() != nil {
		return e.CallChannel(ctx, call.GetChannelCall(), event)
	} else {
		return nil, errors.LogicalError.WithMessage("Both function call and http call is empty")
	}
}

func (e *externalService) CallFunction(ctx context.Context, call *resource_model.FunctionCall, event *core.Event) (*core.Event, error) {
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

	userDetails := jwt_model.GetUserDetailsFromContext(ctx)

	mdMap := map[string]string{}

	if userDetails != nil {
		mdMap["user"] = userDetails.Username
		mdMap["userId"] = userDetails.UserId
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(mdMap))

	result, err := functionService.FunctionCall(ctx, &ext.FunctionCallRequest{
		Name:  call.FunctionName,
		Event: event.ToProtoEvent(),
	})

	if err != nil {
		log.Warn(err.Error())

		if sterr, ok := status.FromError(err); ok {
			return nil, errors.ExternalBackendError.WithMessage(sterr.Message())
		} else {
			return nil, errors.ExternalBackendError.WithMessage(err.Error())
		}
	}

	log.Print(result)

	return core.FromProtoEvent(result.Event), nil
}

func (e *externalService) CallHttp(ctx context.Context, call *resource_model.HttpCall, event *core.Event) (*core.Event, error) {
	body, err := json.Marshal(extramappings.EventFromProto(event))

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

	if resp.StatusCode != 200 {
		var result = new(resource_model.Error)

		err = json.Unmarshal(responseData, result)

		if err != nil {
			log.Print(err)
			return e.reportHttpError(responseData)
		}

		return nil, errors.FromProtoError(extramappings.ErrorToProto(*result))
	}

	if len(responseData) == 0 {
		return nil, nil
	}

	var result = new(resource_model.Event)

	err = json.Unmarshal(responseData, result)

	if err != nil {
		log.Print(err)
		return e.reportHttpError(responseData)
	}

	return extramappings.EventToProto(result), nil
}

func (e *externalService) reportHttpError(responseData []byte) (*core.Event, error) {
	var responseError = &model.Error{}

	err := protojson.Unmarshal(responseData, responseError)

	if err != nil {
		return nil, errors.ExternalBackendCommunicationError.WithMessage(fmt.Sprintf("Error: %s; content: %s", err.Error(), string(responseData)))
	}

	return nil, errors.RecordValidationError.WithDetails(responseError.Message)
}

func (e *externalService) CallChannel(ctx context.Context, call *resource_model.ChannelCall, event *core.Event) (*core.Event, error) {
	return e.eventChannelService.Exec(ctx, call.ChannelKey, event)
}

func NewExternalService(eventChannelService service.EventChannelService) service.ExternalService {
	return &externalService{
		functionClientMap:   make(map[string]ext.FunctionClient),
		eventChannelService: eventChannelService,
	}
}
