package http

import (
	"bytes"
	"encoding/json"
	"github.com/apibrew/apibrew/modules/nano/pkg/addons/util"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/url"
)

type httpObject struct {
	vm *goja.Runtime
}

type HttpRequest struct {
	Headers map[string]string `json:"headers"`
}

type Body struct {
	vm   *goja.Runtime
	data []byte
}

func (b *Body) Json() interface{} {
	var body = new(interface{})

	err := json.Unmarshal(b.data, body)

	if err != nil {
		util.ThrowError(b.vm, err.Error())
	}

	return *body
}

func (b *Body) Text() interface{} {
	return string(b.data)
}

func (b *Body) UrlEncoded() map[string][]string {
	values, err := url.ParseQuery(string(b.data))

	if err != nil {
		util.ThrowError(b.vm, err.Error())
	}

	return values
}

type HttpResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       Body
}

func (h *httpObject) Get(url string, params HttpRequest) HttpResponse {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	return h.makeResponse(resp)
}

func (h *httpObject) Delete(url string, params HttpRequest) HttpResponse {
	req, err := http.NewRequest("DELETE	", url, nil)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	return h.makeResponse(resp)
}

func (h *httpObject) Method(methodName string, url string, body interface{}, params HttpRequest) HttpResponse {
	var r io.Reader

	if bodyStr, ok := body.(string); ok {
		r = bytes.NewReader([]byte(bodyStr))
	} else {
		bodyBytes, err := json.Marshal(body)

		if err != nil {
			util.ThrowError(h.vm, err.Error())
		}

		r = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(methodName, url, r)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	return h.makeResponse(resp)
}

func (h *httpObject) Post(url string, body interface{}, params HttpRequest) HttpResponse {
	return h.Method("POST", url, body, params)
}

func (h *httpObject) Put(url string, body interface{}, params HttpRequest) HttpResponse {
	return h.Method("PUT", url, body, params)
}

func (h *httpObject) makeResponse(resp *http.Response) HttpResponse {
	bodyRaw, err := io.ReadAll(resp.Body)

	if err != nil {
		util.ThrowError(h.vm, err.Error())
	}

	return HttpResponse{
		StatusCode: resp.StatusCode,
		Body: Body{
			data: bodyRaw,
			vm:   h.vm,
		},
	}
}
