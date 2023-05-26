import {ModuleData} from "../model/module-data.ts";
import {RecordService} from "./record.ts";
import * as jsxRuntime from "react/jsx-runtime";
import * as CoreUI from "../index.ts";
import {Module} from "../model/ui/module.ts";
import {ActionComponent} from "../model/component-interfaces.ts";

export namespace ModuleService {
    const modules: ModuleData[] = []

    function require(path: string) {
        switch (path) {
            case 'react/jsx-runtime':
                return jsxRuntime
            case 'core-ui':
                return CoreUI
        }
    }

    export async function loadModule(name: string, pkg: string): Promise<ModuleData> {
        const existingModule = modules.find(m => m.name === name && m.package === pkg)

        if (existingModule) {
            return existingModule
        }

        const module = await RecordService.findByMulti<Module>('ui', 'Module', [
            {property: 'name', value: name},
            {property: 'package', value: pkg},
        ])

        if (!module) {
            throw new Error(`Module ${pkg}/${name} not found`)
        }

        const moduleData: ModuleData = {
            exports: {},
            name: module.name,
            package: module.package,
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

    export async function loadModuleComponent<T>(name: string, pkg: string, componentName: string): Promise<T> {
        const module = await loadModule(name, pkg)

        return module.exports[componentName]
    }

    export function getModuleComponent<T>(name: string, pkg: string, componentName: string): T {
        const module = modules.find(m => m.name === name && m.package === pkg)

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
