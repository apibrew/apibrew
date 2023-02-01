package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	plugin2 "github.com/tislib/data-handler/pkg/plugin"
	"os"
	"plugin"
	"strings"
)

type PluginService interface {
	Init(data *model.InitData)
}

type pluginService struct {
}

func (p pluginService) Init(data *model.InitData) {
	for _, pluginsPath := range strings.Split(data.Config.PluginsPath, ":") {

		files, err := os.ReadDir(pluginsPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".so") {
				p.loadPlugin(pluginsPath + "/" + file.Name())
			}
		}
	}
}

func (p pluginService) loadPlugin(path string) {
	pl, err := plugin.Open(path)

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
