import { Checkbox, Table, TableBody, TableCell, TableHead, TableRow, TextField } from "@mui/material";
import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import { Crud, GridColumnConfig, GridConfig } from "../../model/ui/crud.ts";
import { useRecord } from "../../context/record.ts";
import { useResourceByName } from "../../hooks/resource.ts";
import { useValue } from "../../context/value.ts";

export interface CrudSettingsGridConfigProps {
    config: GridConfig
}

export function CrudSettingsGridConfig() {
    const record = useRecord<Crud>()
    const resource = useResourceByName(record.resource, record.namespace)
    const valueContext = useValue()

    const [columns, setColumns] = useState<GridColumnConfig[]>(valueContext.value?.columns)

    useEffect(() => {
        if (!resource) {
            return
        }
        setColumns(resource.properties.map((property) => {
            const existingColumn = valueContext.value?.columns.find((column: GridColumnConfig) => column.name === property.name)
            return {
                name: property.name,
                disabled: existingColumn?.disabled ?? false,
                hidden: existingColumn?.hidden ?? false,
                sortable: existingColumn?.sortable ?? false,
                filterable: existingColumn?.filterable ?? false,
                width: existingColumn?.width ?? 0,
                flex: existingColumn?.flex ?? 0,
            } as GridColumnConfig
        }))
    }, [resource])

    if (!resource) {
        return <>Loading...</>
    }

    function handleColumnUpdate(column: GridColumnConfig, prop: string, newValue: any) {
        const updatedColumns = columns.map((c: any) => {
            if (c.name === column.name) {
                c[prop] = newValue
            }
            return c
        })
        setColumns(updatedColumns)
        valueContext.onChange({
            ...valueContext.value,
            columns: updatedColumns
        })
    }


    return (
        <Box m={1}>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Property</TableCell>
                        <TableCell>Disabled</TableCell>
                        <TableCell>Hidden</TableCell>
                        <TableCell>Sortable</TableCell>
                        <TableCell>Filterable</TableCell>
                        <TableCell>Width</TableCell>
                        <TableCell>Flex</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {columns.map((column) => (
                        <TableRow key={column.name}>
                            <TableCell padding='none'>{column.name}</TableCell>
                            <TableCell padding='none'>
                                <Checkbox checked={column.disabled} onChange={e => {
                                    handleColumnUpdate(column, 'disabled', e.target.checked);
                                }} />
                            </TableCell>
                            <TableCell padding='none'>
                                <Checkbox checked={column.hidden} onChange={e => {
                                    handleColumnUpdate(column, 'hidden', e.target.checked);
                                }} />
                            </TableCell>

                            <TableCell padding='none'>
                                <Checkbox checked={column.sortable} onChange={e => {
                                    handleColumnUpdate(column, 'sortable', e.target.checked);
                                }} />
                            </TableCell>

                            <TableCell padding='none'>
                                <Checkbox checked={column.filterable} onChange={e => {
                                    handleColumnUpdate(column, 'filterable', e.target.checked);
                                }} />
                            </TableCell>
                            <TableCell padding='none'>
                                <TextField type={'number'} value={column.width} onChange={e => {
                                    handleColumnUpdate(column, 'width', parseInt(e.target.value));
                                }} />
                            </TableCell>
                            <TableCell padding='none'>
                                <TextField type={'number'} value={column.flex} onChange={e => {
                                    handleColumnUpdate(column, 'flex', parseInt(e.target.value));
                                }} />
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Box>
    )
}