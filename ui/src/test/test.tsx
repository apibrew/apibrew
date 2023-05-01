import { BACKEND_URL } from '../config'
import { DashboardLayout } from '../layout/dashboard-layout'

export function Test(): JSX.Element {
    return (
        <>
            Hello World {BACKEND_URL}
        </>
    )
}
