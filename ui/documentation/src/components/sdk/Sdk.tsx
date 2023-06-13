import {Box, MenuItem, Select, Typography} from "@mui/material";
import {Resource} from "../../model";

import React, {JSX} from "react";
import CurdSdkPlatform from "./CurlSdkPlatform.tsx";

export interface SdkProps {
    resource?: Resource
}

export interface SdkComponentProps {
    resource?: Resource
}

export interface SdkPlatform {
    name: string
    component: (props: SdkComponentProps) => JSX.Element
}

export const SdkPlatforms: SdkPlatform[] = [
    CurdSdkPlatform,
]

export function Sdk(props: SdkProps): JSX.Element {

    return <Box display={'flex'} flexDirection={'row'} height={'100%'}>
        <Box display={'flex'} flexDirection={'row'} alignItems={'center'}
             sx={{background: 'white', width: '300px', height: '100%'}}>
            <Typography variant={'h6'}>SDK</Typography>
            <Box flexGrow={1}/>
        </Box>
        <Box display={'flex'} m={1}>
        </Box>
    </Box>
}