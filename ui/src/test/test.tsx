import { BACKEND_URL } from '../config'
import { DashboardLayout } from '../layout/dashboard-layout'
import { PageLayout } from '../layout/PageLayout'

export function Test(): JSX.Element {
    return (
        <DashboardLayout>
            <PageLayout pageTitle={'Test Page'}>
                <>
                    Hello World {BACKEND_URL}
                </>
            </PageLayout>
        </DashboardLayout>
    )
}
