package rest

import (
	"net/http"
)

type HealthEndpoint struct {
}

func (e HealthEndpoint) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
}
