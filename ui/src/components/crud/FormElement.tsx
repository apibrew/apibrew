import { Button, Checkbox, FormControl, FormHelperText, FormLabel, MenuItem, Select, TextField, TextFieldProps, TextFieldVariants } from "@mui/material"
import { Resource, ResourceProperty } from "../../model"
import React, { ReactNode, useEffect, useMemo } from "react"
import { Variant } from "@mui/material/styles/createTypography"
import { ResourceService } from "../../service/resource"
import { Record, RecordService } from "../../service/record"

export interface FormElementProps {
    resource: Resource
    property: ResourceProperty
    readOnly?: boolean
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
    }} />
}

const ReferenceField = (props: FieldProps & { namespace: string, referencedResourceName: string }) => {
    const [records, setRecords] = React.useState<Record[]>([])

    useEffect(() => {
        console.log('loading records');

        (async () => {
            const list = (await RecordService.list<Record>(props.namespace, props.referencedResourceName))
            setRecords(list)
        })();
    }, [])



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
                return <MenuItem value={item.id}>{item.name}</MenuItem>
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
    }} />
}

export function FormElement(props: FormElementProps) {
    let title = props.property.title || props.property.name
    const resource = props.resource
    const referencedResourceName = props.property.reference?.referencedResource ?? ''

    let Field: FieldComponent = useMemo(() => {
        switch (props.property.type) {
            case 'STRING':
                return TextField
            case 'INT32':
            case 'INT64':
                return FieldValueConvertWrapper((props: FieldProps) => <RegexTextField type='number' pattern={/\d+/} {...props} />, (e) => parseInt(e.target.value))
            case 'FLOAT32':
            case 'FLOAT64':
                return FieldValueConvertWrapper((props: FieldProps) => <TextField type='number' {...props} />, (e) => parseFloat(e.target.value))
            case 'BOOL':
                return FieldValueConvertWrapper((props: FieldProps) => <Checkbox {...props} checked={props.value} />, (e) => e.target.checked)
            case 'REFERENCE':
                return (props: FieldProps) => <ReferenceField namespace={resource.namespace ?? 'default'} referencedResourceName={referencedResourceName} {...props} />
        }

        return TextField
    }, [])

    return (
        <>
            <FormControl>
                <FormLabel>{title}</FormLabel>
                <Field required={props.property.required}
                    disabled={props.readOnly}
                    value={props.value}
                    onChange={(e) => {
                        props.setValue(e.target.value)
                    }} />
                {props.property.description && <FormHelperText>{props.property.description}</FormHelperText>}
            </FormControl>
        </>
    )
}