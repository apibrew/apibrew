


export const ExtensionResource = {
    resource: "extension",
    namespace: "system",
};

// Sub Types

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
selector?: object;
order: number;
finalizes: boolean;
sync: boolean;
responds: boolean;
call: object;

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


