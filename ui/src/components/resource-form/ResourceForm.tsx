import { Box, Button, Card, CardActions, CardContent, CardHeader } from "@mui/material";
import React from "react";
import { Resource } from "../../model";
import { ResourceBasicForm } from "./ResourceBasicForm";
import { ResourceAdvancedForm } from "./ResourceAdvancedForm";

export type ResourceFormVariant = 'basic' | 'advanced'

export interface ResourceFormProps {
    initResource: Resource;
    onSave: (resource: Resource) => void;
}

export function ResourceForm(props: ResourceFormProps): JSX.Element {
    const [formVariant, setFormVariant] = React.useState<ResourceFormVariant>('basic')
    const [resource, setResource] = React.useState<Resource>(props.initResource)

    return <>
        <Card>
            <CardHeader title={<Box sx={{ display: 'flex' }}>
                <Box>
                    <span>Resource Form</span>
                </Box>
                <Box sx={{ flexGrow: 5 }} />
                <Box m={0.5}>
                    <Button variant={formVariant === 'basic' ? 'contained' : 'outlined'}
                        size="small"
                        onClick={() => setFormVariant('basic')}
                        color="primary">Basic</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant={formVariant === 'advanced' ? 'contained' : 'outlined'}
                        size="small"
                        onClick={() => setFormVariant('advanced')}
                        color="primary">Advanced</Button>
                </Box>
            </Box>} />
            <CardContent>
                {formVariant === 'basic' && <ResourceBasicForm resource={resource} onChange={setResource} />}
                {formVariant === 'advanced' && <ResourceAdvancedForm resource={resource} onChange={setResource} />}
            </CardContent>
            <CardActions sx={{ display: 'flex' }}>
                <Box sx={{ flexGrow: 1 }} />
                <Box m={0.5}>
                    <Button variant="outlined" size="small" color="warning">Cancel</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant="contained" size="small" color="success" onClick={() => {
                        props.onSave(resource)
                    }}>Save</Button>
                </Box>
            </CardActions>
        </Card>
    </>
}
