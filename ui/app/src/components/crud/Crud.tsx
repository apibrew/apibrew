import React, {JSX, useEffect, Fragment, useContext} from "react"
import {Route, Routes} from "react-router-dom"
import {List} from "./List.tsx"
import {New} from "./New.tsx"
import {View} from "./View.tsx"
import {Update} from "./Update.tsx"
import {Settings} from "./Settings.tsx";
import {RecordService} from "@apibrew/ui-lib";
import {Crud as CrudModel, CrudName} from "../../model/ui/crud.ts";
import {resetCrudForm} from "./helper.ts";
import {useResourceByName} from "../../hooks/resource.ts";
import {Loading} from "../basic/Loading.tsx";
import {ResourceContext} from "../../context/resource.ts";

export interface CrudProps {
    namespace: string
    resource: string
}

export function Crud(props: CrudProps): JSX.Element {
    const resource = useResourceByName(props.resource, props.namespace)


    const [crudConfig, setCrudConfig] = React.useState<CrudModel>()

    console.log('crudConfig', crudConfig)

    useEffect(() => {
        if (resource) {
            const name = `ResourceCrud-${resource.namespace.name}-${resource.name}`
            RecordService.findBy<CrudModel>('ui', CrudName, 'name', name)
                .then((record) => {
                    console.log('RecordService.findBy', record)
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
        <ResourceContext.Provider value={resource}>
            <Routes>
                <Route path="new/*" element={<New crudConfig={crudConfig} resource={resource}/>}/>
                {!crudConfig.hideSettings &&
                    <Route path="settings/*" element={<Settings resource={resource} updateCrud={updatedCrudConfig => {
                        setCrudConfig(updatedCrudConfig)
                    }}/>}/>}
                <Route path=":id/edit/*" element={<Update crudConfig={crudConfig} resource={resource}/>}/>
                <Route path=":id/view/*" element={<View crudConfig={crudConfig} resource={resource}/>}/>
                <Route path="" element={<List crudConfig={crudConfig} resource={resource}/>}/>
            </Routes>
        </ResourceContext.Provider>
    )
}
