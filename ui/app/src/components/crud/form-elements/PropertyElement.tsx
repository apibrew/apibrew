import React, {useContext} from "react";
import {ResourcePropertyContext, useResourceProperty} from "../../../context/property.ts";
import {useValue, ValueContext} from "../../../context/value.ts";
import {StructElement} from "./StructElement.tsx";
import {useResource} from "../../../context/resource.ts";
import {FormItem} from "../../../model/ui/crud.ts";
import {PropertyPathContext} from "../PropertyPathContext.tsx";
import {AuthorizationService, Record} from "@apibrew/ui-lib";
import {useRecord} from "../../../context/record.ts";
import { Property } from "@apibrew/client";

export interface PropertyElementProps {
    properties: Property[]
    config: FormItem
}

export function PropertyElement(props: PropertyElementProps) {
    const value = useValue()
    const parentProperty = useResourceProperty(false)
    const resource = useResource()
    const record = useRecord<Record>()
    const propertyPath = useContext(PropertyPathContext)

    if (!props.properties) {
        throw new Error(`Properties not available in property form item`)
    }

    let property = props.properties.find((property) => property.name === props.config.propertyPath)!

    if (!property && parentProperty?.type as any == 'OBJECT') {
        property = {
            name: props.config.propertyPath!,
            type: 'OBJECT' as any,
        } as Property
    }

    if (!property) {
        throw new Error(`Property ${props.config.propertyPath} not found`)
    }

    let subProperties = property.properties ?? []

    if (property.typeRef) {
        subProperties = resource.types!.find(item => item.name == property.typeRef)?.properties || []
    }

    let newValue = value.value ? value.value[property.name] : undefined

    if (property.type as any == 'STRUCT' && !newValue) {
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