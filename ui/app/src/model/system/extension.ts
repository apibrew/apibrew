


export const ExtensionResource = {
    resource: "Extension",
    namespace: "system",
};

// Sub Types

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
     records?: Record[];
     ids?: string[];
     finalizes?: boolean;
     sync?: boolean;
     time?: string;
     annotations?: object;

}

// Resource Type
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
// Resource and Property Names
export const ExtensionName = "Extension";

export const ExtensionIdName = "Id";

export const ExtensionVersionName = "Version";

export const ExtensionCreatedByName = "CreatedBy";

export const ExtensionUpdatedByName = "UpdatedBy";

export const ExtensionCreatedOnName = "CreatedOn";

export const ExtensionUpdatedOnName = "UpdatedOn";

export const ExtensionNameName = "Name";

export const ExtensionDescriptionName = "Description";

export const ExtensionSelectorName = "Selector";

export const ExtensionOrderName = "Order";

export const ExtensionFinalizesName = "Finalizes";

export const ExtensionSyncName = "Sync";

export const ExtensionRespondsName = "Responds";

export const ExtensionCallName = "Call";

export const ExtensionAnnotationsName = "Annotations";


