


// Sub Types

// Resource Type
export interface Module {
    id: string;
name: string;
package: string;
description?: string;
source: string;
components: string[];
version: number;

}
// Resource and Property Names
export const ModuleName = "Module";

export const ModuleIdName = "Id";

export const ModuleNameName = "Name";

export const ModulePackageName = "Package";

export const ModuleDescriptionName = "Description";

export const ModuleSourceName = "Source";

export const ModuleComponentsName = "Components";

export const ModuleVersionName = "Version";


