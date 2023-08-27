import { Module } from "./module";
import { FunctionExecutionEngine } from "./function-execution-engine";
export declare const FunctionResource: {
    resource: string;
    namespace: string;
};
export interface Options {
    namedParams?: boolean;
}
export interface Argument {
    name: string;
    label?: string;
}
export interface Function {
    id: string;
    package: string;
    name: string;
    script?: string;
    module?: Module;
    engine: FunctionExecutionEngine;
    options?: Options;
    args?: Argument[];
    version: number;
}
export declare const FunctionName = "Function";
export declare const FunctionIdName = "Id";
export declare const FunctionPackageName = "Package";
export declare const FunctionNameName = "Name";
export declare const FunctionScriptName = "Script";
export declare const FunctionModuleName = "Module";
export declare const FunctionEngineName = "Engine";
export declare const FunctionOptionsName = "Options";
export declare const FunctionArgsName = "Args";
export declare const FunctionVersionName = "Version";
