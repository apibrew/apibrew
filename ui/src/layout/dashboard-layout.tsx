import { BaseLayout } from './base-layout'
import { BACKEND_URL } from '../config'
export interface DashboardLayoutProps {
    children: JSX.Element | JSX.Element[]
}

export function DashboardLayout(props: DashboardLayoutProps): JSX.Element {
    return <>
        <BaseLayout>
            <>
            Hello World {BACKEND_URL}
                { props.children}
            </>

        </BaseLayout>
    </>
}
