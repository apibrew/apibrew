import {Box} from "@mui/material"
import React, {JSX, useEffect, Fragment} from "react"
import {Resource} from "../../model"
import {ResourceService} from "../../service/resource"
import {Route, Routes} from "react-router-dom"
import {List} from "./List"
import {New} from "./New"
import {View} from "./View"
import {Update} from "./Update"
import {Settings} from "./Settings";
import {RecordService} from "../../service/record";
import {Crud as CrudModel, CrudName} from "../../model/schema";
import {resetCrudForm} from "./helper";

export interface CrudProps {
    namespace: string
    resource: string
}

export function Crud(props: CrudProps): JSX.Element {
    const [resource, setResource] = React.useState<Resource>()

    useEffect(() => {
        ResourceService.getByName(props.resource, props.namespace).then((resource) => {
            setResource(resource)
        })
    }, [props.namespace, props.resource])

    const [crudConfig, setCrudConfig] = React.useState<CrudModel>()

    useEffect(() => {
        if (resource) {
            const name = `ResourceCrud-${resource.namespace}-${resource.name}`
            RecordService.findBy<CrudModel>('ui', CrudName, 'name', name)
                .then((record) => {
                    if (record) {
                        setCrudConfig(record)
                    }
                }, (e) => {
                    console.warn(e)
                    resetCrudForm(resource).then((newCrudConfig) => {
                        setCrudConfig(newCrudConfig)
                    })
                })
        }
    }, [resource])

    return (
        <Fragment>
            {resource && crudConfig && <Box>
                <Routes>
                    <Route path="new" element={<New crudConfig={crudConfig} resource={resource}/>}/>
                    <Route path="settings" element={<Settings resource={resource} updateCrud={updatedCrudConfig => {
                        setCrudConfig(updatedCrudConfig)
                    }}/>}/>
                    <Route path=":id/edit" element={<Update crudConfig={crudConfig} resource={resource}/>}/>
                    <Route path=":id/view" element={<View crudConfig={crudConfig} resource={resource}/>}/>
                    <Route path="" element={<List crudConfig={crudConfig} resource={resource}/>}/>
                </Routes>
            </Box>}
        </Fragment>
    )
}
