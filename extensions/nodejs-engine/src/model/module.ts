

import { FunctionExecutionEngine } from "./function-execution-engine";


// Sub Types

// Resource Type
export interface Module {
    id: string;
package: string;
content: string;
engine: FunctionExecutionEngine;
version: number;

}
// Resource and Property Names
export const ModuleName = "Module";

export const ModuleIdName = "Id";

export const ModulePackageName = "Package";

export const ModuleContentName = "Content";

export const ModuleEngineName = "Engine";

export const ModuleVersionName = "Version";


