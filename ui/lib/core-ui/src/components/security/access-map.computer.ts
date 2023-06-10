import {Namespace, Resource} from "../../model";
import {PermissionChecks, AccessMap} from "./model.ts";
import {SecurityConstraint} from "../../model/security-constraint.ts";
import {namespacePermissions, resourcePermissions} from "./helper.ts";

function mapConstraintTo(constraint: SecurityConstraint, permissionChecks?: PermissionChecks): boolean {
    if (!permissionChecks) {
        return false
    }

    if (constraint.operation === 'FULL') {
        permissionChecks.full = true
    } else if (constraint.operation === 'OPERATION_TYPE_READ') {
        permissionChecks.read = true
    } else if (constraint.operation === 'OPERATION_TYPE_UPDATE') {
        permissionChecks.update = true
    } else if (constraint.operation === 'OPERATION_TYPE_CREATE') {
        permissionChecks.create = true
    } else if (constraint.operation === 'OPERATION_TYPE_DELETE') {
        permissionChecks.delete = true
    } else {
        throw new Error(`Unknown operation ${constraint.operation}`)
    }

    // if (constraint.property == 'owner' && constraint.propertyMode == 'PROPERTY_MATCH_ANY' && constraint.propertyValue == '$username') {
    //     permissionChecks.allowOwnedOnly = true
    // }

    return true
}

export function prepareAccessMap(accessMap: AccessMap, namespaces: Namespace[], resources: Resource[], constraints: SecurityConstraint[]) {
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
        if (constraint.propertyMode || constraint.permit == 'PERMIT_TYPE_REJECT' || constraint.propertyValue) {
            continue
        }

        if (constraint.property === '*') { // resource or namespace or system level
            if (constraint.resource === '*') { // namespace or system level
                if (constraint.namespace === '*') { // system level
                    if (mapConstraintTo(constraint, updatedAccessMap['system'])) {
                        constraint.localFlags = {
                            imported: true,
                        }
                    }
                } else { // namespace level
                    if (mapConstraintTo(constraint, updatedAccessMap[`namespace-${constraint.namespace}`])) {
                        constraint.localFlags = {
                            imported: true,
                        }
                    }
                }
            } else { // resource level
                if (mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace}/${constraint.resource}`])) {
                    constraint.localFlags = {
                        imported: true,
                    }
                }
            }
        } else { // property level constraint
            if (mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace}/${constraint.resource}-${constraint.property}`])) {
                constraint.localFlags = {
                    imported: true,
                }
            }
        }
    }

    return updatedAccessMap;
}

export function prepareConstraintsFromAccessMap(constraints: SecurityConstraint[], accessMap: AccessMap, namespaces: Namespace[], resources: Resource[]): SecurityConstraint[] {
    const updatedConstraints: SecurityConstraint[] = []

    const systemPermission = accessMap.system

    const systemConstraint = {
        namespace: '*',
        resource: '*',
        property: '*',
        permit: 'PERMIT_TYPE_ALLOW',
        localFlags: {
            imported: true,
        }
    } as SecurityConstraint

    if (systemPermission.full) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'FULL',
        })
    }
    if (systemPermission.read) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'OPERATION_TYPE_READ',
        })
    }
    if (systemPermission.create) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'OPERATION_TYPE_CREATE',
        })
    }
    if (systemPermission.update) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'OPERATION_TYPE_UPDATE',
        })
    }
    if (systemPermission.delete) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'OPERATION_TYPE_DELETE',
        })
    }

    for (const namespace of namespaces) {
        const namespacePermission = namespacePermissions(accessMap, namespace.name)

        const namespaceConstraint = {
            namespace: namespace.name,
            resource: '*',
            property: '*',
            permit: 'PERMIT_TYPE_ALLOW',
            localFlags: {
                imported: true,
            }
        } as SecurityConstraint

        if (namespacePermission.full) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'FULL',
            })
        }
        if (namespacePermission.read) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'OPERATION_TYPE_READ',
            })
        }
        if (namespacePermission.create) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'OPERATION_TYPE_CREATE',
            })
        }
        if (namespacePermission.update) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'OPERATION_TYPE_UPDATE',
            })
        }
        if (namespacePermission.delete) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'OPERATION_TYPE_DELETE',
            })
        }
    }

    for (const resource of resources) {
        const resourcePermission = accessMap[`resource-${resource.namespace}/${resource.name}`]

        const resourceConstraint = {
            namespace: resource.namespace,
            resource: resource.name,
            property: '*',
            permit: 'PERMIT_TYPE_ALLOW',
            localFlags: {
                imported: true,
            }
        } as SecurityConstraint

        // if (resourcePermission.allowOwnedOnly) {
        //     resourceConstraint.propertyMode = 'PROPERTY_MATCH_ANY'
        //     resourceConstraint.property = 'owner'
        //     resourceConstraint.propertyValue = '$username'
        // }

        if (resourcePermission.full) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'FULL',
            })
        }
        if (resourcePermission.read) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'OPERATION_TYPE_READ',
            })
        }
        if (resourcePermission.create) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'OPERATION_TYPE_CREATE',
            })
            console.log(resourcePermission.create)
        }
        if (resourcePermission.update) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'OPERATION_TYPE_UPDATE',
            })
        }
        if (resourcePermission.delete) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'OPERATION_TYPE_DELETE',
            })
        }

        for (const property of resource.properties) {
            const propertyPermission = accessMap[`resource-${resource.namespace}/${resource.name}-${property.name}`]

            const propertyConstraint = {
                namespace: resource.namespace,
                resource: resource.name,
                property: property.name,
                permit: 'PERMIT_TYPE_ALLOW',
                localFlags: {
                    imported: true,
                }
            } as SecurityConstraint

            if (propertyPermission.full) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'FULL',
                })
            }
            if (propertyPermission.read) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'OPERATION_TYPE_READ',
                })
            }
            if (propertyPermission.create) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'OPERATION_TYPE_CREATE',
                })
            }
            if (propertyPermission.update) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'OPERATION_TYPE_UPDATE',
                })
            }
            if (propertyPermission.delete) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'OPERATION_TYPE_DELETE',
                })
            }
        }
    }

    // keep constraints which is not related to access map
    for (const constraint of constraints) {
        if (constraint.localFlags?.imported) {
            continue
        }

        updatedConstraints.push(constraint)
        console.log(constraint)
    }

    return updatedConstraints;
}
