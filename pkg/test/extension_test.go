package test

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/test/setup"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
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
	res := util.FromAny[*model.Resource](request.Request["resource"], &model.Resource{})
	req := util.FromAny[*stub.ListRecordRequest](request.Request["request"], &stub.ListRecordRequest{})

	log.Println(res, req)

	result := &stub.ListRecordResponse{
		Content: simpleVirtualResourceRecords,
		Total:   2,
	}

	return &ext.FunctionCallResponse{
		Response: map[string]*anypb.Any{
			"response": util.ToAny(result),
		},
	}, nil
}

func extensionHandler(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	var reqBody = &model.MapAnyWrap{}

	err = protojson.Unmarshal(bodyBytes, reqBody)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	res := util.FromAny[*model.Resource](reqBody.Content["resource"], &model.Resource{})
	req := util.FromAny[*stub.ListRecordRequest](reqBody.Content["request"], &stub.ListRecordRequest{})

	log.Println(res, req)

	result := &stub.ListRecordResponse{
		Content: simpleVirtualResourceRecords,
		Total:   2,
	}

	respBody, err := protojson.Marshal(&model.MapAnyWrap{
		Content: map[string]*anypb.Any{
			"response": util.ToAny(result),
		},
	})

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
	var te = &model.Extension{
		Name:      "test-extension",
		Namespace: setup.SimpleVirtualResource1.Namespace,
		Resource:  setup.SimpleVirtualResource1.Name,
		Before:    nil,
		Instead: &model.Extension_Instead{
			List: &model.ExternalCall{
				Kind: &model.ExternalCall_FunctionCall{
					FunctionCall: &model.FunctionCall{
						Host:         extensionGrpcHost,
						FunctionName: "testFunc",
					},
				},
			},
		},
		After:     nil,
		AuditData: nil,
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

func TestListResourceWithHttpExtension(t *testing.T) {
	var te = &model.Extension{
		Name:      "test-extension",
		Namespace: setup.SimpleVirtualResource1.Namespace,
		Resource:  setup.SimpleVirtualResource1.Name,
		Before:    nil,
		Instead: &model.Extension_Instead{
			List: &model.ExternalCall{
				Kind: &model.ExternalCall_HttpCall{
					HttpCall: &model.HttpCall{
						Uri:    "http://" + extensionRestHost + "/path-1",
						Method: "POST",
					},
				},
			},
		},
		After:     nil,
		AuditData: nil,
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
