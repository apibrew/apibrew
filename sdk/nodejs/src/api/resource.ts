import axios from 'axios'
import { type Resource } from '../model'
import { ServiceConfig } from './config'
import { isObjectModified } from '../util'

export async function list(config: ServiceConfig): Promise<Resource[]> {

    const result = await axios.get<Resource[]>(`${config.backendUrl}/resources`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function create(config: ServiceConfig, resource: Resource): Promise<Resource> {

    const result = await axios.post<Resource>(`${config.backendUrl}/resources`, resource, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function update(config: ServiceConfig, resource: Resource): Promise<Resource> {

    const result = await axios.put<Resource>(`${config.backendUrl}/resources/${resource.id}`, resource, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function remove(config: ServiceConfig, resource: Resource): Promise<void> {

    await axios.delete(`${config.backendUrl}/resources/${resource.id}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

}

export async function get(config: ServiceConfig, resourceId: string): Promise<Resource> {

    const result = await axios.get<Resource>(`${config.backendUrl}/resources/${resourceId}`, {
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


    const result = await axios.get<Resource>(`${config.backendUrl}/resources/by-name/${namespace}/${resourceName}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

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
        const existingResource = await getByName(config, resource.name, resource.namespace.name)
        resource.id = existingResource.id

        if (!isObjectModified(existingResource, resource)) {
            return existingResource
        }

        console.log('Updating resource', resource.name)
        return await update(config, resource)
    } catch (e) {
        console.log('Creating resource', resource.name)
        return await create(config, resource)
    }
}