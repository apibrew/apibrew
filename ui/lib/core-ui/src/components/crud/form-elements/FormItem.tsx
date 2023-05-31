import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import {ResourceProperty} from "../../../model";
import {useValue} from "../../../context/value.ts";
import {useResourceProperty} from "../../../context/property.ts";
import {useResource} from "../../../context/resource.ts";
import {PropertyElement} from "./PropertyElement.tsx";
import {FormInput} from "./FormInput.tsx";
import React from "react";
import {StructElement} from "./StructElement.tsx";
import Box from "@mui/material/Box";
import {DynamicComponent} from "../../dynamic/DynamicComponent.tsx";

export interface FormItemProps {
    config: CrudFormItem
    property?: ResourceProperty
    properties?: ResourceProperty[]
}

export function FormItem(props: FormItemProps) {
    const value = useValue()
    const parentProperty = useResourceProperty(false)
    const resource = useResource()

    if (props.config.kind === 'property') {
        return <PropertyElement properties={props.properties} config={props.config}/>
    } else if (props.config.kind === 'input') {
        return <FormInput property={parentProperty}
                          readOnly={value.readOnly}
                          config={props.config}
                          value={value.value}
                          setValue={val => {
                              value.onChange(val)
                          }}/>
    } else if (props.config.kind === 'section') {
        return <React.Fragment>
            <h3>{props.config.title}</h3>
            {props.config.children && <StructElement properties={props.properties} config={props.config}/>}
            <hr/>
        </React.Fragment>
    } else if (props.config.kind === 'group') {
        return <React.Fragment>
            <Box sx={{display: 'flex', flex: 1, gap: 3}} flexDirection='row'>
                {props.config.children &&
                    <StructElement properties={props.properties} config={props.config}/>}
            </Box>
        </React.Fragment>
    } else if (props.config.kind === 'custom') {
        return <DynamicComponent component={props.config.component}>
            {props.config.children && <StructElement properties={props.properties} config={props.config}/>}
        </DynamicComponent>
    } else if (props.config.kind === 'container') {
        let properties = props.property.properties
        if (props.property.typeRef) {
            properties = resource.types.find(item => item.name == props.property.typeRef).properties
        }
        return <React.Fragment>
            {props.config.children &&
                <StructElement properties={properties} config={props.config}/>}
        </React.Fragment>
    }

    throw new Error(`Unknown form item kind ${props.config.kind}`)
}