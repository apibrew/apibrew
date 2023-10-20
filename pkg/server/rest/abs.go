package rest

import "github.com/gorilla/mux"

type Api interface {
	ConfigureRouter(router *mux.Router)
}
