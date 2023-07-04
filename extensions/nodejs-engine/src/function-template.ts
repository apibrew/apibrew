import { Function } from "./model/function"
import { Module } from "./model/module"

export const functionInitTemplate = `
const apbrClient = require('@apibrew/client')

const Client = apbrClient.Client

Client.setDefaultClient(global.apbrClient)

`

export const scriptFunctionTemplate = (record: Function) => `
${functionInitTemplate}

const fn = (function () {
    ${record.script}
})

exports.result = fn(params)
`

export const moduleFunctionTemplate = (record: Function) => `
${functionInitTemplate}

const { FunctionParams, defineFunction, setModuleId, getModuleChild } = require('@apibrew/client').LogicDef;

setModuleId('${record.module.id}')

const fnModule = require('./${record.module.id}/index.js')

const fn = fnModule["${record.name}"]  || getModuleChild("${record.name}")

exports.result = fn(params)
`

export const moduleInitTemplate = (module: Module) => `
${functionInitTemplate}

const { FunctionParams, defineFunction, setModuleId } = require('@apibrew/client').LogicDef;

setModuleId('${module.id}')

const fnModule = require('./index.js')

const initModule = fnModule["initModule"]

if (initModule) {
    const result = initModule()

    if (result.then) {
        result.then(() => {
            console.log('Module ${module.id} initialized')
        })
    } else {
        console.log('Module ${module.id} initialized')
    }
}
`