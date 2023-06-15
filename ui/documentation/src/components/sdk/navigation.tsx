import React from "react";

export interface NavigationComponentProps {
    name?: string
}

export type NavigationComponent = React.ComponentType<NavigationComponentProps>

export interface NavigationItem {
    title: string
    name: string
    component?: NavigationComponent
    children?: NavigationItem[]
}

export const navigationItems: NavigationItem[] = [
    {
        title: 'Authentication API',
        name: 'authentication',
        // component: <h1></h1>
        children: [
            {
                title: 'Authenticate',
                name: 'authenticate',
            },
            {
                title: 'Refresh',
                name: 'refresh',
            }
        ]
    },
    {
        title: 'User APIs',
        name: 'user',
        // component: React.lazy(() => import('../pages/user')),
        children: [
            {
                title: 'Country',
                name: 'country',
                // component: React.lazy(() => import('../pages/user/country')),
            },
            {
                title: 'City',
                name: 'city',
                // component: React.lazy(() => import('../pages/user/city')),
            },
        ]
    }
]