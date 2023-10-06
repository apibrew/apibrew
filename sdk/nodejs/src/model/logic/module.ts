

import { FunctionExecutionEngine } from "./function-execution-engine";


export const ModuleResource = {
    resource: "Module",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface Module {
    id: string;
package: string;
content: string;
engine: FunctionExecutionEngine;
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const ModuleName = "Module";

export const ModuleIdName = "Id";

export const ModulePackageName = "Package";

export const ModuleContentName = "Content";

export const ModuleEngineName = "Engine";

export const ModuleAnnotationsName = "Annotations";

export const ModuleCreatedByName = "CreatedBy";

export const ModuleUpdatedByName = "UpdatedBy";

export const ModuleCreatedOnName = "CreatedOn";

export const ModuleUpdatedOnName = "UpdatedOn";

export const ModuleVersionName = "Version";


