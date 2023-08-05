import Box from "@mui/material/Box";
import {Tab, Tabs} from "@mui/material";
import React from "react";
import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import {StructElement} from "./StructElement.tsx";
import {Property} from "@apibrew/client";
import {Navigate, Route, Routes, useNavigate, useParams} from "react-router-dom";

export interface TabElementProps {
    tabs: CrudFormItem[]
    properties?: Property[]
}

export function TabElement(props: TabElementProps) {
    const navigate = useNavigate()
    const params = useParams()

    const tabName = params["*"]

    const slugify = (str: string) => {
        return str.toLowerCase().replace(/ /g, '-')
    }

    const tabIndex = props.tabs.findIndex((tab) => slugify(tab.title!) === tabName)

    return <React.Fragment>
        <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
            <Tabs value={tabIndex} onChange={(_, value) => {
                navigate(slugify(props.tabs[value].title!))
            }} aria-label="basic tabs example">
                {props.tabs.map((tab, index) => <Tab key={index} value={index} label={tab.title}/>)}
            </Tabs>
        </Box>
        <Routes>
            {props.tabs.map((tab, index) => <Route path={slugify(tab.title!)}
                                                   key={tab.title}
                                                   element={<StructElement properties={props.properties}
                                                                           config={tab}/>}/>)}
            <Route path='' element={<Navigate to={slugify(props.tabs[0].title!)}/>}/>
        </Routes>
    </React.Fragment>
}
