import React from "react";
import {ResourcePropertyContext, useResourceProperty} from "../../../context/property.ts";
import {useValue, ValueContext} from "../../../context/value.ts";
import {StructElement} from "./StructElement.tsx";
import {useResource} from "../../../context/resource.ts";
import {ResourceProperty} from "../../../model";
import {FormItem} from "../../../model/ui/crud.ts";

export interface PropertyElementProps {
    properties: ResourceProperty[]
    config: FormItem
}

export function PropertyElement(props: PropertyElementProps) {
    const value = useValue()
    const parentProperty = useResourceProperty(false)
    const resource = useResource()

    if (!props.properties) {
        throw new Error(`Properties not available in property form item`)
    }

    let property = props.properties.find((property) => property.name === props.config.propertyPath)

    if (!property && parentProperty?.type == 'OBJECT') {
        property = {
            name: props.config.propertyPath,
            type: 'OBJECT',
        }
    }

    if (!property) {
        throw new Error(`Property ${props.config.propertyPath} not found`)
    }

    let subProperties = property.properties ?? []

    if (property.typeRef) {
        subProperties = resource.types.find(item => item.name == property.typeRef)?.properties
    }

    let newValue = value.value ? value.value[property.name] : undefined

    if (property.type == 'STRUCT' && !newValue) {
        newValue = {}
    }

    return <React.Fragment>
        <ResourcePropertyContext.Provider value={property}>
            <ValueContext.Provider value={{
                value: newValue,
                readOnly: value.readOnly,
                onChange: (val) => {
                    value.onChange({
                        ...value.value,
                        [property.name]: val
                    })
                }
            }}>
                {props.config.children &&
                    <StructElement properties={subProperties} config={props.config}/>}
            </ValueContext.Provider>
        </ResourcePropertyContext.Provider>
    </React.Fragment>
}