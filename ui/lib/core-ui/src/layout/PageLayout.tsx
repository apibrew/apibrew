import {Breadcrumbs, Card, CardContent, CardHeader} from '@mui/material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext'
import {Link} from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import * as React from 'react'
import {ReactNode} from 'react'
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
    return <Box sx={{width: '100%', height: '100%', padding: '20px'}}>
        <Box sx={{display: 'flex'}}>
            <Box>
                <h3>
                    {props.pageTitle}
                </h3>

                {props.breadcrumbs &&
                    <Breadcrumbs aria-label="breadcrumb" separator={<NavigateNextIcon fontSize="small"/>}>
                        {props.breadcrumbs.map(item => {
                            if (item.to) {
                                return <Link key={item.label} to={item.to}>{item.label}</Link>
                            } else {
                                return <Typography key={item.label}
                                                   color="text.primary">{item.label}</Typography>
                            }
                        })}
                    </Breadcrumbs>}
            </Box>
            <Box sx={{flexGrow: 1}}/>
            {(props.actions != null) && <Box>{props.actions}</Box>}
        </Box>
        <Box sx={{width: '100%', height: '100%', marginTop: '20px'}}>
            {props.children}
        </Box>
    </Box>
}
