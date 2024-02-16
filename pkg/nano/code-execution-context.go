package nano

type codeExecutionContext struct {
	handlerIds    []string
	closeHandlers []func()
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
