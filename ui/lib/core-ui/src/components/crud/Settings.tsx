import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form"
import {Record, RecordService} from "../../service/record"
import React, {useEffect} from "react"
import {Crud, CrudName} from "../../model/schema";
import {ResourceService} from "../../service/resource";
import {resetCrudForm} from "./helper";

export interface SettingsProps {
    resource: Resource
    updateCrud: (crud: Crud) => void
}

export function Settings(props: SettingsProps): JSX.Element {
    const navigate = useNavigate()
    const [selfCrud, setSelfCrud] = React.useState<Crud>()
    const [crudConfig, setCrudConfig] = React.useState<Record>({
        name: `ResourceCrud-${props.resource.namespace}-${props.resource.name}`,
        resource: props.resource.name,
        namespace: props.resource.namespace,
    })
    const [crudResource, setCrudResource] = React.useState<Resource>()

    const params = useParams<{ id: string }>()

    useEffect(() => {
        RecordService.findBy<Crud>('ui', CrudName, 'name', 'CrudSettings')
            .then((record) => {
                setSelfCrud(record)
            })

        RecordService.findBy<Record>('ui', CrudName, 'name', crudConfig.name)
            .then((record) => {
                if (record) {
                    setCrudConfig(record)
                }
            }, console.warn)

        ResourceService.getByName(CrudName, 'ui').then((resource) => {
            setCrudResource(resource)
        })
    }, [params.id])

    if (!selfCrud) {
        return <>Loading...</>
    }

    return (
        <PageLayout breadcrumbs={[
            {label: 'Country', to: '../'},
            {label: 'Crud Settings'}
        ]}
                    actions={<>
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
                    </>}>
            <>
                {crudResource && <Form resource={crudResource}
                                       formConfig={selfCrud.formConfig}
                                       record={crudConfig}
                                       setRecord={setCrudConfig}/>}
            </>
        </PageLayout>
    )
}