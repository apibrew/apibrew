import { Resource } from "./resource";
export declare const ExtensionResource: {
    resource: string;
    namespace: string;
};
export interface BooleanExpression {
}
export interface FunctionCall {
    host: string;
    functionName: string;
}
export interface HttpCall {
    uri: string;
    method: string;
}
export interface ExternalCall {
    functionCall?: FunctionCall;
    httpCall?: HttpCall;
}
export interface EventSelector {
    actions?: 'CREATE' | 'UPDATE' | 'DELETE' | 'GET' | 'LIST' | 'OPERATE'[];
    recordSelector?: BooleanExpression;
    namespaces?: string[];
    resources?: string[];
    ids?: string[];
    annotations?: object;
}
export interface RecordSearchParams {
    query?: BooleanExpression;
    limit?: number;
    offset?: number;
    resolveReferences?: string[];
}
export interface Event {
    id: string;
    action: 'CREATE' | 'UPDATE' | 'DELETE' | 'GET' | 'LIST' | 'OPERATE';
    recordSearchParams?: RecordSearchParams;
    actionSummary?: string;
    actionDescription?: string;
    resource?: Resource;
    records?: any[];
    ids?: string[];
    finalizes?: boolean;
    sync?: boolean;
    time?: string;
    annotations?: object;
}
export interface Extension {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    name: string;
    description?: string;
    selector?: EventSelector;
    order: number;
    finalizes: boolean;
    sync: boolean;
    responds: boolean;
    call: ExternalCall;
    annotations?: object;
}
export declare const ExtensionName = "Extension";
export declare const ExtensionIdName = "Id";
export declare const ExtensionVersionName = "Version";
export declare const ExtensionCreatedByName = "CreatedBy";
export declare const ExtensionUpdatedByName = "UpdatedBy";
export declare const ExtensionCreatedOnName = "CreatedOn";
export declare const ExtensionUpdatedOnName = "UpdatedOn";
export declare const ExtensionNameName = "Name";
export declare const ExtensionDescriptionName = "Description";
export declare const ExtensionSelectorName = "Selector";
export declare const ExtensionOrderName = "Order";
export declare const ExtensionFinalizesName = "Finalizes";
export declare const ExtensionSyncName = "Sync";
export declare const ExtensionRespondsName = "Responds";
export declare const ExtensionCallName = "Call";
export declare const ExtensionAnnotationsName = "Annotations";
