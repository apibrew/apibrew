import { Box, FormControl, FormGroup, FormHelperText, FormLabel, TextField } from "@mui/material";
import { Resource } from "../../model";

export interface ResourceBasicFormProps {
    resource: Resource;
    onChange: (resource: Resource) => void;
}

export function ResourceBasicForm(props: ResourceBasicFormProps): JSX.Element {
    return <>
        <Box>
            <FormControl>
                <FormLabel>Resource Name</FormLabel>
                <TextField value={props.resource.name} onChange={(e) => props.onChange({ ...props.resource, name: e.target.value })} />
                <FormHelperText>Resource name is required</FormHelperText>
            </FormControl>
        </Box>
    </>
}
