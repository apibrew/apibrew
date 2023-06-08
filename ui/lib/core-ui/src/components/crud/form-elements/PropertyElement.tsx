import React, {useContext} from "react";
import {ResourcePropertyContext, useResourceProperty} from "../../../context/property.ts";
import {useValue, ValueContext} from "../../../context/value.ts";
import {StructElement} from "./StructElement.tsx";
import {useResource} from "../../../context/resource.ts";
import {ResourceProperty} from "../../../model";
import {FormItem} from "../../../model/ui/crud.ts";
import {PropertyPathContext} from "../PropertyPathContext.tsx";
import {AuthorizationService, Record} from "../../../service";
import {useRecord} from "../../../context/record.ts";

export interface PropertyElementProps {
    properties: ResourceProperty[]
    config: FormItem
}

export function PropertyElement(props: PropertyElementProps) {
    const value = useValue()
    const parentProperty = useResourceProperty(false)
    const resource = useResource()
    const record = useRecord<Record>()
    const propertyPath = useContext(PropertyPathContext)

    if (!props.properties) {
        console.log(props)
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
        console.log(props)
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

    let readOnly = value.readOnly

    let accessLevel: AuthorizationService.AccessLevel

    // check access
    if (propertyPath === '') { // root property
        accessLevel = AuthorizationService.checkResourcePropertyAccess(resource, property.name, record.id)
    } else {
        accessLevel = AuthorizationService.AccessLevel.READ_WRITE
    }

    if (accessLevel == AuthorizationService.AccessLevel.NONE) {
        return <></>
    }

    if (accessLevel == AuthorizationService.AccessLevel.READ) {
        readOnly = true
    }

    return <React.Fragment>
        <ResourcePropertyContext.Provider value={property}>
            <PropertyPathContext.Provider value={`${propertyPath}.${property.name}`}>
                <ValueContext.Provider value={{
                    value: newValue,
                    readOnly: readOnly,
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
            </PropertyPathContext.Provider>
        </ResourcePropertyContext.Provider>
    </React.Fragment>
}