import React, {Fragment, ReactNode, useEffect} from 'react'
import {type Resource} from '../../model'
import {ResourceService} from '../../service/resource'
import {ResourceElement} from './ResourceElement'
import {Scale} from './Scale'
import {Movable, MovableComponent} from './Movable'
import {SvgContainer} from './SvgContainer'
import {ReferenceLink} from './ReferanceLink'
import {Selectable} from './Selectable'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import {
    Add,
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
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    Menu,
    MenuItem,
    Tooltip
} from "@mui/material";
import {LayoutContext} from "../../context/layout-context";
import Button from "@mui/material/Button";

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
    const [viewMode, setViewMode] = React.useState<'wide' | 'compact'>('wide')

    const [addButtonRef, setAddButtonRef] = React.useState<null | HTMLElement>(null);
    const [flags, setFlags] = React.useState<{
        [key: string]: boolean
    }>({});
    const modules: ReactNode[] = []
    const layoutOptions = React.useContext(LayoutContext)

    useEffect(() => {
        ResourceService.list().then(list => {
            setResources(list.filter(item => item.namespace !== 'system'))
        }, error => {
            console.error(error)
        })
    }, [])

    const handleAdd = () => {

    }

    const handleDelete = () => {
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
    }

    const handleEdit = () => {
        if (selected.length == 0) {
            layoutOptions.showAlert({
                severity: 'error',
                message: 'Please select an item to edit'
            })
        }
    }

    const actionPanel = <Box style={{display: 'flex'}}>
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
                    <Add/>
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
                }}>Add Resource</MenuItem>
            </Menu>}
            <Tooltip title={'Edit Item'}>
                <IconButton onClick={(e) => {
                    handleEdit()
                }}>
                    <Edit textAnchor={'asd'}/>
                </IconButton>
            </Tooltip>
            <Tooltip title={'Delete Item'}>
                <IconButton onClick={(e) => {
                    handleDelete()
                }}>
                    <Delete/>
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
                            {"Use Google's location service?"}
                        </DialogTitle>
                        <DialogContent>
                            <DialogContentText id="alert-dialog-description">
                                Let Google help apps determine location. This means sending anonymous
                                location data to Google, even when no apps are running.
                            </DialogContentText>
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={() => {
                                setFlags({
                                    ...flags,
                                    deleteDialog: false,
                                })
                            }}>Disagree</Button>
                            <Button onClick={() => {
                                setFlags({
                                    ...flags,
                                    deleteDialog: false,
                                })
                            }} autoFocus>
                                Agree
                            </Button>
                        </DialogActions>
                    </Dialog>
                </IconButton>
            </Tooltip>
            <Tooltip title={'Load Existing'}>
                <IconButton onClick={(e) => {

                }}>
                    <GetApp/>
                </IconButton>
            </Tooltip>
        </Box>
        <Box sx={{flexGrow: 5}}/>
        <Box sx={{flexGrow: 1}}/><Box>
        <IconButton value="wide" aria-label="left aligned">
            <Tooltip title={'Wide elements'}>
                <WidthWide/>
            </Tooltip>
        </IconButton>
        <IconButton value="compact" aria-label="left aligned">
            <Tooltip title={'Compact'}>
                <ViewCompact/>
            </Tooltip>
        </IconButton>
    </Box>
        <Box sx={{flexGrow: 1}}/>
        <Box>
            <Tooltip title={`${Math.round(zoomLevel * 100)}%`}>
                <Box>
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
                </Box>
            </Tooltip>
        </Box>
        <Box sx={{flexGrow: 1}}/>
        <Box>
            <IconButton value="wide" aria-label="left aligned">
                <Tooltip title={'Rearrange elements'}>
                    <FormatAlignCenter/>
                </Tooltip>
            </IconButton>
            <Tooltip title={'Reload'}>
                <IconButton onClick={(e) => {

                }}>
                    <Replay/>
                </IconButton>
            </Tooltip>
            <Tooltip title={'Undo'}>
                <IconButton onClick={(e) => {

                }}>
                    <Undo/>
                </IconButton>
            </Tooltip>
            <Tooltip title={'Redo'}>
                <IconButton onClick={(e) => {

                }}>
                    <Redo/>
                </IconButton>
            </Tooltip>
            <Tooltip title={'Save'}>
                <IconButton onClick={(e) => {

                }}>
                    <Save/>
                </IconButton>
            </Tooltip>
        </Box>
        <Box sx={{flexGrow: 1}}/>
        <Box>
            <Tooltip title={'Settings'}>
                <IconButton onClick={(e) => {

                }}>
                    <SettingsApplications/>
                </IconButton>
            </Tooltip>
        </Box>
    </Box>

    const designingArea = <svg className={'designer-parent'}
                               style={{width: '100%', height: '90vh'}}>
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

    return <Box>
        {/* Action Panel */}
        {actionPanel}
        {/* Designing Area */}
        {designingArea}
        {modules}
    </Box>
}
