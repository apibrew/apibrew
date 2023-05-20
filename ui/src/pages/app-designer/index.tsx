import { PageLayout } from '../../layout/PageLayout'
import { Designer } from '../../components/designer/Designer'

export function AppDesigner(): JSX.Element {
    return <PageLayout pageTitle={'App Designer'}>
        <Designer name='Default'/>
    </PageLayout>
}
