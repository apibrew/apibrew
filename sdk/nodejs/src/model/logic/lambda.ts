

import { Function } from "./function";


export const LambdaResource = {
    resource: "Lambda",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface Lambda {
    id: string;
package: string;
name: string;
eventSelectorPattern: string;
function: Function;
version: number;

}
// Resource and Property Names
export const LambdaName = "Lambda";

export const LambdaIdName = "Id";

export const LambdaPackageName = "Package";

export const LambdaNameName = "Name";

export const LambdaEventSelectorPatternName = "EventSelectorPattern";

export const LambdaFunctionName = "Function";

export const LambdaVersionName = "Version";


