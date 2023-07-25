import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type Resource } from '../model'
import * as TokenService from './token'

export namespace ResourceService {

    export async function list(): Promise<Resource[]> {

        const result = await axios.get<Resource[]>(`${BACKEND_URL}/resources`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function create(resource: Resource): Promise<Resource> {

        const result = await axios.post<Resource>(`${BACKEND_URL}/resources`, {
            resources: [resource],
            doMigration: true,
            forceMigration: true
        }, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function update(resource: Resource): Promise<Resource> {

        const result = await axios.put<Resource>(`${BACKEND_URL}/resources`, {
            resources: [resource],
            doMigration: true,
            forceMigration: true
        }, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function remove(resource: Resource, forceMigrate: boolean): Promise<void> {

        await axios.delete(`${BACKEND_URL}/resources`, {
            data: {
                doMigration: true,
                forceMigration: forceMigrate,
                ids: [resource.id]
            },
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

    }

    export async function get(resourceId: string): Promise<Resource> {

        const result = await axios.get<Resource>(`${BACKEND_URL}/resources/${resourceId}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function getByName(resourceName: string, namespace?: string): Promise<Resource> {
        if (!namespace) {
            namespace = 'default'
        }


        const result = await axios.get<{
            resource: Resource
        }>(`${BACKEND_URL}/resources/${namespace}/${resourceName}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data.resource

    }

    export async function save(resource: Resource): Promise<Resource> {
        if (resource.id) {
            return await update(resource)
        } else {
            return await create(resource)
        }
    }

    export async function migrate(resource: Resource): Promise<Resource> {
        try {
            return await create(resource)
        } catch (e) {
            const existingResource = await getByName(resource.name, resource.namespace)
            resource.id = existingResource.id
            return await update(resource)
        }
    }
}
