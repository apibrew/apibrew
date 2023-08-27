import { SecurityConstraint } from "./security-constraint";


export const RoleResource = {
    resource: "Role",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface Role {
    id: string;
version: number;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
name: string;
securityConstraints?: SecurityConstraint[];
details?: object;

}
// Resource and Property Names
export const RoleName = "Role";

export const RoleIdName = "Id";

export const RoleVersionName = "Version";

export const RoleCreatedByName = "CreatedBy";

export const RoleUpdatedByName = "UpdatedBy";

export const RoleCreatedOnName = "CreatedOn";

export const RoleUpdatedOnName = "UpdatedOn";

export const RoleNameName = "Name";

export const RoleSecurityConstraintsName = "SecurityConstraints";

export const RoleDetailsName = "Details";


