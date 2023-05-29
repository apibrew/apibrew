import {ModuleData} from "../model/module-data.ts";
import {RecordService} from "./record.ts";
import * as jsxRuntime from "react/jsx-runtime";
import * as CoreUI from "../index.ts";
import {Module} from "../model/ui/module.ts";
import {ActionComponent} from "../model/component-interfaces.ts";

export namespace ModuleService {
    const modules: ModuleData[] = []
    const awaitedModules: Promise<ModuleData>[] = []

    function require(path: string) {
        switch (path) {
            case 'react/jsx-runtime':
                return jsxRuntime
            case 'core-ui':
                return CoreUI
        }
    }

    export async function loadModule(name: string): Promise<ModuleData> {
        await Promise.all(awaitedModules)

        const existingModule = modules.find(m => m.name === name)

        if (existingModule) {
            return existingModule
        }

        const module = await RecordService.findByMulti<Module>('ui', 'Module', [
            {property: 'name', value: name},
        ])

        if (!module) {
            throw new Error(`Module ${name} not found`)
        }

        const moduleData: ModuleData = {
            exports: {},
            name: module.name,
        }

        const code = atob(module.source)

        const func = new Function('exports', 'require', code)

        func(moduleData.exports, require)

        modules.push(moduleData)

        return moduleData
    }

    export function registerLocalModule(moduleData: ModuleData) {
        modules.push(moduleData)
    }

    export function registerLocalModuleAwait(moduleData: Promise<ModuleData>) {
        moduleData.then(md => {
            registerLocalModule(md)
        })

        awaitedModules.push(moduleData)
    }

    export async function loadModuleComponent<T>(name: string, pkg: string, componentName: string): Promise<T> {
        const module = await loadModule(name)

        return module.exports[componentName]
    }

    export function getModuleComponent<T>(name: string, pkg: string, componentName: string): T {
        const module = modules.find(m => m.name === name)

        if (!module) {
            throw new Error(`Module ${pkg}/${name} not found`)
        }

        return module.exports[componentName]
    }

    export async function executeActionComponent<R>(name: string, pkg: string, componentName: string, ...args: any) {
        const component = await loadModuleComponent<ActionComponent<R>>(name, pkg, componentName)

        return component.execute(...args)
    }
}
