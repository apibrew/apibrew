import { ENGINE_REMOTE_ADDR, EXTENSION_NAME, TOKEN } from "./config";
import { initFunctionRegistry } from "./function-registry";
import { extensionRepository, functionExecutionEngineRepository } from "./client";
import { EventSelector, Extension, ExternalCall, FunctionExecutionEngine, HttpCall } from "@apibrew/client";

export let engineId: string

async function registerFunctionEngine() {
    await functionExecutionEngineRepository.apply({
        name: EXTENSION_NAME
    } as FunctionExecutionEngine)

    const result = await functionExecutionEngineRepository.findBy('name', EXTENSION_NAME)

    engineId = result.id

    console.log('Engine registration Id:', engineId)
    await initFunctionRegistry(engineId)
}

export async function registerExtensions(extensions: Extension[]) {
    console.log('registerExtensions called[first line]')
    const existingExtensions = (await extensionRepository.list()).content

    if (extensions.length === 0) {
        console.log('No extensions to register')
        return
    }

    for (const extension of extensions) {
        console.log('registerExtensions called', extension.name)
        let found = false

        for (const existing of existingExtensions) {
            if (existing.name === extension.name) {
                console.log('updating extension', extension.name)
                const result = await extensionRepository.apply(extension)
                console.log('updating extension done', result)
                found = true
            }
        }

        if (!found) {
            console.log('creating extension', extension.name)
            const result = await extensionRepository.create(extension)
            console.log('creating extension done', result)
        }
    }
}

export async function initExtensions() {
    await registerFunctionEngine();

    const extensions = prepareExtensions()

    await registerExtensions(extensions);
}

function prepareExtensions(): Extension[] {
    return [
        prepareFunctionExtension(),
        prepareFunctionExecutionExtension()
    ]
}

function prepareFunctionExtension(): Extension {
    const extension = {} as Extension
    extension.name = (EXTENSION_NAME)
    extension.sync = (true)
    const call = {} as ExternalCall
    const hCall = {} as HttpCall
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/reload`
    call.httpCall = hCall
    extension.call = (call)
    extension.order = (10000)
    extension.responds = (false)
    extension.sync = (false)
    const eventSelector = {} as EventSelector
    eventSelector.namespaces = (['logic'])
    eventSelector.actions = ['CREATE', 'UPDATE', 'DELETE'] as any
    extension.selector = (eventSelector)

    return extension
}

function prepareFunctionExecutionExtension(): Extension {
    const extension = {} as Extension
    extension.name = (`${EXTENSION_NAME}-execution`)
    extension.sync = (true)
    const call = {} as ExternalCall
    const hCall = {} as HttpCall
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/call/function`
    call.httpCall = hCall
    extension.call = (call)
    extension.order = (10)
    extension.responds = (true)
    extension.finalizes = (true)
    const eventSelector = {} as EventSelector
    eventSelector.resources = (['FunctionExecution'])
    eventSelector.namespaces = (['logic'])
    extension.selector = (eventSelector)

    return extension
}
