package generator

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"go/format"
	"reflect"
	"strings"
)

type GenerateResourceCodeParams struct {
	Package string
}

func GenerateResourceCode(resource *model.Resource, params GenerateResourceCodeParams) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("package %s\n", params.Package))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("import \"time\" \n"))
	sb.WriteString(fmt.Sprintf("import \"github.com/tislib/data-handler/pkg/model\" \n"))
	sb.WriteString(fmt.Sprintf("import \"github.com/google/uuid\" \n"))
	sb.WriteString(fmt.Sprintf("import \"github.com/tislib/data-handler/pkg/types\" \n"))
	sb.WriteString(fmt.Sprintf("import \"google.golang.org/protobuf/types/known/structpb\" \n"))
	sb.WriteRune('\n')

	writeResourceStruct(&sb, resource, params)
	sb.WriteRune('\n')
	writeResourceStructGetIdFunc(&sb, resource, params)
	sb.WriteRune('\n')
	writeResourceStructToRecordFunc(&sb, resource, params)
	sb.WriteRune('\n')
	writeResourceStructFromRecordFunc(&sb, resource, params)
	sb.WriteRune('\n')

	formatted, err := format.Source([]byte(sb.String()))
	if err != nil {
		log.Print(sb.String())
		panic(err)
	}

	return string(formatted)
}

func writeResourceStructGetIdFunc(sb *strings.Builder, resource *model.Resource, params GenerateResourceCodeParams) {
	sb.WriteString(fmt.Sprintf("func (s*%s) GetId() string {\n", dashCaseToCamelCase(resource.Name)))
	sb.WriteString("return s.Id\n")
	sb.WriteString("}\n")
}

func writeResourceStructToRecordFunc(sb *strings.Builder, resource *model.Resource, params GenerateResourceCodeParams) {
	sb.WriteString(fmt.Sprintf("func (s*%s) ToRecord() *model.Record {\n", dashCaseToCamelCase(resource.Name)))
	sb.WriteString(" var rec = &model.Record{} \n")
	sb.WriteString(" rec.Id = s.Id \n")
	sb.WriteString(" rec.Properties = make(map[string]*structpb.Value) \n")
	sb.WriteRune('\n')

	for i, prop := range resource.Properties {
		if !prop.Required {
			sb.WriteString(fmt.Sprintf("if s.%s != nil { \n", dashCaseToCamelCase(prop.Name)))
		}
		sb.WriteString(fmt.Sprintf("val%d, _ := types.ByResourcePropertyType(model.ResourcePropertyType_%s).Pack(s.%s) \n", i, prop.Type.String(), dashCaseToCamelCase(prop.Name)))
		sb.WriteString(fmt.Sprintf("rec.Properties[\"%s\"] = val%d\n", prop.Name, i))
		if !prop.Required {
			sb.WriteString("}\n")
		}
		sb.WriteString("\n")
	}
	sb.WriteString("return rec\n")
	sb.WriteString("}\n")
}

func writeResourceStructFromRecordFunc(sb *strings.Builder, resource *model.Resource, params GenerateResourceCodeParams) {
	sb.WriteString(fmt.Sprintf("func (s*%s) FromRecord(record *model.Record) {\n", dashCaseToCamelCase(resource.Name)))
	sb.WriteString(" s.Id = record.Id \n")
	sb.WriteRune('\n')

	for i, prop := range resource.Properties {
		sb.WriteString(fmt.Sprintf("if record.Properties[\"%s\"] != nil { \n", prop.Name))
		sb.WriteString(fmt.Sprintf("val%d, _ := types.ByResourcePropertyType(model.ResourcePropertyType_%s).UnPack(record.Properties[\"%s\"]) \n", i, prop.Type.String(), prop.Name))
		if prop.Required {
			sb.WriteString(fmt.Sprintf("s.%s = val%d.(%s)\n", dashCaseToCamelCase(prop.Name), i, getGoType(prop.Type)))
		} else {

			sb.WriteString(fmt.Sprintf("s.%s = new(%s)\n", dashCaseToCamelCase(prop.Name), getGoType(prop.Type)))
			sb.WriteString(fmt.Sprintf("*s.%s = val%d.(%s)\n", dashCaseToCamelCase(prop.Name), i, getGoType(prop.Type)))
		}
		sb.WriteString("}\n\n")
	}

	sb.WriteString("}\n")
}

func writeResourceStruct(sb *strings.Builder, resource *model.Resource, params GenerateResourceCodeParams) {
	sb.WriteString(fmt.Sprintf("type %s struct {\n", dashCaseToCamelCase(resource.Name)))

	if !annotations.IsEnabled(resource, annotations.DoPrimaryKeyLookup) {
		sb.WriteString(fmt.Sprintf("    Id string\n"))
	}

	for _, prop := range resource.Properties {
		typeDef := getGoType(prop.Type)
		if !prop.Required {
			typeDef = "*" + typeDef
		}
		sb.WriteString(fmt.Sprintf("    %s %s\n", dashCaseToCamelCase(prop.Name), typeDef))
	}

	if !annotations.IsEnabled(resource, annotations.DisableAudit) {
		sb.WriteString(fmt.Sprintf("    Version uint64\n"))
		sb.WriteString(fmt.Sprintf("    CreatedBy string\n"))
		sb.WriteString(fmt.Sprintf("    UpdatedBy *string\n"))
		sb.WriteString(fmt.Sprintf("    CreatedOn time.Time\n"))
		sb.WriteString(fmt.Sprintf("    UpdatedOn *time.Time\n"))
	}

	sb.WriteString("}\n")
}

func dashCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	//snake_case to camelCase

	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '-' {
					isToUpper = true
				} else if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return

}

func getGoType(propertyType model.ResourcePropertyType) string {
	typ := types.ByResourcePropertyType(propertyType)

	return reflect.TypeOf(typ.Default()).String()
}
