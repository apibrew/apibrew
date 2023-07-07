import { ENGINE_REMOTE_ADDR, EXTENSION_NAME, TOKEN } from "./config";
import { initFunctionRegistry } from "./function-registry";
import { FunctionExecutionEngine } from "./model/function-execution-engine";
import { Extension } from "./model";
import { components } from "./model/base-schema";
import { extensionRepository, functionExecutionEngineRepository } from "./client";

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
    const existingExtensions = (await extensionRepository.list()).content

    for (const extension of extensions) {
        let found = false

        for (const existing of existingExtensions) {
            if (existing.name === extension.name) {
                console.log('updating extension')
                await extensionRepository.apply(extension)
                found = true
            }
        }

        if (!found) {
            console.log('creating extension')
            await extensionRepository.create(extension)
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
    const call = {} as components['schemas']['ExternalCall']
    const hCall = {} as components['schemas']['HttpCall']
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/reload`
    call.httpCall = hCall
    extension.call = (call)
    extension.order = (10000)
    extension.responds = (false)
    extension.sync = (false)
    const eventSelector = {} as components['schemas']['EventSelector']
    eventSelector.namespaces = (['logic'])
    eventSelector.actions = ['CREATE', 'UPDATE', 'DELETE']
    extension.selector = (eventSelector)

    return extension
}

function prepareFunctionExecutionExtension(): Extension {
    const extension = {} as Extension
    extension.name = (`${EXTENSION_NAME}-execution`)
    extension.sync = (true)
    const call = {} as components['schemas']['ExternalCall']
    const hCall = {} as components['schemas']['HttpCall']
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/call/function`
    call.httpCall = hCall
    extension.call = (call)
    extension.order = (10)
    extension.responds = (true)
    extension.finalizes = (true)
    const eventSelector = {} as components['schemas']['EventSelector']
    eventSelector.resources = (['FunctionExecution'])
    eventSelector.namespaces = (['logic'])
    extension.selector = (eventSelector)

    return extension
}
