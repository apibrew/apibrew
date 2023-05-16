import {
    Box,
    Checkbox,
    FormControl,
    FormGroup,
    FormHelperText,
    FormLabel,
    IconButton,
    MenuItem,
    Select,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableFooter,
    TableHead,
    TableRow,
    TextField
} from '@mui/material'
import { type Resource, type ResourceProperty } from '../../model'
import { Add, Delete, Edit, Save } from '@mui/icons-material'
import React from 'react'

export interface ResourceBasicFormProps {
    resources: Resource[]
    resource: Resource
    onChange: (resource: Resource) => void
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
    const [propertyFlags, setPropertyFlags] = React.useState<Record<string, boolean>>({})

    return <>
        <Box>
            <FormGroup>
                <FormControl>
                    <FormLabel>Resource Name</FormLabel>
                    <TextField value={props.resource.name}
                        onChange={(e) => {
                            props.onChange({ ...props.resource, name: e.target.value })
                        }}/>
                    <FormHelperText>Resource name is required</FormHelperText>
                </FormControl>
                <Box m={1}>
                    <IconButton onClick={() => {
                        const propName = `prop-${((props.resource.properties?.length ?? 0) + 1)}`
                        props.onChange({
                            ...props.resource,
                            properties: [...props.resource.properties ?? [], {
                                name: propName,
                                type: 'STRING',
                                required: false,
                                primary: false,
                                unique: false,
                                length: 255
                            }]
                        })

                        setPropertyFlags({
                            ...propertyFlags,
                            [propName]: true
                        })
                    }}>
                        <Add/>
                    </IconButton>
                </Box>
                <TableContainer>
                    <Table size='small'>
                        <TableHead>
                            <TableRow>
                                <TableCell sx={{ width: '230px' }}>Property Name</TableCell>
                                <TableCell sx={{ width: '150px' }}>Type</TableCell>
                                <TableCell sx={{ width: '50px' }}>Required</TableCell>
                                <TableCell sx={{ width: '50px' }}>Primary</TableCell>
                                <TableCell sx={{ width: '50px' }}>Unique</TableCell>
                                <TableCell sx={{ width: '120px' }}>Actions</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {props.resource.properties?.map((property, index) => {
                                return <TableRow key={index}>
                                    <TableCell>
                                        {!propertyFlags[property.name] && <span>{property.name}</span>}
                                        {propertyFlags[property.name] && <TextField size='small' value={property.name}
                                            onChange={(e) => {
                                                setPropertyFlags({
                                                    ...propertyFlags,
                                                    [e.target.value ?? '']: propertyFlags[property.name],
                                                    [property.name]: false
                                                })

                                                handlePropertyFieldOnChange(props.resource, props.onChange, index, 'name', e.target.value)
                                            }}/>}

                                        {property.type === 'REFERENCE' && <>
                                            &nbsp;
                                            {!propertyFlags[property.name] && (
                                                <span>[{property.reference?.referencedResource}]</span>
                                            )}
                                            {propertyFlags[property.name] && (
                                                <Select
                                                    size='small'
                                                    value={property.reference?.referencedResource}
                                                    onChange={(e) => {
                                                        handlePropertyFieldOnChange(
                                                            props.resource,
                                                            props.onChange,
                                                            index,
                                                            'reference',
                                                            {
                                                                referencedResource: e.target.value
                                                            }
                                                        )
                                                    }
                                                    }
                                                >
                                                    {props.resources.map(item => <MenuItem key={item.name}
                                                        value={item.name}>{item.name}</MenuItem>)}
                                                </Select>
                                            )}
                                        </>}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name] && (
                                            <span>{property.type?.toLowerCase()}</span>
                                        )}
                                        {propertyFlags[property.name] && (
                                            <Select
                                                size='small'
                                                value={property.type}
                                                onChange={(e) => {
                                                    if (e.target.value === 'STRING' && ((property.length ?? 0) <= 0)) {
                                                        handlePropertyFieldOnChange(
                                                            props.resource,
                                                            props.onChange,
                                                            index,
                                                            'length',
                                                            255
                                                        )
                                                    }

                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        'type',
                                                        e.target.value as ResourceProperty['type']
                                                    )
                                                }}
                                            >
                                                <MenuItem value="BOOL">bool</MenuItem>
                                                <MenuItem value="STRING">string</MenuItem>
                                                <MenuItem value="FLOAT32">float32</MenuItem>
                                                <MenuItem value="FLOAT64">float64</MenuItem>
                                                <MenuItem value="INT32">int32</MenuItem>
                                                <MenuItem value="INT64">int64</MenuItem>
                                                <MenuItem value="BYTES">bytes</MenuItem>
                                                <MenuItem value="UUID">uuid</MenuItem>
                                                <MenuItem value="DATE">date</MenuItem>
                                                <MenuItem value="TIME">time</MenuItem>
                                                <MenuItem value="TIMESTAMP">timestamp</MenuItem>
                                                <MenuItem value="OBJECT">object</MenuItem>
                                                <MenuItem value="MAP">map</MenuItem>
                                                <MenuItem value="LIST">list</MenuItem>
                                                <MenuItem value="REFERENCE">reference</MenuItem>
                                                <MenuItem value="ENUM">enum</MenuItem>
                                            </Select>
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name] && (
                                            <span>{property.required ? 'Yes' : 'No'}</span>
                                        )}
                                        {propertyFlags[property.name] && (
                                            <Checkbox
                                                size='small'
                                                checked={property.required}
                                                onChange={(e) => {
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        'required',
                                                        e.target.checked
                                                    )
                                                }
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name] && (
                                            <span>{property.primary ? 'Yes' : 'No'}</span>
                                        )}
                                        {propertyFlags[property.name] && (
                                            <Checkbox
                                                size='small'
                                                checked={property.primary}
                                                onChange={(e) => {
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        'primary',
                                                        e.target.checked
                                                    )
                                                }
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name] && (
                                            <span>{property.unique ? 'Yes' : 'No'}</span>
                                        )}
                                        {propertyFlags[property.name] && (
                                            <Checkbox
                                                size='small'
                                                checked={property.unique}
                                                onChange={(e) => {
                                                    handlePropertyFieldOnChange(
                                                        props.resource,
                                                        props.onChange,
                                                        index,
                                                        'unique',
                                                        e.target.checked
                                                    )
                                                }
                                                }
                                            />
                                        )}
                                    </TableCell>
                                    <TableCell>
                                        {!propertyFlags[property.name] && <IconButton onClick={() => {
                                            setPropertyFlags({
                                                ...propertyFlags,
                                                [property.name]: true
                                            })
                                        }}>
                                            <Edit/>
                                        </IconButton>}
                                        {propertyFlags[property.name] && <IconButton onClick={() => {
                                            setPropertyFlags({
                                                ...propertyFlags,
                                                [property.name]: false
                                            })
                                        }}>
                                            <Save/>
                                        </IconButton>}
                                        <IconButton onClick={() => {
                                            props.onChange({
                                                ...props.resource,
                                                properties: props.resource.properties?.filter((_, i) => i !== index)
                                            })
                                        }}>
                                            <Delete/>
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
