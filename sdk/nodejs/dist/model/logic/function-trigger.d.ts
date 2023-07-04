import { Function } from "./function";
export declare const FunctionTriggerResource: {
    resource: string;
    namespace: string;
};
export interface FunctionTrigger {
    id: string;
    name: string;
    resource: string;
    namespace: string;
    action: 'create' | 'update' | 'delete' | 'list' | 'get';
    order?: 'before' | 'after' | 'instead';
    async: boolean;
    function: Function;
    version: number;
}
export declare const FunctionTriggerName = "FunctionTrigger";
export declare const FunctionTriggerIdName = "Id";
export declare const FunctionTriggerNameName = "Name";
export declare const FunctionTriggerResourceName = "Resource";
export declare const FunctionTriggerNamespaceName = "Namespace";
export declare const FunctionTriggerActionName = "Action";
export declare const FunctionTriggerOrderName = "Order";
export declare const FunctionTriggerAsyncName = "Async";
export declare const FunctionTriggerFunctionName = "Function";
export declare const FunctionTriggerVersionName = "Version";
