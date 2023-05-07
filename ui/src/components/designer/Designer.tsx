import React, {useEffect} from 'react';
import {JointPaper} from "../joint/JointPaper";
import {JointGraph} from "../joint/JointGraph";
import {Resource} from "../../model";
import {ResourceService} from "../../service/resource";
import Button from "@mui/material/Button";
import {ResourceElement} from "./ResourceElement";

// React component to render the diagram
export const Designer: React.FC = () => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)

    useEffect(() => {
        ResourceService.list().then(setResources)
    }, [])

    return <div>
        <Button variant='outlined' color='info' size='small' onClick={() => {
            setZoomLevel(Math.min(3, zoomLevel + 0.2))
        }}>Zoom in</Button> &nbsp;
        <Button variant='outlined' color='info' size='small' onClick={() => {
            setZoomLevel(1)
        }}>Zoom Reset</Button> &nbsp;
        <Button variant='outlined' color='info' size='small' onClick={() => {
            setZoomLevel(Math.max(0.2, zoomLevel - 0.2))
        }}>Zoom out</Button> &nbsp;

        Zoom: {zoomLevel}
        <JointGraph>
            <JointPaper
                preventCollision={true}
                zoomLevel={zoomLevel}
                options={{
                    width: '100%',
                    height: '600px',
                    gridSize: 10,
                }}>
                {resources.map((resource, index) => {
                    const x = 210 * index
                    const y = 10
                    return <ResourceElement position={{x: x, y: y}} key={(resource.namespace ?? '') + resource.name} resource={resource}/>
                })}
            </JointPaper>
        </JointGraph>
    </div>
}

