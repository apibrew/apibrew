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

let moduleId: string

let moduleRegistry: Record<string, any> = {}

export async function setModuleId(_moduleId: string) {
    moduleId = _moduleId
}

export function getModuleId() {
    if (!moduleId) {
        throw new Error('Module not inited')
    }

    return moduleId
}

export function registerModuleChild(name: string, element: any) {
    moduleRegistry[name] = element
}

export function getModuleChild(name: string) {
    return moduleRegistry[name]
}

