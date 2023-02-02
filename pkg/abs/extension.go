package abs

import (
	"context"
	"github.com/tislib/data-handler/pkg/ext"
	"github.com/tislib/data-handler/pkg/model"
)

type Extension interface {
	Init(container Container)
	GetExtensionConfig() *model.ExtensionConfig

	BeforeList(ctx context.Context, in *ext.BeforeListRecordRequest) (*ext.BeforeListRecordResponse, error)
	List(ctx context.Context, in *ext.ListRecordRequest) (*ext.ListRecordResponse, error)
	AfterList(ctx context.Context, in *ext.AfterListRecordRequest) (*ext.AfterListRecordResponse, error)
	BeforeCreate(ctx context.Context, in *ext.BeforeCreateRecordRequest) (*ext.BeforeCreateRecordResponse, error)
	Create(ctx context.Context, in *ext.CreateRecordRequest) (*ext.CreateRecordResponse, error)
	AfterCreate(ctx context.Context, in *ext.AfterCreateRecordRequest) (*ext.AfterCreateRecordResponse, error)
	BeforeUpdate(ctx context.Context, in *ext.BeforeUpdateRecordRequest) (*ext.BeforeUpdateRecordResponse, error)
	Update(ctx context.Context, in *ext.UpdateRecordRequest) (*ext.UpdateRecordResponse, error)
	AfterUpdate(ctx context.Context, in *ext.AfterUpdateRecordRequest) (*ext.AfterUpdateRecordResponse, error)
	BeforeDelete(ctx context.Context, in *ext.BeforeDeleteRecordRequest) (*ext.BeforeDeleteRecordResponse, error)
	Delete(ctx context.Context, in *ext.DeleteRecordRequest) (*ext.DeleteRecordResponse, error)
	AfterDelete(ctx context.Context, in *ext.AfterDeleteRecordRequest) (*ext.AfterDeleteRecordResponse, error)
	BeforeGet(ctx context.Context, in *ext.BeforeGetRecordRequest) (*ext.BeforeGetRecordResponse, error)
	Get(ctx context.Context, in *ext.GetRecordRequest) (*ext.GetRecordResponse, error)
	AfterGet(ctx context.Context, in *ext.AfterGetRecordRequest) (*ext.AfterGetRecordResponse, error)
}
