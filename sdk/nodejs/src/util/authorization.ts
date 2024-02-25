import {Permission, Resource} from "../model";
import {Entity} from "../entity";

export enum AccessLevel {
    NONE,
    READ,
    READ_WRITE
}

export function checkResourceAccess(permissions: Permission[], resource: Resource, operation: Permission["operation"], recordId?: string): boolean {
    const matchingConstraints = permissions.filter(constraint => {
        if (constraint.namespace && constraint.namespace !== resource.namespace.name) {
            return false
        }

        if (constraint.resource && constraint.resource !== resource.name) {
            return false
        }

        return true
    })

    const readCompatibleConstraints = matchingConstraints.filter(constraint => {
        return constraint.operation === 'FULL' || constraint.operation === operation;
    })

    // checking can read
    const hasAllow = readCompatibleConstraints.some(constraint => !constraint.permit || constraint.permit === 'ALLOW')
    const hasReject = readCompatibleConstraints.some(constraint => constraint.permit === 'REJECT')

    return hasAllow && !hasReject
}

export function checkResourcePropertyAccess(permissions: Permission[], resource: Resource): AccessLevel {
    const matchingConstraints = permissions.filter(constraint => {
        if (constraint.namespace && constraint.namespace !== resource.namespace.name) {
            return false
        }

        if (constraint.resource && constraint.resource !== resource.name) {
            return false
        }

        return true
    })

    const readCompatibleConstraints = matchingConstraints.filter(constraint => {
        return !constraint.operation || constraint.operation === 'READ' || constraint.operation === 'FULL';
    })

    const writeCompatibleConstraints = matchingConstraints.filter(constraint => {
        return constraint.operation === 'UPDATE' || constraint.operation === 'CREATE' || constraint.operation === 'FULL';
    })

    // checking can read
    const hasReadAllow = readCompatibleConstraints.some(constraint => !constraint.permit || constraint.permit === 'ALLOW')
    const hasReadReject = readCompatibleConstraints.some(constraint => constraint.permit === 'REJECT')

    if (hasReadReject || !hasReadAllow) {
        return AccessLevel.NONE
    }

    // checking can write
    const hasWriteAllow = writeCompatibleConstraints.some(constraint => !constraint.permit || constraint.permit === 'ALLOW')
    const hasWriteReject = writeCompatibleConstraints.some(constraint => constraint.permit === 'REJECT')

    if (hasWriteReject || !hasWriteAllow) {
        return AccessLevel.READ
    }

    return AccessLevel.READ_WRITE
}

export function filterRecordForUpdate<T extends Entity>(permissions: Permission[], resource: Resource, record: T): T {
    const recordForUpdate = {...record}

    for (const name in resource.properties) {
        if (name === 'id') {
            continue
        }
        const propertyAccess = checkResourcePropertyAccess(permissions, resource)

        if (propertyAccess !== AccessLevel.READ_WRITE) {
            delete (recordForUpdate as any)[name]
        }
    }

    return recordForUpdate
}
