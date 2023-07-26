import {useEffect, useState} from "react";
import {ResourceService} from "../service";
import {useErrorHandler} from "./error-handler.tsx";
import { Resource } from "@apibrew/client";

export function useResourceByName(resourceName: string, namespace = 'default'): Resource | undefined {
    const errorHandler = useErrorHandler()
    const [resource, setResource] = useState<Resource>()

    useEffect(() => {
        ResourceService.getByName(resourceName, namespace).then(setResource, errorHandler)
    }, [resourceName, namespace])

    return resource
}