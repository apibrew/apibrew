package test

import (
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestScriptConcurrency(t *testing.T) {
	t.Log("Testing concurrency")

	var wg sync.WaitGroup
	var count = 1000
	var mu sync.Mutex

	// This test is to test the concurrency of the application
	var result []unstructured.Unstructured

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var res = runScript(t, `let a = 5 + 6; sleep(100); a`)

			mu.Lock()
			result = append(result, res)
			mu.Unlock()
		}()
	}

	if t.Failed() {
		return
	}

	wg.Wait()

	for _, r := range result {
		assert.NotNil(t, r["output"])

		if t.Failed() {
			return
		}

		output := r["output"]

		assert.Equal(t, float64(11), output)
	}

	assert.Equal(t, count, len(result))

	t.Log("Tested concurrency")
}

func TestCodeConcurrency(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	defer func() {
		log.SetLevel(log.DebugLevel)
	}()

	t.Log("Testing concurrency")

	var wg sync.WaitGroup
	var count = 1000

	// This test is to test the concurrency of the application

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

			 sleep(100)

			return req
		})
`,
	})

	var result []unstructured.Unstructured

	if err != nil {
		t.Error(err)
		return
	}

	var mux sync.Mutex

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := api.Create(util.SystemContext, unstructured.Unstructured{
				"type": "TestAction1",
			})

			if err != nil {
				t.Error(err)
				return
			}

			mux.Lock()
			result = append(result, res)
			mux.Unlock()
		}()
	}

	wg.Wait()

	if t.Failed() {
		return
	}

	for _, r := range result {
		assert.NotNil(t, r["output"])

		if t.Failed() {
			return
		}

		output := r["output"]

		assert.Equal(t, float64(11), output)
	}

	assert.Len(t, result, count)

	t.Log("Tested concurrency")
}
