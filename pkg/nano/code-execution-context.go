package nano

type codeExecutionContext struct {
	handlerIds    []string
	closeHandlers []func()
}
