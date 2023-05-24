import {Resource} from "../../model"
import {Record} from "../../service/record"
import {CrudFormConfig, CrudFormItem} from "../../model/schema";
import {ResourceContext, useResource} from "../../context/resource";
import {useValue, ValueContext} from "../../context/value";
import {RecordContext} from "../../context/record";
import React from "react";
import {ResourcePropertyContext, useResourceProperty} from "../../context/property";
import {FormElement} from "./FormElement";
import Box from "@mui/material/Box";
import {Tab, Tabs} from "@mui/material";
import {CrudSettingsFormConfig} from "./form-elements/CrudSettingsFormConfig";
import {CrudSettingsGridConfig} from "./form-elements/CrudSettingsGridConfig";

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

/*
          - property
          - input
          - tab
          - section
          - group
          - custom
 */

export interface FormItemCollectionProps {
    items: CrudFormItem[]
}

export function FormItemCollection(props: FormItemCollectionProps) {
    // tabs will be combined
    const tabs = props.items.filter((item) => item.kind === 'tab')
    const other = props.items.filter((item) => item.kind !== 'tab')

    const [value, setValue] = React.useState(0);

    return <>
        {tabs.length > 0 && <>
            <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
                <Tabs value={value} onChange={(_, value) => setValue(value)} aria-label="basic tabs example">
                    {tabs.map((tab, index) => <Tab key={index} value={index} label={tab.title}/>)}
                </Tabs>
            </Box>
            {tabs[value].children && <FormItemCollection items={tabs[value].children}/>}
        </>}
        <Box display='flex' flexDirection='column'>
            {other.map((child) => (
                <FormItem config={child}/>
            ))}
        </Box>
    </>
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

        return <>
            <ResourcePropertyContext.Provider value={property}>
                <ValueContext.Provider value={{
                    value: value.value[property.name],
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
        </>
    } else if (props.config.kind === 'input') {
        return <FormElement resource={resource}
                            property={property!}
                            readOnly={value.readOnly}
                            {...props.config.params}
                            value={value.value}
                            setValue={val => {
                                value.onChange(val)
                            }}/>
    } else if (props.config.kind === 'section') {
        return <>
            <h3>{props.config.title}</h3>
            {props.config.children && <FormItemCollection items={props.config.children}/>}
            <hr/>
        </>
    } else if (props.config.kind === 'group') {
        return <>
            <h1>Group: {props.config.title}</h1>
            {props.config.children && <FormItemCollection items={props.config.children}/>}
        </>
    } else if (props.config.kind === 'custom') {
        switch (props.config.component) {
            case 'CrudSettingsFormConfig':
                return <CrudSettingsFormConfig config={props.config}/>
            case 'CrudSettingsGridConfig':
                return <CrudSettingsGridConfig config={props.config}/>
        }
    }

    return <>
        Unknown form item kind {props.config.kind} {props.config.component}
    </>
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
                    <FormItemCollection items={props.formConfig.children}/>
                </ValueContext.Provider>
            </RecordContext.Provider>
        </ResourceContext.Provider>
    );
    // const isNew = props.record.id === undefined
    //
    // return (
    //     <Box sx={{ display: 'flex', flexDirection: 'column' }}>
    //         {props.resource.properties.filter(not(isSpecialProperty)).map(property => (
    //             <Box m={1}>
    //                 <FormElement resource={props.resource}
    //                     property={property}
    //                     value={props.record[property.name]}
    //                     readOnly={props.readOnly}
    //                     setValue={value => {
    //                         props.setRecord({
    //                             ...props.record,
    //                             [property.name]: value
    //                         })
    //                     }} />
    //             </Box>
    //         ))}
    //         {!isNew && <Box>
    //             <Typography variant='h6'>System Properties</Typography>
    //             {props.resource.properties.filter(isSpecialProperty).map(property => (
    //                 <Box m={1}>
    //                     <FormElement resource={props.resource}
    //                         property={property}
    //                         readOnly={true}
    //                         value={props.record[property.name]}
    //                         setValue={value => {
    //                             props.setRecord({
    //                                 ...props.record,
    //                                 [property.name]: value
    //                             })
    //                         }} />
    //                 </Box>
    //             ))}
    //         </Box>}
    //     </Box>
}