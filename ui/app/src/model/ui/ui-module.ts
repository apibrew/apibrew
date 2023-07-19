


export const UIModuleResource = {
    resource: "UIModule",
    namespace: "ui",
};

// Sub Types

// Resource Type
export interface UIModule {
    id: string;
name: string;
description?: string;
source: string;
components: string[];
version: number;

}
// Resource and Property Names
export const UIModuleName = "UIModule";

export const UIModuleIdName = "Id";

export const UIModuleNameName = "Name";

export const UIModuleDescriptionName = "Description";

export const UIModuleSourceName = "Source";

export const UIModuleComponentsName = "Components";

export const UIModuleVersionName = "Version";


