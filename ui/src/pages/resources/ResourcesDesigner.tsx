import {PageLayout} from '../../layout/PageLayout'
import {Designer} from "../../components/designer/Designer2";

export interface ResourcesDesignerProps {

}

export function ResourcesDesigner(props: ResourcesDesignerProps): JSX.Element {
    return <PageLayout pageTitle={'Resources Designer'}>
        <Designer/>
    </PageLayout>
}
