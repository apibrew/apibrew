import {useEffect, useState} from "react";
import {Resource} from "../model";
import {ResourceService} from "../service";

export function useResourceByName(resourceName: string, namespace = 'default'): Resource | undefined {
    const [resource, setResource] = useState<Resource>()

    useEffect(() => {
        ResourceService.getByName(resourceName, namespace).then(setResource)
    }, [resourceName, namespace])

    return resource
}