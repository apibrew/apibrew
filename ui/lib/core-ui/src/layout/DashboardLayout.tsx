import * as React from 'react'
import {Fragment, ReactNode} from 'react'
import Box from '@mui/material/Box'
import Container from "@mui/material/Container";
import {Drawer, IconButton, Menu, MenuItem, Toolbar, Tooltip} from "@mui/material";
import {Icon} from "../components/Icon.tsx";
import Divider from "@mui/material/Divider";
import {useRecordByName} from "../hooks/record.ts";
import {Menu as MenuRecord, MenuItem as RecordMenuItem} from "../model/ui/menu";
import {Loading} from "../components/basic/Loading.tsx";
import {Route, Routes, useLocation, useNavigate} from "react-router-dom";
import {useRoute} from "../hooks/route.ts";
import {SxProps} from "@mui/system";
import {Theme} from "@mui/material/styles";
import {DynamicComponent} from "../components/dynamic/DynamicComponent.tsx";

export interface DashboardLayoutProps {
    children: ReactNode
}

const DashboardLayoutSettings = {
    drawer: {
        width: 60,
        color: 'rgb(16,16,19)',
        textColor: 'rgb(200,200,200)',
        mainMenuItemSpacing: 1.5,
        mainMenuItemColor: 'rgb(156,156,159)',
        mainMenuItemActiveColor: 'rgb(226,226,229)',
        mainMenuItemActiveWeight: 900,
        mainMenuItemSize: 24,
    },
    secondDrawer: {
        width: 200,
        color: 'rgb(31,31,35)',
        textColor: 'rgb(200,200,200)',
        mainMenuItemSpacing: 1,
        mainMenuItemColor: 'rgb(156,156,159)',
        mainMenuItemActiveColor: 'rgb(226,226,229)',
        mainMenuItemActiveWeight: 600,
        mainMenuItemSize: 24,
    },
}

function normalizePath(parent: string, path: string) {
    let result = path

    if (!result.startsWith('/')) {
        if (parent.endsWith('/')) {
            result = parent + result
        } else {
            result = parent + '/' + result
        }
    }

    if (result.endsWith('/')) {
        result = result.substring(0, result.length - 1)
    }

    return result
}

function pathInside(parent: string, path: string) {
    if (parent == path) {
        return true
    }

    if (parent.endsWith('/')) {
        return path.startsWith(parent)
    } else {
        return path.startsWith(parent + '/')
    }
}

function MainDrawer(props: { menu: MenuRecord }) {
    const navigate = useNavigate()
    const route = useRoute()

    return <Drawer variant='permanent'
                   PaperProps={{
                       sx: {
                           width: DashboardLayoutSettings.drawer.width,
                           backgroundColor: DashboardLayoutSettings.drawer.color,
                           color: DashboardLayoutSettings.drawer.textColor,
                       }
                   }}>
        <Box marginTop={1.5}/>
        <IconButton>
            <Icon color='white' name='webhook' size={30}/>
        </IconButton>
        <Box marginTop={0.5}/>
        <Divider style={{background: 'rgb(50,50,50)'}}/>
        <Box marginTop={5}/>
        {props.menu.children.map((item, index) => {
            const active = pathInside(normalizePath(route.path, item.link), route.fullPath)

            return <Tooltip key={item.title} title={item.title} placement='right'>
                <IconButton sx={{
                    marginTop: DashboardLayoutSettings.drawer.mainMenuItemSpacing,
                    backgroundColor: active ? 'rgba(30, 50, 30, 0.5)' : 'transparent',
                    borderRadius: '3px',
                    "&:hover": {
                        backgroundColor: 'rgba(30, 50, 30, 0.5)',
                    },
                }} onClick={() => {
                    navigate(item.link)
                }}>
                    <Icon
                        color={active ? DashboardLayoutSettings.drawer.mainMenuItemActiveColor : DashboardLayoutSettings.drawer.mainMenuItemColor}
                        name={item.icon}
                        weight={active ? DashboardLayoutSettings.drawer.mainMenuItemActiveWeight : 400}
                        size={DashboardLayoutSettings.drawer.mainMenuItemSize}/>
                </IconButton>
            </Tooltip>
        })}

        <Box flexGrow={1}/>
        <Box marginTop={5}/>
        <Tooltip title={'User Pages'} placement='right'>
            <IconButton>
                <Icon color='white' name='person' size={24}/>
            </IconButton>
        </Tooltip>
        <Box marginBottom={3}/>
    </Drawer>;
}

function SecondDrawer(props: { menu: MenuRecord }) {
    const navigate = useNavigate()
    const route = useRoute()

    let activeMenuItem = props.menu.children.find(item => pathInside(normalizePath(route.path, item.link), route.fullPath))

    if (!activeMenuItem) {
        activeMenuItem = props.menu.children[0]
    }

    const activeMenuPath = normalizePath(route.path, activeMenuItem.link)

    activeMenuItem.children.forEach((item, index) => {
        if (index == 0 && activeMenuPath == route.fullPath && item.link) {
            setTimeout(() => {
                navigate(normalizePath(activeMenuPath, item.link))
            })
        }
    })

    function prepareMenuItemSx(active: boolean, item: RecordMenuItem) {
        const sx: SxProps<Theme> = {
            "&:hover": {
                backgroundColor: 'green',
            },
            borderRadius: '5px',
            marginBottom: DashboardLayoutSettings.secondDrawer.mainMenuItemSpacing,
            backgroundColor: 'transparent'
        }

        let itemSx = {...sx}

        if (active) {
            itemSx.backgroundColor = 'green'
        }

        if (item.children) {
            itemSx["&:hover"] = {
                backgroundColor: 'rgba(0, 250, 0, 0.1)'
            }
            itemSx.backgroundColor = 'rgba(0, 250, 0, 0.1)'
        }

        return itemSx
    }

    return <Drawer variant='persistent'
                   open={true}
                   PaperProps={{
                       sx: {
                           marginLeft: DashboardLayoutSettings.drawer.width + 'px',
                           width: DashboardLayoutSettings.secondDrawer.width,
                           backgroundColor: DashboardLayoutSettings.secondDrawer.color,
                           color: DashboardLayoutSettings.secondDrawer.textColor,
                       }
                   }}>
        <Box marginTop={0.7}/>
        <Box paddingLeft={0.9} paddingRight={2}>
            <h4>
                {activeMenuItem.title}
            </h4>
            {activeMenuItem.children.map((item, index) => {
                const menuItemPath = normalizePath(activeMenuPath, item.link)
                let active = pathInside(route.fullPath, menuItemPath)

                const itemSx = prepareMenuItemSx(active, item)

                return <Fragment key={item.link}>
                    <MenuItem sx={itemSx} onClick={() => {
                        if (item.children) {
                            return
                        }
                        navigate(menuItemPath)
                    }}>
                        <span>{item.title}</span>
                    </MenuItem>
                    <Box paddingLeft={1.6}>
                        {item.children && item.children.map(subItem => {
                            const subMenuItemPath = normalizePath(normalizePath(activeMenuPath, item.link), subItem.link)
                            let active = pathInside(route.fullPath, subMenuItemPath)

                            const subItemSx = prepareMenuItemSx(active, subItem)

                            return <MenuItem key={subItem.link} sx={subItemSx} onClick={() => {
                                navigate(subMenuItemPath)
                            }}>
                                <span>{subItem.title}</span>
                            </MenuItem>
                        })}
                    </Box>
                </Fragment>
            })}
        </Box>
    </Drawer>;
}

function DashboardPage(props: { menuItem: RecordMenuItem }) {
    return <DynamicComponent component={props.menuItem.component}
                             componentProps={props.menuItem.params}/>
}

function DashboardRoutes(props: { menu: MenuRecord }) {
    return <>
        <Routes>
            {props.menu.children.map(item => (
                <Route key={item.link} path={item.link + '/*'}
                       element={<Routes>
                           {item.children.map(subItem => (
                               <Route key={subItem.link}
                                      path={subItem.link + '/*'}
                                      element={
                                          subItem.children ? <Routes>
                                              {subItem.children.map(subSubItem => (
                                                  <Route key={subSubItem.link}
                                                         path={subSubItem.link + '/*'}
                                                         element={<DashboardPage menuItem={subSubItem}/>}/>
                                              ))}
                                          </Routes> : <DashboardPage menuItem={subItem}/>
                                      }/>
                           ))}

                       </Routes>}></Route>
            ))}
        </Routes>
    </>;
}

export function DashboardLayout(props: DashboardLayoutProps): React.JSX.Element {
    const menu = useRecordByName<MenuRecord>('Menu', 'ui', 'main')

    if (!menu) {
        return <Loading/>
    }

    return <Container maxWidth={false} disableGutters
                      sx={{background: 'rgb(210,240,240)', height: '100vh', display: 'flex'}}>
        <MainDrawer menu={menu}/>
        <SecondDrawer menu={menu}/>
        <Box marginLeft={(DashboardLayoutSettings.drawer.width + DashboardLayoutSettings.secondDrawer.width) + 'px'}
             flexGrow={1} height='100%' width='100%'>
            <DashboardRoutes menu={menu}/>
        </Box>
    </Container>
}
