import React, {useEffect} from 'react';
import {Resource} from "../../model";
import {ResourceService} from "../../service/resource";
import Button from "@mui/material/Button";
import {ResourceElement} from "./ResourceElement";
import { Arrow } from './Arrow';

// React component to render the diagram
export const Designer: React.FC = () => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)

    useEffect(() => {
        ResourceService.list().then(list => {
            console.log(list)
            setResources(list.filter(item => item.namespace !== 'system'))
        })
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
        <br/>
        <svg style={{width: '100%', height: '600px'}}>
            {resources.map((resource, index) => {
                const x = 410 * index
                const y = 10
                return <g key={(resource.namespace ?? '') + resource.name} transform={`translate(${x}, ${y})`}>
                    <ResourceElement resource={resource}/>
                </g>
            })}
            <Arrow
                isHighlighted={true}
                startPoint={{x: 186, y: 191}}
                endPoint={{x: 416, y: 41}}
            />
        </svg>
    </div>
}

