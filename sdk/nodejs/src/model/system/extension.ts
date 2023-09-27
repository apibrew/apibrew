import {Resource} from "./resource";


export const ExtensionResource = {
    resource: "Extension",
    namespace: "system",
};

// Sub Types

export interface BooleanExpression {
     and?: BooleanExpression[];
     or?: BooleanExpression[];
     not?: BooleanExpression;
     equal?: PairExpression;
     lessThan?: PairExpression;
     greaterThan?: PairExpression;
     lessThanOrEqual?: PairExpression;
     greaterThanOrEqual?: PairExpression;
     in?: PairExpression;
     isNull?: Expression;
     regexMatch?: RegexMatchExpression;

}

export interface PairExpression {
     left?: Expression;
     right?: Expression;

}

export interface RefValue {
     namespace?: string;
     resource?: string;
     properties?: object;

}

export interface RegexMatchExpression {
     pattern?: string;
     expression?: Expression;

}

export interface Expression {
     property?: string;
     value?: object;
     refValue?: RefValue;

}

export interface FunctionCall {
     host: string;
     functionName: string;

}

export interface HttpCall {
     uri: string;
     method: string;

}

export interface ChannelCall {
     channelKey: string;

}

export interface ExternalCall {
     functionCall?: FunctionCall;
     httpCall?: HttpCall;
     channelCall?: ChannelCall;

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
     error?: Error;

}

export interface ErrorField {
     recordId?: string;
     property?: string;
     message?: string;
     value?: object;

}

export interface Error {
     code?: 'UNKNOWN_ERROR' | 'RECORD_NOT_FOUND' | 'UNABLE_TO_LOCATE_PRIMARY_KEY' | 'INTERNAL_ERROR' | 'PROPERTY_NOT_FOUND' | 'RECORD_VALIDATION_ERROR' | 'RESOURCE_VALIDATION_ERROR' | 'AUTHENTICATION_FAILED' | 'ALREADY_EXISTS' | 'ACCESS_DENIED' | 'BACKEND_ERROR' | 'UNIQUE_VIOLATION' | 'REFERENCE_VIOLATION' | 'RESOURCE_NOT_FOUND' | 'UNSUPPORTED_OPERATION' | 'EXTERNAL_BACKEND_COMMUNICATION_ERROR' | 'EXTERNAL_BACKEND_ERROR';
     message?: string;
     fields?: ErrorField[];

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


