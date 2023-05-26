import {Module} from "../../model/ui/module";
import React from "react";
import {ModuleData} from '../../model/module-data'
import {ModuleService} from "../../service/module";

export interface DynamicComponentProps {
    module?: Module
    moduleName?: string
    modulePackage?: string
    componentName: string

    componentProps?: any
    children?: React.ReactNode
}


export function DynamicComponent(props: DynamicComponentProps) {
    const [moduleData, setModuleData] = React.useState<ModuleData>(undefined)
    const [loaded, setLoaded] = React.useState(false)

    React.useEffect(() => {
        ModuleService.loadModule(props.moduleName || props.module?.name || '', props.modulePackage || props.module?.package || '').then((moduleData) => {
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
