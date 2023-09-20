package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/util"
	jwt_model "github.com/apibrew/apibrew/pkg/util/jwt-model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"strings"
	"time"
)

type externalService struct {
	functionClientMap map[string]ext.FunctionClient
}

func (e *externalService) Call(ctx context.Context, call resource_model.ExtensionExternalCall, event *model.Event) (*model.Event, errors.ServiceError) {
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
	} else {
		return nil, errors.LogicalError.WithMessage("Both function call and http call is empty")
	}
}

func (e *externalService) CallFunction(ctx context.Context, call *resource_model.ExtensionFunctionCall, event *model.Event) (*model.Event, errors.ServiceError) {
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

	log.Print(result)

	return result.Event, nil
}

func (e *externalService) CallHttp(ctx context.Context, call *resource_model.ExtensionHttpCall, event *model.Event) (*model.Event, errors.ServiceError) {
	body, err := json.Marshal(e.eventTo(event))

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
		var result = new(resource_model.ExtensionError)

		err = json.Unmarshal(responseData, result)

		if err != nil {
			return e.reportHttpError(err, responseData)
		}

		return nil, e.toServiceError(result)
	}

	if len(responseData) == 0 {
		return nil, nil
	}

	var result = new(resource_model.ExtensionEvent)

	err = json.Unmarshal(responseData, result)

	if err != nil {
		return e.reportHttpError(err, responseData)
	}

	return e.eventFrom(result), nil
}

func (e *externalService) toServiceError(result *resource_model.ExtensionError) errors.ServiceError {
	serviceErr := errors.NewServiceError(model.ErrorCode(model.ErrorCode_value[string(util.DePointer(result.Code, "UNKNOWN_ERROR"))]), util.DePointer(result.Message, ""), codes.Aborted)

	for _, field := range result.Fields {
		serviceErr = serviceErr.WithErrorFields([]*model.ErrorField{
			{
				Property: util.DePointer(field.Property, ""),
				Message:  util.DePointer(field.Message, ""),
			},
		})
	}

	return serviceErr
}

func (e *externalService) reportHttpError(err error, responseData []byte) (*model.Event, errors.ServiceError) {
	var responseError = &model.Error{}

	err = protojson.Unmarshal(responseData, responseError)

	if err != nil {
		return nil, errors.ExternalBackendCommunicationError.WithMessage(fmt.Sprintf("Error: %s; content: %s", err.Error(), string(responseData)))
	}

	return nil, errors.RecordValidationError.WithDetails(responseError.Message)
}

func (e *externalService) eventFrom(result *resource_model.ExtensionEvent) *model.Event {
	var event = new(model.Event)

	event.Id = result.Id
	event.Action = model.Event_Action(model.Event_Action_value[string(result.Action)])
	event.Ids = result.Ids
	if result.Time != nil {
		event.Time = timestamppb.New(*result.Time)
	}
	if result.ActionDescription != nil {
		event.ActionDescription = *result.ActionDescription
	}
	if result.ActionSummary != nil {
		event.ActionSummary = *result.ActionSummary
	}
	event.Annotations = result.Annotations
	if result.Finalizes != nil {
		event.Finalizes = *result.Finalizes
	}
	if result.Sync != nil {
		event.Sync = *result.Sync
	}
	// event.RecordSearchParams = // fixme
	event.Records = util.ArrayMapX(result.Records, func(t *resource_model.Record) *model.Record {
		return &model.Record{
			Properties: resource_model.RecordMapperInstance.ToRecord(t).Properties["properties"].GetStructValue().Fields,
		}
	})

	event.Resource = e.resourceFrom(result.Resource)

	return event
}

func (e *externalService) eventTo(event *model.Event) *resource_model.ExtensionEvent {
	extensionEvent := new(resource_model.ExtensionEvent)
	extensionEvent.Id = event.Id
	extensionEvent.Action = resource_model.ExtensionAction(model.Event_Action_name[int32(event.Action)])
	extensionEvent.Ids = event.Ids
	if event.Time != nil {
		extensionEvent.Time = new(time.Time)
		*extensionEvent.Time = event.Time.AsTime()
	}
	extensionEvent.ActionDescription = &event.ActionDescription
	extensionEvent.ActionSummary = &event.ActionSummary
	extensionEvent.Annotations = event.Annotations
	extensionEvent.Finalizes = &event.Finalizes
	extensionEvent.Sync = &event.Sync
	// extensionEvent.RecordSearchParams = // fixme
	extensionEvent.Records = util.ArrayMapX(event.Records, func(item *model.Record) *resource_model.Record {
		return resource_model.RecordMapperInstance.FromRecord(&model.Record{
			Id: item.Id,
			Properties: map[string]*structpb.Value{
				"id":         item.Properties["id"],
				"properties": structpb.NewStructValue(&structpb.Struct{Fields: item.Properties}),
			},
		})
	})

	extensionEvent.Resource = e.resourceTo(event.Resource)

	return extensionEvent
}

func (e *externalService) resourceTo(resource *model.Resource) *resource_model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := mapping.ResourceToRecord(resource)
	return resource_model.ResourceMapperInstance.FromRecord(resourceRec)
}

func (e *externalService) resourceFrom(resource *resource_model.Resource) *model.Resource {
	if resource == nil {
		return nil
	}
	resourceRec := resource_model.ResourceMapperInstance.ToRecord(resource)
	return mapping.ResourceFromRecord(resourceRec)
}

func NewExternalService() service.ExternalService {
	return &externalService{
		functionClientMap: make(map[string]ext.FunctionClient),
	}
}
