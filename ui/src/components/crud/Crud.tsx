import { Box } from "@mui/material"
import React, { useEffect, JSX } from "react"
import { Resource } from "../../model"
import { ResourceService } from "../../service/resource"
import { Route, Routes } from "react-router-dom"
import { List } from "./List"
import { New } from "./New"
import { View } from "./View"
import { Update } from "./Update"
import {Settings} from "./Settings";

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

    return (
        <>
            {resource && <Box>
                <Routes>
                    <Route path="new" element={<New resource={resource} />} />
                    <Route path="settings" element={<Settings resource={resource} />} />
                    <Route path=":id/edit" element={<Update resource={resource} />} />
                    <Route path=":id/view" element={<View resource={resource} />} />
                    <Route path="" element={<List resource={resource} />} />
                </Routes>
            </Box>}
        </>
    )
}