import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Cancel, Save} from "@mui/icons-material"
import {Resource} from "../../model"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form"
import {Record, RecordService} from "../../service/record"
import React, {useEffect} from "react"
import {Crud, CrudName} from "../../model/ui/crud";
import {ResourceService} from "../../service/resource";
import {resetCrudForm} from "./helper";
import {useResourceByName} from "../../hooks/resource.ts";
import {useRecordByName} from "../../hooks/record.ts";
import {useBreadCramps} from "../../context/layout-context.ts";

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
        <PageLayout breadcrumbs={[
            {label: 'Country', to: '../'},
            {label: 'Crud Settings'}
        ]}
                    actions={<React.Fragment>
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
                    </React.Fragment>}>
            <React.Fragment>
                {crudResource && <Form resource={crudResource}
                                       formConfig={selfCrud.formConfig}
                                       record={crudConfig}
                                       setRecord={setCrudConfig}/>}
            </React.Fragment>
        </PageLayout>
    )
}
