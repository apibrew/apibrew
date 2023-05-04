import * as React from 'react'
import {useState} from 'react'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import IconButton from '@mui/material/IconButton'
import MenuIcon from '@mui/icons-material/Menu'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Drawer from '@mui/material/Drawer'
import Divider from '@mui/material/Divider'
import List from '@mui/material/List'
import {Collapse, Stack} from '@mui/material'
import AccountPopover from './AccountPopover'
import {MenuList, menuLists} from './menu-items'
import ListItemButton from '@mui/material/ListItemButton'
import {Link} from 'react-router-dom'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemText from '@mui/material/ListItemText'
import ListItem from '@mui/material/ListItem'
import {ExpandLess, ExpandMore} from "@mui/icons-material";

const drawerWidth = 260

const drawerStyle = {
    background: '#233044',
    color: 'rgba(238, 238, 238, 0.7)'
}

export interface DashboardLayoutProps {
    children: JSX.Element | JSX.Element[]
}

export function DashboardLayout(props: DashboardLayoutProps): JSX.Element {
    const [mobileOpen, setMobileOpen] = React.useState(false)

    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen)
    }

    const drawer = (
        <Box sx={{display: 'flex', flexDirection: 'column', height: '100%', ...drawerStyle}}>
            <div>
                <Toolbar>
                    <img style={{textAlign: 'center'}} src="/logo-small-white.svg"></img>
                </Toolbar>
                <Divider style={{background: '#AAA'}}/>
                {menuLists.map((menuList, index) => <>
                    <NavList key={menuList.title} menuList={menuList}/>
                    <Divider style={{background: '#FFF'}}/>
                </>)}
            </div>
            <div style={{flexGrow: 1}}/>
        </Box>
    )

    return <>
        <Box sx={{display: 'flex'}}>
            <AppBar
                position="fixed"
                sx={{
                    width: {sm: `calc(100% - ${drawerWidth}px)`},
                    ml: {sm: `${drawerWidth}px`}
                }}
            >
                <Toolbar>
                    <IconButton
                        color="inherit"
                        aria-label="open drawer"
                        edge="start"
                        onClick={handleDrawerToggle}
                        sx={{mr: 2, display: {sm: 'none'}}}
                    >
                        <MenuIcon/>
                    </IconButton>
                    <Typography variant="h6" noWrap component="div">

                    </Typography>
                    <Box sx={{flexGrow: 1}}/>
                    <Stack
                        direction="row"
                        alignItems="center"
                        spacing={{
                            xs: 0.5,
                            sm: 1
                        }}>
                        <AccountPopover/>
                    </Stack>
                </Toolbar>
            </AppBar>
            <Box
                component="nav"
                sx={{width: {sm: drawerWidth}, flexShrink: {sm: 0}}}
                aria-label="mailbox folders"
            >
                {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
                <Drawer
                    variant="temporary"
                    open={mobileOpen}
                    onClose={handleDrawerToggle}
                    ModalProps={{
                        keepMounted: true // Better open performance on mobile.
                    }}
                    sx={{
                        display: {xs: 'block', sm: 'none'},
                        '& .MuiDrawer-paper': {boxSizing: 'border-box', width: drawerWidth}
                    }}
                >
                    {drawer}
                </Drawer>
                <Drawer
                    variant="permanent"
                    sx={{
                        display: {xs: 'none', sm: 'block'},
                        '& .MuiDrawer-paper': {boxSizing: 'border-box', width: drawerWidth}
                    }}
                    open
                >
                    {drawer}
                </Drawer>
            </Box>
            <Box
                component="main"
                sx={{flexGrow: 1, p: 3, width: {sm: `calc(100% - ${drawerWidth}px)`}}}
            >
                <Toolbar/>
                {props.children}
            </Box>
        </Box>
    </>
}

export interface NavListProps {
    menuList: MenuList
}

function NavList(props: NavListProps): JSX.Element {
    const [open, setOpen] = useState<{
        [key: string]: boolean
    }>({});

    return <List subheader={props.menuList.title && <Box sx={{ml: 1, mt: 1}}>
        <Typography>{props.menuList.title}</Typography>
    </Box>}>
        {props.menuList.items.map((menuItem, index) => {
            const key = `${menuItem.title}`
            return <>
                <ListItem key={menuItem.title} disablePadding>
                    {!menuItem.children && <ListItemButton component={Link} to={menuItem.link!}>
                        {(menuItem.icon != null) &&
                            <ListItemIcon style={drawerStyle}>{menuItem.icon} </ListItemIcon>}
                        <ListItemText primary={menuItem.title}/>
                    </ListItemButton>}
                    {menuItem.children && <ListItemButton onClick={() => {
                        setOpen({...open, [key]: !open[key]})
                    }}>
                        {(menuItem.icon != null) &&
                            <ListItemIcon style={drawerStyle}>{menuItem.icon} </ListItemIcon>}
                        <ListItemText primary={menuItem.title}/>
                        {open[key] ? <ExpandLess/> : <ExpandMore/>}
                    </ListItemButton>}
                </ListItem>
                {menuItem.children && <Collapse in={open[key]} timeout="auto" unmountOnExit>
                    <Box sx={{ml:3}}>
                        <NavList menuList={{items: menuItem.children!}}/>
                    </Box>
                </Collapse>}
            </>
        })}
    </List>
}
