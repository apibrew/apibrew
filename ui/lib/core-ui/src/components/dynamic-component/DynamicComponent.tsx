import {Module} from "../../model/schema.ts";
import React from "react";
import * as jsxRuntime from 'react/jsx-runtime'
import * as CoreUI from '../../index.ts'
import {RecordService} from '../../index.ts'

interface ExportOptions {
    [key: string]: any
}

interface ModuleData {
    exports: ExportOptions
    module: Module
}

export interface DynamicComponentProps {
    module?: Module
    moduleName?: string
    modulePackage?: string
    componentName: string

    componentProps?: any
    children?: React.ReactNode
}

const modules: ModuleData[] = []

async function loadModule(name: string, pkg: string): Promise<ModuleData> {
    const existingModule = modules.find(m => m.module.name === name && m.module.package === pkg)

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
        module,
    }

    const code = atob(module.source)

    const func = new Function('exports', 'require', code)

    func(moduleData.exports, (path: string) => {
        switch (path) {
            case 'react/jsx-runtime':
                return jsxRuntime
            case 'core-ui':
                return CoreUI
        }
    })

    modules.push(moduleData)

    return moduleData
}

export function DynamicComponent(props: DynamicComponentProps) {
    const [moduleData, setModuleData] = React.useState<ModuleData>(undefined)
    const [loaded, setLoaded] = React.useState(false)

    React.useEffect(() => {
        loadModule(props.moduleName || props.module?.name || '', props.modulePackage || props.module?.package || '').then((moduleData) => {
            setModuleData(moduleData)
            setLoaded(true)
        })
    })

    if (!loaded) {
        return <div>Loading...</div>
    }

    if (!moduleData) {
        return <div>Module {props.moduleName || props.module?.name || ''} not found</div>
    }

    const Component = moduleData.exports[props.componentName]

    if (!Component) {
        return <div>Component {props.componentName} not found</div>
    }

    return <Component {...props.componentProps} />
}
