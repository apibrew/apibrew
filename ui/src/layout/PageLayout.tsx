import { Breadcrumbs } from '@mui/material'
import NavigateNextIcon from '@mui/icons-material/NavigateNext'
import { Link } from 'react-router-dom'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import * as React from 'react'
import Divider from '@mui/material/Divider'

export interface PageLayoutProps {
    pageTitle: string
    children: JSX.Element | JSX.Element[]
}

export function PageLayout(props: PageLayoutProps) {
    return <>
        <Typography sx={{ m: 3 }} component={'h3'}>{props.pageTitle}</Typography>
        <Box sx={{ m: 3 }}>
            <Breadcrumbs aria-label="breadcrumb" separator={<NavigateNextIcon fontSize="small"/>}>
                <Link color="inherit" to="/">
                    MUI
                </Link>
                <Link
                    color="inherit"
                    to="/material-ui/getting-started/installation/"
                >
                    Core
                </Link>
                <Typography color="text.primary">Breadcrumbs</Typography>
            </Breadcrumbs>
        </Box>
        <Divider/>
        <Box sx={{ m: 3 }}>
            {props.children}
        </Box>
    </>
}
