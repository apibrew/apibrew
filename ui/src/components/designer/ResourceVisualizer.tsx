import { Card } from "@mui/material";
import { Resource } from "../../model";

export interface ResourceVisualizerProps {
    resource?: Resource
}

export const ResourceVisualizer: React.FC<ResourceVisualizerProps> = ({ resource }) => {
    return (
        <Card>
            <h2>Resource Visualizer</h2>
            <p>Resource: {resource?.name}</p>
        </Card>
    )
}
