import React, { useState } from 'react'
import { alpha } from '@mui/material/styles'
import { Box, Divider, Typography, Stack, MenuItem, Avatar, IconButton, Popover } from '@mui/material'
import { useNavigate } from 'react-router-dom'
import { TokenService } from '../../service/token'
const MENU_OPTIONS = [
    {
        label: 'Home',
        icon: 'eva:home-fill'
    },
    {
        label: 'Profile',
        icon: 'eva:person-fill'
    },
    {
        label: 'Settings',
        icon: 'eva:settings-2-fill'
    }
]

export default function AccountPopover() {
    const navigate = useNavigate()
    const [open, setOpen] = useState(false)
    const handleOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }

    return (
        <React.Fragment>
            <IconButton
                onClick={handleOpen}
                sx={{
                    p: 0,
                    ...(open && {
                        '&:before': {
                            zIndex: 1,
                            content: "''",
                            width: '100%',
                            height: '100%',
                            borderRadius: '50%',
                            position: 'absolute',
                            bgcolor: (theme) => alpha(theme.palette.grey[900], 0.8)
                        }
                    })
                }}
            >
                <Avatar alt="photoURL" />
            </IconButton>
            <Popover
                open={Boolean(open)}
                onClose={handleClose}
                anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
                transformOrigin={{ vertical: 'top', horizontal: 'right' }}
                PaperProps={{
                    sx: {
                        p: 0,
                        mt: 0,
                        ml: 0.75,
                        width: 180,
                        '& .MuiMenuItem-root': {
                            typography: 'body2',
                            borderRadius: 0.75
                        }
                    }
                }}
            >
                <Box sx={{ my: 1.5, px: 2.5 }}>
                    <Typography variant="subtitle2" noWrap>
                    </Typography>
                    <Typography variant="body2" sx={{ color: 'text.secondary' }} noWrap>
                    </Typography>
                </Box>
                <Divider sx={{ borderStyle: 'dashed' }} />
                <Stack sx={{ p: 1 }}>
                    {MENU_OPTIONS.map((option) => (
                        <MenuItem key={option.label} onClick={handleClose}>
                            {option.label}
                        </MenuItem>
                    ))}
                </Stack>
                <Divider sx={{ borderStyle: 'dashed' }} />
                <MenuItem onClick={() => {
                    TokenService.removeToken()
                    navigate('/login')
                }} sx={{ m: 1 }}>
                    Logout
                </MenuItem>
            </Popover>
        </React.Fragment>
    )
}
