package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"testing"
)

func BenchmarkScriptConcurrency(b *testing.B) {
	log.SetLevel(log.InfoLevel)

	defer func() {
		log.SetLevel(log.DebugLevel)
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runScript(b, `let a = 5 + 6; a`)
		}
	})
}

func BenchmarkCodeConcurrency(b *testing.B) {
	log.SetLevel(log.InfoLevel)

	defer func() {
		log.SetLevel(log.DebugLevel)
	}()

	// registering code
	var api = api.NewInterface(container)

	_, err := api.Apply(util.SystemContext, unstructured.Unstructured{
		"type":             "nano/Code",
		"name":             "TestAction1.js",
		"concurrencyLevel": 100,
		"content": `
		const TestAction1 = resource({
			name: "TestAction1",
			virtual: true,
			properties: {
				output: {
					type: "int32",
				}
			}
		})

		TestAction1.beforeCreate(req => {
			req.output = 5 + 6

			return req
		})
`,
	})

	if err != nil {
		b.Error(err)
		return
	}
	b.SetParallelism(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := api.Create(util.SystemContext, unstructured.Unstructured{
				"type": "TestAction1",
			})

			if err != nil {
				b.Error(err)
				return
			}
		}
	})
}
