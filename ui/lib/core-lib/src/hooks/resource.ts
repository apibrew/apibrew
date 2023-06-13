import {useEffect, useState} from "react";
import {Resource} from "../model";
import {ResourceService} from "../service";
import {useErrorHandler} from "./error-handler.tsx";

export function useResourceByName(resourceName: string, namespace = 'default'): Resource | undefined {
    const errorHandler = useErrorHandler()
    const [resource, setResource] = useState<Resource>()

    useEffect(() => {
        ResourceService.getByName(resourceName, namespace).then(setResource, errorHandler)
    }, [resourceName, namespace])

    return resource
}