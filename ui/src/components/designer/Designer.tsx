import React, { Fragment, useEffect } from 'react'
import { type Resource } from '../../model'
import { ResourceService } from '../../service/resource'
import { ResourceElement } from './ResourceElement'
import { Scale } from './Scale'
import { Movable, MovableComponent } from './Movable'
import { SvgContainer } from './SvgContainer'
import { ReferenceLink } from './ReferanceLink'
import { Selectable } from './Selectable'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import { Search, ZoomIn, ZoomOut } from '@mui/icons-material'

export interface Selection {
    type: string
    identifier: string
    data: object
}

// React component to render the diagram
export const Designer: React.FC = () => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)
    const [selected, setSelected] = React.useState<Selection[]>([])

    useEffect(() => {
        ResourceService.list().then(list => {
            setResources(list.filter(item => item.namespace !== 'system'))
        }, error => {
            console.error(error)
        })
    }, [])

    console.log('selected', selected)

    return <div>
        {/* Action Panel */}
        <Box style={{ display: 'flex' }}>
            <div style={{ flexGrow: 1 }}/>
            <IconButton onClick={() => {
                setZoomLevel(Math.min(3, zoomLevel + 0.2))
            }}>
                <ZoomIn/>
            </IconButton>
            <IconButton onClick={() => {
                setZoomLevel(1)
            }}>
                <Search/>
            </IconButton>
            <IconButton onClick={() => {
                setZoomLevel(Math.max(0.2, zoomLevel - 0.2))
            }}>
                <ZoomOut/>
            </IconButton>
            {/* {Math.round(zoomLevel * 100)}% &nbsp; */}
        </Box>
        {/* Designing Area */}
        <svg className={'designer-parent'}
            style={{ width: '100%', height: '600px' }}>
            <SvgContainer>
                <Scale level={zoomLevel}>
                    <Movable>
                        {resources.map((resource, index) => {
                            const x = 20 + 410 * index
                            const y = 20
                            return <g key={`${(resource.namespace ?? '')}-${resource.name ?? ''}`}
                                transform={`translate(${x}, ${y})`}>
                                <MovableComponent>
                                    <Selectable onSelected={isSelected => {
                                        if (isSelected) {
                                            setSelected([...selected, {
                                                type: 'resource',
                                                identifier: resource.name ?? '',
                                                data: resource
                                            }])
                                        } else {
                                            setSelected(selected.filter(item => item.type === 'resource' && item.identifier !== resource.name))
                                        }
                                    }}>
                                        <ResourceElement resource={resource}/>
                                    </Selectable>
                                </MovableComponent>
                            </g>
                        })}

                        {resources.map((resource, index) => {
                            return <Fragment key={resource.name ?? ''}>
                                {resource.properties?.filter(item => item.type === 'REFERENCE')?.filter(item => item.reference?.referencedResource)?.map((property, index) => {
                                    return <ReferenceLink key={`${resource.name ?? ''}-${property.name ?? ''}`}
                                        resource={resource}
                                        property={property}/>
                                })}
                            </Fragment>
                        })}
                    </Movable>
                </Scale>
            </SvgContainer>
        </svg>
    </div>
}
