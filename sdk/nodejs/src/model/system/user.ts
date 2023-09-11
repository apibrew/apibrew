import { Permission } from "./permission";
import { Role } from "./role";



export const UserResource = {
    resource: "User",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface User {
    id: string;
version: number;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
username: string;
password?: string;
roles?: Role[];
permissions?: Permission[];
details?: object;

}
// Resource and Property Names
export const UserName = "User";

export const UserIdName = "Id";

export const UserVersionName = "Version";

export const UserCreatedByName = "CreatedBy";

export const UserUpdatedByName = "UpdatedBy";

export const UserCreatedOnName = "CreatedOn";

export const UserUpdatedOnName = "UpdatedOn";

export const UserUsernameName = "Username";

export const UserPasswordName = "Password";

export const UserRolesName = "Roles";

export const UserPermissionsName = "Permissions";

export const UserDetailsName = "Details";


