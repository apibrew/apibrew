import React from "react";
import {SecurityConstraintsInputSimple} from "./SecurityConstraintsInputSimple.tsx";
import {SecurityConstraintsInputAdvanced} from "./SecurityConstraintsInputAdvanced.tsx";
import {Tab, Tabs} from "@mui/material";

export interface SecurityConstraintsInputProps {
    mode: 'role' | 'resource' | 'namespace'
}

export function SecurityConstraintsInput(props: SecurityConstraintsInputProps) {
    const [tab, setTab] = React.useState(0)
    return <>
        <Tabs value={tab} onChange={(_, value) => setTab(value)}>
            <Tab label="Simple"/>
            <Tab label="Advanced"/>
        </Tabs>
        {tab === 0 && <SecurityConstraintsInputSimple mode={props.mode}/>}
        {tab === 1 && <SecurityConstraintsInputAdvanced mode={props.mode}/>}
    </>
}