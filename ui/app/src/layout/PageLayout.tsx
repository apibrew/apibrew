import Box from '@mui/material/Box'
import {ReactNode} from 'react'

export interface PageLayoutProps {
    children: ReactNode
}

export function PageLayout(props: PageLayoutProps) {
    return <Box sx={{width: '100%', height: '100%', padding: '20px', paddingTop: '10px'}}>
        {props.children}
    </Box>
}
