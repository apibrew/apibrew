import { Link } from './Link'
import { type Resource, type Property } from '@apibrew/client'

export interface ReferenceLinkProps {
    resource: Resource
    property: Property
}

export function ReferenceLink(props: ReferenceLinkProps) {
    return (
        <Link sourceSelector={`.resource-${props.resource.name} .resource-property-${props.property.name} .right-ref`}
            targetSelector={`.resource-${props.property.reference?.resource?.name ?? ''} .resource-head`} />
    )
}
