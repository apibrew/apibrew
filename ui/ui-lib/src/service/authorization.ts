import {Resource} from "../model/system/resource.ts";
import {get, getBody} from "./token.ts";
import {Record} from "./record.ts";

export enum AccessLevel {
    NONE,
    READ,
    READ_WRITE
}

export function checkResourcePropertyAccess(resource: Resource, property: string, recordId?: string): AccessLevel {
    const userConstraints = getBody().securityConstraints;

    console.log('checkResourcePropertyAccess for', resource.name, property, recordId)

    const matchingConstraints = userConstraints.filter(constraint => {
        if (constraint.namespace && constraint.namespace.name != resource.namespace.name) {
            console.log('could not pass namespace gate')
            return false
        }

        if (constraint.resource && constraint.resource.name != resource.name) {
            console.log('could not pass resource gate')
            return false
        }

        if (constraint.property && constraint.property != property) {
            console.log('could not pass property gate')
            return false
        }

        /*
        if (recordId && constraint.recordIds && constraint.recordIds.length > 0 && !constraint.recordIds.includes(recordId)) {
            return false
        }
        */

        return true
    })

    console.log(matchingConstraints)

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

    console.log(hasWriteAllow, hasWriteReject)

    if (hasWriteReject || !hasWriteAllow) {
        return AccessLevel.READ
    }

    return AccessLevel.READ_WRITE
}

export function filterRecordForUpdate<T extends Record>(resource: Resource, record: T): T {
    const recordForUpdate = {...record}

    for (const property of resource.properties) {
        if (property.name === 'id') {
            continue
        }
        const propertyAccess = checkResourcePropertyAccess(resource, property.name, record.id)
        if (propertyAccess !== AccessLevel.READ_WRITE) {
            console.log(`Removing property ${property.name} from record ${record.id} because of access level ${propertyAccess}`)
            delete recordForUpdate[property.name]
        }
    }

    return recordForUpdate
}
