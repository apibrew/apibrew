package nano

type codeExecutionContext struct {
	handlerIds    []string
	closeHandlers []func()
}

func (c *codeExecutionContext) AddHandlerId(id string) {
	c.handlerIds = append(c.handlerIds, id)
}
