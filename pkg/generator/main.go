//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates

package main

import (
	"github.com/apibrew/apibrew/pkg/generator/golang"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
)

func main() {
	//names := []string{"Kate", "Go", "John", "Brad"}
	//
	//// qtc creates Write* function for each template function.
	//// Such functions accept io.Writer as first parameter:
	//var buf bytes.Buffer
	//templates.WriteGreetings(&buf, names)
	//
	//fmt.Printf("buf=\n%s", buf.Bytes())

	//fmt.Printf("%s\n", templates.Hello("Foo"))

	golang.GenerateGoResourceCode(golang.GenerateResourceCodeParams{
		Package: "modelnew",
		Resources: []*model.Resource{
			resources.DataSourceResource,
		},
		Path:      "/Users/taleh/Projects/apibrew/pkg/modelnew",
		Namespace: "system",
	})
}
