import PropTypes from 'prop-types'
import { styled } from '@mui/material/styles'
import { Box, Stack, AppBar, Toolbar } from '@mui/material'
import AccountPopover from './AccountPopover'

const NAV_WIDTH = 280
const HEADER_MOBILE = 64
const HEADER_DESKTOP = 92
const StyledRoot = styled(AppBar)(({ theme }) => ({
    boxShadow: 'none',
    [theme.breakpoints.up('lg')]: {
        width: `calc(100% - ${NAV_WIDTH + 1}px)`
    }
}))
const StyledToolbar = styled(Toolbar)(({ theme }) => ({
    minHeight: HEADER_MOBILE,
    [theme.breakpoints.up('lg')]: {
        minHeight: HEADER_DESKTOP,
        padding: theme.spacing(0, 5)
    }
}))

Header.propTypes = {
    onOpenNav: PropTypes.func
}

export default function Header() {
    return (
        <StyledRoot>
            <StyledToolbar>
                <Box sx={{ flexGrow: 1 }} />
                <Stack
                    direction="row"
                    alignItems="center"
                    spacing={{
                        xs: 0.5,
                        sm: 1
                    }}>
                    <AccountPopover />
                </Stack>
            </StyledToolbar>
        </StyledRoot>
    )
}
