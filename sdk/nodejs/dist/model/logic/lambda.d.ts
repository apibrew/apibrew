import { Function } from "./function";
export declare const LambdaResource: {
    resource: string;
    namespace: string;
};
export interface Lambda {
    id: string;
    package: string;
    name: string;
    eventSelectorPattern: string;
    function: Function;
    version: number;
}
export declare const LambdaName = "Lambda";
export declare const LambdaIdName = "Id";
export declare const LambdaPackageName = "Package";
export declare const LambdaNameName = "Name";
export declare const LambdaEventSelectorPatternName = "EventSelectorPattern";
export declare const LambdaFunctionName = "Function";
export declare const LambdaVersionName = "Version";
