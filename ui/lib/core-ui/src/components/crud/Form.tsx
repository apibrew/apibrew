import {Resource} from "../../model"
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
}

export interface FormItemCollectionProps {
    items: CrudFormItem[]
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
            {tabs[value].children && <FormItemCollection items={tabs[value].children}/>}
        </React.Fragment>}
        {other.map((child, index) => (
            <Box key={index} flex={1} style={{width: '100%', margin: '5px'}}>
                <FormItem config={child}/>
            </Box>
        ))}
    </React.Fragment>
}

export function FormItem(props: FormItemProps) {
    const resource = useResource()
    const value = useValue()
    const property = useResourceProperty(false)

    if (props.config.kind === 'property') {
        const property = resource.properties.find((property) => property.name === props.config.propertyPath)

        if (!property) {
            throw new Error(`Property ${props.config.propertyPath} not found`)
        }

        return <React.Fragment>
            <ResourcePropertyContext.Provider value={property}>
                <ValueContext.Provider value={{
                    value: value.value[property.name],
                    readOnly: value.readOnly,
                    onChange: (val) => {
                        value.onChange({
                            ...value.value,
                            [property.name]: val
                        })
                    }
                }}>
                    {props.config.children && <FormItemCollection items={props.config.children}/>}
                </ValueContext.Provider>
            </ResourcePropertyContext.Provider>
        </React.Fragment>
    } else if (props.config.kind === 'input') {
        return <FormElement resource={resource}
                            property={property}
                            readOnly={value.readOnly}
                            config={props.config}
                            value={value.value}
                            setValue={val => {
                                value.onChange(val)
                            }}/>
    } else if (props.config.kind === 'section') {
        return <React.Fragment>
            <h3>{props.config.title}</h3>
            {props.config.children && <FormItemCollection items={props.config.children}/>}
            <hr/>
        </React.Fragment>
    } else if (props.config.kind === 'group') {
        return <React.Fragment>
            <Box sx={{display: 'flex', width: '100%'}} flexDirection='row' letterSpacing='10px'>
                {props.config.children && <FormItemCollection items={props.config.children}/>}
            </Box>
        </React.Fragment>
    } else if (props.config.kind === 'custom') {
        return <DynamicComponent component={props.config.component}>
            {props.config.children && <FormItemCollection items={props.config.children}/>}
        </DynamicComponent>
    }


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
                    {props.formConfig.children && <FormItemCollection items={props.formConfig.children}/>}
                </ValueContext.Provider>
            </RecordContext.Provider>
        </ResourceContext.Provider>
    );
}
