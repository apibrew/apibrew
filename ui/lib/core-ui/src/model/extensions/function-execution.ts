import {Function} from "./function";


// Sub Types

// Resource Type
export interface FunctionExecution {
    error?: object;
    status?: string;
    id: string;
    function: Function;
    input?: object;
    output?: object;
    version: number;

}

// Resource and Property Names
export const FunctionExecutionName = "FunctionExecution";

export const FunctionExecutionErrorName = "Error";

export const FunctionExecutionStatusName = "Status";

export const FunctionExecutionIdName = "Id";

export const FunctionExecutionFunctionName = "Function";

export const FunctionExecutionInputName = "Input";

export const FunctionExecutionOutputName = "Output";

export const FunctionExecutionVersionName = "Version";


