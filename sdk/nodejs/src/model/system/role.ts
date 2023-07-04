


export const RoleResource = {
    resource: "role",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface Role {
    id: string;
name: string;
securityConstraints?: object[];
details?: object;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const RoleName = "Role";

export const RoleIdName = "Id";

export const RoleNameName = "Name";

export const RoleSecurityConstraintsName = "SecurityConstraints";

export const RoleDetailsName = "Details";

export const RoleCreatedByName = "CreatedBy";

export const RoleUpdatedByName = "UpdatedBy";

export const RoleCreatedOnName = "CreatedOn";

export const RoleUpdatedOnName = "UpdatedOn";

export const RoleVersionName = "Version";


