import {FunctionExecutionEngine} from "./function-execution-engine";


// Sub Types

export interface Options {
    namedParams: boolean;

}

export interface Argument {
    name: string;
    direction: string;

}

// Resource Type
export interface Function {
    options: Options;
    args: Argument[];
    id: string;
    version: number;
    package: string;
    name: string;
    engine: FunctionExecutionEngine;
    script: string;

}

// Resource and Property Names
export const FunctionName = "Function";

export const FunctionOptionsName = "Options";

export const FunctionArgsName = "Args";

export const FunctionIdName = "Id";

export const FunctionVersionName = "Version";

export const FunctionPackageName = "Package";

export const FunctionNameName = "Name";

export const FunctionEngineName = "Engine";

export const FunctionScriptName = "Script";


