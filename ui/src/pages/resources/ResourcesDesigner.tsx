import {PageLayout} from '../../layout/PageLayout'
import {Designer} from "../../components/designer/Designer";

export interface ResourcesDesignerProps {

}

export function ResourcesDesigner(props: ResourcesDesignerProps): JSX.Element {
    return <PageLayout pageTitle={'Resources Designer'}>
        <Designer/>
    </PageLayout>
}
