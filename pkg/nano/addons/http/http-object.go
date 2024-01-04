package http

import (
	"encoding/json"
	"github.com/dop251/goja"
	"io"
	"net/http"
)

type httpObject struct {
	vm *goja.Runtime
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

	return h.makeResponse(resp)
}

func (h *httpObject) makeResponse(resp *http.Response) HttpResponse {
	bodyRaw, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       bodyRaw,
	}
}
