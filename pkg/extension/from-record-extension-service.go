package extension

import (
	"context"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
)

type recordExtensionServiceClient struct {
	client ext.RecordExtensionServiceClient
	config *model.ExtensionConfig
}

func (r recordExtensionServiceClient) Init(container abs.Container) {

}

func (r recordExtensionServiceClient) GetExtensionConfig() *model.ExtensionConfig {
	return r.config
}

func (r recordExtensionServiceClient) BeforeList(ctx context.Context, in *ext.BeforeListRecordRequest) (*ext.BeforeListRecordResponse, error) {
	return r.client.BeforeList(ctx, in)
}

func (r recordExtensionServiceClient) List(ctx context.Context, in *ext.ListRecordRequest) (*ext.ListRecordResponse, error) {
	return r.client.List(ctx, in)
}

func (r recordExtensionServiceClient) AfterList(ctx context.Context, in *ext.AfterListRecordRequest) (*ext.AfterListRecordResponse, error) {
	return r.client.AfterList(ctx, in)
}

func (r recordExtensionServiceClient) BeforeCreate(ctx context.Context, in *ext.BeforeCreateRecordRequest) (*ext.BeforeCreateRecordResponse, error) {
	return r.client.BeforeCreate(ctx, in)
}

func (r recordExtensionServiceClient) Create(ctx context.Context, in *ext.CreateRecordRequest) (*ext.CreateRecordResponse, error) {
	return r.client.Create(ctx, in)
}

func (r recordExtensionServiceClient) AfterCreate(ctx context.Context, in *ext.AfterCreateRecordRequest) (*ext.AfterCreateRecordResponse, error) {
	return r.client.AfterCreate(ctx, in)
}

func (r recordExtensionServiceClient) BeforeUpdate(ctx context.Context, in *ext.BeforeUpdateRecordRequest) (*ext.BeforeUpdateRecordResponse, error) {
	return r.client.BeforeUpdate(ctx, in)
}

func (r recordExtensionServiceClient) Update(ctx context.Context, in *ext.UpdateRecordRequest) (*ext.UpdateRecordResponse, error) {
	return r.client.Update(ctx, in)
}

func (r recordExtensionServiceClient) AfterUpdate(ctx context.Context, in *ext.AfterUpdateRecordRequest) (*ext.AfterUpdateRecordResponse, error) {
	return r.client.AfterUpdate(ctx, in)
}

func (r recordExtensionServiceClient) BeforeDelete(ctx context.Context, in *ext.BeforeDeleteRecordRequest) (*ext.BeforeDeleteRecordResponse, error) {
	return r.client.BeforeDelete(ctx, in)
}

func (r recordExtensionServiceClient) Delete(ctx context.Context, in *ext.DeleteRecordRequest) (*ext.DeleteRecordResponse, error) {
	return r.client.Delete(ctx, in)
}

func (r recordExtensionServiceClient) AfterDelete(ctx context.Context, in *ext.AfterDeleteRecordRequest) (*ext.AfterDeleteRecordResponse, error) {
	return r.client.AfterDelete(ctx, in)
}

func (r recordExtensionServiceClient) BeforeGet(ctx context.Context, in *ext.BeforeGetRecordRequest) (*ext.BeforeGetRecordResponse, error) {
	return r.client.BeforeGet(ctx, in)
}

func (r recordExtensionServiceClient) Get(ctx context.Context, in *ext.GetRecordRequest) (*ext.GetRecordResponse, error) {
	return r.client.Get(ctx, in)
}

func (r recordExtensionServiceClient) AfterGet(ctx context.Context, in *ext.AfterGetRecordRequest) (*ext.AfterGetRecordResponse, error) {
	return r.client.AfterGet(ctx, in)
}

func FromRecordExtensionServiceClient(client ext.RecordExtensionServiceClient, config *model.ExtensionConfig) abs.Extension {
	return &recordExtensionServiceClient{client: client, config: config}
}
