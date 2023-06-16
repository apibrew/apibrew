import React, {useState} from "react";
import Box from "@mui/material/Box";
import {Checkbox, MenuItem, Select, Table, TableBody, TableCell, TableHead, TableRow, TextField} from "@mui/material";
import {Operation, Permit, PropertyMode, SecurityConstraint} from "../../model/security-constraint.ts";
import IconButton from "@mui/material/IconButton";
import {Add, Delete, Edit, Save} from "@mui/icons-material";

export interface SecurityConstraintsInputAdvancedProps {
    mode: 'role' | 'resource' | 'namespace'
    constraints: SecurityConstraint[]
    setConstraints: (constraints: SecurityConstraint[]) => void
}

export function SecurityConstraintsInputAdvanced(props: SecurityConstraintsInputAdvancedProps) {
    return <Box>
        <IconButton onClick={() => {
            props.setConstraints([...props.constraints, {
                namespace: 'namespace-1',
                resource: 'resource-1',
                property: '*',
                operation: 'FULL',
            } as SecurityConstraint])
        }}>
            <Add/>
        </IconButton>
        <Table size='small'>
            <TableHead>
                <TableRow>
                    <TableCell>Namespace</TableCell>
                    <TableCell>Resource</TableCell>
                    <TableCell>Property</TableCell>
                    <TableCell>PropertyValue</TableCell>
                    <TableCell>PropertyMode</TableCell>
                    <TableCell>Operation</TableCell>
                    <TableCell>Record(s)</TableCell>
                    {/*<TableCell>Before</TableCell>*/}
                    {/*<TableCell>After</TableCell>*/}
                    {props.mode === 'resource' && <TableCell>Username</TableCell>}
                    {props.mode === 'resource' && <TableCell>Role</TableCell>}
                    <TableCell style={{width: '50px'}}>Permit</TableCell>
                    <TableCell style={{width: '50px'}}>Actions</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {props.constraints.map((constraint, index) => <TableRow key={index}>
                    <TableCell sx={{padding: 1}}>
                        <TextField sx={{margin: 0}} disabled={props.mode === 'namespace' || props.mode == 'resource'} size='small'
                                   variant='outlined' value={constraint.namespace} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].namespace = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}} >
                        <TextField disabled={props.mode === 'resource'} size='small' variant='outlined'
                                   value={constraint.resource} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].resource = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.property} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].property = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.propertyValue} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].propertyValue = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <Select sx={{width: '100%'}} size='small' variant='outlined' value={constraint.propertyMode} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].propertyMode = e.target.value as string as PropertyMode
                            props.setConstraints(updatedConstraints)
                        }}>
                            <MenuItem value='PROPERTY_MATCH_ONLY'>Only</MenuItem>
                            <MenuItem value='PROPERTY_MATCH_ANY'>Any</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <Select sx={{width: '100%'}} size='small' variant='outlined' value={constraint.operation} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].operation = e.target.value as string as Operation
                            props.setConstraints(updatedConstraints)
                        }}>
                            <MenuItem value='FULL'>full</MenuItem>
                            <MenuItem value='OPERATION_TYPE_READ'>read</MenuItem>
                            <MenuItem value='OPERATION_TYPE_UPDATE'>update</MenuItem>
                            <MenuItem value='OPERATION_TYPE_CREATE'>create</MenuItem>
                            <MenuItem value='OPERATION_TYPE_DELETE'>delete</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <TextField size='small'
                                   variant='outlined'
                                   value={constraint.recordIds.join(',')}
                                   onChange={e => {
                                       const updatedConstraints = [...props.constraints]
                                       updatedConstraints[index].recordIds = e.target.value.split(',')
                                       props.setConstraints(updatedConstraints)
                                   }}/>
                    </TableCell>
                    {props.mode === 'resource' && <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.username} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].username = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>}
                    {props.mode === 'resource' && <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.role} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].role = e.target.value
                            props.setConstraints(updatedConstraints)
                        }}/>
                    </TableCell>}
                    <TableCell sx={{padding: 1}}>
                        <Select sx={{width: '100%'}} size='small' variant='outlined' value={constraint.permit} onChange={e => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints[index].permit = e.target.value as string as Permit
                            props.setConstraints(updatedConstraints)
                        }}>
                            <MenuItem value='PERMIT_TYPE_ALLOW'>allow</MenuItem>
                            <MenuItem value='PERMIT_TYPE_REJECT'>reject</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <IconButton onClick={() => {
                            const updatedConstraints = [...props.constraints]
                            updatedConstraints.splice(index, 1)
                            props.setConstraints(updatedConstraints)
                        }}>
                            <Delete/>
                        </IconButton>
                    </TableCell>
                </TableRow>)}
            </TableBody>
        </Table>
    </Box>
}