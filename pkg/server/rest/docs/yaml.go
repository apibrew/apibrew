package docs

import (
	"github.com/apibrew/apibrew/pkg/abs"
	yamlformat "github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func yaml(resourceService abs.ResourceService) func(r *mux.Router) {
	return func(r *mux.Router) {
		r.PathPrefix("/docs/yaml").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// strip prefix
			r.URL.Path = r.URL.Path[len("/docs/yaml"):]

			// extract namespace and resource (in namespace-name format)
			namespace, resourceName, extension := extractNamespaceAndResource(r.URL.Path)

			if extension != "yaml" && extension != "yml" {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Format not found"))
				return
			}

			// locate resource
			resource := resourceService.GetResourceByName(r.Context(), namespace, resourceName)

			if resource == nil {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Resource Not found"))
				return
			}

			if err := yamlformat.NewWriter(w, nil).WriteResource(resource); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal error"))
				return
			}
		})
	}
}

func extractNamespaceAndResource(path string) (string, string, string) {
	parts := strings.Split(path, ".")

	if len(parts) < 2 {
		return "", "", ""
	}

	subParts := strings.Split(parts[0], "-")

	if len(subParts) < 2 {
		return "", "", ""
	}

	return subParts[0], subParts[1], parts[1]
}
