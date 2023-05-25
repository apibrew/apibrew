package typescript

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/apibrew/apibrew/pkg/generator/typescript/statik"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/iancoleman/strcase"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type GenerateResourceCodeParams struct {
	Package   string
	Resources []*model.Resource
	Path      string
	Namespace string
}

func GenerateResourceCode(params GenerateResourceCodeParams) error {
	for _, resource := range params.Resources {
		err := generateResource(params, resource)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateResource(params GenerateResourceCodeParams, resource *model.Resource) error {
	var b bytes.Buffer
	br := io.Writer(&b)

	err := resourceTmpl.ExecuteTemplate(br, "resource", map[string]interface{}{
		"params":   params,
		"resource": resource,
	})

	if err != nil {
		return err
	}

	if err := os.MkdirAll(params.Path, 0777); err != nil {
		log.Fatal(err)
	}

	fileName := params.Path + "/" + util.ToDashCase(resource.Name) + ".ts"

	existingFile, err := os.ReadFile(fileName)

	if err == nil {
		if util.StripSpaces(string(existingFile)) == util.StripSpaces(string(b.Bytes())) {
			return nil
		}
	}

	err = os.WriteFile(fileName, b.Bytes(), 0777)

	if err != nil {
		return err
	}
	return nil
}

func PropNodejsType(resource *model.Resource, prop *model.ResourceProperty) string {
	if prop.Type == model.ResourceProperty_REFERENCE {
		return strcase.ToCamel(prop.Reference.ReferencedResource)
	}

	if prop.Type == model.ResourceProperty_LIST {
		return strings.TrimSpace(PropNodejsType(resource, prop.Item)) + "[]"
	}

	if prop.Type == model.ResourceProperty_STRUCT {
		if prop.TypeRef != nil {
			return strcase.ToCamel(*prop.TypeRef)
		}

		var b bytes.Buffer
		br := io.Writer(&b)

		err := structTmpl.ExecuteTemplate(br, "struct", map[string]interface{}{
			"Properties": prop.Properties,
		})

		if err != nil {
			panic(err)
		}

		return string(b.Bytes())
	}

	return util.ResourcePropertyTypeToJsonSchemaType(prop.Type).Type
}

func IsNullable(prop *model.ResourceProperty) bool {
	return !prop.Required || prop.Type == model.ResourceProperty_REFERENCE
}

func ReferenceProps(prop *model.Resource) []*model.ResourceProperty {
	return util.ArrayFilter(prop.Properties, func(elem *model.ResourceProperty) bool {
		return elem.Type == model.ResourceProperty_REFERENCE
	})
}

var resourceTmpl *template.Template
var structTmpl *template.Template

func init() {
	statikFS, err := fs.NewWithNamespace(statik.GeneratorNodejs)

	if err != nil {
		panic(err)
	}

	resourceTmpl, err = loadTemplate(statikFS, "resource")

	if err != nil {
		panic(err)
	}

	structTmpl, err = loadTemplate(statikFS, "struct")

	if err != nil {
		panic(err)
	}
}

func loadTemplate(statikFS http.FileSystem, templateName string) (*template.Template, error) {
	entityExistsFile, err := statikFS.Open("/" + templateName + ".tmpl")

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(entityExistsFile)

	if err != nil {
		return nil, err
	}

	tmplData := string(data)

	return template.Must(template.New(templateName).
		Funcs(sprig.FuncMap()).
		Funcs(map[string]any{
			"ToLowerCamel":   strcase.ToLowerCamel,
			"ToCamel":        strcase.ToCamel,
			"Lower":          strings.ToLower,
			"PropNodejsType": PropNodejsType,
			"IsNullable":     IsNullable,
			"IsPrimitive":    types.IsPrimitive,
			"ReferenceProps": ReferenceProps,
			"ToDash":         util.ToDashCase,
		}).
		Parse(tmplData)), nil
}
