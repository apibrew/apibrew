import * as React from 'react'
import {Fragment, ReactNode, useContext} from 'react'
import Box from '@mui/material/Box'
import Container from "@mui/material/Container";
import {Breadcrumbs, Drawer, IconButton, Menu, MenuItem, Toolbar, Tooltip} from "@mui/material";
import {Icon} from "../components/Icon.tsx";
import Divider from "@mui/material/Divider";
import {useRecordByName} from "../hooks/record.ts";
import {Menu as MenuRecord, MenuItem as RecordMenuItem} from "../model/ui/menu.ts";
import {Loading} from "../components/basic/Loading.tsx";
import {Link, Route, Routes, useLocation, useNavigate} from "react-router-dom";
import {useRoute} from "../hooks/route.ts";
import {SxProps} from "@mui/system";
import {Theme} from "@mui/material/styles";
import {DynamicComponent, LayoutContext} from "@apibrew/ui-lib";
import NavigateNextIcon from "@mui/icons-material/NavigateNext";
import Typography from "@mui/material/Typography";

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

    if (!result) {
        result = ''
    }

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
                let active = pathInside(menuItemPath, route.fullPath)

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

function DashboardPage(props: { menuItem: RecordMenuItem, parents: RecordMenuItem[], selfLink: string }) {
    const layoutContext = useContext(LayoutContext)
    const route = useRoute()

    return <Box display='flex' flexDirection='column' sx={{height: '100%'}}>
        <Box sx={{
            height: '70px',
            background: 'rgb(255,255,255)',
            padding: '20px',
            borderBottom: '1px solid rgb(200,200,200)',
        }}>
            <Breadcrumbs aria-label="breadcrumb" separator={<NavigateNextIcon fontSize="small"/>}>
                <Typography color="text.primary">{props.parents[0].title}</Typography>
                {props.parents[1] && <Typography color="text.primary">{props.parents[1].title}</Typography>}
                <Link style={{
                    textDecoration: 'underline',
                    color: 'rgb(0, 0, 0)',
                }} to={props.selfLink} color="text.primary">{props.menuItem.title}</Link>

                {layoutContext.breadCramps && layoutContext.breadCramps.map((item, index) => {
                    if (item.link) {
                        return <Link key={item.link} to={item.link}/>
                    } else {
                        return <Typography key={index} color="text.primary">{item.title}</Typography>
                    }
                })}

            </Breadcrumbs>
        </Box>
        <Box flexGrow={1} display='flex'>
            <DynamicComponent component={props.menuItem.component}
                              componentProps={props.menuItem.params}/>
        </Box>
    </Box>
}

function DashboardRoutes(props: { menu: MenuRecord }) {
    const route = useRoute()

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
                                                         element={<DashboardPage menuItem={subSubItem}
                                                                                 selfLink={route.path + item.link + '/' + subItem.link + '/' + subSubItem.link}
                                                                                 parents={[item, subItem]}/>}/>
                                              ))}
                                          </Routes> : <DashboardPage menuItem={subItem}
                                                                     selfLink={route.path + item.link + '/' + subItem.link}
                                                                     parents={[item]}/>
                                      }/>
                           ))}

                       </Routes>}></Route>
            ))}
        </Routes>
    </>;
}

export function DashboardLayout(props: DashboardLayoutProps): React.JSX.Element {
    const route = useRoute()
    const navigate = useNavigate()

    if (route.fullPath.endsWith('/')) {
        navigate(route.fullPath.slice(0, -1))
        return <Loading/>
    }

    const menu = useRecordByName<MenuRecord>('Menu', 'ui', 'main')

    if (!menu) {
        return <Loading/>
    }

    return <Container maxWidth={false} disableGutters
                      sx={{background: 'rgb(251,251,251)', height: '100vh', display: 'flex'}}>
        <MainDrawer menu={menu}/>
        <SecondDrawer menu={menu}/>
        <Box paddingLeft={(DashboardLayoutSettings.drawer.width + DashboardLayoutSettings.secondDrawer.width) + 'px'}
             flexGrow={1} height='100%' width='100%'>
            <DashboardRoutes menu={menu}/>
        </Box>
    </Container>
}
