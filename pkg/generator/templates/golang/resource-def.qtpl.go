// Code generated by qtc from "resource-def.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/golang/resource-def.qtpl:1
package golang

//line templates/golang/resource-def.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/golang/resource-def.qtpl:2
import "strings"

//line templates/golang/resource-def.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/golang/resource-def.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/golang/resource-def.qtpl:4
func StreamGenerateResourceDefCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource) {
//line templates/golang/resource-def.qtpl:5
	pkgParts := strings.Split(pkg, "/")

//line templates/golang/resource-def.qtpl:6
	pkgName := pkgParts[len(pkgParts)-1]

//line templates/golang/resource-def.qtpl:6
	qw422016.N().S(`package `)
//line templates/golang/resource-def.qtpl:7
	qw422016.E().S(pkgName)
//line templates/golang/resource-def.qtpl:7
	qw422016.N().S(`

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	sub_types "github.com/apibrew/apibrew/pkg/resources/sub-types"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var `)
//line templates/golang/resource-def.qtpl:16
	StreamGoName(qw422016, resource.Name)
//line templates/golang/resource-def.qtpl:16
	qw422016.N().S(`Resource = &model.Resource{
	Name:      "`)
//line templates/golang/resource-def.qtpl:17
	qw422016.E().S(resource.Name)
//line templates/golang/resource-def.qtpl:17
	qw422016.N().S(`",
	Namespace: "`)
//line templates/golang/resource-def.qtpl:18
	qw422016.E().S(resource.Namespace)
//line templates/golang/resource-def.qtpl:18
	qw422016.N().S(`",
`)
//line templates/golang/resource-def.qtpl:19
	if resource.SourceConfig != nil {
//line templates/golang/resource-def.qtpl:19
		qw422016.N().S(`	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "`)
//line templates/golang/resource-def.qtpl:21
		qw422016.E().S(resource.SourceConfig.DataSource)
//line templates/golang/resource-def.qtpl:21
		qw422016.N().S(`",
		Entity:     "`)
//line templates/golang/resource-def.qtpl:22
		qw422016.E().S(resource.SourceConfig.Entity)
//line templates/golang/resource-def.qtpl:22
		qw422016.N().S(`",
	},
`)
//line templates/golang/resource-def.qtpl:24
	}
//line templates/golang/resource-def.qtpl:24
	qw422016.N().S(`	Types: []*model.ResourceSubType{
		sub_types.SecurityConstraint,
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "username",
			Mapping:  "username",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name: "password",

			Mapping:  "password",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name:    "roles",
			Mapping: "roles",
			Type:    model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_STRING,
			},
		},
		special.SecurityConstraintsProperty,
		{
			Name:    "details",
			Mapping: "details",
			Type:    model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}


}`)
//line templates/golang/resource-def.qtpl:76
}

//line templates/golang/resource-def.qtpl:76
func WriteGenerateResourceDefCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource) {
//line templates/golang/resource-def.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/resource-def.qtpl:76
	StreamGenerateResourceDefCode(qw422016, pkg, resource)
//line templates/golang/resource-def.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/resource-def.qtpl:76
}

//line templates/golang/resource-def.qtpl:76
func GenerateResourceDefCode(pkg string, resource *model.Resource) string {
//line templates/golang/resource-def.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/resource-def.qtpl:76
	WriteGenerateResourceDefCode(qb422016, pkg, resource)
//line templates/golang/resource-def.qtpl:76
	qs422016 := string(qb422016.B)
//line templates/golang/resource-def.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/resource-def.qtpl:76
	return qs422016
//line templates/golang/resource-def.qtpl:76
}
