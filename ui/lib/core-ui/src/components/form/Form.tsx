import {useResource} from "../../context/resource.ts";
import {Button, Card, CardActions, CardContent, CardHeader} from "@mui/material";
import React, {ReactNode, useState} from "react";
import Box from "@mui/material/Box";
import {RecordContext} from "../../context/record.ts";
import {ResourceAutoForm} from "./resource-auto-form.tsx";
import {ValueContext} from "../../context/value.ts";

export interface FormProps {
    isNew?: boolean
    viewOnly?: boolean
    children: ReactNode
}

export function Form(props: FormProps) {
    const resource = useResource()
    const [record, setRecord] = useState({})

    let children = props.children

    if (!children) {
        children = <ResourceAutoForm/>
    }

    return <Card>
        <CardHeader title={'Form'}/>
        <CardContent>
            <RecordContext.Provider value={record}>
                <ValueContext.Provider value={{
                    value: record,
                    onChange: setRecord,
                    readOnly: props.viewOnly,
                }}>
                    {children}
                </ValueContext.Provider>
            </RecordContext.Provider>
        </CardContent>
        <CardActions>
            <Box flexGrow={1}/>
            <Button color='warning'>cancel</Button>
            <Button color='success'>save</Button>
        </CardActions>
    </Card>
}
