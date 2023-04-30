import PropTypes from 'prop-types';
import { forwardRef, Ref } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Box, BoxProps, Link } from '@mui/material';

type LogoProps = {
    disabledLink?: boolean;
    sx?: BoxProps['sx'];
} & BoxProps;

export const Logo = forwardRef<HTMLDivElement, LogoProps>(
    ({ disabledLink = false, sx, ...other }, ref: Ref<HTMLDivElement>) => {

        const logo = (
            <Box
                ref={ref}
                component="div"
                sx={{
                    width: 250,
                    height: 100,
                    marginLeft: 8,
                    marginTop: -2,
                    marginBottom: -5,
                    display: 'inline-flex',
                    ...sx,
                }}
                {...other}>
                <img src="/logo.svg"></img>
            </Box>
        );

        if (disabledLink) {
            return <>{logo}</>;
        }
        return (
            <Link to="/" component={RouterLink} sx={{ display: 'contents' }}>
                {logo}
            </Link>
        );
    });

Logo.propTypes = {
    sx: PropTypes.object,
    disabledLink: PropTypes.bool,
};