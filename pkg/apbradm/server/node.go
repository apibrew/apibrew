package server

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/resources/mapping"
	"github.com/tislib/apibrew/pkg/service"
	"github.com/tislib/apibrew/pkg/stub"
	. "github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
)

type nodeService struct {
	stub.NodeServer
	container *service.App
}

func (n *nodeService) Create(ctx context.Context, req *CreateNodeRequest) (*CreateNodeResponse, error) {
	records, _, err := n.container.GetRecordService().Create(ctx, abs.RecordCreateParams{
		Namespace:      NodeResource.Namespace,
		Resource:       NodeResource.Name,
		Records:        mapping.MapToRecord(req.Nodes, NodeToRecord),
		IgnoreIfExists: false,
	})

	return &stub.CreateNodeResponse{
		Nodes: mapping.MapFromRecord(records, RecordToNode),
	}, util.ToStatusError(err)
}

func (n *nodeService) List(ctx context.Context, req *ListNodeRequest) (*ListNodeResponse, error) {
	records, _, err := n.container.GetRecordService().List(ctx, abs.RecordListParams{
		Namespace: NodeResource.Namespace,
		Resource:  NodeResource.Name,
		Limit:     10000,
	})

	return &stub.ListNodeResponse{
		Content: mapping.MapFromRecord(records, RecordToNode),
	}, util.ToStatusError(err)
}

func (n *nodeService) Update(ctx context.Context, req *UpdateNodeRequest) (*UpdateNodeResponse, error) {
	records, err := n.container.GetRecordService().Update(ctx, abs.RecordUpdateParams{
		Namespace: NodeResource.Namespace,
		Resource:  NodeResource.Name,
		Records:   mapping.MapToRecord(req.Nodes, NodeToRecord),
	})

	return &stub.UpdateNodeResponse{
		Nodes: mapping.MapFromRecord(records, RecordToNode),
	}, util.ToStatusError(err)

}
func (n *nodeService) Delete(ctx context.Context, req *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	err := n.container.GetRecordService().Delete(ctx, abs.RecordDeleteParams{
		Namespace: NodeResource.Namespace,
		Resource:  NodeResource.Name,
		Ids:       req.Ids,
	})

	return &stub.DeleteNodeResponse{}, util.ToStatusError(err)
}

func (n *nodeService) Get(ctx context.Context, req *GetNodeRequest) (*GetNodeResponse, error) {
	record, err := n.container.GetRecordService().Get(ctx, abs.RecordGetParams{
		Namespace: NodeResource.Namespace,
		Resource:  NodeResource.Name,
		Id:        req.Id,
	})

	return &stub.GetNodeResponse{
		Node: RecordToNode(record),
	}, util.ToStatusError(err)
}

// Status will return connection status of data source
func (n *nodeService) NodeStatus(ctx context.Context, req *NodeStatusRequest) (*NodeStatusResponse, error) {
	return &stub.NodeStatusResponse{}, nil
}

func (n *nodeService) InstallNewNode(context.Context, *InstallNewNodeRequest) (*InstallNewNodeResponse, error) {
	return nil, nil
}

func (n *nodeService) UninstallNode(context.Context, *UninstallNodeRequest) (*UninstallNodeResponse, error) {
	return nil, nil
}
