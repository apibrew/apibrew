import {Namespace, Resource} from "../../model";
import {PermissionChecks} from "./SecurityConstraintsInputSimple.tsx";
import {SecurityConstraint} from "../../model/security-constraint.ts";

function mapConstraintTo(constraints: SecurityConstraint, permissionChecks?: PermissionChecks) {
    if (!permissionChecks) {
        return
    }

    if (constraints.operation === 'FULL') {
        permissionChecks.full = true
    } else if (constraints.operation === 'OPERATION_TYPE_READ') {
        permissionChecks.read = true
    } else if (constraints.operation === 'OPERATION_TYPE_UPDATE') {
        permissionChecks.update = true
    } else if (constraints.operation === 'OPERATION_TYPE_CREATE') {
        permissionChecks.create = true
    } else if (constraints.operation === 'OPERATION_TYPE_DELETE') {
        permissionChecks.delete = true
    } else {
        throw new Error(`Unknown operation ${constraints.operation}`)
    }
}

export function prepareAccessMap(accessMap: {
    [p: string]: PermissionChecks
}, namespaces: Namespace[], resources: Resource[], constraints: SecurityConstraint[]) {
    let updatedAccessMap = {...accessMap}

    if (!namespaces.some(item => item.name === 'system')) {
        namespaces.push({
            name: 'system',
        })
    }

    updatedAccessMap = {
        ...updatedAccessMap,
        "system": {
            full: false,
            read: false,
            create: false,
            update: false,
            delete: false
        },
    }

    for (const namespace of namespaces) {
        updatedAccessMap = {
            ...updatedAccessMap,
            [`namespace-${namespace.name}`]: {
                full: false,
                read: false,
                create: false,
                update: false,
                delete: false
            },
        }
    }
    for (const resource of resources) {
        updatedAccessMap = {
            ...updatedAccessMap,
            [`resource-${resource.namespace}/${resource.name}`]: {
                full: false,
                read: false,
                create: false,
                update: false,
                delete: false
            },
        }

        for (const property of resource.properties) {
            updatedAccessMap = {
                ...updatedAccessMap,
                [`resource-${resource.namespace}/${resource.name}-${property.name}`]: {
                    full: false,
                    read: false,
                    create: false,
                    update: false,
                    delete: false
                },
            }
        }
    }

    for (const constraint of constraints) {
        if (constraint.property === '*') { // resource or namespace or system level
            if (constraint.resource === '*') { // namespace or system level
                if (constraint.namespace === '*') { // system level
                    mapConstraintTo(constraint, updatedAccessMap['system'])
                } else { // namespace level
                    mapConstraintTo(constraint, updatedAccessMap[`namespace-${constraint.namespace}`])
                }
            } else { // resource level
                mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace}/${constraint.resource}`])
            }
        } else { // property level constraint
            mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace}/${constraint.resource}-${constraint.property}`])
        }
    }

    return updatedAccessMap;
}