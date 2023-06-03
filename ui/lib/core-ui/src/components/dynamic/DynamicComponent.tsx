import React, {Suspense} from "react";
import {ModuleData} from '../../model/module-data'
import {ModuleService} from "../../service/module";
import Box from "@mui/material/Box";
import {Loading} from "../basic/Loading.tsx";

export interface DynamicComponentProps {
    component: string

    componentProps?: any
    children?: React.ReactNode
}


export function DynamicComponent(props: DynamicComponentProps) {
    const [moduleData, setModuleData] = React.useState<ModuleData>(undefined)
    const [loaded, setLoaded] = React.useState(false)

    const parts = props.component.split('/')

    if (parts.length !== 2) {
        throw new Error(`DynamicComponent component should module/componentName`)
    }

    const moduleName = parts[0]
    const componentName = parts[1]

    React.useEffect(() => {
        ModuleService.loadModule(moduleName).then((moduleData) => {
            setModuleData(moduleData)
            setLoaded(true)
        })
    })

    if (!loaded) {
        return <Loading/>
    }

    if (!moduleData) {
        return <div>Module {moduleName} not found</div>
    }

    const Component = moduleData.exports[componentName]

    if (!Component) {
        return <div>Component {componentName} not found</div>
    }

    if (props.componentProps?.debug) {
        return <Box sx={{border: '1px solid red'}} m={0.5}>
            <h6>{props.component} [{JSON.stringify(props.componentProps)}]</h6>
            <Component {...props.componentProps}>
                {props.children}
            </Component>
        </Box>
    }

    return <Suspense fallback={<Loading/>}>
        <Component {...props.componentProps}>
            {props.children}
        </Component>
    </Suspense>
}
