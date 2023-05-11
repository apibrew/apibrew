import React, {Fragment, useEffect} from 'react';
import {Resource} from "../../model";
import {ResourceService} from "../../service/resource";
import Button from "@mui/material/Button";
import {ResourceElement} from "./ResourceElement";
import {Link} from "./Link";
import {Scale} from "./Scale";
import {Movable, MovableComponent} from "./Movable";
import {SvgContainer} from "./SvgContainer";

// React component to render the diagram
export const Designer: React.FC = () => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)

    useEffect(() => {
        ResourceService.list().then(list => {
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

        Zoom: {Math.round(zoomLevel * 100)}% &nbsp;
        <br/>
        <svg className={'designer-parent'}
             style={{width: '100%', height: '600px'}}>
            <SvgContainer>
                <Scale level={zoomLevel}>
                    <Movable>
                        {resources.map((resource, index) => {
                            const x = 410 * index
                            const y = 10
                            return <MovableComponent key={(resource.namespace ?? '') + resource.name}>
                                <g transform={`translate(${x}, ${y})`}>
                                    <ResourceElement resource={resource}/>
                                </g>
                            </MovableComponent>
                        })}

                        {resources.map((resource, index) => {
                            return <Fragment key={resource.name}>
                                {resource.properties?.filter(item => item.type === 'REFERENCE').filter(item => item.reference && item.reference.referencedResource).map((property, index) => {
                                    return <Link key={`${resource.name}-${property.name}`}
                                                 sourceSelector={`.resource-${resource.name} .resource-property-${[property.name]} .right-ref`}
                                                 targetSelector={`.resource-${property.reference!.referencedResource} .resource-head`}/>
                                })}
                            </Fragment>
                        })}
                    </Movable>
                </Scale>
            </SvgContainer>
        </svg>
    </div>
}

