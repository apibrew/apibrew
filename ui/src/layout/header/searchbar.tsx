import { useState } from 'react';
import { styled } from '@mui/material/styles';
import { Input, Slide, Button, IconButton, InputAdornment, ClickAwayListener } from '@mui/material';

const HEADER_MOBILE = 64;
const HEADER_DESKTOP = 92;
const StyledSearchbar = styled('div')(({ theme }) => ({

    top: 0,
    left: 0,
    zIndex: 99,
    width: '100%',
    display: 'flex',
    position: 'absolute',
    alignItems: 'center',
    height: HEADER_MOBILE,
    padding: theme.spacing(0, 3),

    [theme.breakpoints.up('md')]: {
        height: HEADER_DESKTOP,
        padding: theme.spacing(0, 5),
    },
}));

export default function Searchbar() {
    const [open, setOpen] = useState(false);
    const handleOpen = () => {
        setOpen(!open);
    };
    const handleClose = () => {
        setOpen(false);
    };

    return (
        <ClickAwayListener onClickAway={handleClose}>
            <div>
                {!open && (
                    <IconButton onClick={handleOpen}>
                    </IconButton>
                )}
                <Slide direction="down" in={open} mountOnEnter unmountOnExit>
                    <StyledSearchbar>
                        <Input
                            autoFocus
                            fullWidth
                            disableUnderline
                            placeholder="Search…"
                            startAdornment={
                                <InputAdornment position="start">
                                </InputAdornment>
                            }
                            sx={{ mr: 1, fontWeight: 'fontWeightBold' }}
                        />
                        <Button variant="contained" onClick={handleClose}>
                            Search
                        </Button>
                    </StyledSearchbar>
                </Slide>
            </div>
        </ClickAwayListener>
    );
}