package docs

import (
	"github.com/gorilla/mux"
	"net/http"
)

func index(r *mux.Router) {
	r.PathPrefix("/docs").Handler(http.StripPrefix("/docs/", http.FileServer(statikFS)))
}
