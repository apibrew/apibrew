import * as React from 'react'
import {
    CategoryOutlined,
    DatasetOutlined,
    ExtensionOutlined,
    PersonOutline,
    TableRowsOutlined
} from '@mui/icons-material'

export interface MenuList {
    title: string
    items: Array<MenuItem | DividerItem>
}

export interface MenuItem {
    title: string
    link: string
    icon?: JSX.Element
}

export interface DividerItem {
    divider: true
}

export const menuLists: MenuList[] = [
    {
        title: 'Develop',
        items: [
            {
                title: 'Resources',
                link: '/dashboard/resources',
                icon: <TableRowsOutlined/>
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
                divider: true
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
