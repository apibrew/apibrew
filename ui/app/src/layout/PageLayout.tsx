import {Breadcrumbs, Card, CardContent, CardHeader} from '@mui/material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext'
import {Link} from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import * as React from 'react'
import {type ReactNode} from 'react'
import Divider from '@mui/material/Divider'

export interface Breadcrumb {
    label: string
    to?: string
}

export interface PageLayoutProps {
    pageTitle?: ReactNode
    children: ReactNode
    actions?: ReactNode
    breadcrumbs?: Breadcrumb[]
}

export function PageLayout(props: PageLayoutProps) {
    return <>
        <Card>
            <CardHeader title={
                <Box>
                    <Box sx={{display: 'flex'}}>
                        <Box>
                            {props.pageTitle}

                            {props.breadcrumbs &&
                                <Breadcrumbs aria-label="breadcrumb" separator={<NavigateNextIcon fontSize="small"/>}>
                                    {props.breadcrumbs.map(item => <>
                                        {item.to && <Link color="black"
                                                          to={item.to}
                                                          style={{textDecoration: 'none', color: 'black'}}>
                                            {item.label}
                                        </Link>}
                                        {!item.to && <Typography color="text.primary">{item.label}</Typography>}
                                    </>)}
                                </Breadcrumbs>}
                        </Box>
                        <Box sx={{flexGrow: 1}}/>
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
