import { Box, Button, Drawer } from "@mui/material"
import { PageLayout } from "../../layout/PageLayout"
import { Api, Delete, Edit, PlusOneOutlined, Search, Title } from "@mui/icons-material"
import { Resource } from "../../model"
import { useNavigate } from "react-router-dom"
import { DataGrid, GridColDef, GridRowParams, GridActionsCellItem, GridValueGetterParams } from '@mui/x-data-grid';
import { Record, RecordService } from "../../service/record"
import React, { useEffect, useState } from "react"
import { SdkDrawer } from "../sdk/SdkDrawer"
import {Crud} from "../../model/schema";

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

    useEffect(() => {
        load();
    }, [props.resource])

    const columns: GridColDef[] = props.resource.properties.map((property) => {
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

    columns.push({
        field: 'actions',
        type: 'actions',
        width: 150,
        headerName: 'Actions',
        getActions: (params: GridRowParams) => [
            <GridActionsCellItem label='Edit' icon={<Edit />} onClick={() => {
                navigate(`${params.id}/edit`)
            }} />,
            <GridActionsCellItem label='Delete' icon={<Delete />} onClick={() => {
                RecordService.remove(props.resource.namespace ?? 'default', props.resource.name, params.id as string).then(() => {
                    load()
                })
            }} />,
            <GridActionsCellItem label='Delete' icon={<Search />} onClick={() => {
                navigate(`${params.id}/view`)
            }} />,
        ],
    })

    const rows = list;

    return (
        <PageLayout pageTitle={props.resource.name} actions={<React.Fragment>
            <Button variant={'contained'} color='success' onClick={() => {
                navigate('new')
            }} startIcon={<PlusOneOutlined />}>New {props.resource.name}</Button>
            <Button variant={'contained'} color='primary' onClick={() => {
                setShowSdk(true)
            }} startIcon={<Api />}>sdk</Button>
            <Button variant={'contained'} color='secondary' onClick={() => {
                navigate('settings')
            }} startIcon={<Api />}>Crud Settings</Button>
        </React.Fragment>}>
            <SdkDrawer resource={props.resource} open={showSdk} onClose={() => { setShowSdk(false) }} />
            <div>
                Hello world3!
            </div>
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
