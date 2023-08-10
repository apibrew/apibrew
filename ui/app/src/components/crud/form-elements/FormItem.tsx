import { FormItem as CrudFormItem } from "../../../model/ui/crud.ts";
import { useValue } from "../../../context/value.ts";
import { useResourceProperty } from "../../../context/property.ts";
import { useResource } from "../../../context/resource.ts";
import { PropertyElement } from "./PropertyElement.tsx";
import { FormInput } from "./FormInput.tsx";
import React from "react";
import { StructElement } from "./StructElement.tsx";
import Box from "@mui/material/Box";
import { DynamicComponent } from "@apibrew/ui-lib";
import { Property } from "@apibrew/client";

export interface FormItemProps {
    config: CrudFormItem
    properties?: Property[]
}

export function FormItem(props: FormItemProps) {
    const value = useValue()
    const resource = useResource()

    const parentProperty = useResourceProperty(true) as Property

    let properties = props.properties || parentProperty?.properties || []

    if (parentProperty?.typeRef) {
        const type = resource.types?.find(item => item.name == parentProperty.typeRef)
        properties = type?.properties || []
    }

    if (props.config.kind === 'property') {
        return <PropertyElement properties={properties} config={props.config} />
    } else if (props.config.kind === 'input') {
        return <FormInput property={parentProperty}
            readOnly={value.readOnly}
            config={props.config}
            value={value.value}
            setValue={val => {
                value.onChange(val)
            }} />
    } else if (props.config.kind === 'section') {
        return <React.Fragment>
            <h3>{props.config.title}</h3>
            {props.config.children && <StructElement properties={properties} config={props.config} />}
            <hr />
        </React.Fragment>
    } else if (props.config.kind === 'group') {
        return <React.Fragment>
            <Box sx={{ display: 'flex', flex: 1, gap: 3 }} flexDirection='row'>
                {props.config.children &&
                    <StructElement properties={properties} config={props.config} />}
            </Box>
        </React.Fragment>
    } else if (props.config.kind === 'custom') {
        return <DynamicComponent component={props.config.component!} componentProps={props.config.params}>
            {props.config.children && <StructElement properties={properties} config={props.config} />}
        </DynamicComponent>
    } else if (props.config.kind === 'container') {
        return <React.Fragment>
            {props.config.children &&
                <StructElement properties={properties} config={props.config} />}
        </React.Fragment>
    }

    throw new Error(`Unknown form item kind ${props.config.kind}`)
}