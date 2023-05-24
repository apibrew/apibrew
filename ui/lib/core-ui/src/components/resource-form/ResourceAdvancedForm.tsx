import { type Resource } from '../../model'

export interface ResourceAdvancedProps {
    resource: Resource
    onChange: (resource: Resource) => void
}

export function ResourceAdvancedForm(props: ResourceAdvancedProps): JSX.Element {
    return <>
        <div>Resource Advanced Form</div>
    </>
}
