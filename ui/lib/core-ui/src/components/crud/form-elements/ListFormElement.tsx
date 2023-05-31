import {useResourceProperty} from "../../../context/property.ts";
import IconButton from "@mui/material/IconButton";
import {Icon} from "../../Icon.tsx";
import Box from "@mui/material/Box";
import {FormItem as CrudFormItem} from "../../../model/ui/crud.ts";
import React from "react";
import {FormItem} from "../Form";
import {ValueContext} from "../../../context/value.ts";
import {Paper} from "@mui/material";
import {DeleteForever} from "@mui/icons-material";

export interface ListFormElementProps {
    required?: boolean
    disabled?: boolean
    value: any
    onChange: (e: any) => void
    useTable?: boolean
    config: CrudFormItem
}

export function ListFormElements(props: ListFormElementProps) {
    const property = useResourceProperty(true)
    // const [items, setItems] = useState<any[]>(props.value ?? [])

    if (property.type !== 'LIST') {
        throw new Error('ListFormElements can only be used with a list property')
    }

    if (props.useTable) {

    }

    console.log(props.value)

    const items = props.value ?? []

    const children = props.config.children ?? []

    if (children.length !== 1) {
        throw new Error('ListFormElements can only have one child')
    }

    const subConfig = children[0]

    console.log(items)

    return <>
        <Paper>
            <Box m={1}>
                <IconButton onClick={() => {
                    props.onChange({
                        target: {
                            value: [...items, {}]
                        }
                    })
                }}><Icon name={'add'}/></IconButton>
                {items.map((item, index) => {
                    return <ValueContext.Provider value={{
                        value: item,
                        onChange: (val: any) => {
                            const newItems = [...items]
                            newItems[index] = val
                            props.onChange({
                                target: {
                                    value: newItems
                                }
                            })
                            console.log(newItems)
                        },
                        readOnly: props.disabled,
                    }}>
                        <Box display='flex'>
                            <Box flex={1}>
                                <FormItem property={property.item} config={subConfig}/>
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