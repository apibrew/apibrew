import { SecurityConstraint } from "./security-constraint";
export declare const RoleResource: {
    resource: string;
    namespace: string;
};
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
export declare const RoleName = "Role";
export declare const RoleIdName = "Id";
export declare const RoleVersionName = "Version";
export declare const RoleCreatedByName = "CreatedBy";
export declare const RoleUpdatedByName = "UpdatedBy";
export declare const RoleCreatedOnName = "CreatedOn";
export declare const RoleUpdatedOnName = "UpdatedOn";
export declare const RoleNameName = "Name";
export declare const RoleSecurityConstraintsName = "SecurityConstraints";
export declare const RoleDetailsName = "Details";
