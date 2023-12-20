package nano

import (
	"encoding/base64"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

type codeExecutorService struct {
}

func (s codeExecutorService) registerCode(code *Code) {
	decodedBytes, err := base64.StdEncoding.DecodeString(code.Content)

	if err == nil {
		code.Content = string(decodedBytes)
	}

	log.Debug("Registering code: " + code.Name + " / " + code.Content)

	vm := goja.New()

	s.registerBuiltIns(vm)

	_, err = vm.RunString(code.Content)
	if err != nil {
		panic(err)
	}
}

func (s codeExecutorService) registerBuiltIns(vm *goja.Runtime) {
	err := vm.Set("console", map[string]interface{}{
		"log": func(args ...interface{}) {
			log.Print(args...)
		},
	})

	if err != nil {
		panic(err)
	}
}

func newCodeExecutorService() *codeExecutorService {
	return &codeExecutorService{}
}
