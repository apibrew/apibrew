import { Breadcrumbs, Card, CardActions, CardContent, CardHeader } from '@mui/material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext'
import { Link } from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import * as React from 'react'
import Divider from '@mui/material/Divider'

export interface PageLayoutProps {
    pageTitle: string
    children: React.ReactNode
    headerActions?: React.ReactNode
    bottomActions?: React.ReactNode
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
                        {(props.headerActions != null) && <Box>{props.headerActions}</Box>}
                    </Box>
                </Box>
            }></CardHeader>
            <Divider/>
            <CardContent>
                {props.children}
            </CardContent>
            {(props.bottomActions != null) && <CardActions>{props.bottomActions}</CardActions>}
        </Card>
    </>
}
