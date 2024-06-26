package rest

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LogApi interface {
	ConfigureRouter(r *mux.Router)
}

type logApi struct {
	handler              *chan *logrus.Entry
	authorizationService service.AuthorizationService
}

func (r *logApi) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (r *logApi) Fire(entry *logrus.Entry) error {
	if r.handler != nil {
		*r.handler <- entry
	}

	return nil
}

func (r *logApi) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/_logs", r.pollLogs).Methods("GET")
}

func (r *logApi) pollLogs(writer http.ResponseWriter, request *http.Request) {
	if err := r.authorizationService.CheckIsExtensionController(request.Context()); err != nil {
		handleError(writer, err)
		return
	}

	levelStr := request.URL.Query().Get("level")

	format := request.URL.Query().Get("format")

	if levelStr == "" {
		levelStr = "debug"
	}

	level, err := logrus.ParseLevel(levelStr)

	if err != nil {
		handleError(writer, err)
		return
	}

	writer.WriteHeader(200)

	handler := make(chan *logrus.Entry, 100)

	r.handler = &handler

	defer func() {
		r.handler = nil
		close(handler)
	}()

	for {
		select {
		case <-request.Context().Done():
			return
		case entry := <-handler:
			if entry.Level <= level {
				if format == "json" {
					data, err := json.Marshal(map[string]interface{}{
						"caller":  entry.Caller,
						"time":    entry.Time,
						"level":   entry.Level.String(),
						"message": entry.Message,
						"fields":  entry.Data,
					})

					if err != nil {
						handleError(writer, err)
						return
					}

					data = append(data, []byte("\n")...)

					_, _ = writer.Write(data)
				} else {
					data, err := entry.Bytes()

					if err != nil {
						handleError(writer, err)
						return
					}

					_, _ = writer.Write(data)
				}

				if f, ok := writer.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	}
}

func (r *logApi) InitHook() {
	logrus.AddHook(r)
}

func NewLogApi(container service.Container) LogApi {
	la := &logApi{
		authorizationService: container.GetAuthorizationService(),
	}

	la.InitHook()

	return la
}
