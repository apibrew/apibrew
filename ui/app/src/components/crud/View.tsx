import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout.tsx"
import {Cancel} from "@mui/icons-material"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form.tsx"
import {Record, RecordService, useBreadCramps} from "@apibrew/ui-lib"
import React, {useEffect} from "react"
import {Crud} from "../../model/ui/crud.ts";
import { Resource } from "@apibrew/client"

export interface ViewProps {
    resource: Resource
    crudConfig: Crud
}

export function View(props: ViewProps): JSX.Element {
    const navigate = useNavigate()
    const [record, setRecord] = React.useState<Record>()

    const params = useParams<{ id: string }>()

    useBreadCramps({title: params.id!}, {title: 'View'})

    useEffect(() => {
        RecordService.get<Record>(props.resource.namespace.name ?? 'default', props.resource.name, params.id!)
            .then((record) => {
                setRecord(record)
            })
    }, [params.id])

    if (!record) {
        return <>Loading...</>
    }

    return (
        <PageLayout>
            <Box sx={{display: 'flex', paddingBottom: '10px', width: '100%'}}>
                <Button variant={'outlined'}
                        color='primary'
                        size='small'
                        onClick={() => {
                            navigate('../')
                        }}
                        startIcon={<Cancel/>}>Cancel</Button>
            </Box>
            <Form resource={props.resource} readOnly={true} record={record} setRecord={setRecord}
                  formConfig={props.crudConfig.formConfig!}/>
        </PageLayout>
    )
}
