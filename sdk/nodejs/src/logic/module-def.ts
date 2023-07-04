import { Module } from "../model"

export interface ModuleBackendParams {
    authentication?: {
        username: string
        password: string
    },
    backendUrl?: string
}

export interface ModuleInitParams {

}

export interface ModuleInitFunctionParams {
    moduleId: string
}

export type ModuleInitFunction = (params: ModuleInitFunctionParams) => Promise<void>

let module: Module

let moduleRegistry: Record<string, any> = {}

export async function setModule(_module: Module) {
    module = _module
}

export function getModule() {
    if (!module) {
        throw new Error('Module not inited')
    }

    return module
}

export function registerModuleChild(name: string, element: any) {
    moduleRegistry[name] = element
}

export function getModuleChild(name: string) {
    return moduleRegistry[name]
}

