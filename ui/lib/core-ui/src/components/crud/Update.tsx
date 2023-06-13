import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form"
import {Record, RecordService} from "../../service/record"
import React, {useContext, useEffect, useState} from "react"
import {Crud} from "../../model/ui/crud.ts";
import {LayoutContext, useBreadCramps} from "../../context/layout-context.ts";
import {useErrorHandler} from "../../hooks/error-handler.tsx";
import {useResource} from "../../context/resource.ts";
import {filterRecordForUpdate} from "../../service/authorization.ts";

export interface UpdateProps {
    resource: Resource
    crudConfig: Crud
}

export function Update(props: UpdateProps): JSX.Element {
    const resource = useResource()
    const navigate = useNavigate()
    const errorHandler = useErrorHandler()
    const [record, setRecord] = React.useState<Record>()
    const layoutContext = useContext(LayoutContext)
    const [loading, setLoading] = useState(true)

    const params = useParams<{ id: string }>()

    useBreadCramps({title: params.id}, {title: 'Update'})

    const load = async () => {
        setLoading(true)
        return RecordService.get<Record>(props.resource.namespace ?? 'default', props.resource.name, params.id!)
            .then((record) => {
                setRecord(record)
                setLoading(false)
            }, errorHandler)
    }

    useEffect(() => {
        load()
    }, [params.id])

    if (!record) {
        return <span>Loading...</span>
    }

    return (
        <PageLayout pageTitle={props.resource.name} actions={<React.Fragment>
            <Box sx={{display: 'flex'}}>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                            color='warning'
                            size='small'
                            onClick={() => {
                                load().then(() => {
                                    layoutContext.showAlert({
                                        severity: 'success',
                                        message: 'Form reloaded'
                                    })
                                })
                            }}
                            startIcon={<Cancel/>}>Reset</Button>
                </Box>
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
                                const updateFilteredRecord = filterRecordForUpdate(resource, record)

                                RecordService.update(props.resource.namespace ?? 'default', props.resource.name, updateFilteredRecord).then(() => {
                                    layoutContext.showAlert({
                                        severity: 'success',
                                        message: resource.name + ' updated'
                                    })
                                }, errorHandler)
                            }}
                            startIcon={<Save/>}>Save</Button>
                </Box>
            </Box>
        </React.Fragment>}>
            <React.Fragment>
                {!loading && <Form resource={props.resource} record={record} setRecord={setRecord}
                                   formConfig={props.crudConfig.formConfig}/>}
            </React.Fragment>
        </PageLayout>
    )
}
