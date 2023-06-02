import {Box, LinearProgress} from "@mui/material";
import React from "react";

export function Loading() {
    return <Box sx={{width: '100%'}}>
        <LinearProgress/>
    </Box>
}