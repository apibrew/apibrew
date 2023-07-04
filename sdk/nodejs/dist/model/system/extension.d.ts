export declare const ExtensionResource: {
    resource: string;
    namespace: string;
};
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
export declare const ExtensionName = "Extension";
export declare const ExtensionIdName = "Id";
export declare const ExtensionVersionName = "Version";
export declare const ExtensionCreatedByName = "CreatedBy";
export declare const ExtensionUpdatedByName = "UpdatedBy";
export declare const ExtensionCreatedOnName = "CreatedOn";
export declare const ExtensionUpdatedOnName = "UpdatedOn";
export declare const ExtensionNameName = "Name";
export declare const ExtensionDescriptionName = "Description";
export declare const ExtensionSelectorName = "Selector";
export declare const ExtensionOrderName = "Order";
export declare const ExtensionFinalizesName = "Finalizes";
export declare const ExtensionSyncName = "Sync";
export declare const ExtensionRespondsName = "Responds";
export declare const ExtensionCallName = "Call";
