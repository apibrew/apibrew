import { Box, Drawer, IconButton, Typography } from "@mui/material"
import React, { useState } from "react"
import { Resource } from "@apibrew/client"
import { Close, KeyboardArrowLeft, KeyboardArrowRight } from "@mui/icons-material"
import { Sdk } from "./Sdk"

type DrawerSize = 'small' | 'medium' | 'large'

export interface SdkDrawerProps {
    resource?: Resource
    open: boolean
    size?: DrawerSize
    onClose: () => void
}

export function SdkDrawer(props: SdkDrawerProps): JSX.Element {
    const [size, setSize] = useState<DrawerSize>(props.size ?? 'large')

    const width = size === 'small' ? 400 : size === 'medium' ? 600 : 800

    return <React.Fragment>
        <Drawer anchor={'right'}
            BackdropProps={{ invisible: true }}
            ModalProps={{ sx: { '& .MuiDrawer-paper': { width: `${width}px`, top: '88px' } } }}
            onClose={props.onClose}
            open={props.open}

        >
            <Box display='flex' flexDirection={'column'} padding='10px'>
                <Box display='flex' width='100%'>
                    {size !== 'large' && <IconButton onClick={() => {
                        setSize(size === 'small' ? 'medium' : 'large')
                    }}>
                        <KeyboardArrowLeft />
                    </IconButton>}
                    {size != 'small' && <IconButton onClick={() => {
                        setSize(size === 'medium' ? 'small' : 'medium')
                    }}>
                        <KeyboardArrowRight />
                    </IconButton>}
                    <Box flexGrow={1} />
                    <Typography variant='h6'>SDK</Typography>
                    <Box flexGrow={1} />
                    <IconButton onClick={props.onClose}>
                        <Close />
                    </IconButton>
                </Box>
                <Box>
                    <Sdk />
                </Box>
            </Box>
        </Drawer>
    </React.Fragment>
}
