package main

import "github.com/tislib/data-handler/pkg/ext"

type recordExtensionService struct {
}

func main() {
	ext.record
	stub.RegisterResourceServiceServer(g.grpcServer, NewResourceServiceServer(g.resourceService))
}
