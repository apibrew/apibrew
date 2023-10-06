

import { User } from "./user";

import { Role } from "./role";


export const PermissionResource = {
    resource: "Permission",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface Permission {
    id: string;
version: number;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
namespace?: string;
resource?: string;
property?: string;
propertyValue?: string;
propertyMode?: 'PROPERTY_MATCH_ONLY' | 'PROPERTY_MATCH_ANY';
operation: 'READ' | 'CREATE' | 'UPDATE' | 'DELETE' | 'FULL';
recordIds?: string[];
before?: string;
after?: string;
user?: User;
role?: Role;
permit: 'ALLOW' | 'REJECT';
localFlags?: object;

}
// Resource and Property Names
export const PermissionName = "Permission";

export const PermissionIdName = "Id";

export const PermissionVersionName = "Version";

export const PermissionCreatedByName = "CreatedBy";

export const PermissionUpdatedByName = "UpdatedBy";

export const PermissionCreatedOnName = "CreatedOn";

export const PermissionUpdatedOnName = "UpdatedOn";

export const PermissionNamespaceName = "Namespace";

export const PermissionResourceName = "Resource";

export const PermissionPropertyName = "Property";

export const PermissionPropertyValueName = "PropertyValue";

export const PermissionPropertyModeName = "PropertyMode";

export const PermissionOperationName = "Operation";

export const PermissionRecordIdsName = "RecordIds";

export const PermissionBeforeName = "Before";

export const PermissionAfterName = "After";

export const PermissionUserName = "User";

export const PermissionRoleName = "Role";

export const PermissionPermitName = "Permit";

export const PermissionLocalFlagsName = "LocalFlags";


