import {Function} from "./function";


// Sub Types

// Resource Type
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

// Resource and Property Names
export const FunctionTriggerName = "FunctionTrigger";

export const FunctionTriggerIdName = "Id";

export const FunctionTriggerNameName = "Name";

export const FunctionTriggerResourceName = "Resource";

export const FunctionTriggerNamespaceName = "Namespace";

export const FunctionTriggerActionName = "Action";

export const FunctionTriggerOrderName = "Order";

export const FunctionTriggerAsyncName = "Async";

export const FunctionTriggerFunctionName = "Function";

export const FunctionTriggerVersionName = "Version";
