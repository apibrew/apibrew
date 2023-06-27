

import { FunctionExecutionEngine } from "./function-execution-engine";


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
script: string;
engine: FunctionExecutionEngine;
options?: Options;
args?: Argument[];
version: number;
startFunction?: string

}
// Resource and Property Names
export const FunctionName = "Function";

export const FunctionIdName = "Id";

export const FunctionPackageName = "Package";

export const FunctionNameName = "Name";

export const FunctionScriptName = "Script";

export const FunctionEngineName = "Engine";

export const FunctionOptionsName = "Options";

export const FunctionArgsName = "Args";

export const FunctionVersionName = "Version";


