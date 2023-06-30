export type Operation =
    'OPERATION_TYPE_READ'
    | 'OPERATION_TYPE_UPDATE'
    | 'OPERATION_TYPE_CREATE'
    | 'OPERATION_TYPE_DELETE'
    | 'FULL'

export type PropertyMode =
    'PROPERTY_MATCH_ONLY'
    | 'PROPERTY_MATCH_ANY'

export type Permit = 'PERMIT_TYPE_ALLOW' | 'PERMIT_TYPE_REJECT'

export interface SecurityConstraint {
    namespace: string
    resource: string
    property: string
    propertyValue?: string
    propertyMode?: PropertyMode
    operation: Operation
    recordIds?: string[]
    username?: string
    role?: string
    permit: Permit
    localFlags?: any
}