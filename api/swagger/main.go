package swagger

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func ConfigureRouter(r *mux.Router) {
	swaggerFiles.Handler.Prefix = "/docs/"
	loader := openapi3.NewLoader()

	doc, err := loader.LoadFromFile("openapi.yml")

	if err != nil {
		panic(err)
	}

	r.HandleFunc("/docs/api.json", func(writer http.ResponseWriter, request *http.Request) {
		doc, err = loader.LoadFromFile("openapi.yml")

		if err != nil {
			panic(err)
		}

		data, err := doc.MarshalJSON()

		if err != nil {
			writer.WriteHeader(500)
			return
		}

		writer.Write(data)
	})

	r.PathPrefix("/docs").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("/docs/api.json"), //The url pointing to API definition
	))
}
