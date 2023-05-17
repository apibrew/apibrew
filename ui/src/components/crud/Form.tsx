import { Box, Typography } from "@mui/material"
import { Resource } from "../../model"
import { Record } from "../../service/record"
import { FormElement } from "./FormElement"
import { isSpecialProperty } from "../../util/property"
import { not } from "../../util/lambda"

export interface FormProps {
    resource: Resource
    record: Record
    readOnly?: boolean
    setRecord: (record: Record) => void
}

export function Form(props: FormProps) {
    const isNew = props.record.id === undefined

    return (
        <Box sx={{ display: 'flex', flexDirection: 'column' }}>
            {props.resource.properties.filter(not(isSpecialProperty)).map(property => (
                <Box m={1}>
                    <FormElement resource={props.resource}
                        property={property}
                        value={props.record[property.name]}
                        readOnly={props.readOnly}
                        setValue={value => {
                            props.setRecord({
                                ...props.record,
                                [property.name]: value
                            })
                        }} />
                </Box>
            ))}
            {!isNew && <Box>
                <Typography variant='h6'>System Properties</Typography>
                {props.resource.properties.filter(isSpecialProperty).map(property => (
                    <Box m={1}>
                        <FormElement resource={props.resource}
                            property={property}
                            readOnly={true}
                            value={props.record[property.name]}
                            setValue={value => {
                                props.setRecord({
                                    ...props.record,
                                    [property.name]: value
                                })
                            }} />
                    </Box>
                ))}
            </Box>}
        </Box>
    )
}