import {
    Box,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Collapse
} from "@mui/material";
import React from "react";
import { ChevronRight, ExpandMore, Link } from "@mui/icons-material";
import IconButton from "@mui/material/IconButton";
import { OpenAPIV3_1 } from "openapi-types";
import { preparePropertyList, resolveSchema } from "../../util/openapi";

export interface SchemaElementTableCellsProps {
    doc: OpenAPIV3_1.Document;
    schema: OpenAPIV3_1.SchemaObject | OpenAPIV3_1.ReferenceObject;
    depth: number
    ignoredTypeNames: string[]
}

function SchemaElementTableCells(props: SchemaElementTableCellsProps) {
    const [open, setOpen] = React.useState<{ [key: string]: boolean }>({});

    const properties = preparePropertyList(props.doc, props.schema)

    if (properties.length == 0) {
        return <TableRow><TableCell>Empty</TableCell></TableRow>
    }

    return <>
        {properties.map(property => {
            const schema = resolveSchema(props.doc, property.schema)

            const hasSubPart = schema.type === 'object' || schema.type === 'array'

            return <React.Fragment key={property.name}>
                <TableRow>
                    <TableCell sx={{
                        padding: 0,
                        paddingLeft: `${props.depth * 40}px`
                    }}>
                        {property.name}
                        {hasSubPart && <IconButton sx={{ padding: 0 }} onClick={() => setOpen({
                            ...open,
                            [property.name]: !open[property.name]
                        })}>
                            {open[property.name] ? <ExpandMore fontSize='small' /> : <ChevronRight fontSize='small' />}
                        </IconButton>}
                    </TableCell>
                    <TableCell sx={{ padding: 0 }}>{schema.type}</TableCell>
                    <TableCell sx={{ padding: 0 }}>
                        {schema.description}
                    </TableCell>
                </TableRow>
                {open[property.name] && <>
                    {schema.type === 'object' && <SchemaElementTableCells
                        key={property.name}
                        depth={props.depth + 1}
                        doc={props.doc}
                        ignoredTypeNames={[...props.ignoredTypeNames, property.name]}
                        schema={schema}
                    />}
                    {schema.type === 'array' && <SchemaElementTableCells
                        key={property.name}
                        depth={props.depth + 1}
                        doc={props.doc}
                        ignoredTypeNames={[...props.ignoredTypeNames, property.name]}
                        schema={(schema as OpenAPIV3_1.ArraySchemaObject).items}
                    />}
                </>}
            </React.Fragment>
        })}
    </>
}

export interface SchemaElementProps {
    name: string
    doc: OpenAPIV3_1.Document;
    open?: boolean
}

export function SchemaElement(props: SchemaElementProps): JSX.Element {
    const [open, setOpen] = React.useState(props.open)
    const schema = props.doc.components!.schemas![props.name]

    return <Box>
        <h4 style={{ margin: 0 }}>
            {props.name}
            <IconButton onClick={() => setOpen(!open)}>
                {open ? <ExpandMore fontSize='small' /> : <ChevronRight fontSize='small' />}
            </IconButton>
            <a href={`#element-${props.name.toLowerCase()}`}>
                <Link fontSize='small' />
            </a>
        </h4>
        {schema.description && <p>{schema.description}</p>}
        <Collapse in={open} timeout="auto" unmountOnExit>
            <Table sx={{ marginLeft: '10px' }}>
                <TableHead>
                    <TableRow>
                        <TableCell sx={{ width: '45%', padding: 0, fontWeight: 'bold' }}>Property</TableCell>
                        <TableCell sx={{ width: '100px', padding: 0, fontWeight: 'bold' }}>Type</TableCell>
                        <TableCell sx={{ padding: 0, fontWeight: 'bold' }}>Description</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <SchemaElementTableCells depth={0}
                        ignoredTypeNames={[props.name]}
                        doc={props.doc}
                        schema={schema} />
                </TableBody>
            </Table>
        </Collapse>
    </Box>
}