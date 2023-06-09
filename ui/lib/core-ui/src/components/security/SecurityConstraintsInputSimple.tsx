import React, {useEffect, useState} from "react";
import {ChevronRight, ExpandMore} from "@mui/icons-material";
import {Checkbox, Table, TableBody, TableCell, TableHead, TableRow, TextField} from "@mui/material";
import {useRecords} from "../../hooks/record.ts";
import {Namespace, Resource} from "../../model";
import Box from "@mui/material/Box";
import {ResourceService} from "../../service";
import {useErrorHandler} from "../../hooks/error-handler.tsx";
import IconButton from "@mui/material/IconButton";
import {prepareAccessMap} from "./access-map.computer.ts";
import {SecurityConstraint} from "../../model/security-constraint.ts";

export interface SecurityConstraintsInputSimpleProps {
    mode: 'role' | 'resource' | 'namespace'
    constraints: SecurityConstraint[]
    setConstraints: (constraints: SecurityConstraint[]) => void
}

export interface PermissionChecks {
    full: boolean
    read: boolean
    create: boolean
    update: boolean
    delete: boolean
}

export function SecurityConstraintsInputSimple(props: SecurityConstraintsInputSimpleProps) {
    const namespaces = useRecords<Namespace>('namespace', 'system')
    const [resources, setResources] = useState<Resource[]>([])
    const errorHandler = useErrorHandler()

    useEffect(() => {
        ResourceService.list().then(setResources, errorHandler)
    }, [])

    const [open, setOpen] = useState<{
        [k: string]: boolean
    }>({
        'namespace-default': true
    })

    const [accessMap, setAccessMap] = useState<{
        [k: string]: PermissionChecks
    }>({})

    useEffect(() => {
        let updatedAccessMap = prepareAccessMap(accessMap, namespaces, resources, props.constraints);

        setAccessMap(updatedAccessMap)
    }, [
        namespaces, resources
    ])

    return <>
        <Table size={'small'}>
            <TableHead>
                <TableRow>
                    <TableCell width={'300px'}><b>Name</b></TableCell>
                    <TableCell><b>Full</b></TableCell>
                    <TableCell><b>Read</b></TableCell>
                    <TableCell><b>Create</b></TableCell>
                    <TableCell><b>Update</b></TableCell>
                    <TableCell><b>Delete</b></TableCell>
                    <TableCell><b>Allow only record(s)</b></TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                <TableRow>
                    <TableCell>
                        <b>All</b>
                    </TableCell>
                    {accessMap[`system`] && <PermissionCheckBoxGroup value={accessMap[`system`]}
                                                                     indeterminate={anyOf(namespaces, item => {
                                                                         return accessMap[`namespace-${item.name}`]
                                                                     })}
                                                                     onChange={value => {
                                                                         setAccessMap({
                                                                             ...accessMap,
                                                                             [`system`]: value
                                                                         })
                                                                     }}/>}
                    <TableCell/>
                </TableRow>
                {namespaces.map(namespace => <React.Fragment key={namespace.name}>
                    <TableRow>
                        <TableCell>
                            <span>Namespace: <b>{namespace.name}</b></span>
                            <IconButton onClick={() => setOpen({
                                ...open,
                                [`namespace-${namespace.name}`]: !open[`namespace-${namespace.name}`]
                            })}>
                                {open[`namespace-${namespace.name}`] ? <ExpandMore/> : <ChevronRight/>}
                            </IconButton>
                        </TableCell>
                        {accessMap[`namespace-${namespace.name}`] &&
                            <PermissionCheckBoxGroup
                                value={combine(accessMap[`system`], accessMap[`namespace-${namespace.name}`])}
                                indeterminate={anyOf(resources.filter(item => item.namespace === namespace.name), resource => {
                                    return combine(accessMap[`resource-${resource.namespace}/${resource.name}`], anyOf(resource.properties, property => {
                                        return accessMap[`resource-${resource.namespace}/${resource.name}-${property.name}`]
                                    }))
                                })}
                                onChange={value => {
                                    setAccessMap({
                                        ...accessMap,
                                        [`namespace-${namespace.name}`]: value
                                    })
                                }}/>}
                        <TableCell/>
                    </TableRow>
                    {open[`namespace-${namespace.name}`] && resources.filter(item => item.namespace === namespace.name)
                        .map(resource => <React.Fragment key={resource.name}>
                            <TableRow>
                                <TableCell>
                                    <Box sx={{marginLeft: '30px'}}>
                                        <span>Resource: <b>{resource.name}</b></span>
                                        <IconButton onClick={() => setOpen({
                                            ...open,
                                            [`resource-${resource.namespace}/${resource.name}`]: !open[`resource-${resource.namespace}/${resource.name}`]
                                        })}>
                                            {open[`resource-${resource.namespace}/${resource.name}`] ? <ExpandMore/> :
                                                <ChevronRight/>}
                                        </IconButton>
                                    </Box>
                                </TableCell>
                                {accessMap[`resource-${resource.namespace}/${resource.name}`] &&
                                    <PermissionCheckBoxGroup
                                        value={combine(accessMap[`system`], accessMap[`namespace-${namespace.name}`], accessMap[`resource-${resource.namespace}/${resource.name}`])}
                                        indeterminate={anyOf(resource.properties, property => {
                                            return accessMap[`resource-${resource.namespace}/${resource.name}-${property.name}`]
                                        })}
                                        onChange={value => {
                                            setAccessMap({
                                                ...accessMap,
                                                [`resource-${resource.namespace}/${resource.name}`]: isolate(value, combine(accessMap[`system`], accessMap[`namespace-${namespace.name}`]), accessMap[`resource-${resource.namespace}/${resource.name}`])
                                            })
                                        }}/>}
                                <TableCell>
                                    <TextField
                                        size={'small'}
                                        fullWidth
                                        variant={'outlined'}
                                    />
                                </TableCell>
                            </TableRow>

                            {open[`resource-${resource.namespace}/${resource.name}`] && resource.properties.map(property =>
                                <TableRow key={property.name}>
                                    <TableCell>
                                        <Box sx={{marginLeft: '70px'}}>
                                            <span>Property: <b>{property.name}</b></span>
                                        </Box>
                                    </TableCell>
                                    {accessMap[`resource-${resource.namespace}/${resource.name}-${property.name}`] &&
                                        <PermissionCheckBoxGroup
                                            value={combine(accessMap[`system`], accessMap[`namespace-${namespace.name}`], accessMap[`resource-${resource.namespace}/${resource.name}`], accessMap[`resource-${resource.namespace}/${resource.name}-${property.name}`])}
                                            onChange={value => {
                                                setAccessMap({
                                                    ...accessMap,
                                                    [`resource-${resource.namespace}/${resource.name}-${property.name}`]: isolate(value, combine(accessMap[`system`], accessMap[`namespace-${namespace.name}`], accessMap[`resource-${resource.namespace}/${resource.name}`]), accessMap[`resource-${resource.namespace}/${resource.name}`]),
                                                })
                                            }}/>}
                                    <TableCell/>
                                </TableRow>)}
                        </React.Fragment>)}
                </React.Fragment>)}
            </TableBody>
        </Table>
    </>
}

export interface PermissionCheckBoxGroupProps {
    value: PermissionChecks
    indeterminate?: PermissionChecks
    onChange: (value: PermissionChecks) => void
}

export function PermissionCheckBoxGroup(props: PermissionCheckBoxGroupProps) {
    const controls = ['full', 'read', 'create', 'update', 'delete']

    return <>
        {controls.map(item => {
            return <TableCell key={item}>
                <Checkbox checked={props.value['full'] || props.value[item]}
                          indeterminate={!(props.value['full'] || props.value[item]) && props?.indeterminate && (props.indeterminate['full'] || props.indeterminate[item])}
                          onChange={e => {
                              if (item != 'full' && props.value['full']) {
                                  return
                              }

                              props.onChange({
                                  ...props.value,
                                  [item]: e.target.checked
                              })
                          }}/>
            </TableCell>
        })}
    </>
}

function combine(...permissions: PermissionChecks[]): PermissionChecks {
    permissions = permissions.filter(item => item)
    return {
        full: permissions.some(item => item.full),
        read: permissions.some(item => item.read),
        create: permissions.some(item => item.create),
        update: permissions.some(item => item.update),
        delete: permissions.some(item => item.delete),
    }
}

function isolate(value: PermissionChecks, combiner: PermissionChecks, actual: PermissionChecks): PermissionChecks {
    return {
        full: value.full || combiner.full && actual.full,
        read: value.read || combiner.read && actual.read,
        create: value.create || combiner.create && actual.create,
        update: value.update || combiner.update && actual.update,
        delete: value.delete || combiner.delete && actual.delete,
    }
}

function anyOf<T>(resources: T[], param2: (T) => PermissionChecks): PermissionChecks {
    return combine(...resources.map(param2))
}