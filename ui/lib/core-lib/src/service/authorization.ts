import {Resource} from "../model";
import {get, getBody} from "./token.ts";
import {Record} from "./record.ts";
import {roundToNearestMinutesWithOptions} from "date-fns/fp";

export enum AccessLevel {
    NONE,
    READ,
    READ_WRITE
}

export function checkResourcePropertyAccess(resource: Resource, property: string, recordId?: string): AccessLevel {
    const userConstraints = getBody().securityConstraints;

    const matchingConstraints = userConstraints.filter(constraint => {
        if (constraint.namespace !== '*' && constraint.namespace != resource.namespace) {
            return false
        }

        if (constraint.resource !== '*' && constraint.resource != resource.name) {
            return false
        }

        if (constraint.property !== '*' && constraint.property != property) {
            return false
        }

        /*
        if (recordId && constraint.recordIds && constraint.recordIds.length > 0 && !constraint.recordIds.includes(recordId)) {
            return false
        }
        */

        return true
    })

    const readCompatibleConstraints = matchingConstraints.filter(constraint => {
        return !constraint.operation || constraint.operation === 'OPERATION_TYPE_READ' || constraint.operation === 'FULL';
    })

    const writeCompatibleConstraints = matchingConstraints.filter(constraint => {
        return constraint.operation === 'OPERATION_TYPE_UPDATE' || constraint.operation === 'OPERATION_TYPE_CREATE' || constraint.operation === 'FULL';
    })

    // checking can read
    const hasReadAllow = readCompatibleConstraints.some(constraint => !constraint.permit || constraint.permit === 'PERMIT_TYPE_ALLOW')
    const hasReadReject = readCompatibleConstraints.some(constraint => constraint.permit === 'PERMIT_TYPE_REJECT')

    if (hasReadReject || !hasReadAllow) {
        return AccessLevel.NONE
    }

    // checking can write
    const hasWriteAllow = writeCompatibleConstraints.some(constraint => !constraint.permit || constraint.permit === 'PERMIT_TYPE_ALLOW')
    const hasWriteReject = writeCompatibleConstraints.some(constraint => constraint.permit === 'PERMIT_TYPE_REJECT')

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
