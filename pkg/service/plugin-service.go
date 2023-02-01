package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	plugin2 "github.com/tislib/data-handler/pkg/plugin"
	"os"
	"plugin"
)

type PluginService interface {
	Init(data *model.InitData)
}

type pluginService struct {
}

func (p pluginService) Init(data *model.InitData) {
	if data.Config.PluginsFolder == "" {
		return
	}

	files, err := os.ReadDir(data.Config.PluginsFolder)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		p.loadPlugin(file)
	}
}

func (p pluginService) loadPlugin(file os.DirEntry) {
	pl, err := plugin.Open(file.Name())

	if err != nil {
		panic(err)
	}

	symbol, err := pl.Lookup(plugin2.MetaDataKey)

	if err != nil {
		panic(err)
	}

	if meta, ok := symbol.(plugin2.MetaData); ok {
		log.Print(meta.Handlers)
	}
}

func NewPluginService() PluginService {
	return &pluginService{}
}
