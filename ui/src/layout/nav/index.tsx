import PropTypes from 'prop-types';
import { useEffect } from 'react';
import { useLocation } from 'react-router-dom';
// @mui
import { styled, alpha } from '@mui/material/styles';
import { Box, Link, Button, Drawer, Typography, Avatar, Stack } from '@mui/material';
// mock
// import account from '../../../_mock/account';
// // hooks

// // components
// import Logo from '../../../components/logo';
// import Scrollbar from '../../../components/scrollbar';
// import NavSection from '../../../components/nav-section';
// //
import navConfig from './config';
import useResponsive from '../../hooks/use-responsive';
import { Scrollbar } from '../../components/scroll-bar'
import { Logo } from '../../components/logo';
import { account } from '../../components/account';
import NavSection from '../../components/navsection';

// ----------------------------------------------------------------------

const NAV_WIDTH = 280;

// ----------------------------------------------------------------------

Nav.propTypes = {
    openNav: PropTypes.bool,
    onCloseNav: PropTypes.func,
};

export interface NavProps {
    openNav: boolean
    onCloseNav: () => void
}

export default function Nav({ openNav, onCloseNav }: NavProps) {
    const { pathname } = useLocation();

    const isDesktop = useResponsive('up', 'lg');

    useEffect(() => {
        if (openNav) {
            onCloseNav();
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [pathname]);

    const renderContent = (
        <Scrollbar
            sx={{
                height: 1,
                '& .simplebar-content': { height: 1, display: 'flex', flexDirection: 'column' },
            }}
        >
            <Box sx={{ px: 2.5, py: 3, display: 'inline-flex' }}>
                <Logo />
            </Box>

            <NavSection data={navConfig} />

            <Box sx={{ flexGrow: 1 }} />
        </Scrollbar>
    );

    return (
        <Box
            component="nav"
            sx={{
                flexShrink: { lg: 0 },
                width: { lg: NAV_WIDTH },
            }}
        >
            {isDesktop ? (
                <Drawer
                    open
                    variant="permanent"
                    PaperProps={{
                        sx: {
                            width: NAV_WIDTH,
                            bgcolor: 'background.default',
                            borderRightStyle: 'dashed',
                        },
                    }}
                >
                     {renderContent}
                </Drawer>
            ) : (
                <Drawer
                    open={openNav}
                    onClose={onCloseNav}
                    ModalProps={{
                        keepMounted: true,
                    }}
                    PaperProps={{
                        sx: { width: NAV_WIDTH },
                    }}
                >
                     {renderContent}
                </Drawer>
            )}
        </Box>
    );
}