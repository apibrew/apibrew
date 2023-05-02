package backend_event_handler

type BackendEventHandler interface {
}

type backendEventHandler struct {
}

func NewBackendEventHandler() BackendEventHandler {
	return &backendEventHandler{}
}
