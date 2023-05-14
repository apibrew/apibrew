import React, { Fragment, ReactNode, useEffect } from 'react'
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
import {
    Add,
    Brush,
    Delete,
    Edit,
    FormatAlignCenter,
    GetApp,
    Redo,
    Replay,
    Save,
    Search,
    SettingsApplications,
    Undo,
    ViewCompact,
    WidthWide,
    ZoomIn,
    ZoomOut
} from '@mui/icons-material'
import {
    Alert,
    Checkbox,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    Menu,
    MenuItem,
    Tooltip
} from "@mui/material";
import { LayoutContext } from "../../context/layout-context";
import Button from "@mui/material/Button";
import { ResourceForm } from '../resource-form/ResourceForm'
import { ResourceVisualizer } from './ResourceVisualizer'

export interface Selection {
    type: string
    identifier: string
    data: object
}

type ViewMode = 'wide' | 'compact'

// React component to render the diagram
export const Designer: React.FC = () => {
    const [resources, setResources] = React.useState<Resource[]>([])
    const [zoomLevel, setZoomLevel] = React.useState<number>(1)
    const [selected, setSelected] = React.useState<Selection[]>([])
    const [viewMode, setViewMode] = React.useState<ViewMode>('wide')
    const [forceDelete, setForceDelete] = React.useState<boolean>(false)

    const [addButtonRef, setAddButtonRef] = React.useState<null | HTMLElement>(null);
    const [flags, setFlags] = React.useState<{
        [key: string]: boolean
    }>({});
    const modules: ReactNode[] = []
    const layoutOptions = React.useContext(LayoutContext)

    const load = async () => {
        try {
            const list = await ResourceService.list()
            setResources(list.filter(item => item.namespace !== 'system'))
        } catch (error) {
            console.error(error)
        }
    }

    useEffect(() => {
        load()
    }, [])

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
                            position: 'absolute' as 'absolute',
                            top: '50%',
                            left: '50%',
                            transform: 'translate(-50%, -50%)',
                            width: 800,
                        }}>
                            <ResourceForm resources={resources} initResource={{
                                name: '',
                                namespace: '',
                                properties: [],
                                version: 1,
                                virtual: false,
                            }}
                                onCancel={() => {
                                    modal.close()
                                }}
                                onSave={() => {
                                    load()
                                    modal.close()
                                }} />
                        </Box>
                    })
                }}>Add Resource</MenuItem>
            </Menu>}
            <Tooltip title={'Edit Item'}>
                <IconButton onClick={(e) => {
                    if (selected.length == 0) {
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
                            width: 800,
                        }}>
                            <ResourceForm resources={resources} initResource={selected[0].data as Resource}
                                onCancel={() => {
                                    modal.close()
                                }}
                                onSave={(updatedResource) => {
                                    load()
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
                    if (selected.length == 0) {
                        layoutOptions.showAlert({
                            severity: 'error',
                            message: 'Please select an item to delete'
                        })
                    }

                    setFlags({
                        ...flags,
                        deleteDialog: true,
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
                        deleteDialog: false,
                    })
                }}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    {"Are you sure, you want to delete this item?\n This operation is not reversible"}
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
                            deleteDialog: false,
                        })
                    }}>No</Button>
                    <Button variant='contained' onClick={() => {
                        setFlags({
                            ...flags,
                            deleteDialog: false,
                        })

                        ResourceService.remove(selected[0].data as Resource, forceDelete).then(() => {
                            layoutOptions.showAlert({
                                severity: 'success',
                                message: 'Item deleted successfully'
                            })

                            load()
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
            <Tooltip title={'Update visualisation'}>
                <IconButton onClick={(e) => {
                    layoutOptions.showModal({
                        content: <ResourceVisualizer />
                    })
                }}>
                    <Brush />
                </IconButton>
            </Tooltip>
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
            <IconButton aria-label="left aligned" onClick={() => load()}>
                <Tooltip title={'Reload'}>
                    <Replay />
                </Tooltip>
            </IconButton>
            <IconButton value="wide" aria-label="left aligned">
                <Tooltip title={'Rearrange elements'}>
                    <FormatAlignCenter />
                </Tooltip>
            </IconButton>
            <Tooltip title={'Settings'}>
                <IconButton onClick={(e) => {

                }}>
                    <SettingsApplications />
                </IconButton>
            </Tooltip>
        </Box>
    </Box>

    const designingArea = <svg className={'designer-parent'}
        style={{ width: '100%', height: '90vh' }}>
        <SvgContainer>
            <Scale level={zoomLevel}>
                <Movable>
                    {resources.map((resource, index) => {
                        const x = 20 + 410 * index
                        const y = 20
                        return <g key={`${(resource.namespace ?? '')}-${resource.name}`}
                            transform={`translate(${x}, ${y})`}>
                            <MovableComponent>
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
                        </g>
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

    return <Box>
        {/* Action Panel */}
        {actionPanel}
        {/* Designing Area */}
        {designingArea}
        {modules}
    </Box>
}
