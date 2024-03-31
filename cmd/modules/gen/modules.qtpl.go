// Code generated by qtc from "modules.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line modules.qtpl:1
package gen

//line modules.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line modules.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line modules.qtpl:1
func StreamGenerateModulesContent(qw422016 *qt422016.Writer, modules map[string]string) {
//line modules.qtpl:1
	qw422016.N().S(`// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis
package module
`)
//line modules.qtpl:8
	var i = 0

//line modules.qtpl:8
	qw422016.N().S(`import (
        "github.com/apibrew/apibrew/pkg/service/impl"
`)
//line modules.qtpl:11
	for module := range modules {
//line modules.qtpl:12
		i++

//line modules.qtpl:12
		qw422016.N().S(`        module`)
//line modules.qtpl:13
		qw422016.N().D(i)
//line modules.qtpl:13
		qw422016.N().S(` "`)
//line modules.qtpl:13
		qw422016.N().S(module)
//line modules.qtpl:13
		qw422016.N().S(`"
`)
//line modules.qtpl:14
	}
//line modules.qtpl:14
	qw422016.N().S(`)

var Modules = map[string]string{
`)
//line modules.qtpl:18
	for module, version := range modules {
//line modules.qtpl:18
		qw422016.N().S(`	"`)
//line modules.qtpl:19
		qw422016.N().S(module)
//line modules.qtpl:19
		qw422016.N().S(`": "`)
//line modules.qtpl:19
		qw422016.N().S(version)
//line modules.qtpl:19
		qw422016.N().S(`",
`)
//line modules.qtpl:20
	}
//line modules.qtpl:20
	qw422016.N().S(`}

func RegisterModules(app *impl.App) {
`)
//line modules.qtpl:24
	i = 0

//line modules.qtpl:25
	for _ = range modules {
//line modules.qtpl:26
		i++

//line modules.qtpl:26
		qw422016.N().S(`    app.RegisterModule(module`)
//line modules.qtpl:27
		qw422016.N().D(i)
//line modules.qtpl:27
		qw422016.N().S(`.NewModule)
`)
//line modules.qtpl:28
	}
//line modules.qtpl:28
	qw422016.N().S(`}


`)
//line modules.qtpl:32
}

//line modules.qtpl:32
func WriteGenerateModulesContent(qq422016 qtio422016.Writer, modules map[string]string) {
//line modules.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line modules.qtpl:32
	StreamGenerateModulesContent(qw422016, modules)
//line modules.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line modules.qtpl:32
}

//line modules.qtpl:32
func GenerateModulesContent(modules map[string]string) string {
//line modules.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line modules.qtpl:32
	WriteGenerateModulesContent(qb422016, modules)
//line modules.qtpl:32
	qs422016 := string(qb422016.B)
//line modules.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line modules.qtpl:32
	return qs422016
//line modules.qtpl:32
}
