export interface PermissionChecks {
    full: boolean
    read: boolean
    create: boolean
    update: boolean
    delete: boolean
    allowOwnedOnly?: boolean
}

export interface AccessMap {
    [k: string]: PermissionChecks
}
