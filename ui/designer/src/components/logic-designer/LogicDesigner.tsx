// eslint disable
import React, { Fragment, type ReactNode, useState } from 'react'
import { type Resource, ResourceService, LayoutContext } from "@apibrew/ui-lib"
import { ResourceElement } from '../designer/ResourceElement'
import { Scale } from '../designer/Scale'
import { Movable, MovableComponent } from '../designer/Movable'
import { SvgContainer } from '../designer/SvgContainer'
import { ReferenceLink } from '../designer/ReferanceLink'
import { Selectable } from '../designer/Selectable'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import {
    Add,
    Delete,
    Edit,
    Search,
    SettingsApplications,
    ZoomIn,
    ZoomOut
} from '@mui/icons-material'
import { Checkbox, Dialog, DialogActions, DialogContent, DialogTitle, Menu, MenuItem, Tooltip } from '@mui/material'
import Button from '@mui/material/Button'
import { ResourceForm } from '../resource-form/ResourceForm'
import { Point } from '../designer/point'
import { Expression, LogicDesignerBoard, VarStatement } from '../../model/ui/logic-designer-board'
import { Function } from '../../model/logic/function'
import { ResourceRule } from '../../model/logic/resource-rule'
// import {SdkDrawer} from '../sdk/SdkDrawer'
// import { checkResourceAllowedOnBoard } from "./util";

export interface Selection {
    type: string
    identifier: string
    data: object
}

export interface LogicDesignerProps {
    id?: string
    name?: string
}

// React component to render the diagram
export const LogicDesigner: React.FC<LogicDesignerProps> = (props: LogicDesignerProps) => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)
    const [selected, setSelected] = React.useState<Selection[]>([])
    const [forceDelete, setForceDelete] = React.useState<boolean>(false)

    const [addButtonRef, setAddButtonRef] = React.useState<null | HTMLElement>(null)
    const [flags, setFlags] = React.useState<Record<string, boolean>>({})
    const modules: ReactNode[] = []
    const layoutOptions = React.useContext(LayoutContext)
    const [locationMap, setLocationMap] = React.useState<Record<string, Point>>({})
    const [showSdk, setShowSdk] = useState(false)
    const [sdkResource, setSdkResource] = useState<Resource>()

    console.log(resources, locationMap)

    const actionPanel = <Box style={{ display: 'flex' }}>
        <Box>
            <Tooltip title={'Add New Item'}>
                <IconButton onClick={(e) => {
                    if (!addButtonRef) {
                        setAddButtonRef(e.currentTarget)
                    }
                    setFlags({
                        ...flags,
                        addMenuOpen: true
                    })
                }}>
                    <Add />
                </IconButton>
            </Tooltip>
            {addButtonRef && <Menu anchorEl={addButtonRef}
                onClose={() => {
                    setFlags({
                        ...flags,
                        addMenuOpen: false
                    })
                }}
                open={flags.addMenuOpen}
                id="hooks-menu"
            >
                <MenuItem onClick={() => {
                    setFlags({
                        ...flags,
                        addMenuOpen: false
                    })

                    const modal = layoutOptions.showModal({
                        content: <Box sx={{
                            position: 'absolute',
                            top: '50%',
                            left: '50%',
                            transform: 'translate(-50%, -50%)',
                            width: 800
                        }}>
                            <ResourceForm resources={resources} initResource={{
                                name: '',
                                namespace: '',
                                properties: [],
                                version: 1,
                                virtual: false
                            }}
                                onCancel={() => {
                                    modal.close()
                                }}
                                onSave={() => {
                                    // load()
                                    modal.close()
                                }} />
                        </Box>
                    })
                }}>Add Resource</MenuItem>
            </Menu>}
            <Tooltip title={'Edit Item'}>
                <IconButton onClick={(e) => {
                    if (selected.length === 0) {
                        layoutOptions.showAlert({
                            severity: 'error',
                            message: 'Please select an item to edit'
                        })
                        return
                    }

                    if (selected.length > 1) {
                        layoutOptions.showAlert({
                            severity: 'error',
                            message: 'You need to select only one item to edit'
                        })
                        return
                    }

                    const modal = layoutOptions.showModal({
                        content: <Box sx={{
                            position: 'absolute' as 'absolute',
                            top: '50%',
                            left: '50%',
                            transform: 'translate(-50%, -50%)',
                            width: 800
                        }}>
                            <ResourceForm resources={resources} initResource={selected[0].data as Resource}
                                onCancel={() => {
                                    modal.close()
                                }}
                                onSave={(updatedResource) => {
                                    // load()
                                    modal.close()
                                }} />
                        </Box>
                    })
                }}>
                    <Edit textAnchor={'asd'} />
                </IconButton>
            </Tooltip>
            <Tooltip title={'Delete Item'}>
                <IconButton onClick={(e) => {
                    if (selected.length === 0) {
                        layoutOptions.showAlert({
                            severity: 'error',
                            message: 'Please select an item to delete'
                        })
                    }

                    setFlags({
                        ...flags,
                        deleteDialog: true
                    })
                }}>
                    <Delete />
                </IconButton>
            </Tooltip>
            <Dialog
                open={flags.deleteDialog}
                onClose={() => {
                    setFlags({
                        ...flags,
                        deleteDialog: false
                    })
                }}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    {'Are you sure, you want to delete this item?\n This operation is not reversible'}
                </DialogTitle>
                <DialogContent>
                    Force Delete? <Checkbox value={forceDelete} onChange={e => {
                        setForceDelete(e.target.checked)
                    }}></Checkbox>
                </DialogContent>
                <DialogActions>
                    <Button variant='contained' onClick={() => {
                        setFlags({
                            ...flags,
                            deleteDialog: false
                        })
                    }}>No</Button>
                    <Button variant='contained' onClick={() => {
                        setFlags({
                            ...flags,
                            deleteDialog: false
                        })

                        ResourceService.remove(selected[0].data as Resource, forceDelete).then(() => {
                            layoutOptions.showAlert({
                                severity: 'success',
                                message: 'Item deleted successfully'
                            })

                            // load()
                        }).catch(error => {
                            layoutOptions.showAlert({
                                severity: 'error',
                                message: error.message
                            })
                        })
                    }} autoFocus>
                        Yes
                    </Button>
                </DialogActions>
            </Dialog>
        </Box>
        <Box sx={{ flexGrow: 5 }} />
        <Box>
            <Tooltip title={`${Math.round(zoomLevel * 100)}%`}>
                <Box>
                    <IconButton onClick={() => {
                        setZoomLevel(Math.min(3, zoomLevel + 0.2))
                    }}>
                        <ZoomIn />
                    </IconButton>
                    <IconButton onClick={() => {
                        setZoomLevel(1)
                    }}>
                        <Search />
                    </IconButton>
                    <IconButton onClick={() => {
                        setZoomLevel(Math.max(0.2, zoomLevel - 0.2))
                    }}>
                        <ZoomOut />
                    </IconButton>
                </Box>
            </Tooltip>
        </Box>
        <Box>
            <Tooltip title={'Settings'}>
                <IconButton onClick={(e) => {

                }}>
                    <SettingsApplications />
                </IconButton>
            </Tooltip>
        </Box>
    </Box>

    const board: LogicDesignerBoard = {
        id: '',
        version: 1,
        functionDefs: [
            {
                function: {
                    name: 'countryNameCapitalCheck',
                    package: 'LogicDesignerBoard'
                } as Function,
                parameters: ['entity'],
                statements: [
                    {
                        type: 'var',
                        var: {
                            name: 'name',
                            value: {

                            } as Expression
                        } as VarStatement
                    }
                ]
            }
        ],
        resourceRules: [
            {
                name: 'countryNameCapitalCheck',
                conditionFunction: {
                    name: 'countryNameCapitalCheck',
                    package: 'LogicDesignerBoard'
                },
                resource: 'country',
                namespace: 'default'
            } as ResourceRule
        ],
    }

    const designingArea = <svg className={'designer-parent'}
        style={{ width: '100%', height: '100vh', overflow: 'auto' }}>
        <SvgContainer>
            <Scale level={zoomLevel}>
                <Movable>
                    {resources.map((resource, index) => {
                        return <MovableComponent key={resource.name}
                            location={locationMap[resource.name]}
                            updateLocation={location => {
                                // if (location.x === locationMap[resource.name].x && location.y === locationMap[resource.name].y) {
                                //     return
                                // }

                                // setLocationMap({
                                //     ...locationMap,
                                //     [resource.name]: location
                                // })

                                // const resourceVisual = board?.resourceVisuals?.find(item => item.resource === resource.name)

                                // if (resourceVisual) {
                                //     resourceVisual.location = location
                                // } else {
                                //     board?.resourceVisuals?.push({
                                //         resource: resource.name,
                                //         allowRecordsOnBoard: false,
                                //         location
                                //     })
                                // }

                                // save()
                            }}>
                            <Selectable onSelected={isSelected => {
                                if (isSelected) {
                                    setSelected([...selected, {
                                        type: 'resource',
                                        identifier: resource.name,
                                        data: resource
                                    }])
                                } else {
                                    setSelected(selected.filter(item => item.type === 'resource' && item.identifier !== resource.name))
                                }
                            }}>
                                <ResourceElement resource={resource} />
                            </Selectable>
                        </MovableComponent>
                    })}

                    {resources.map((resource, index) => {
                        return <Fragment key={resource.name}>
                            {resource.properties?.filter(item => item.type === 'REFERENCE')?.filter(item => item.reference?.referencedResource)?.map((property, index) => {
                                return <ReferenceLink key={`${resource.name}-${property.name}`}
                                    resource={resource}
                                    property={property} />
                            })}
                        </Fragment>
                    })}
                </Movable>
            </Scale>
        </SvgContainer>
    </svg>

    return <Box sx={{ display: 'flex', flexDirection: 'column', width: '100%' }}>
        {/* Action Panel */}
        {actionPanel}
        {/* Designing Area */}
        <Box sx={{ flexGrow: 1 }}>
            {board && designingArea}
        </Box>
        {/*<SdkDrawer resource={sdkResource} open={showSdk} onClose={() => {*/}
        {/*    setShowSdk(false)*/}
        {/*}}/>*/}
        {modules}
    </Box>
}
