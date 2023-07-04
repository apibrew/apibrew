


export const NamespaceResource = {
    resource: "namespace",
    namespace: "system",
};

// Sub Types

// Resource Type
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
securityConstraints?: object[];

}
// Resource and Property Names
export const NamespaceName = "Namespace";

export const NamespaceIdName = "Id";

export const NamespaceVersionName = "Version";

export const NamespaceCreatedByName = "CreatedBy";

export const NamespaceUpdatedByName = "UpdatedBy";

export const NamespaceCreatedOnName = "CreatedOn";

export const NamespaceUpdatedOnName = "UpdatedOn";

export const NamespaceNameName = "Name";

export const NamespaceDescriptionName = "Description";

export const NamespaceDetailsName = "Details";

export const NamespaceSecurityConstraintsName = "SecurityConstraints";


