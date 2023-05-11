import {Link} from "./Link";
import React from "react";
import {Resource, ResourceProperty} from "../../model";

export interface ReferenceLinkProps {
    resource: Resource
    property: ResourceProperty
}

export function ReferenceLink(props: ReferenceLinkProps) {
    return (
        <Link sourceSelector={`.resource-${props.resource.name} .resource-property-${[props.property.name]} .right-ref`}
              targetSelector={`.resource-${props.property.reference!.referencedResource} .resource-head`}/>
    )
}
