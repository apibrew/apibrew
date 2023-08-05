import React, {JSX, ReactNode, useEffect} from "react";
import {ResourceContext as ResourceCtx} from "../../context/resource.ts";
import {ResourceService} from "@apibrew/ui-lib";
import { Resource } from "@apibrew/client";

export interface ResourceContextComponentProps {
    resource$?: Resource
    namespace?: string
    resource?: string
    children?: ReactNode
}

export function ResourceContextComponent(props: ResourceContextComponentProps): JSX.Element {
    const [resource, setResource] = React.useState<Resource | undefined>(props.resource$)

    useEffect(() => {
        if (resource) {
            return
        }

        if (!props.namespace || !props.resource) {
            throw new Error('ResourceContext: namespace and resource must be provided')
        }

        ResourceService.getByName(props.resource!, props.namespace).then((resource) => {
            setResource(resource)
        })

    }, [props.namespace, props.resource])

    return <>
        {resource && <ResourceCtx.Provider value={resource}>
            {props.children}
        </ResourceCtx.Provider>}
    </>
}
