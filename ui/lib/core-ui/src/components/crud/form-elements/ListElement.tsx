import {ResourcePropertyContext, useResourceProperty} from "../../../context/property.ts";
import IconButton from "@mui/material/IconButton";
import {Icon} from "../../Icon.tsx";
import Box from "@mui/material/Box";
import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import React from "react";
import {FormItem} from "./FormItem.tsx";
import {ValueContext} from "../../../context/value.ts";
import {Paper} from "@mui/material";
import {DeleteForever} from "@mui/icons-material";
import {GenericEvent} from "../../../model/event.ts";

export interface ListElementProps {
    required?: boolean
    disabled?: boolean
    value: object[]
    onChange: (e: GenericEvent<object[]>) => void
    useTable?: boolean
    config: CrudFormItem
}

export function ListElement(props: ListElementProps) {
    const property = useResourceProperty(true)
    // const [items, setItems] = useState<any[]>(props.value ?? [])

    if (property.type !== 'LIST') {
        throw new Error('ListFormElements can only be used with a list property')
    }

    const items = props.value ?? []

    return <>
        <Paper>
            <Box m={1}>
                <IconButton onClick={() => {
                    let newItem = {}

                    if (property.item.type === 'STRUCT') {
                        newItem = {}
                    } else {
                        newItem = ''
                    }

                    props.onChange({
                        target: {
                            value: [...items, newItem]
                        }
                    })
                }}>
                    <Icon name={'add'}/>
                </IconButton>
                {items.map((item, index) => {
                    return <ValueContext.Provider key={index} value={{
                        value: item,
                        onChange: (val: any) => {
                            const newItems = [...items]
                            newItems[index] = val
                            props.onChange({
                                target: {
                                    value: newItems
                                }
                            })
                        },
                        readOnly: props.disabled,
                    }}>
                        <Box display='flex'>
                            <Box flex={1}>
                                <ResourcePropertyContext.Provider value={property.item}>
                                    <FormItem config={props.config}/>
                                </ResourcePropertyContext.Provider>
                            </Box>
                            <IconButton onClick={() => {
                                const newItems = [...items]
                                newItems.splice(index, 1)
                                props.onChange({
                                    target: {
                                        value: newItems
                                    }
                                })
                            }}>
                                <DeleteForever/>
                            </IconButton>
                        </Box>
                    </ValueContext.Provider>
                })}
            </Box>
        </Paper>
    </>
}