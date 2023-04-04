package golang

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/generator/golang/statik"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/types"
	"go/format"
	"html/template"
	"io"
	"reflect"
	"strings"
)

type GenerateResourceCodeParams struct {
	Package   string
	Resources []*model.Resource
}

func GenerateGoResourceCode(resource *model.Resource, params GenerateResourceCodeParams) (string, error) {
	var b bytes.Buffer
	br := io.Writer(&b)

	err := tmpl.ExecuteTemplate(br, "resource.tmpl", map[string]interface{}{
		"resource": resource,
	})

	if err != nil {
		return "", err
	}

	rawCode := b.String()
	formatted, err := format.Source([]byte(rawCode))
	if err != nil {
		log.Print(err)
		//log.Print(b.String())
		return rawCode, nil
	}

	return string(formatted), nil
}

func PropGoType(prop *model.ResourceProperty) string {
	if prop.Required || prop.Type == model.ResourceProperty_REFERENCE {
		return PropPureGoType(prop)
	} else {
		return "*" + PropPureGoType(prop)
	}
}

func PropPureGoType(prop *model.ResourceProperty) string {
	typeDef := GoTypeRaw(prop)

	if prop.Type == model.ResourceProperty_REFERENCE {
		typeDef = "*" + strcase.ToCamel(prop.Reference.ReferencedResource)
	}

	return typeDef
}

func GoTypeRaw(prop *model.ResourceProperty) string {
	typ := types.ByResourcePropertyType(prop.Type)

	typeDef := reflect.TypeOf(typ.Default()).String()
	return typeDef
}

func IsNullable(prop *model.ResourceProperty) bool {
	return !prop.Required || prop.Type == model.ResourceProperty_REFERENCE
}

var tmpl *template.Template

func init() {
	statikFS, err := fs.NewWithNamespace(statik.GeneratorGolang)

	if err != nil {
		panic(err)
	}

	entityExistsFile, err := statikFS.Open("/resource.tmpl")

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(entityExistsFile)
	tmplData := string(data)

	tmpl = template.Must(template.New("resource").
		Funcs(sprig.FuncMap()).
		Funcs(map[string]any{
			"ToLowerCamel":   strcase.ToLowerCamel,
			"ToCamel":        strcase.ToCamel,
			"Lower":          strings.ToLower,
			"PropGoType":     PropGoType,
			"PropPureGoType": PropPureGoType,
			"GoTypeRaw":      GoTypeRaw,
			"IsNullable":     IsNullable,
			"IsPrimitive":    types.IsPrimitive,
		}).
		Parse(tmplData))

	return
}
