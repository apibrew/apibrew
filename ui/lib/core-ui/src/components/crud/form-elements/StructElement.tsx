import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import {ResourceProperty} from "../../../model";
import React from "react";
import Box from "@mui/material/Box";
import {Tab, Tabs} from "@mui/material";
import {FormItem} from "./FormItem.tsx";

export interface StructElementProps {
    config: CrudFormItem
    properties?: ResourceProperty[]
}

export function StructElement(props: StructElementProps) {
    // tabs will be combined
    const tabs = props.config.children.filter((item) => item.kind === 'tab')
    const other = props.config.children.filter((item) => item.kind !== 'tab')

    const [value, setValue] = React.useState(0);

    return <React.Fragment>
        {tabs.length > 0 && <React.Fragment>
            <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
                <Tabs value={value} onChange={(_, value) => setValue(value)} aria-label="basic tabs example">
                    {tabs.map((tab, index) => <Tab key={index} value={index} label={tab.title}/>)}
                </Tabs>
            </Box>
            {tabs[value].children && <StructElement properties={props.properties} config={tabs[value]}/>}
        </React.Fragment>}
        {other.map((child, index) => (
            <Box key={index} flex={1} style={{flex: 1}}>
                <FormItem properties={props.properties} config={child}/>
            </Box>
        ))}
    </React.Fragment>
}