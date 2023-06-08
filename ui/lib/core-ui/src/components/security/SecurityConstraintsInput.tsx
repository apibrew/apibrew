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

export interface SecurityConstraintsInputProps {
    mode: 'role' | 'resource' | 'namespace'
}

export function SecurityConstraintsInput(props: SecurityConstraintsInputProps) {
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
        <Table>
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
                    <TableCell>Username</TableCell>
                    <TableCell>Role</TableCell>
                    <TableCell>Require</TableCell>
                    <TableCell>Permit</TableCell>
                    <TableCell style={{width: '150px'}}>Actions</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {constraints.map((constraint, index) => <TableRow key={index}>
                    <TableCell>
                        <TextField disabled={props.mode === 'namespace' || props.mode == 'resource'} size='small'
                                   variant='outlined' value={constraint.namespace} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].namespace = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <TextField disabled={props.mode === 'resource'} size='small' variant='outlined'
                                   value={constraint.resource} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].resource = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <TextField size='small' variant='outlined'
                                   value={constraint.property} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].property = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <Select size='small' variant='outlined' value={constraint.operation} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].operation = e.target.value as string as Operation
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}>
                            <MenuItem value='FULL'>FULL</MenuItem>
                            <MenuItem value='OPERATION_TYPE_READ'>OPERATION_TYPE_READ</MenuItem>
                            <MenuItem value='OPERATION_TYPE_UPDATE'>OPERATION_TYPE_UPDATE</MenuItem>
                            <MenuItem value='OPERATION_TYPE_CREATE'>OPERATION_TYPE_CREATE</MenuItem>
                            <MenuItem value='OPERATION_TYPE_DELETE'>OPERATION_TYPE_DELETE</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell>
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
                    <TableCell>
                        <TextField disabled={props.mode === 'role'} size='small' variant='outlined'
                                   value={constraint.username} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].username = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <TextField disabled={props.mode === 'role'} size='small' variant='outlined'
                                   value={constraint.role} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].role = e.target.value
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <Checkbox size='small'
                                  value={constraint.requirePass} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].requirePass = e.target.checked
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}/>
                    </TableCell>
                    <TableCell>
                        <Select size='small' variant='outlined' value={constraint.permit} onChange={e => {
                            const updatedConstraints = [...constraints]
                            updatedConstraints[index].permit = e.target.value as string as Permit
                            setConstraints(updatedConstraints)
                            valueContext.onChange(updatedConstraints)
                        }}>
                            <MenuItem value='PERMIT_TYPE_ALLOW'>PERMIT_TYPE_ALLOW</MenuItem>
                            <MenuItem value='PERMIT_TYPE_REJECT'>PERMIT_TYPE_REJECT</MenuItem>
                        </Select>
                    </TableCell>
                    <TableCell>
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