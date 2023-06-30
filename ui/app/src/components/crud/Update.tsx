import {Box, Button} from "@mui/material"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model/index.ts"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form.tsx"
import {Record, RecordService} from "@apibrew/ui-lib"
import React, {useContext, useEffect, useState} from "react"
import {Crud} from "../../model/ui/crud.ts";
import {useErrorHandler} from "../../hooks/error-handler.tsx";
import {useResource} from "../../context/resource.ts";
import {AuthorizationService, LayoutContext, useBreadCramps} from "@apibrew/ui-lib"
import {Loading} from "../basic/Loading.tsx";

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

    if (loading) {
        return <Loading/>
    }

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
                                const updateFilteredRecord = AuthorizationService.filterRecordForUpdate(resource, record)

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
        </Box>
    )
}
