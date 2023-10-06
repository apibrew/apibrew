

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
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const LambdaName = "Lambda";

export const LambdaIdName = "Id";

export const LambdaPackageName = "Package";

export const LambdaNameName = "Name";

export const LambdaEventSelectorPatternName = "EventSelectorPattern";

export const LambdaFunctionName = "Function";

export const LambdaAnnotationsName = "Annotations";

export const LambdaCreatedByName = "CreatedBy";

export const LambdaUpdatedByName = "UpdatedBy";

export const LambdaCreatedOnName = "CreatedOn";

export const LambdaUpdatedOnName = "UpdatedOn";

export const LambdaVersionName = "Version";


