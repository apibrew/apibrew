import { Box, Button, Card, CardActions, CardContent, CardHeader } from '@mui/material'
import React, { useContext } from 'react'
import { type Resource } from '../../model'
import { ResourceBasicForm } from './ResourceBasicForm'
import { ResourceAdvancedForm } from './ResourceAdvancedForm'
import { ResourceService } from "@apibrew/core-lib"
import { LayoutContext } from '../../context/layout-context'
import { AxiosError } from 'axios'

export type ResourceFormVariant = 'basic' | 'advanced'

export interface ResourceFormProps {
    resources: Resource[]
    initResource: Resource
    onSave?: (resource: Resource) => void
    onCancel?: () => void
}

export function ResourceForm(props: ResourceFormProps): JSX.Element {
    const [formVariant, setFormVariant] = React.useState<ResourceFormVariant>('basic')
    const [resource, setResource] = React.useState<Resource>(props.initResource)

    const layoutOptions = useContext(LayoutContext)

    return <React.Fragment>
        <Card>
            <CardHeader title={<Box sx={{ display: 'flex' }}>
                <Box>
                    <span>Resource Form</span>
                </Box>
                <Box sx={{ flexGrow: 5 }} />
                <Box m={0.5}>
                    <Button variant={formVariant === 'basic' ? 'contained' : 'outlined'}
                        size="small"
                        onClick={() => { setFormVariant('basic') }}
                        color="primary">Basic</Button>
                </Box>
                <Box m={0.5}>
                    <Button variant={formVariant === 'advanced' ? 'contained' : 'outlined'}
                        size="small"
                        onClick={() => { setFormVariant('advanced') }}
                        color="primary">Advanced</Button>
                </Box>
            </Box>} />
            <CardContent>
                {formVariant === 'basic' && <ResourceBasicForm resources={props.resources} resource={resource} onChange={setResource} />}
                {formVariant === 'advanced' && <ResourceAdvancedForm resource={resource} onChange={setResource} />}
            </CardContent>
            <CardActions sx={{ display: 'flex' }}>
                <Box sx={{ flexGrow: 1 }} />
                <Box m={0.5}>
                    {props.onCancel && <Button variant="outlined" size="small" color="warning" onClick={() => {
                        if (props.onCancel) {
                            props.onCancel()
                        }
                    }}>Cancel</Button>}
                </Box>
                <Box m={0.5}>
                    <Button variant="contained" size="small" color="success" onClick={() => {
                        ResourceService.save(resource).then(() => {
                            layoutOptions.showAlert({ severity: 'success', message: 'Resource saved successfully' })

                            if (props.onSave) {
                                props.onSave(resource)
                            }
                        }).catch((error) => {
                            if (error instanceof AxiosError && error.response?.status === 400) {
                                layoutOptions.showAlert({ severity: 'error', message: error.response.data.message })
                                return
                            }

                            layoutOptions.showAlert({ severity: 'error', message: error.message })
                        })
                    }}>Save</Button>
                </Box>
            </CardActions>
        </Card>
    </React.Fragment>
}
