import { Box, Button } from "@mui/material"
import { PageLayout } from "../../layout/PageLayout"
import { Cancel, PlusOneOutlined, Save } from "@mui/icons-material"
import { Resource } from "../../model"
import { useNavigate, useParams } from "react-router-dom"
import { Form } from "./Form"
import { Record, RecordService } from "../../service/record"
import React, { useEffect } from "react"

export interface UpdateProps {
    resource: Resource
}

export function Update(props: UpdateProps): JSX.Element {
    const navigate = useNavigate()
    const [record, setRecord] = React.useState<Record>()

    const params = useParams<{ id: string }>()

    useEffect(() => {
        RecordService.get<Record>(props.resource.namespace ?? 'default', props.resource.name, params.id!)
            .then((record) => {
                setRecord(record)
            })
    }, [params.id])

    if (!record) {
        return <>Loading...</>
    }

    return (
        <PageLayout pageTitle={props.resource.name} actions={<>
            <Box sx={{ display: 'flex' }}>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                        color='primary'
                        size='small'
                        onClick={() => {
                            navigate('../')
                        }}
                        startIcon={<Cancel />}>Cancel</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                        color='success'
                        size='small'
                        onClick={() => {
                            RecordService.update(props.resource.namespace ?? 'default', props.resource.name, record).then(() => {
                                navigate('../')
                            })
                        }}
                        startIcon={<Save />}>Save</Button>
                </Box>
            </Box>
        </>}>
            <>
                <Form resource={props.resource} record={record} setRecord={setRecord} />
            </>
        </PageLayout>
    )
}