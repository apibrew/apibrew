import React, {useEffect} from 'react';
import {JointPaper} from "../joint/JointPaper";
import {JointGraph} from "../joint/JointGraph";
import {Link} from "../joint/Link";
import {Resource} from "../../model";
import {ResourceService} from "../../service/resource";
import {RectContainer} from "../joint/RectContainer";
import Button from "@mui/material/Button";
import {isAnnotationEnabled, SpecialProperty} from "../../model/annotations";

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
                {resources.map((resource, index) => <RectContainer
                    key={(resource.namespace ?? '') + resource.name}
                    position={{x: index * 210, y: 10}}
                    size={{width: 200, height: 205}}
                    attrs={{
                        rect: {
                            cursor: 'crosshair',
                            strokeWidth: 0,
                            fill: '#DFDFDF',
                            rx: 1,
                        }
                    }}
                    name={(resource.namespace ?? '') + resource.name}>
                    <div style={{margin: '10px', pointerEvents: 'auto', background: '#CFCFCF', textAlign: 'center'}}>
                        <h3>{resource.name}</h3>
                        {resource.properties?.filter(item => !isAnnotationEnabled(item.annotations, SpecialProperty))
                            .slice(0, 4)
                            .map((property, index) =>
                                <div key={property.name}>
                                    <span key={index}>{property.name}:{property.type!.toLowerCase()}</span>
                                    <br/>
                                </div>
                            )}
                        <Button variant='outlined' color='info' size='small'>Details</Button>
                    </div>
                </RectContainer>)}
                <Link source={'item-1'} target={'item-2'}/>
            </JointPaper>
        </JointGraph>
    </div>
}

