import { Box, FormControl, FormGroup, FormHelperText, FormLabel, IconButton, MenuItem, Select, Table, TableBody, TableCell, TableContainer, TableFooter, TableHead, TableRow, TextField, Toolbar, Typography } from "@mui/material";
import { Resource, ResourceProperty } from "../../model";
import { Add, Delete, Edit, Save, TableRows } from "@mui/icons-material";
import React from "react";

export interface ResourceBasicFormProps {
    resource: Resource;
    onChange: (resource: Resource) => void;
}

const handlePropertyFieldOnChange = <K extends keyof ResourceProperty>(resource: Resource, onChange: (resource: Resource) => void, index: number, field: K, value: ResourceProperty[K]) => {
    onChange({
        ...resource,
        properties: resource.properties?.map((p, i) => {
            if (i === index) {
                return {
                    ...p,
                    [field]: value
                }
            }
            return p
        })
    })
}


export function ResourceBasicForm(props: ResourceBasicFormProps): JSX.Element {
    const [propertyFlags, setPropertyFlags] = React.useState<{
        [key: string]: boolean
    }>({})

    return <>
        <Box>
            <FormGroup>
                <FormControl>
                    <FormLabel>Resource Name</FormLabel>
                    <TextField value={props.resource.name}
                        onChange={(e) => props.onChange({ ...props.resource, name: e.target.value })} />
                    <FormHelperText>Resource name is required</FormHelperText>
                </FormControl>
                <Box m={1}>
                    <IconButton onClick={() => {
                        props.onChange({
                            ...props.resource,
                            properties: [...props.resource.properties ?? [], {
                                name: 'prop-' + (props.resource.properties?.length ?? 0 + 1),
                                type: 'STRING',
                                required: false,
                                primary: false,
                                unique: false
                            }]
                        })
                    }}>
                        <Add />
                    </IconButton>
                </Box>
                <TableContainer>
                    <Table size='small' >
                        <TableHead>
                            <TableRow>
                                <TableCell>Property Name</TableCell>
                                <TableCell>Type</TableCell>
                                <TableCell>Required</TableCell>
                                <TableCell>Primary</TableCell>
                                <TableCell>Unique</TableCell>
                                <TableCell>Actions</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {props.resource.properties?.map((property, index) => {
                                return <TableRow key={index}>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ''] && <span>{property.name}</span>}
                                        {propertyFlags[property.name ?? ''] && <TextField value={property.name}
                                            onChange={(e) => {
                                                propertyFlags[e.target.value ?? ''] = propertyFlags[property.name ?? '']
                                                delete (propertyFlags[property.name ?? ''])
                                                handlePropertyFieldOnChange(props.resource, props.onChange, index, 'name', e.target.value)
                                            }} />}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ""] && (
                                            <span>{property.type}</span>
                                        )}
                                        {propertyFlags[property.name ?? ""] && (
                                            <Select
                                                value={property.type}
                                                onChange={(e) =>
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        "type",
                                                        e.target.value as ResourceProperty["type"]
                                                    )
                                                }
                                            >
                                                <MenuItem value="BOOL">BOOL</MenuItem>
                                                <MenuItem value="STRING">STRING</MenuItem>
                                                <MenuItem value="FLOAT32">FLOAT32</MenuItem>
                                                <MenuItem value="FLOAT64">FLOAT64</MenuItem>
                                                <MenuItem value="INT32">INT32</MenuItem>
                                                <MenuItem value="INT64">INT64</MenuItem>
                                                <MenuItem value="BYTES">BYTES</MenuItem>
                                                <MenuItem value="UUID">UUID</MenuItem>
                                                <MenuItem value="DATE">DATE</MenuItem>
                                                <MenuItem value="TIME">TIME</MenuItem>
                                                <MenuItem value="TIMESTAMP">TIMESTAMP</MenuItem>
                                                <MenuItem value="OBJECT">OBJECT</MenuItem>
                                                <MenuItem value="MAP">MAP</MenuItem>
                                                <MenuItem value="LIST">LIST</MenuItem>
                                                <MenuItem value="REFERENCE">REFERENCE</MenuItem>
                                                <MenuItem value="ENUM">ENUM</MenuItem>
                                            </Select>
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ""] && (
                                            <span>{property.required ? "Yes" : "No"}</span>
                                        )}
                                        {propertyFlags[property.name ?? ""] && (
                                            <TextField
                                                value={property.required ? "Yes" : "No"}
                                                onChange={(e) =>
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        "required",
                                                        e.target.value === "Yes" ? true : false
                                                    )
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ""] && (
                                            <span>{property.primary ? "Yes" : "No"}</span>
                                        )}
                                        {propertyFlags[property.name ?? ""] && (
                                            <TextField
                                                value={property.primary ? "Yes" : "No"}
                                                onChange={(e) =>
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        "primary",
                                                        e.target.value === "Yes" ? true : false
                                                    )
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ""] && (
                                            <span>{property.unique ? "Yes" : "No"}</span>
                                        )}
                                        {propertyFlags[property.name ?? ""] && (
                                            <TextField
                                                value={property.unique ? "Yes" : "No"}
                                                onChange={(e) =>
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        "unique",
                                                        e.target.value === "Yes" ? true : false
                                                    )
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name ?? ''] && <IconButton onClick={() => {
                                            setPropertyFlags({
                                                ...propertyFlags,
                                                [property.name ?? '']: true
                                            })
                                        }}>
                                            <Edit />
                                        </IconButton>}
                                        {propertyFlags[property.name ?? ''] && <IconButton onClick={() => {
                                            setPropertyFlags({
                                                ...propertyFlags,
                                                [property.name ?? '']: false
                                            })
                                        }}>
                                            <Save />
                                        </IconButton>}
                                        <IconButton onClick={() => {
                                            props.onChange({
                                                ...props.resource,
                                                properties: props.resource.properties?.filter((_, i) => i !== index)
                                            })
                                        }}>
                                            <Delete />
                                        </IconButton>
                                    </TableCell>
                                </TableRow>
                            })}
                        </TableBody>
                        <TableFooter>

                        </TableFooter>
                    </Table>
                </TableContainer>
            </FormGroup>
        </Box>
    </>
}
