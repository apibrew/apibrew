import {Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Api, Delete, Edit, PlusOneOutlined, Search} from "@mui/icons-material"
import {Resource, ResourceProperty} from "../../model"
import {useNavigate} from "react-router-dom"
import {DataGrid, GridActionsCellItem, GridColDef, GridRowParams, GridValueGetterParams} from '@mui/x-data-grid';
import {Record, RecordService} from "../../service/record"
import React, {useEffect, useState} from "react"
import {SdkDrawer} from "../sdk/SdkDrawer"
import {Crud} from "../../model/ui/crud.ts";

export interface ListProps {
    resource: Resource
    crudConfig: Crud
}

export function List(props: ListProps) {
    const navigate = useNavigate()
    const [list, setList] = useState<Record[]>([])
    const [showSdk, setShowSdk] = useState(false)

    const load = () => {
        RecordService.list<Record>(props.resource.namespace ?? 'default', props.resource.name).then((data) => {
            setList(data)
        })
    }

    const resourcePropertyMap = new Map<string, ResourceProperty>()

    props.resource.properties.forEach((property) => {
        resourcePropertyMap.set(property.name, property)
    })

    useEffect(() => {
        load();
    }, [props.resource])

    let columns: GridColDef[] = props.resource.properties.map((property) => {
        return {
            field: property.name,
            type: 'string',
            headerName: property.name,
            width: 150,
            editable: false,
            valueGetter: (params: GridValueGetterParams<any, any>) => {
                const prop = props.resource.properties.find((p) => p.name === property.name)!

                if (prop.type === 'REFERENCE' && params.row[property.name]) {
                    return params.row[property.name].name
                }

                return params.row[property.name]
            }
        } as GridColDef
    })

    if (props.crudConfig?.gridConfig?.columns) {
        columns = props.crudConfig.gridConfig.columns.map((column) => {
            return {
                field: column.name,
                type: column.type ?? 'string',
                headerName: column.title,
                width: column.width ?? 150,
                editable: false,
                valueGetter: (params: GridValueGetterParams<any, any>) => {
                    const prop = resourcePropertyMap.get(column.name)

                    if (prop === undefined) {
                        return 'Unknown column!'
                    }

                    if (prop.type === 'REFERENCE' && params.row[column.name]) {
                        return params.row[column.name].name
                    }

                    return params.row[column.name]
                }
            } as GridColDef
        })
    }

    columns.push({
        field: 'actions',
        type: 'actions',
        width: 150,
        headerName: 'Actions',
        getActions: (params: GridRowParams) => [
            <GridActionsCellItem label='Edit' icon={<Edit/>} onClick={() => {
                navigate(`${params.id}/edit`)
            }}/>,
            <GridActionsCellItem label='Delete' icon={<Delete/>} onClick={() => {
                RecordService.remove(props.resource.namespace ?? 'default', props.resource.name, params.id as string).then(() => {
                    load()
                })
            }}/>,
            <GridActionsCellItem label='Delete' icon={<Search/>} onClick={() => {
                navigate(`${params.id}/view`)
            }}/>,
        ],
    })

    const rows = list;

    return (
        <PageLayout pageTitle={props.resource.name} actions={<React.Fragment>
            <Button variant={'contained'} color='success' onClick={() => {
                navigate('new')
            }} startIcon={<PlusOneOutlined/>}>New {props.resource.name}</Button>
            <Button variant={'contained'} color='primary' onClick={() => {
                setShowSdk(true)
            }} startIcon={<Api/>}>sdk</Button>
            <Button variant={'contained'} color='secondary' onClick={() => {
                navigate('settings')
            }} startIcon={<Api/>}>Crud Settings</Button>
        </React.Fragment>}>
            <SdkDrawer resource={props.resource} open={showSdk} onClose={() => {
                setShowSdk(false)
            }}/>
            <DataGrid
                rows={rows}
                columns={columns}
                initialState={{
                    pagination: {
                        paginationModel: {
                            pageSize: 5,
                        },
                    },
                }}
                pageSizeOptions={[5]}
                disableRowSelectionOnClick
            />

        </PageLayout>
    )
}
