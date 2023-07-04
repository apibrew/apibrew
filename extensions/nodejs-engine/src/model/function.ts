

import { FunctionExecutionEngine } from "./function-execution-engine";

import { Module } from "./module";


export const FunctionResource = {
    resource: "Function",
    namespace: "logic",
};

// Sub Types

export interface Options {
     namedParams?: boolean;

}

export interface Argument {
     name: string;
     label?: string;

}

// Resource Type
export interface Function {
    id: string;
package: string;
name: string;
script?: string;
engine: FunctionExecutionEngine;
options?: Options;
version: number;
args?: Argument[];
module?: Module;

}
// Resource and Property Names
export const FunctionName = "Function";

export const FunctionIdName = "Id";

export const FunctionPackageName = "Package";

export const FunctionNameName = "Name";

export const FunctionScriptName = "Script";

export const FunctionEngineName = "Engine";

export const FunctionOptionsName = "Options";

export const FunctionVersionName = "Version";

export const FunctionArgsName = "Args";

export const FunctionModuleName = "Module";


