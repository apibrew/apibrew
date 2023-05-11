import { PageLayout } from '../../layout/PageLayout'
import { Designer } from '../../components/designer/Designer'

export function ResourcesDesigner(): JSX.Element {
    return <PageLayout pageTitle={'Resources Designer'}>
        <Designer/>
    </PageLayout>
}
