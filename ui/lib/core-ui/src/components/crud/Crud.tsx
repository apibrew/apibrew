import {Box, CircularProgress, LinearProgress} from "@mui/material"
import React, {JSX, useEffect, Fragment, useContext} from "react"
import {Resource} from "../../model"
import {ResourceService} from "../../service/resource"
import {Route, Routes} from "react-router-dom"
import {List} from "./List"
import {New} from "./New"
import {View} from "./View"
import {Update} from "./Update"
import {Settings} from "./Settings";
import {RecordService} from "../../service/record";
import {Crud as CrudModel, CrudName} from "../../model/ui/crud.ts";
import {resetCrudForm} from "./helper";
import {useResourceByName} from "../../hooks/resource.ts";
import {Loading} from "../basic/Loading.tsx";

export interface CrudProps {
    namespace: string
    resource: string
}

export function Crud(props: CrudProps): JSX.Element {
    const resource = useResourceByName(props.resource, props.namespace)


    const [crudConfig, setCrudConfig] = React.useState<CrudModel>()

    useEffect(() => {
        if (resource) {
            const name = `ResourceCrud-${resource.namespace}-${resource.name}`
            RecordService.findBy<CrudModel>('ui', CrudName, 'name', name)
                .then((record) => {
                    if (record) {
                        setCrudConfig(record)
                    } else {
                        resetCrudForm(resource).then((newCrudConfig) => {
                            setCrudConfig(newCrudConfig)
                        })
                    }
                }, (e) => {
                    console.warn(e)
                })
        }
    }, [resource])

    if (!resource || !crudConfig) {
        return <Loading/>
    }

    return (
        <Box>
            <Routes>
                <Route path="new" element={<New crudConfig={crudConfig} resource={resource}/>}/>
                {!crudConfig.hideSettings &&
                    <Route path="settings" element={<Settings resource={resource} updateCrud={updatedCrudConfig => {
                        setCrudConfig(updatedCrudConfig)
                    }}/>}/>}
                <Route path=":id/edit" element={<Update crudConfig={crudConfig} resource={resource}/>}/>
                <Route path=":id/view" element={<View crudConfig={crudConfig} resource={resource}/>}/>
                <Route path="" element={<List crudConfig={crudConfig} resource={resource}/>}/>
            </Routes>
        </Box>
    )
}
