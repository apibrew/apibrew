import React, {JSX, useContext} from "react";
import {Layout, LayoutComponent} from "../../model/ui/layout.ts";
import {DynamicComponent} from "./DynamicComponent.tsx";
import {LayoutParamsContext} from "../../context/layout-params.ts";

export interface LayoutInnerProps {
    children?: React.ReactNode
    layoutComponent: LayoutComponent
}

export function LayoutInner(props: LayoutInnerProps) {
    let children = props.layoutComponent.children && props.layoutComponent.children.map((item, index) =>
        <LayoutInner key={index} layoutComponent={item}/>
    )

    let existingParams = useContext(LayoutParamsContext)

    let params = {...existingParams, ...props.layoutComponent.params}

    for (const key in params) {
        if (params[key].startsWith && params[key].startsWith('$')) {
            const lookupKey = params[key].substr(1)
            params[key] = existingParams[lookupKey]
        }
    }

    return <LayoutParamsContext.Provider value={params}>
        <DynamicComponent component={props.layoutComponent.component} componentProps={params}>
            {children || props.children}
        </DynamicComponent>
    </LayoutParamsContext.Provider>
}

export function prepareLayoutComponent(layout: Layout): (props: LayoutInnerProps) => JSX.Element {
    return (props: LayoutInnerProps) => {

        return LayoutInner({layoutComponent: layout.layoutComponent, ...props})
    }
}
