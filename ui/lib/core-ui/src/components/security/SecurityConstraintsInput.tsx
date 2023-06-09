import React, {useState} from "react";
import {SecurityConstraintsInputSimple} from "./SecurityConstraintsInputSimple.tsx";
import {SecurityConstraintsInputAdvanced} from "./SecurityConstraintsInputAdvanced.tsx";
import {Tab, Tabs} from "@mui/material";
import {useValue} from "../../context/value.ts";
import {useRecord} from "../../context/record.ts";
import {SecurityConstraint} from "../../model/security-constraint.ts";

export interface SecurityConstraintsInputProps {
    mode: 'role' | 'resource' | 'namespace'
}

export function SecurityConstraintsInput(props: SecurityConstraintsInputProps) {
    const [tab, setTab] = React.useState(0)

    const valueContext = useValue()
    const record = useRecord<{ id: string, name: string, namespace: string }>()

    const [constraints, setConstraints] = useState<SecurityConstraint[]>(valueContext.value ?? [])

    constraints.forEach((constraint) => {
        switch (props.mode) {
            case 'role':
                constraint.role = record.name
                break
            case 'resource':
                constraint.resource = record.name
                constraint.namespace = record.namespace
                break
            case 'namespace':
                constraint.namespace = record.name
                break
        }

        if (!constraint.operation) {
            constraint.operation = 'FULL'
        }

        if (!constraint.permit) {
            constraint.permit = 'PERMIT_TYPE_ALLOW'
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