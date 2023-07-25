import { Namespace } from "./namespace";
import { Resource } from "./resource";
import { User } from "./user";
import { Role } from "./role";


export const SecurityConstraintResource = {
    resource: "SecurityConstraint",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface SecurityConstraint {
    id: string;
version: number;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
namespace?: Namespace;
resource?: Resource;
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
export const SecurityConstraintName = "SecurityConstraint";

export const SecurityConstraintIdName = "Id";

export const SecurityConstraintVersionName = "Version";

export const SecurityConstraintCreatedByName = "CreatedBy";

export const SecurityConstraintUpdatedByName = "UpdatedBy";

export const SecurityConstraintCreatedOnName = "CreatedOn";

export const SecurityConstraintUpdatedOnName = "UpdatedOn";

export const SecurityConstraintNamespaceName = "Namespace";

export const SecurityConstraintResourceName = "Resource";

export const SecurityConstraintPropertyName = "Property";

export const SecurityConstraintPropertyValueName = "PropertyValue";

export const SecurityConstraintPropertyModeName = "PropertyMode";

export const SecurityConstraintOperationName = "Operation";

export const SecurityConstraintRecordIdsName = "RecordIds";

export const SecurityConstraintBeforeName = "Before";

export const SecurityConstraintAfterName = "After";

export const SecurityConstraintUserName = "User";

export const SecurityConstraintRoleName = "Role";

export const SecurityConstraintPermitName = "Permit";

export const SecurityConstraintLocalFlagsName = "LocalFlags";


