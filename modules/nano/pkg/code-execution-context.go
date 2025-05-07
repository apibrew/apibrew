package nano

import (
	"context"
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/apibrew/apibrew/modules/nano/pkg/util"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

var ctxParentContextKey struct{}

type codeExecutionContext struct {
	id                     string
	codeCtx                context.Context
	localCtx               context.Context
	cancel                 context.CancelFunc
	vm                     *goja.Runtime
	identifier             string
	scriptMode             bool
	insideTransaction      bool
	transactionRollbackBag []func() error
	handlerMap             util.Map[string, *abs.HandlerData]
}

func (c *codeExecutionContext) WithContext(ctx context.Context) func() {
	c.localCtx = context.WithValue(ctx, ctxParentContextKey, c.localCtx)

	return func() {
		if c.localCtx != nil {
			if c.localCtx.Value(ctxParentContextKey) == nil {
				c.localCtx = nil
			} else {
				c.localCtx = c.localCtx.Value(ctxParentContextKey).(context.Context)
			}
		}
	}
}

func (c *codeExecutionContext) HandlerMap() util.Map[string, *abs.HandlerData] {
	return c.handlerMap
}

func (c *codeExecutionContext) BeginTransaction() error {
	c.insideTransaction = true
	return nil
}

func (c *codeExecutionContext) CommitTransaction() error {
	c.insideTransaction = false
	c.transactionRollbackBag = nil
	return nil
}

func (c *codeExecutionContext) RollbackTransaction() error {
	for _, f := range c.transactionRollbackBag {
		if err := f(); err != nil {
			log.Error(err)
		}
	}

	c.insideTransaction = false
	c.transactionRollbackBag = nil
	return nil
}

func (c *codeExecutionContext) RegisterRevert(f func() error) {
	c.transactionRollbackBag = append(c.transactionRollbackBag, f)
}

func (c *codeExecutionContext) TransactionalEnabled() bool {
	return c.insideTransaction
}

func (c *codeExecutionContext) CodeContext() context.Context {
	return c.codeCtx
}

func (c *codeExecutionContext) LocalContext() context.Context {
	return c.localCtx
}

func (c *codeExecutionContext) GetCodeIdentifier() string {
	return c.identifier
}

func (c *codeExecutionContext) IsScriptMode() bool {
	return c.scriptMode
}
