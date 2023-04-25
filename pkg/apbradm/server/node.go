package server

import (
	"github.com/tislib/apibrew/pkg/service"
	"github.com/tislib/apibrew/pkg/stub"
)

type node struct {
	stub.NodeServer
	container *service.App
}
