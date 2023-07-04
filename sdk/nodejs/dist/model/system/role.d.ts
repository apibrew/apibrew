export declare const RoleResource: {
    resource: string;
    namespace: string;
};
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
export declare const RoleName = "Role";
export declare const RoleIdName = "Id";
export declare const RoleNameName = "Name";
export declare const RoleSecurityConstraintsName = "SecurityConstraints";
export declare const RoleDetailsName = "Details";
export declare const RoleCreatedByName = "CreatedBy";
export declare const RoleUpdatedByName = "UpdatedBy";
export declare const RoleCreatedOnName = "CreatedOn";
export declare const RoleUpdatedOnName = "UpdatedOn";
export declare const RoleVersionName = "Version";
