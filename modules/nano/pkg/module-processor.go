package nano

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/model"
)

type moduleProcessor struct {
	codeExecutor *codeExecutorService
}

func (f moduleProcessor) MapperTo(record *model.Record) *model2.Module {
	return model2.ModuleMapperInstance.FromRecord(record)
}

func (f moduleProcessor) Register(ctx context.Context, entity *model2.Module) error {
	return f.codeExecutor.registerModule(ctx, entity)
}

func (f moduleProcessor) Update(ctx context.Context, entity *model2.Module) error {
	if err := f.codeExecutor.updateModule(ctx, entity); err != nil {
		return err
	}

	f.codeExecutor.restartCodeContext(ctx)

	return nil
}

func (f moduleProcessor) UnRegister(ctx context.Context, entity *model2.Module) error {
	if err := f.codeExecutor.unRegisterModule(ctx, entity); err != nil {
		return err
	}

	f.codeExecutor.restartCodeContext(ctx)

	return nil
}
