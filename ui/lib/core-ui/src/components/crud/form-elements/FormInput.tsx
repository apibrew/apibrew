import {
    Checkbox,
    FormControl,
    FormHelperText,
    FormLabel,
    MenuItem,
    Select,
    TextField,
    TextFieldProps
} from "@mui/material"
import {ResourceProperty} from "../../../model"
import React, {useEffect, useMemo, useState} from "react"
import {Record, RecordService} from "@apibrew/core-lib"
import {FormItem} from "../../../model/ui/crud.ts";
import {useResource} from "../../../context/resource.ts";
import {ListElement} from "./ListElement.tsx";
import {StructElement} from "./StructElement.tsx";

export interface FormElementProps {
    property: ResourceProperty
    readOnly?: boolean
    config: FormItem
    value: any
    setValue: (value: any) => void
}

interface FieldProps {
    required?: boolean
    disabled?: boolean
    value: any
    onChange: (e: any) => void
}

export type FieldComponent = (props: FieldProps) => JSX.Element

const RegexTextField = (props: TextFieldProps & { pattern: RegExp | string }) => {
    const expr = new RegExp(props.pattern)

    return <TextField {...props} onChange={e => {
        if (props.onChange) {
            if (expr.test(e.target.value)) {
                props.onChange(e)
            }
        }
    }}/>
}

const ReferenceField = (props: FieldProps & { namespace: string, referencedResourceName: string }) => {
    const [records, setRecords] = React.useState<Record[]>([])

    useEffect(() => {
        (async () => {
            const list = (await RecordService.list<Record>(props.namespace, props.referencedResourceName))
            setRecords(list)
        })();
    }, [])

    if (records.length == 0) {
        return <></>
    }

    return (
        <Select
            disabled={props.disabled}
            value={props.value?.id ?? ''}
            onChange={e => {
                props.onChange({
                    target: {
                        value: {
                            id: e.target.value,
                        }
                    }
                })
            }}>
            {records.map(item => {
                return <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
            })}
        </Select>
    )
}

const FieldValueConvertWrapper = (Field: FieldComponent, converter: (val: any) => any) => (props: FieldProps) => {
    return <Field {...props} onChange={e => {
        if (props.onChange) {
            props.onChange({
                target: {
                    value: converter(e)
                }
            })
        }
    }}/>
}

export function FormInput(props: FormElementProps) {
    const title = props.property.title || props.property.name
    const resource = useResource()

    const referencedResourceName = props.property.reference?.referencedResource ?? ''
    const [value, setValue] = useState(props.value)

    const Field = useMemo(() => {
        switch (props.property.type) {
            case 'STRING':
                return TextField
            case 'INT32':
            case 'INT64':
                return FieldValueConvertWrapper((props: FieldProps) => <RegexTextField type='number'
                                                                                       pattern={/\d+/} {...props} />, (e) => parseInt(e.target.value))
            case 'FLOAT32':
            case 'FLOAT64':
                return FieldValueConvertWrapper((props: FieldProps) => <TextField
                    type='number' {...props} />, (e) => parseFloat(e.target.value))
            case 'BOOL':
                return FieldValueConvertWrapper((props: FieldProps) => <Checkbox {...props}
                                                                                 checked={props.value}/>, (e) => e.target.checked)
            case 'REFERENCE':
                return (props: FieldProps) => <ReferenceField namespace={resource.namespace ?? 'default'}
                                                              referencedResourceName={referencedResourceName} {...props} />
            case 'LIST':
                return (_props: FieldProps) => <ListElement config={props.config} {..._props}/>
            case 'STRUCT':
                let properties = props.property.properties

                if (props.property.typeRef) {
                    const type = resource.types.find(item => item.name == props.property.typeRef)
                    properties = type?.properties || []
                }

                return (_props: FieldProps) => <StructElement properties={props.property.properties}
                                                              config={props.config}/>
        }

        return TextField
    }, [props.property.type, referencedResourceName, resource.namespace])

    const field = <Field required={props.property.required}
                         disabled={props.readOnly}
                         {...props.config.params}
                         value={value}
                         onChange={(e) => {
                             setValue(e.target.value)
                             props.setValue(e.target.value)
                         }}/>

    const hideLabel = (props.config.params && props.config.params['hideLabel'])

    if (hideLabel) {
        return <FormControl style={{width: '100%'}}>
            {field}
        </FormControl>
    }

    return (
        <FormControl style={{width: '100%'}}>
            <FormLabel>{title}</FormLabel>
            {field}
            {props.property.description && <FormHelperText>{props.property.description}</FormHelperText>}
        </FormControl>
    )
}
