

import { Function } from "./function";


export const FunctionExecutionResource = {
    resource: "FunctionExecution",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface FunctionExecution {
    id: string;
function: Function;
input?: object;
output?: object;
error?: object;
status?: 'pending' | 'success' | 'error';
annotations?: object;
version: number;

}
// Resource and Property Names
export const FunctionExecutionName = "FunctionExecution";

export const FunctionExecutionIdName = "Id";

export const FunctionExecutionFunctionName = "Function";

export const FunctionExecutionInputName = "Input";

export const FunctionExecutionOutputName = "Output";

export const FunctionExecutionErrorName = "Error";

export const FunctionExecutionStatusName = "Status";

export const FunctionExecutionAnnotationsName = "Annotations";

export const FunctionExecutionVersionName = "Version";


