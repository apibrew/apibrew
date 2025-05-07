package test

import (
	nano "github.com/apibrew/apibrew/modules/nano/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicNativeFunction(t *testing.T) {
	getNativeRegistry().RegisterNative("add", func(a, b float64) float64 {
		return a + b
	})

	result := runScript(t, `add(9, 6)`)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(15), output)
}

func getNativeRegistry() nano.NativeRegistry {
	for _, module := range container.GetModules() {
		if module, ok := module.(nano.NativeRegistry); ok {
			return module
		}
	}

	panic("NativeRegistry not found")
}
