import {BrowserRouter as Router, Route, Routes, useParams} from 'react-router-dom'
import React, {JSX} from 'react'
import {BaseLayout} from "./layout";
import {useRecordByName} from "./hooks/record.ts";
import {Route as RouteItem, Router as RouterModel, RouterName} from "./model/ui/router.ts";
import {DynamicComponent} from "./components/dynamic/DynamicComponent.tsx";

export interface RouteElementComponentProps {
    route: RouteItem
}

export function RouteElementComponent(props: RouteElementComponentProps) {
    const params = useParams()

    const route = props.route

    const componentProps = {
        ...params,
        ...route.params,
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

export function App(): JSX.Element {
    const router = useRecordByName<RouterModel>(RouterName, 'ui', 'main')

    return (
        <BaseLayout>
            <Router>
                {router && <RouterComponent routes={router.routes}/>}
            </Router>
        </BaseLayout>
    )
}

