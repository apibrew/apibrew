import React, {useState} from "react";
import {SecurityConstraintsInputSimple} from "./SecurityConstraintsInputSimple.tsx";
import {SecurityConstraintsInputAdvanced} from "./SecurityConstraintsInputAdvanced.tsx";
import {Tab, Tabs} from "@mui/material";
import {useValue} from "../../context/value.ts";
import { SecurityConstraint } from "@apibrew/client";

export interface SecurityConstraintsInputProps {
    mode: 'role' | 'resource' | 'namespace'
}

export function SecurityConstraintsInput(props: SecurityConstraintsInputProps) {
    const [tab, setTab] = React.useState(0)

    const valueContext = useValue()
    const [constraints, setConstraints] = useState<SecurityConstraint[]>(valueContext.value ?? [])

    constraints.forEach((constraint) => {
        if (!constraint.operation) {
            constraint.operation = 'FULL'
        }

        if (!constraint.permit) {
            constraint.permit = 'ALLOW'
        }

        if (!constraint.recordIds) {
            constraint.recordIds = []
        }
    })

    return <>
        <Tabs value={tab} onChange={(_, value) => setTab(value)}>
            <Tab label="Simple"/>
            <Tab label="Advanced"/>
        </Tabs>
        {tab === 0 && <SecurityConstraintsInputSimple
            constraints={constraints}
            setConstraints={value => {
                setConstraints(value)
                valueContext.onChange(value)
            }}
            mode={props.mode}/>}
        {tab === 1 && <SecurityConstraintsInputAdvanced
            constraints={constraints}
            setConstraints={value => {
                setConstraints(value)
                valueContext.onChange(value)
            }}
            mode={props.mode}/>}
    </>
}