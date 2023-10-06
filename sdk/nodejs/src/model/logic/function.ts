

import { Module } from "./module";

import { FunctionExecutionEngine } from "./function-execution-engine";


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
module?: Module;
engine: FunctionExecutionEngine;
options?: Options;
args?: Argument[];
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const FunctionName = "Function";

export const FunctionIdName = "Id";

export const FunctionPackageName = "Package";

export const FunctionNameName = "Name";

export const FunctionScriptName = "Script";

export const FunctionModuleName = "Module";

export const FunctionEngineName = "Engine";

export const FunctionOptionsName = "Options";

export const FunctionArgsName = "Args";

export const FunctionAnnotationsName = "Annotations";

export const FunctionCreatedByName = "CreatedBy";

export const FunctionUpdatedByName = "UpdatedBy";

export const FunctionCreatedOnName = "CreatedOn";

export const FunctionUpdatedOnName = "UpdatedOn";

export const FunctionVersionName = "Version";


