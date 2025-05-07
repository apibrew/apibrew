package console

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type consoleObject struct {
	Log      func(args ...interface{}) `json:"log"`
	Debug    func(args ...interface{}) `json:"debug"`
	Warn     func(args ...interface{}) `json:"warn"`
	Error    func(args ...interface{}) `json:"error"`
	Trace    func(args ...interface{}) `json:"trace"`
	Info     func(args ...interface{}) `json:"info"`
	codeName string
}

func (c *consoleObject) log(level log.Level) func(args ...interface{}) {
	return func(args ...interface{}) {
		log.WithField("CodeName", c.codeName).Logln(level, util.ArrayMap(args, func(item interface{}) interface{} {
			data, err := json.Marshal(item)

			if err != nil {
				return err.Error()
			}

			return string(data)
		})...)
	}
}

func Register(vm *goja.Runtime, codeName string) error {
	obj := &consoleObject{codeName: codeName}

	obj.Log = obj.log(log.DebugLevel)
	obj.Debug = obj.log(log.DebugLevel)
	obj.Trace = obj.log(log.TraceLevel)
	obj.Info = obj.log(log.InfoLevel)
	obj.Warn = obj.log(log.WarnLevel)
	obj.Error = obj.log(log.ErrorLevel)

	return vm.Set("console", obj)
}
