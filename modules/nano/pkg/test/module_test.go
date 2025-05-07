package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicModule(t *testing.T) {
	var api = api.NewInterface(container)

	var module1 = new(model.Module)
	module1.Name = `module1`
	module1.Source = `
		exports.add = function(a, b) {
			return a + b;
		}
`
	module1.Language = model.ModuleLanguage_JAVASCRIPT
	module1.ContentFormat = model.ModuleContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

	if err != nil {
		t.Error(err)
		return
	}

	result := runScript(t, `
		const module1 = require('module1');

		module1.add(5, 6);
`)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestBasicModuleTypescript(t *testing.T) {
	var api = api.NewInterface(container)

	var module1 = new(model.Module)
	module1.Name = `module1`
	module1.Source = `
		export function add(a: number, b: number): number {
			return a + b;
		}
`
	module1.Language = model.ModuleLanguage_TYPESCRIPT
	module1.ContentFormat = model.ModuleContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

	if err != nil {
		t.Error(err)
		return
	}

	result := runScriptWithLanguage(t, `
		import {add} from 'module1';

		add(5, 6);
`, model.ScriptLanguage_TYPESCRIPT)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestComplexModuleTypescript(t *testing.T) {
	var api = api.NewInterface(container)

	var module1 = new(model.Module)
	module1.Name = `module3`
	module1.Source = `
		export class Person {
			add(a: number, b: number): number {
				return a + b;
			}
		}
`
	module1.Language = model.ModuleLanguage_TYPESCRIPT
	module1.ContentFormat = model.ModuleContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

	if err != nil {
		t.Error(err)
		return
	}

	result := runScriptWithLanguage(t, `
		import {Person} from 'module3';
		
		const person = new Person();

		person.add(5, 6);
`, model.ScriptLanguage_TYPESCRIPT)
	if t.Failed() {
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestComplexModuleOnTypescriptWithCode(t *testing.T) {
	var api = api.NewInterface(container)

	var module1 = new(model.Module)
	module1.Name = `module3`
	module1.Source = `
		export class Person {
			add(a: number, b: number): number {
				return a + b;
			}
		}
`
	module1.Language = model.ModuleLanguage_TYPESCRIPT
	module1.ContentFormat = model.ModuleContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

	if err != nil {
		t.Error(err)
		return
	}

	_, err = api.Apply(util.SystemContext, unstructured.Unstructured{
		"type":             "nano/Code",
		"name":             "TestAction1.ts",
		"concurrencyLevel": 1,
		"language":         "TYPESCRIPT",
		"content": `
		import {Person} from './module3'

		import {resource} from '@apibrew/nano'

		const TestAction1 = resource({
			name: "TestComplexModuleOnTypescriptWithCodeResource",
			virtual: true,
			properties: {
				output: {
					type: "int32",
				}
			}
		})

		TestAction1.on(req => {
			const person = new Person()
			req.output = person.add(5, 6)
			return req
		})
`,
	})

	if err != nil {
		t.Error(err)
		return
	}

	result, err := api.Create(util.SystemContext, unstructured.Unstructured{
		"type": "TestComplexModuleOnTypescriptWithCodeResource",
	})

	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(11), output)
}

func TestModuleReload(t *testing.T) {
	var api = api.NewInterface(container)

	for I := 0; I < 2; I++ {
		var module1 = new(model.Module)
		module1.Name = `module4`
		module1.Language = model.ModuleLanguage_TYPESCRIPT
		module1.ContentFormat = model.ModuleContentFormat_TEXT

		module1.Source = `
		export class Person {
			add(a: number, b: number): number {
				return a + b;
			}
		}
	`
		_, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

		if err != nil {
			t.Error(err)
			return
		}

		_, err = api.Apply(util.SystemContext, unstructured.Unstructured{
			"type":             "nano/Code",
			"name":             "TestAction1.ts",
			"concurrencyLevel": 8,
			"language":         "TYPESCRIPT",
			"content": `
		import {Person} from './module4'

		import {resource} from '@apibrew/nano'

		const TestAction1 = resource({
			name: "TestModuleReloadResource",
			virtual: true,
			properties: {
				output: {
					type: "int32",
				}
			}
		})

		TestAction1.on(req => {
			const person = new Person()
			req.output = person.add(5, 6)
			return req
		})
`,
		})

		if err != nil {
			t.Error(err)
			return
		}

		result, err := api.Create(util.SystemContext, unstructured.Unstructured{
			"type": "TestModuleReloadResource",
		})

		if err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, result["output"])

		if t.Failed() {
			return
		}

		output := result["output"]

		assert.Equal(t, float64(11), output)

		if t.Failed() {
			return
		}

		module1.Source = `
		export class Person {
			add(a: number, b: number): number {
				return a - b;
			}
		}
	`
		_, err = api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

		if err != nil {
			t.Error(err)
			return
		}

		result, err = api.Create(util.SystemContext, unstructured.Unstructured{
			"type": "TestModuleReloadResource",
		})

		if err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, result["output"])

		if t.Failed() {
			return
		}

		output = result["output"]

		assert.Equal(t, float64(-1), output)
	}
}

func TestModuleReload2(t *testing.T) {
	var api = api.NewInterface(container)

	for I := 0; I < 2; I++ {
		var module1 = new(model.Module)
		module1.Name = `module4`
		module1.Language = model.ModuleLanguage_TYPESCRIPT
		module1.ContentFormat = model.ModuleContentFormat_TEXT

		module1.Source = `
		export class Person {
			add(a: number, b: number): number {
				return a + b;
			}
		}
	`
		moduleResult, err := api.Apply(util.SystemContext, model.ModuleMapperInstance.ToUnstructured(module1))

		if err != nil {
			t.Error(err)
			return
		}

		_, err = api.Apply(util.SystemContext, unstructured.Unstructured{
			"type":             "nano/Code",
			"name":             "TestAction1.ts",
			"concurrencyLevel": 8,
			"language":         "TYPESCRIPT",
			"content": `
		import {Person} from './module4'

		import {resource} from '@apibrew/nano'

		const TestAction1 = resource({
			name: "TestModuleReloadResource",
			virtual: true,
			properties: {
				output: {
					type: "int32",
				}
			}
		})

		TestAction1.on(req => {
			const person = new Person()
			req.output = person.add(5, 6)
			return req
		})
`,
		})

		if err != nil {
			t.Error(err)
			return
		}

		result, err := api.Create(util.SystemContext, unstructured.Unstructured{
			"type": "TestModuleReloadResource",
		})

		if err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, result["output"])

		if t.Failed() {
			return
		}

		output := result["output"]

		assert.Equal(t, float64(11), output)

		if t.Failed() {
			return
		}

		_, err = api.Apply(util.SystemContext, unstructured.Unstructured{
			"type": "nano/Module",
			"id":   moduleResult["id"],
			"source": `export class Person {
			add(a: number, b: number): number {
				return a - b;
			}
		}`,
		})

		if err != nil {
			t.Error(err)
			return
		}

		result, err = api.Create(util.SystemContext, unstructured.Unstructured{
			"type": "TestModuleReloadResource",
		})

		if err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, result["output"])

		if t.Failed() {
			return
		}

		output = result["output"]

		assert.Equal(t, float64(-1), output)
	}
}
