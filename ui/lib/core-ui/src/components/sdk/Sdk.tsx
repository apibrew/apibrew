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
    const [platform, setPlatform] = React.useState<string>(SdkPlatforms[0].name)

    const Component = (SdkPlatforms.find(p => p.name === platform)?.component)

    return <Box display={'flex'} flexDirection={'column'}>
        <Box display={'flex'} m={1}>
            <h2>Platform: </h2>
            <Select sx={{width: '400px', m: 1}}
                    value={platform}
                    onChange={e => setPlatform(e.target.value)}>
                {SdkPlatforms.map((platform) => {
                    return <MenuItem value={platform.name}>{platform.name}</MenuItem>
                })}
            </Select>
        </Box>

        <Box display={'flex'} m={1}>
            <Component resource={props.resource}/>
        </Box>
    </Box>
}