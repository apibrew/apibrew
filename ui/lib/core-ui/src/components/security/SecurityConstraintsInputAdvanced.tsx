import React, {useState} from "react";
import Box from "@mui/material/Box";
import {Checkbox, MenuItem, Select, Table, TableBody, TableCell, TableHead, TableRow, TextField} from "@mui/material";
import {Operation, Permit, SecurityConstraint} from "../../model/security-constraint.ts";
import {useRecord} from "../../context/record.ts";
import {useResourceByName} from "../../hooks/resource.ts";
import {useValue} from "../../context/value.ts";
import {FormItem as CrudFormItem} from "../../model/ui/crud.ts";
import IconButton from "@mui/material/IconButton";
import {Add, Delete, Edit, Save} from "@mui/icons-material";

export interface SecurityConstraintsInputAdvancedProps {
    mode: 'role' | 'resource' | 'namespace'
}

export function SecurityConstraintsInputAdvanced(props: SecurityConstraintsInputAdvancedProps) {
    const valueContext = useValue()
    const record = useRecord<{ id: string, name: string, namespace: string }>()

    const [constraints, setConstraints] = useState<SecurityConstraint[]>(valueContext.value ?? [])

    constraints.forEach((constraint) => {
        switch (props.mode) {
            case 'role':
                constraint.role = record.name
                break
            case 'resource':
                constraint.resource = record.name
                constraint.namespace = record.namespace
                break
            case 'namespace':
                constraint.namespace = record.name
                break
        }

        if (!constraint.operation) {
            constraint.operation = 'FULL'
        }

        if (!constraint.permit) {
            constraint.permit = 'PERMIT_TYPE_ALLOW'
        }

        if (!constraint.namespace) {
            constraint.namespace = '*'
        }

        if (!constraint.resource) {
            constraint.resource = '*'
        }

        if (!constraint.username) {
            constraint.username = '*'
        }

        if (!constraint.role) {
            constraint.role = '*'
        }

        if (!constraint.recordIds) {
            constraint.recordIds = []
        }
    })

    return <Box>
        <IconButton onClick={() => {
            setConstraints([...constraints, {} as SecurityConstraint])
        }}>
            <Add/>
        </IconButton>
        <Table size='small'>
            <TableHead>
                <TableRow>
                    <TableCell>Namespace</TableCell>
                    <TableCell>Resource</TableCell>
                    <TableCell>Property</TableCell>
                    {/*<TableCell>Value</TableCell>*/}
                    <TableCell>Operation</TableCell>
                    <TableCell>Record(s)</TableCell>
                    {/*<TableCell>Before</TableCell>*/}
                    {/*<TableCell>After</TableCell>*/}
                    {props.mode === 'resource' && <TableCell>Username</TableCell>}
                    {props.mode === 'resource' && <TableCell>Role</TableCell>}
                    {props.mode === 'resource' && <TableCell style={{width: '50px'}}>Require</TableCell>}
                    <TableCell style={{width: '50px'}}>Permit</TableCell>
                    <TableCell style={{width: '50px'}}>Actions</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {constraints.map((constraint, index) => <TableRow key={index}>
                    <TableCell sx={{padding: 1}}>
                        <TextField sx={{margin: 0}} disabled={props.mode === 'namespace' || props.mode == 'resource'} size='small'
                                   variant='outlined' value={constraint.namespace} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].namespace = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}} >
                        <TextField disabled={props.mode === 'resource'} size='small' variant='outlined'
                                   value={constraint.resource} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].resource = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.property} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].property = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <Select sx={{width: '100%'}} size='small' variant='outlined' value={constraint.operation} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].operation = e.target.value as string as Operation
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
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
                                       const updatedConstraints = [...constraints]
                                       updatedConstraints[index].recordIds = e.target.value.split(',')
                                       setConstraints(updatedConstraints)
                                       valueContext.onChange(updatedConstraints)
                                   }}/>
                    </TableCell>
                    {props.mode === 'resource' && <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.username} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].username = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>}
                    {props.mode === 'resource' && <TableCell sx={{padding: 1}}>
                        <TextField size='small' variant='outlined'
                                   value={constraint.role} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].role = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>}
                    {props.mode === 'resource' && <TableCell sx={{padding: 1}}>
                        <Checkbox size='small'
                                  value={constraint.requirePass} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].requirePass = e.target.checked
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>}
                    <TableCell sx={{padding: 1}}>
                        <Select sx={{width: '100%'}} size='small' variant='outlined' value={constraint.permit} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].permit = e.target.value as string as Permit
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}>
                            <MenuItem value='PERMIT_TYPE_ALLOW'>allow</MenuItem>
                            <MenuItem value='PERMIT_TYPE_REJECT'>reject</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell sx={{padding: 1}}>
                        <IconButton onClick={() => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints.splice(index, 1)
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}>
                            <Delete/>
                        </IconButton>
                    </TableCell>
                </TableRow>)}
            </TableBody>
        </Table>
    </Box>
}