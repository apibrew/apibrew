import { Box, BoxProps } from '@mui/material';
import { StyledRootScrollbar, StyledScrollbar } from './styles';

type ScrollbarProps = {
    sx?: BoxProps['sx'];
    children?: React.ReactNode;
};

export function Scrollbar({ children, sx, ...other }: ScrollbarProps) {
    const userAgent = typeof navigator === 'undefined' ? 'SSR' : navigator.userAgent;

    const isMobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(userAgent);

    if (isMobile) {
        return (
            <Box sx={{ overflowX: 'auto', ...sx }} {...other}>
                {children}
            </Box>
        );
    }

    return (
        <StyledRootScrollbar>
            <StyledScrollbar clickOnTrack={false} sx={sx} {...other}>
                {children}
            </StyledScrollbar>
        </StyledRootScrollbar>
    );
}