export interface ServiceConfig {
    backendUrl: string;
    token: string;
}
export type ServiceConfigProvider = () => ServiceConfig;
