import { Box, Button, Drawer } from "@mui/material"
import { PageLayout } from "../../layout/PageLayout"
import { Delete, Edit, PlusOneOutlined, Search } from "@mui/icons-material"
import { Resource } from "../../model"
import { useNavigate } from "react-router-dom"
import { DataGrid, GridColDef, GridRowParams, GridActionsCellItem, GridValueGetterParams } from '@mui/x-data-grid';
import { Record, RecordService } from "../../service/record"
import { useEffect, useState } from "react"

export interface ListProps {
    resource: Resource
}

export function List(props: ListProps) {
    const navigate = useNavigate()
    const [list, setList] = useState<Record[]>([])

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

    const [showSdk, setShowSdk] = useState(false)

    return (
        <PageLayout pageTitle={props.resource.name} actions={<>
            <Button variant={'contained'} color='success' onClick={() => {
                navigate('new')
            }} startIcon={<PlusOneOutlined />}>New Item</Button>
            <Button variant={'contained'} color='success' onClick={() => {
                setShowSdk(!showSdk)
            }} startIcon={<PlusOneOutlined />}>sdk</Button>
        </>}>
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
            <Drawer anchor={'right'}
                BackdropProps={{ invisible: true }}
                ModalProps={{ sx: { '& .MuiDrawer-paper': { width: '600px', top: '88px' } } }}
                onClose={() => {
                    setShowSdk(false)
                }}
                open={showSdk}

            >
                <Box>
                    <h1>asdadas</h1>
                </Box>
            </Drawer>
        </PageLayout>
    )
}