import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout.tsx"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model/index.ts"
import {useNavigate} from "react-router-dom"
import {Form} from "./Form.tsx"
import {Record, RecordService, useBreadCramps} from "@apibrew/ui-lib"
import React, {useEffect} from "react"
import {Crud} from "../../model/ui/crud.ts";
import {useErrorHandler} from "../../hooks/error-handler.tsx";

export interface NewProps {
    resource: Resource
    crudConfig: Crud
}

export function New(props: NewProps) {
    const navigate = useNavigate()
    const [record, setRecord] = React.useState<Record>({})
    const errorHandler = useErrorHandler()

    useBreadCramps({title: 'New'})

    return (
        <Box flexDirection='column' display='flex' width='100%' height='100%' padding='20px'>
            <Box flexGrow={1}>
                <Form resource={props.resource} record={record} setRecord={setRecord}
                      formConfig={props.crudConfig.formConfig}/>
            </Box>
            <Box sx={{display: 'flex', paddingBottom: '10px', width: '100%'}}>
                <Box flexGrow={1}/>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                            color='primary'
                            size='small'
                            onClick={() => {
                                navigate('../')
                            }}
                            startIcon={<Cancel/>}>Cancel</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                            color='success'
                            size='small'
                            onClick={() => {
                                RecordService.create(props.resource.namespace ?? 'default', props.resource.name, record).then(() => {
                                    navigate('../')
                                }, errorHandler)
                            }}
                            startIcon={<Save/>}>Save</Button>
                </Box>
            </Box>
        </Box>
    )
}
