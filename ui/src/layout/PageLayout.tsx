import { Breadcrumbs, Card, CardContent, CardHeader } from '@mui/material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext'
import { Link } from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import * as React from 'react'
import Divider from '@mui/material/Divider'
import {ReactNode} from "react";

export interface PageLayoutProps {
    pageTitle: string
    children: ReactNode
    actions?: ReactNode
}

export function PageLayout(props: PageLayoutProps) {
    return <>
        <Card>
            <CardHeader title={
                <Box>
                    <Box sx={{ display: 'flex' }}>
                        <Box>
                            {props.pageTitle}

                            <Breadcrumbs aria-label="breadcrumb" separator={<NavigateNextIcon fontSize="small"/>}>
                                <Link color="black"
                                    to="/"
                                    style={{ textDecoration: 'none', color: 'black' }}>
                                   MUI
                                </Link>
                                <Link
                                    color="black"
                                    style={{ textDecoration: 'none', color: 'black' }}
                                    to="/material-ui/getting-started/installation/">
                                   Core
                                </Link>
                                <Typography color="text.primary">Breadcrumbs</Typography>
                            </Breadcrumbs>
                        </Box>
                        <Box sx={{ flexGrow: 1 }}/>
                        {(props.actions != null) && <Box>{props.actions}</Box>}
                    </Box>
                </Box>
            }></CardHeader>
            <Divider/>
            <CardContent>
                {props.children}
            </CardContent>
        </Card>
    </>
}
