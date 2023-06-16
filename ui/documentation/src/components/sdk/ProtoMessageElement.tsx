import {resolveByTypeName, translateType, locateDescription} from "../../proto";
import {MessageType, File, ProtoElement} from "../../proto/image";
import {
    Box,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Collapse
} from "@mui/material";
import React, {ReactNode} from "react";
import {ChevronRight, ExpandMore, Link} from "@mui/icons-material";
import IconButton from "@mui/material/IconButton";

export interface ProtoMessageTableCellsProps {
    messageType: MessageType
    protoFile: File
    depth: number
    ignoredMessageTypes: string[]
}

function ProtoMessageTableCells(props: ProtoMessageTableCellsProps) {
    const [open, setOpen] = React.useState<{ [key: string]: boolean }>({});

    const messageIndex = props.protoFile.messageType?.indexOf(props.messageType);

    return <>
        {(props.messageType?.field ?? []).map((field, index) => {
            let subElements: ReactNode = null;
            let sub: {
                element: ProtoElement
                file: File,
                special?: boolean,
            } | null = null;

            if (field.type == 'TYPE_MESSAGE') {
                if (props.ignoredMessageTypes.includes(field.typeName as string)) {
                    subElements = null
                } else {
                    sub = resolveByTypeName('message', props.protoFile, field.typeName as string)

                    if (sub.special) {
                        subElements = null
                    } else {
                        const subMessage = sub.element as MessageType;
                        subElements = <ProtoMessageTableCells depth={props.depth + 1}
                                                              ignoredMessageTypes={[...props.ignoredMessageTypes, field.typeName as string]}
                                                              messageType={subMessage}
                                                              protoFile={sub.file}/>
                    }
                }
            }

            console.log(`[4, messageIndex!, 2, index, 1]`, [4, messageIndex!, 2, index, 1])

            return <>
                <TableRow>
                    <TableCell sx={{
                        padding: 0,
                        paddingLeft: `${props.depth * 40}px`
                    }}>
                        {field.name}
                        {subElements && <IconButton sx={{padding: 0}} onClick={() => setOpen({
                            ...open,
                            [field.name]: !open[field.name]
                        })}>
                            {open[field.name] ? <ExpandMore fontSize='small'/> : <ChevronRight fontSize='small'/>}
                        </IconButton>}
                    </TableCell>
                    <TableCell sx={{padding: 0}}>{translateType(props.protoFile, field)}</TableCell>
                    <TableCell sx={{padding: 0}}>
                        {locateDescription(props.protoFile, [4, messageIndex!, 2, index])}
                    </TableCell>
                </TableRow>
                <Collapse in={open[field.name]} timeout="auto" unmountOnExit>
                    {subElements}
                </Collapse>
            </>
        })}
    </>;
}

export interface ProtoMessageElementProps {
    messageType: MessageType
    protoFile: File
    open?: boolean
}

export function ProtoMessageElement(props: ProtoMessageElementProps): JSX.Element {
    const [open, setOpen] = React.useState(props.open)

    return <Box>
        <h4 id={`element-${props.messageType.name.toLowerCase()}`}>
            {props.messageType.name}
            <IconButton onClick={() => setOpen(!open)}>
                {open ? <ExpandMore fontSize='small'/> : <ChevronRight fontSize='small'/>}
            </IconButton>
            <a href={`#element-${props.messageType.name.toLowerCase()}`}>
                <Link fontSize='small'/>
            </a>
        </h4>
        <Collapse in={open} timeout="auto" unmountOnExit>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell sx={{width: '45%'}}>Property</TableCell>
                        <TableCell sx={{width: '100px'}}>Type</TableCell>
                        <TableCell>Description</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <ProtoMessageTableCells depth={0}
                                            ignoredMessageTypes={[props.messageType.name]}
                                            messageType={props.messageType}
                                            protoFile={props.protoFile}/>
                </TableBody>
            </Table>
        </Collapse>
    </Box>
}