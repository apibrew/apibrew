import { Namespace, Resource, Property } from "@apibrew/client";
import { AccessMap, PermissionChecks } from "./model.ts";

export const computeSystemIndeterminate = (accessMap: AccessMap, namespaces: Namespace[]) => {
    return combine(...namespaces.map(item => {
        return accessMap[`namespace-${item.name}`]
    }))
}

export const computeSystemValue = (accessMap: AccessMap) => {
    return accessMap['system']
}

export const computeNamespaceValue = (accessMap: AccessMap, namespace: Namespace) => {
    return combine(computeSystemValue(accessMap), namespacePermissions(accessMap, namespace.name))
}

export const computeNamespaceIndeterminate = (accessMap: AccessMap, namespace: Namespace, resources: Resource[]) => {
    return combine(...resources.filter(item => item.namespace.name === namespace.name).map(resource => {
        return combine(resourcePermissions(accessMap, resource), ...resource.properties.map(property => {
            return propertyPermissions(accessMap, resource, property)
        }))
    }))
}

export const computeResourceValue = (accessMap: AccessMap, resource: Resource) => {
    return combine(computeSystemValue(accessMap), namespacePermissions(accessMap, resource.namespace.name), resourcePermissions(accessMap, resource))
}

export const computeResourceIndeterminate = (accessMap: AccessMap, resource: Resource) => {
    return combine(...resource.properties.map(property => propertyPermissions(accessMap, resource, property)))
}

export const computeResourcePropertyValue = (accessMap: AccessMap, resource: Resource, property: Property) => {
    return combine(computeSystemValue(accessMap), namespacePermissions(accessMap, resource.namespace.name), resourcePermissions(accessMap, resource), propertyPermissions(accessMap, resource, property))
}

export const propertyPermissions = (accessMap: AccessMap, resource: Resource, property: Property) => {
    return accessMap[`resource-${resource.namespace.name}/${resource.name}-${property.name}`]
}

export const resourcePermissions = (accessMap: AccessMap, resource: Resource) => {
    return accessMap[`resource-${resource.namespace.name}/${resource.name}`]
}

export const namespacePermissions = (accessMap: AccessMap, namespaceName: string) => {
    return accessMap[`namespace-${namespaceName}`]
}

export function combine(...permissions: PermissionChecks[]): PermissionChecks {
    permissions = permissions.filter(item => item)
    return {
        full: permissions.some(item => item.full),
        read: permissions.some(item => item.read),
        create: permissions.some(item => item.create),
        update: permissions.some(item => item.update),
        delete: permissions.some(item => item.delete),
    }
}

export function isolate(value: PermissionChecks, combiner: PermissionChecks, actual: PermissionChecks): PermissionChecks {
    return {
        full: value.full || combiner.full && actual.full,
        read: value.read || combiner.read && actual.read,
        create: value.create || combiner.create && actual.create,
        update: value.update || combiner.update && actual.update,
        delete: value.delete || combiner.delete && actual.delete,
    }
}
