export type Operation =
    'OPERATION_TYPE_READ'
    | 'OPERATION_TYPE_UPDATE'
    | 'OPERATION_TYPE_CREATE'
    | 'OPERATION_TYPE_DELETE'
    | 'FULL'

export type Permit = 'PERMIT_TYPE_ALLOW' | 'PERMIT_TYPE_REJECT'

export interface SecurityConstraint {
    namespace: string
    resource: string
    property: string
    propertyValue: string
    operation: Operation
    recordIds: string[]
    username: string
    role: string
    requirePass: boolean
    permit: Permit
}