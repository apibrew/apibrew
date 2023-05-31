import {Resource, ResourceProperty} from "../../model"
import {Record} from "../../service/record"
import {FormConfig as CrudFormConfig, FormItem as CrudFormItem} from "../../model/ui/crud.ts";
import {ResourceContext, useResource} from "../../context/resource";
import {useValue, ValueContext} from "../../context/value";
import {RecordContext} from "../../context/record";
import React from "react";
import {ResourcePropertyContext, useResourceProperty} from "../../context/property";
import {FormElement} from "./FormElement";
import Box from "@mui/material/Box";
import {Tab, Tabs} from "@mui/material";
import {DynamicComponent} from "../dynamic/DynamicComponent.tsx";

export interface FormProps {
    resource: Resource
    record: Record
    readOnly?: boolean
    setRecord: (record: Record) => void
    formConfig: CrudFormConfig
}

export interface FormItemProps {
    config: CrudFormItem
    property?: ResourceProperty
    properties?: ResourceProperty[]
}

export interface FormItemCollectionProps {
    items: CrudFormItem[]
    properties?: ResourceProperty[]
}

export function FormItemCollection(props: FormItemCollectionProps) {
    // tabs will be combined
    const tabs = props.items.filter((item) => item.kind === 'tab')
    const other = props.items.filter((item) => item.kind !== 'tab')

    const [value, setValue] = React.useState(0);

    return <React.Fragment>
        {tabs.length > 0 && <React.Fragment>
            <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
                <Tabs value={value} onChange={(_, value) => setValue(value)} aria-label="basic tabs example">
                    {tabs.map((tab, index) => <Tab key={index} value={index} label={tab.title}/>)}
                </Tabs>
            </Box>
            {tabs[value].children && <FormItemCollection properties={props.properties} items={tabs[value].children}/>}
        </React.Fragment>}
        {other.map((child, index) => (
            <Box key={index} flex={1} style={{flex: 1}}>
                <FormItem properties={props.properties} config={child}/>
            </Box>
        ))}
    </React.Fragment>
}

export function FormItem(props: FormItemProps) {
    const value = useValue()
    const parentProperty = useResourceProperty(false)
    const resource = useResource()

    if (props.config.kind === 'property') {
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
            console.log('props.properties', props.properties)
            throw new Error(`Property ${props.config.propertyPath} not found`)
        }

        let subProperties = property.properties

        if (property.typeRef) {
            console.log(property.typeRef, resource.types)
            // @ts-ignore
            subProperties = resource.types.find(item => item.name == property.typeRef).properties!
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
                        <FormItemCollection properties={subProperties} items={props.config.children}/>}
                </ValueContext.Provider>
            </ResourcePropertyContext.Provider>
        </React.Fragment>
    } else if (props.config.kind === 'input') {
        return <FormElement property={parentProperty}
                            readOnly={value.readOnly}
                            config={props.config}
                            value={value.value}
                            setValue={val => {
                                value.onChange(val)
                            }}/>
    } else if (props.config.kind === 'section') {
        return <React.Fragment>
            <h3>{props.config.title}</h3>
            {props.config.children && <FormItemCollection properties={props.properties} items={props.config.children}/>}
            <hr/>
        </React.Fragment>
    } else if (props.config.kind === 'group') {
        return <React.Fragment>
            <Box sx={{display: 'flex', flex: 1, gap: 3}} flexDirection='row'>
                {props.config.children &&
                    <FormItemCollection properties={props.properties} items={props.config.children}/>}
            </Box>
        </React.Fragment>
    } else if (props.config.kind === 'custom') {
        return <DynamicComponent component={props.config.component}>
            {props.config.children && <FormItemCollection properties={props.properties} items={props.config.children}/>}
        </DynamicComponent>
    } else if (props.config.kind === 'container') {
        let properties = props.property.properties
        if (props.property.typeRef) {
            // @ts-ignore
            properties = resource.types.find(item => item.name == props.property.typeRef).properties
        }
        return <React.Fragment>
            {props.config.children &&
                <FormItemCollection properties={properties} items={props.config.children}/>}
        </React.Fragment>
    }

    throw new Error(`Unknown form item kind ${props.config.kind}`)
}

export function Form(props: FormProps) {
    return (
        <ResourceContext.Provider value={props.resource}>
            <RecordContext.Provider value={props.record}>
                <ValueContext.Provider value={{
                    value: props.record,
                    onChange: props.setRecord,
                    readOnly: props.readOnly,
                }}>
                    {props.formConfig.children &&
                        <FormItemCollection properties={props.resource.properties} items={props.formConfig.children}/>}
                </ValueContext.Provider>
            </RecordContext.Provider>
        </ResourceContext.Provider>
    );
}
