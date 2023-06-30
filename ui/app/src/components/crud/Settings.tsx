import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout.tsx"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model/index.ts"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form.tsx"
import {Record, RecordService, useBreadCramps} from "@apibrew/ui-lib"
import React, {useEffect} from "react"
import {Crud, CrudName} from "../../model/ui/crud.ts";
import {resetCrudForm} from "./helper.ts";
import {useResourceByName} from "../../hooks/resource.ts";
import {useRecordByName} from "../../hooks/record.ts";

export interface SettingsProps {
    resource: Resource
    updateCrud: (crud: Crud) => void
}

export function Settings(props: SettingsProps): JSX.Element {
    const navigate = useNavigate()
    const [crudConfig, setCrudConfig] = React.useState<Record>({
        name: `ResourceCrud-${props.resource.namespace}-${props.resource.name}`,
        resource: props.resource.name,
        namespace: props.resource.namespace,
    })
    const crudResource = useResourceByName(CrudName, 'ui')
    const selfCrud = useRecordByName<Crud>(CrudName, 'ui', 'CrudSettings')

    const params = useParams<{ id: string }>()

    useBreadCramps({title: 'Settings'})

    useEffect(() => {
        RecordService.findBy<Record>('ui', CrudName, 'name', crudConfig.name)
            .then((record) => {
                if (record) {
                    setCrudConfig(record)
                }
            }, console.warn)
    }, [params.id])

    if (!selfCrud) {
        return <span>Loading...</span>
    }

    return (
        <Box flexDirection='column' display='flex' width='100%' height='100%' padding='20px'>
            <Box flexGrow={1}>
                {crudResource && <Form resource={crudResource}
                                       formConfig={selfCrud.formConfig}
                                       record={crudConfig}
                                       setRecord={setCrudConfig}/>}
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
                                RecordService.apply('ui', CrudName, crudConfig).then(() => {
                                    props.updateCrud(crudConfig as Crud)
                                    navigate('../')
                                })
                            }}
                            startIcon={<Save/>}>Save</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant={'outlined'}
                            color='success'
                            size='small'
                            onClick={() => {
                                resetCrudForm(props.resource).then((newCrudConfig) => {
                                    props.updateCrud(newCrudConfig)
                                    navigate('../')
                                })
                            }}
                            startIcon={<Save/>}>Reset to Defaults</Button>
                </Box>
            </Box>
        </Box>
    )
}
