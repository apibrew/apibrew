import { ServiceConfig } from "@apibrew/client/dist/service/config";
import { BACKEND_URL } from "../config";
import { TokenService } from ".";

export const ServiceConfigProvider = () => {
    return {
        backendUrl: BACKEND_URL,
        token: TokenService.get()
    } as ServiceConfig
}

export const ServiceConfigProviderWithoutToken = () => {
    return {
        backendUrl: BACKEND_URL
    } as ServiceConfig
}
