package nodejs

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/generator/nodejs/statik"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"html/template"
	"io"
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
	var b bytes.Buffer
	br := io.Writer(&b)

	err := tmpl.ExecuteTemplate(br, "resource", map[string]interface{}{
		"params": params,
	})

	if err != nil {
		return err
	}

	if err := os.MkdirAll(params.Path, 0777); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(params.Path+"/schema.ts", b.Bytes(), 0777)

	if err != nil {
		return err
	}

	return nil
}

func PropNodejsType(prop *model.ResourceProperty) string {
	if prop.Type == model.ResourceProperty_REFERENCE {
		return strcase.ToCamel(prop.Reference.ReferencedResource)
	}
	return util.ResourcePropertyTypeToJsonSchemaType(prop.Type).Type
}

func IsNullable(prop *model.ResourceProperty) bool {
	return !prop.Required || prop.Type == model.ResourceProperty_REFERENCE
}

var tmpl *template.Template

func init() {
	statikFS, err := fs.NewWithNamespace(statik.GeneratorNodejs)

	if err != nil {
		panic(err)
	}

	entityExistsFile, err := statikFS.Open("/resource.tmpl")

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(entityExistsFile)

	if err != nil {
		panic(err)
	}

	tmplData := string(data)

	tmpl = template.Must(template.New("resource").
		Funcs(sprig.FuncMap()).
		Funcs(map[string]any{
			"ToLowerCamel":   strcase.ToLowerCamel,
			"ToCamel":        strcase.ToCamel,
			"Lower":          strings.ToLower,
			"PropNodejsType": PropNodejsType,
			"IsNullable":     IsNullable,
			"IsPrimitive":    types.IsPrimitive,
		}).
		Parse(tmplData))
}
