package test

import (
	"context"
	"github.com/apibrew/apibrew/pkg/ext"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/test/setup"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"net"
	"net/http"
	"testing"
)

var simpleVirtualResourceRecords = []*model.Record{
	{
		Id: "5429846c-a309-11ed-a8fc-0242ac120002",
		Properties: map[string]*structpb.Value{
			"name":        structpb.NewStringValue("rec-1"),
			"description": structpb.NewStringValue("rec-1-desc"),
		},
	},
	{
		Id: "54298994-a309-11ed-a8fc-0242ac120002",
		Properties: map[string]*structpb.Value{
			"name":        structpb.NewStringValue("rec-2"),
			"description": structpb.NewStringValue("rec-2-desc"),
		},
	},
}

const extensionGrpcHost = "127.0.0.1:47182"
const extensionRestHost = "127.0.0.1:37182"

type TestFunctionBackend struct {
	ext.FunctionServer
}

func (t TestFunctionBackend) FunctionCall(ctx context.Context, request *ext.FunctionCallRequest) (*ext.FunctionCallResponse, error) {
	log.Println(request.Event)

	event := request.Event
	event.Records = simpleVirtualResourceRecords

	return &ext.FunctionCallResponse{
		Event: event,
	}, nil
}

func extensionHandler(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	var event = &model.Event{}
	err = protojson.Unmarshal(bodyBytes, event)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	log.Println(event)

	event.Records = simpleVirtualResourceRecords

	respBody, err := protojson.Marshal(event)

	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	writer.WriteHeader(200)
	_, _ = writer.Write(respBody)
}

func TestMain(m *testing.M) {
	server := grpc.NewServer()

	ext.RegisterFunctionServer(server, &TestFunctionBackend{})

	go func() {
		l, err := net.Listen("tcp", extensionGrpcHost)
		if err != nil {
			log.Fatal(err)
		}

		err = server.Serve(l)

		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		srv := &http.Server{
			Addr:    extensionRestHost,
			Handler: http.HandlerFunc(extensionHandler),
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	m.Run()
}

func TestListResourceWithFunctionCallExtension(t *testing.T) {
	var id = uuid.New()
	var te = &resource_model.Extension{
		Id:   &id,
		Name: "test-extension",

		Selector: &resource_model.ExtensionEventSelector{
			Namespaces: []string{setup.SimpleVirtualResource1.Namespace},
			Resources:  []string{setup.SimpleVirtualResource1.Name},
		},
		Call: resource_model.ExtensionExternalCall{
			FunctionCall: &resource_model.ExtensionFunctionCall{
				Host:         extensionGrpcHost,
				FunctionName: "testFunc",
			},
		},

		Sync:      true,
		Finalizes: true,
		Responds:  true,

		Version: 1,
		Order:   91,
	}

	container.GetExtensionService().RegisterExtension(te)
	defer container.GetExtensionService().UnRegisterExtension(te)

	resp, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
		Namespace: setup.SimpleVirtualResource1.Namespace,
		Resource:  setup.SimpleVirtualResource1.Name,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if resp.Total != 2 {
		t.Error("resp.Total should be 2")
		return
	}

	if resp.Content[0].Id != simpleVirtualResourceRecords[0].Id {
		t.Error("record[0].id does not match")
		return
	}

	if resp.Content[1].Id != simpleVirtualResourceRecords[1].Id {
		t.Error("record[1].id does not match")
		return
	}

	if resp.Content[0].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["name"].GetStringValue() {
		t.Error("record[0].name does not match")
		return
	}

	if resp.Content[1].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["name"].GetStringValue() {
		t.Error("record[1].name does not match")
		return
	}

	if resp.Content[0].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["description"].GetStringValue() {
		t.Error("record[0].description does not match")
		return
	}

	if resp.Content[1].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["description"].GetStringValue() {
		t.Error("record[1].description does not match")
		return
	}

}

func TestListResourceWithHttpExtension(t *testing.T) {
	var id = uuid.New()
	var te = &resource_model.Extension{
		Id:   &id,
		Name: "test-extension",
		Selector: &resource_model.ExtensionEventSelector{
			Namespaces: []string{setup.SimpleVirtualResource1.Namespace},
			Resources:  []string{setup.SimpleVirtualResource1.Name},
		},
		Call: resource_model.ExtensionExternalCall{
			HttpCall: &resource_model.ExtensionHttpCall{
				Uri:    "http://" + extensionRestHost + "/path-1",
				Method: "POST",
			},
		},

		Sync:      true,
		Finalizes: true,
		Responds:  true,
		Order:     90,
		Version:   1,
	}

	container.GetExtensionService().RegisterExtension(te)
	defer container.GetExtensionService().UnRegisterExtension(te)

	resp, err := recordClient.List(setup.Ctx, &stub.ListRecordRequest{
		Namespace: setup.SimpleVirtualResource1.Namespace,
		Resource:  setup.SimpleVirtualResource1.Name,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if resp.Total != 2 {
		t.Error("resp.Total should be 2")
		return
	}

	if resp.Content[0].Id != simpleVirtualResourceRecords[0].Id {
		t.Error("record[0].id does not match")
		return
	}

	if resp.Content[1].Id != simpleVirtualResourceRecords[1].Id {
		t.Error("record[1].id does not match")
		return
	}

	if resp.Content[0].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["name"].GetStringValue() {
		t.Error("record[0].name does not match")
		return
	}

	if resp.Content[1].Properties["name"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["name"].GetStringValue() {
		t.Error("record[1].name does not match")
		return
	}

	if resp.Content[0].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[0].Properties["description"].GetStringValue() {
		t.Error("record[0].description does not match")
		return
	}

	if resp.Content[1].Properties["description"].GetStringValue() != simpleVirtualResourceRecords[1].Properties["description"].GetStringValue() {
		t.Error("record[1].description does not match")
		return
	}

}
