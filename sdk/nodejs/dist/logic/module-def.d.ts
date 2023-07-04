import { Module } from "../model";
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
export declare function setModule(_module: Module): Promise<void>;
export declare function getModule(): Module;
export declare function registerModuleChild(name: string, element: any): void;
export declare function getModuleChild(name: string): any;
