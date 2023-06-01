import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model"
import {useNavigate} from "react-router-dom"
import {Form} from "./Form"
import {Record, RecordService} from "../../service/record"
import React from "react"
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

    return (
        <PageLayout actions={<React.Fragment>
            <Box sx={{display: 'flex'}}>
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
        </React.Fragment>}>
            <React.Fragment>
                <Form resource={props.resource} record={record} setRecord={setRecord}
                      formConfig={props.crudConfig.formConfig}/>
            </React.Fragment>
        </PageLayout>
    )
}
