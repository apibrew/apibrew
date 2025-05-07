package test

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionalCreate(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `
const Book3 = resource({
	name: 'Book3',
	properties: {
		name: {
			type: 'string',
		},
		description: {
			type: 'string',
		}
	}
})

Book3.deleteAll()

try {
	transactional(() => {
		const book = Book3.create({ name: 'Book3'})
	
		throw new Error('error for rollback')
	})()
} catch (e) {
	console.error(e)
}

const books = Book3.list()

books.total`
	testFn.Language = model.ScriptLanguage_JAVASCRIPT
	testFn.ContentFormat = model.ScriptContentFormat_TEXT

	result, err := api.Apply(util.SystemContext, model.ScriptMapperInstance.ToUnstructured(testFn))

	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, float64(0), output)
}

func TestTransactionalUpdate(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `
const Book3 = resource({
	name: 'Book3',
	properties: {
		name: {
			type: 'string',
		},
		description: {
			type: 'string',
		}
	}
})

Book3.deleteAll()

const book = Book3.create({ name: 'Book3', description: 'aaa'})

try {
	transactional(() => {
		book.description = 'bbb'

		Book3.update(book)
	
		throw new Error('error for rollback')
	})()
} catch (e) {
	console.error(e.message)
}

const books = Book3.list()

books.content[0].description`
	testFn.Language = model.ScriptLanguage_JAVASCRIPT
	testFn.ContentFormat = model.ScriptContentFormat_TEXT

	result, err := api.Apply(util.SystemContext, model.ScriptMapperInstance.ToUnstructured(testFn))

	if err != nil {
		t.Error(err)
		return
	}

	assert.NotNil(t, result["output"])

	if t.Failed() {
		return
	}

	output := result["output"]

	assert.Equal(t, "aaa", output)
}

func TestTransactionalDelete(t *testing.T) {
	var api = api.NewInterface(container)

	var testFn = new(model.Script)
	testFn.Source = `
const Book3 = resource({
	name: 'Book3',
	properties: {
		name: {
			type: 'string',
		},
		description: {
			type: 'string',
		}
	}
})

Book3.deleteAll()

const book = Book3.create({ name: 'Book3', description: 'aaa'})

try {
	transactional(() => {
		Book3.deleteAll()
	
		throw new Error('error for rollback')
	})()
} catch (e) {
	console.error(e.message)
}

const books = Book3.list()

const actual = books.content[0]

if (actual.id !== book.id) {
	throw new Error('book id is wrong')
}

if (actual.name !== book.name) {
	throw new Error('book id is wrong')
}

if (actual.description !== book.description) {
	throw new Error('book id is wrong')
}
`
	testFn.Language = model.ScriptLanguage_JAVASCRIPT
	testFn.ContentFormat = model.ScriptContentFormat_TEXT

	_, err := api.Apply(util.SystemContext, model.ScriptMapperInstance.ToUnstructured(testFn))

	if err != nil {
		t.Error(err)
		return
	}
}
