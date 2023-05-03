package backend_event_handler

type Action string

const (
	ActionAdd    Action = "add"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Event struct {
}
