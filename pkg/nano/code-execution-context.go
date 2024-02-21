package nano

import "context"

type codeExecutionContext struct {
	handlerIds    []string
	closeHandlers []func()
	ctx           context.Context
}

func (c *codeExecutionContext) AddHandlerId(id string) {
	c.handlerIds = append(c.handlerIds, id)
}

func (c *codeExecutionContext) RemoveHandlerId(id string) {
	for i, handlerId := range c.handlerIds {
		if handlerId == id {
			c.handlerIds = append(c.handlerIds[:i], c.handlerIds[i+1:]...)
			return
		}
	}
}

func (c *codeExecutionContext) Context() context.Context {
	return c.ctx
}
