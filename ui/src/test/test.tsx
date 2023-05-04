import { BACKEND_URL } from '../config'
import { PageLayout } from '../layout/PageLayout'
import Button from '@mui/material/Button'
import { PlusOneOutlined } from '@mui/icons-material'

export function Test(): JSX.Element {
    return (
        <PageLayout pageTitle={'Test Page'} actions={<>
            <Button variant={'contained'} color='success' startIcon={<PlusOneOutlined />}>New Item</Button>
        </>}>
            <>
                Hello World {BACKEND_URL}
            </>
        </PageLayout>
    )
}
