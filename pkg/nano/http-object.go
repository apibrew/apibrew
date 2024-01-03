package nano

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/dop251/goja"
	"io"
	"net/http"
)

type httpObject struct {
	goja.DynamicObject
	container service.Container
	resource  *model.Resource

	Fire   func(event map[string]interface{})       `json:"fire"`
	Listen func(func(event map[string]interface{})) `json:"listen"`

	vm                  *goja.Runtime
	cec                 *codeExecutionContext
	backendEventHandler backend_event_handler.BackendEventHandler
}

type HttpRequest struct {
	Headers map[string]string `json:"headers"`
}

type Body []byte

func (b *Body) Json() interface{} {
	var body = new(interface{})

	err := json.Unmarshal(*b, body)

	if err != nil {
		panic(err)
	}

	return *body
}

type HttpResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       Body
}

func (h *httpObject) Get(url string, params HttpRequest) HttpResponse {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	return h.makeResponse(resp, err)
}

func (h *httpObject) makeResponse(resp *http.Response, err error) HttpResponse {
	bodyRaw, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       bodyRaw,
	}
}

func NewHttpObject(vm *goja.Runtime) goja.Value {
	return vm.ToValue(&httpObject{vm: vm})
}
