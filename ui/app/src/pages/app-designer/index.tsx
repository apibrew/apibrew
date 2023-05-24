import { PageLayout, Designer } from 'core-ui'

export function AppDesigner(): JSX.Element {
    return <PageLayout pageTitle={'App Designer'}>
        <Designer name='Default'/>
    </PageLayout>
}
