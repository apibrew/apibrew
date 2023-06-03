import {useResource} from "../../context/resource.ts";
import {useValue} from "../../context/value.ts";
import {TextField} from "@mui/material";

export function FunctionScriptInput() {
    const valueProvider = useValue()

    return <>
        <TextField rows={20}
                   sx={{fontSize: 10}}
                   multiline={true}
                   value={valueProvider.value ?? ''}
                   onChange={e => {
                       valueProvider.onChange(e.target.value)
                   }}/>
    </>
}