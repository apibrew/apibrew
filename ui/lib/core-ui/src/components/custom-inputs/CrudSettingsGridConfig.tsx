import {CrudGridConfig} from "../../../model/schema";
import {Checkbox, Table, TableBody, TableCell, TableHead, TableRow} from "@mui/material";
import {useResource} from "../../context/resource.ts";
import React from "react";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";

export interface CrudSettingsGridConfigProps {
    config: CrudGridConfig
}

export function CrudSettingsGridConfig(props: CrudSettingsGridConfigProps) {
    const resource = useResource()

    return (
        <Box m={1}>
            <Typography>Grid Column Config</Typography>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Property</TableCell>
                        <TableCell>Disabled</TableCell>
                        <TableCell>Hidden</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {resource.properties.map((property) => (
                        <TableRow key={property.name}>
                            <TableCell>{property.name}</TableCell>
                            <TableCell>
                                <Checkbox />
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Box>
    )
}