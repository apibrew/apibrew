package rest

import (
	"github.com/gorilla/mux"
	"net/http/pprof"
)

type pprofApi struct {
}

func (p pprofApi) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func NewPprofApi() Api {
	return &pprofApi{}
}
