import {BrowserRouter as Router, Route, Routes, useNavigate, useParams} from 'react-router-dom'
import React, {JSX, useEffect} from 'react'
import {BaseLayout} from "./layout";
import {useRecordByName} from "./hooks/record.ts";
import {Route as RouteItem, Router as RouterModel, RouterName} from "./model/ui/router.ts";
import {DynamicComponent} from "./components/dynamic/DynamicComponent.tsx";
import {Loading} from "./components/basic/Loading.tsx";

export interface RouteElementComponentProps {
    route: RouteItem
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
            <RouterComponent routes={route.routes}/>
        </DynamicComponent>
    } else {
        return <DynamicComponent component={route.component} componentProps={componentProps}></DynamicComponent>
    }
}

export interface RouterComponentProps {
    routes: RouteItem[]
}

function RouterComponent(props: RouterComponentProps) {
    return <Routes>
        {props.routes.map(route => {
            return <Route key={route.path} path={route.path} element={<RouteElementComponent route={route}/>}></Route>
        })}
    </Routes>
}

function RouterComponentWithRouterName(props: {routerName: string}) {
    const router = useRecordByName<RouterModel>(RouterName, 'ui', props.routerName)

    if (!router) {
        return <Loading/>
    }

    return <RouterComponent routes={router.routes}/>
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

