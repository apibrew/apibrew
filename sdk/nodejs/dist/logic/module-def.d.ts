export interface ModuleBackendParams {
    authentication?: {
        username: string;
        password: string;
    };
    backendUrl?: string;
}
export interface ModuleInitParams {
}
export interface ModuleInitFunctionParams {
    moduleId: string;
}
export type ModuleInitFunction = (params: ModuleInitFunctionParams) => Promise<void>;
export declare function setModuleId(_moduleId: string): Promise<void>;
export declare function getModuleId(): string;
export declare function registerModuleChild(name: string, element: any): void;
export declare function getModuleChild(name: string): any;
