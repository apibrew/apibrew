export declare const UserResource: {
    resource: string;
    namespace: string;
};
export interface User {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    username: string;
    password?: string;
    securityConstraints?: object[];
    details?: object;
    roles?: string[];
}
export declare const UserName = "User";
export declare const UserIdName = "Id";
export declare const UserVersionName = "Version";
export declare const UserCreatedByName = "CreatedBy";
export declare const UserUpdatedByName = "UpdatedBy";
export declare const UserCreatedOnName = "CreatedOn";
export declare const UserUpdatedOnName = "UpdatedOn";
export declare const UserUsernameName = "Username";
export declare const UserPasswordName = "Password";
export declare const UserSecurityConstraintsName = "SecurityConstraints";
export declare const UserDetailsName = "Details";
export declare const UserRolesName = "Roles";
