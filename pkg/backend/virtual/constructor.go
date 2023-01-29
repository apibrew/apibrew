package virtual

import (
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/model"
)

type virtualBackend struct {
	options *model.VirtualOptions
}

func NewVirtualBackend(connectionDetails backend.DataSourceConnectionDetails) backend.Backend {
	return &virtualBackend{
		//options: connectionDetails.(*model.VirtualOptions),
		options: &model.VirtualOptions{Mode: model.VirtualOptions_DISCARD},
	}
}
