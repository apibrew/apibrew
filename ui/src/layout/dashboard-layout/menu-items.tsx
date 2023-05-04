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
        title: 'Develop',
        items: [
            {
                title: 'Resources',
                icon: <TableRowsOutlined/>,
                children: [
                    {
                        icon: <TableRowsOutlined/>,
                        title: 'Designer',
                        link: '/dashboard/resources/designer'
                    },
                    {
                        icon: <TableRowsOutlined/>,
                        title: 'List',
                        link: '/dashboard/resources'
                    }
                ]
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
                title: 'Extensions',
                link: '/dashboard/extensions',
                icon: <ExtensionOutlined/>
            },
            {
                title: 'Records',
                link: '/dashboard/records',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Action Designer',
                link: '/dashboard/action-designer',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Python Extension',
                link: '/dashboard/action-designer',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Nodejs Extension',
                link: '/dashboard/action-designer',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Golang Extension',
                link: '/dashboard/action-designer',
                icon: <TableRowsOutlined/>
            },
            {
                title: 'Java Extension',
                link: '/dashboard/action-designer',
                icon: <TableRowsOutlined/>
            }
        ]
    },
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
    }
]
