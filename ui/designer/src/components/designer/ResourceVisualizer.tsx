import { Card } from '@mui/material'
import { type Resource } from '@apibrew/client'

export interface ResourceVisualizerProps {
    resource?: Resource
}

export const ResourceVisualizer = ({ resource }: ResourceVisualizerProps) => {
    return (
        <Card>
            <h2>Resource Visualizer</h2>
            <p>Resource: {resource?.name}</p>
        </Card>
    )
}
