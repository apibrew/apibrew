import { FunctionExecutionEngine } from "./function-execution-engine";
export declare const ModuleResource: {
    resource: string;
    namespace: string;
};
export interface Module {
    id: string;
    package: string;
    content: string;
    engine: FunctionExecutionEngine;
    version: number;
}
export declare const ModuleName = "Module";
export declare const ModuleIdName = "Id";
export declare const ModulePackageName = "Package";
export declare const ModuleContentName = "Content";
export declare const ModuleEngineName = "Engine";
export declare const ModuleVersionName = "Version";
