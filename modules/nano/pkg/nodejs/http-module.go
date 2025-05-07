package nodejs

import (
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

func HTTPSModule(runtime *goja.Runtime, object *goja.Object) {

	object.Set("get", func(call goja.FunctionCall) goja.Value {
		log.Print("GET request")
		return nil
	})

	object.Set("post", func(call goja.FunctionCall) goja.Value {
		log.Print("POST request")
		return nil
	})

	object.Set("put", func(call goja.FunctionCall) goja.Value {
		log.Print("PUT request")
		return nil
	})

	object.Set("delete", func(call goja.FunctionCall) goja.Value {
		log.Print("DELETE request")
		return nil
	})
	log.Print("HTTPS module loaded")
}
