import axios from 'axios'
import { type Resource } from '../model'
import { ServiceConfig } from './config'

interface ResourceListContainer {
    resources: Resource[]
}

export async function list(config: ServiceConfig): Promise<Resource[]> {

    const result = await axios.get<ResourceListContainer>(`${config.backendUrl}/system/resources`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data.resources

}

export async function create(config: ServiceConfig, resource: Resource): Promise<Resource> {

    const result = await axios.post<Resource>(`${config.backendUrl}/system/resources`, {
        resources: [resource],
        doMigration: true,
        forceMigration: true
    }, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function update(config: ServiceConfig, resource: Resource): Promise<Resource> {

    const result = await axios.put<Resource>(`${config.backendUrl}/system/resources`, {
        resources: [resource],
        doMigration: true,
        forceMigration: true
    }, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function remove(config: ServiceConfig, resource: Resource, forceMigrate: boolean): Promise<void> {

    await axios.delete(`${config.backendUrl}/system/resources`, {
        data: {
            doMigration: true,
            forceMigration: forceMigrate,
            ids: [resource.id]
        },
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

}

export async function get(config: ServiceConfig, resourceId: string): Promise<Resource> {

    const result = await axios.get<Resource>(`${config.backendUrl}/system/resources/${resourceId}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function getByName(config: ServiceConfig, resourceName: string, namespace?: string): Promise<Resource> {
    if (!namespace) {
        namespace = 'default'
    }


    const result = await axios.get<{
        resource: Resource
    }>(`${config.backendUrl}/system/resources/${namespace}/${resourceName}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data.resource

}

export async function save(config: ServiceConfig, resource: Resource): Promise<Resource> {
    if (resource.id) {
        return await update(config, resource)
    } else {
        return await create(config, resource)
    }
}

export async function apply(config: ServiceConfig, resource: Resource): Promise<Resource> {
    try {
        const existingResource = await getByName(config, resource.name, resource.namespace)
        resource.id = existingResource.id

        return await update(config, resource)
    } catch (e) {
        return await create(config, resource)
    }
}