import { useEffect, useState } from "react";

import { useErrorHandler } from "./error-handler.tsx";
import { RecordApi, Resource } from "@apibrew/client";
import { ServiceConfig } from "@apibrew/client/dist/service/config";
import { BACKEND_URL } from "../config.ts";
import { TokenService } from "@apibrew/ui-lib";

export function useResourceByName(resourceName: string, namespace = 'default'): Resource | undefined {
    const errorHandler = useErrorHandler()
    const [resource, setResource] = useState<Resource>()

    useEffect(() => {
        const config = {
            backendUrl: BACKEND_URL,
            token: TokenService.get(),
        } as ServiceConfig
        RecordApi.resource(config, namespace, resourceName).then(setResource, errorHandler)
    }, [resourceName, namespace])

    return resource
}