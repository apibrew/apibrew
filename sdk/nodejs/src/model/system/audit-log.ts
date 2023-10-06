


export const AuditLogResource = {
    resource: "AuditLog",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface AuditLog {
    id: string;
version: number;
namespace: string;
resource: string;
recordId: string;
time: string;
username: string;
operation: 'CREATE' | 'UPDATE' | 'DELETE';
properties?: object;
annotations?: object;

}
// Resource and Property Names
export const AuditLogName = "AuditLog";

export const AuditLogIdName = "Id";

export const AuditLogVersionName = "Version";

export const AuditLogNamespaceName = "Namespace";

export const AuditLogResourceName = "Resource";

export const AuditLogRecordIdName = "RecordId";

export const AuditLogTimeName = "Time";

export const AuditLogUsernameName = "Username";

export const AuditLogOperationName = "Operation";

export const AuditLogPropertiesName = "Properties";

export const AuditLogAnnotationsName = "Annotations";


