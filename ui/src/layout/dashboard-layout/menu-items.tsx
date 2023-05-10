import * as React from 'react'
import {
    CategoryOutlined,
    DatasetOutlined,
    ExtensionOutlined,
    PersonOutline,
    TableRowsOutlined
} from '@mui/icons-material'

export interface MenuList {
    title?: string
    items: MenuItem[]
}

export interface MenuItem {
    title: string
    link?: string
    icon?: React.ReactNode
    children?: MenuItem[]
}

export const menuLists: MenuList[] = [
    {
        title: 'User',
        items: [
            {
                title: 'Country',
                link: '/dashboard/country',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'City',
                link: '/dashboard/city',
                icon: <TableRowsOutlined/>
            }
        ]
    },
    {
        title: 'Develop',
        items: [
            {
                icon: <TableRowsOutlined/>,
                title: 'App Designer',
                link: '/dashboard/resources/designer'
            },
            {
                title: 'Resources',
                icon: <TableRowsOutlined/>,
                link: '/dashboard/resources',
            },
            {
                title: 'Namespaces',
                link: '/dashboard/namespaces',
                icon: <CategoryOutlined/>
            },
            {
                title: 'Data Sources',
                link: '/inbox',
                icon: <DatasetOutlined/>
            },
            {
                title: 'Users',
                link: '/dashboard/users',
                icon: <PersonOutline/>
            },
            {
                title: 'Records',
                link: '/dashboard/records',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Logic',
                icon: <ExtensionOutlined/>,
                children: [
                    {
                        title: 'Functions',
                        link: '/dashboard/functions',
                        icon: <TableRowsOutlined/>
                    },
                    {
                        title: 'Triggers',
                        link: '/dashboard/triggers',
                        icon: <TableRowsOutlined/>
                    }, {
                        title: 'Jobs',
                        link: '/dashboard/jobs',
                        icon: <TableRowsOutlined/>
                    }, {
                        title: 'Schedules',
                        link: '/dashboard/schedules',
                        icon: <TableRowsOutlined/>
                    }, {
                        title: 'Events',
                        link: '/dashboard/events',
                        icon: <TableRowsOutlined/>
                    }, {
                        title: 'Rules',
                        link: '/dashboard/rules',
                        icon: <TableRowsOutlined/>
                    }
                ]
            }
        ]
    },
]
