import {BrowserRouter as Router, Route, Routes, useNavigate, useParams} from 'react-router-dom'
import React, {JSX, useEffect} from 'react'
import {BaseLayout} from "./layout";
import {useRecordByName} from "./hooks/record.ts";
import {Route as RouteItem, Router as RouterModel, RouterName} from "./model/ui/router.ts";
import {DynamicComponent} from "./components/dynamic/DynamicComponent.tsx";
import {Loading} from "./components/basic/Loading.tsx";
import {RouteContext} from "./context/route-context.ts";

export interface RouteElementComponentProps {
    route: RouteItem
    path: string
}

export function RouteElementComponent(props: RouteElementComponentProps) {
    const params = useParams()
    const navigate = useNavigate()

    const route = props.route

    const componentProps = {
        ...params,
        ...route.params,
    }

    useEffect(() => {
        if (route.component === 'Router/Forward') {
            navigate((route.params as any)?.to)
        }
    }, [route])

    if (route.component === 'Router/Forward') {
        return <Loading/>
    }

    if (route.routes) {
        return <DynamicComponent component={route.component} componentProps={componentProps}>
            <RouterComponent routes={route.routes} path={props.path}/>
        </DynamicComponent>
    } else {
        return <DynamicComponent component={route.component} componentProps={componentProps}></DynamicComponent>
    }
}

export interface RouterComponentProps {
    routes: RouteItem[]
    path: string
}

function RouterComponent(props: RouterComponentProps) {
    return <Routes>
        {props.routes.map(route => {
            let path = props.path + '/' + route.path

            if (props.path === '') {
                path = route.path
            }

            return <Route key={route.path} path={route.path} element={
                <RouteContext.Provider value={path}>
                    <RouteElementComponent route={route} path={path}/>
                </RouteContext.Provider>
            }></Route>
        })}
    </Routes>
}

function RouterComponentWithRouterName(props: { routerName: string }) {
    const router = useRecordByName<RouterModel>(RouterName, 'ui', props.routerName)

    if (!router) {
        return <Loading/>
    }

    return <RouterComponent routes={router.routes} path=''/>
}

export function App(): JSX.Element {
    return (
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path={'/login'} element={<DynamicComponent component={'CoreUI/LoginPage'}/>}></Route>
                    <Route path={'*'} element={<RouterComponentWithRouterName routerName={'main'}/>}></Route>
                </Routes>
            </Router>
        </BaseLayout>
    )
}

