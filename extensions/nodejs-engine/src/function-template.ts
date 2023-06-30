import { Function } from "./model/function"

export const scriptFunctionTemplate = (record: Function) => `
const fn = (function () {
    ${record.script}
})

exports.result = fn(params)
`

export const moduleFunctionTemplate = (record: Function) => `
const fnModule = require('./${record.module.id}/index.js')

const fn = fnModule["${record.name}"]

exports.result = fn(params)
`