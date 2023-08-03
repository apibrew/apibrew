export declare const NamespaceResource: {
    resource: string;
    namespace: string;
};
export interface Namespace {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    name: string;
    description?: string;
    details?: object;
}
export declare const NamespaceName = "Namespace";
export declare const NamespaceIdName = "Id";
export declare const NamespaceVersionName = "Version";
export declare const NamespaceCreatedByName = "CreatedBy";
export declare const NamespaceUpdatedByName = "UpdatedBy";
export declare const NamespaceCreatedOnName = "CreatedOn";
export declare const NamespaceUpdatedOnName = "UpdatedOn";
export declare const NamespaceNameName = "Name";
export declare const NamespaceDescriptionName = "Description";
export declare const NamespaceDetailsName = "Details";
