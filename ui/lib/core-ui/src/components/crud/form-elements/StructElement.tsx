import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import {ResourceProperty} from "../../../model";
import React from "react";
import Box from "@mui/material/Box";
import {Tab, Tabs} from "@mui/material";
import {FormItem} from "./FormItem.tsx";
import {TabElement} from "./TabElement.tsx";
import {useResource} from "../../../context/resource.ts";
import {PropertyPathContext} from "../PropertyPathContext.tsx";

export interface StructElementProps {
    config: CrudFormItem
    properties?: ResourceProperty[]
}

export function StructElement(props: StructElementProps) {
    if (!props.config.children) {
        return <Box/>
    }
    // tabs will be combined
    const tabs = props.config.children.filter((item) => item.kind === 'tab')
    const other = props.config.children.filter((item) => item.kind !== 'tab')

    return <React.Fragment>
        {tabs.length > 0 && <TabElement tabs={tabs} properties={props.properties}/>}
        {other.map((child, index) => (
            <Box key={index} flex={1} style={{flex: 1}}>
                <FormItem properties={props.properties} config={child}/>
            </Box>
        ))}
    </React.Fragment>
}