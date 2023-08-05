import {Box, Typography} from '@mui/material';

export interface ErrorProps {
    message: string
}

export function Error(props: ErrorProps) {
    return (
        <Box
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                minHeight: '100vh',
                backgroundColor: 'gray',
                color: 'black'
            }}
        >
            <Typography variant="h1" style={{color: 'white'}}>
                {props.message}
            </Typography>
        </Box>
    );
}