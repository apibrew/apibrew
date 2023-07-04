import { Function } from "./function";
export declare const FunctionExecutionResource: {
    resource: string;
    namespace: string;
};
export interface FunctionExecution {
    id: string;
    function: Function;
    input?: object;
    output?: object;
    error?: object;
    status?: 'pending' | 'success' | 'error';
    version: number;
}
export declare const FunctionExecutionName = "FunctionExecution";
export declare const FunctionExecutionIdName = "Id";
export declare const FunctionExecutionFunctionName = "Function";
export declare const FunctionExecutionInputName = "Input";
export declare const FunctionExecutionOutputName = "Output";
export declare const FunctionExecutionErrorName = "Error";
export declare const FunctionExecutionStatusName = "Status";
export declare const FunctionExecutionVersionName = "Version";
