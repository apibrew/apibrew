import axios, { AxiosError } from 'axios'
import { BACKEND_URL } from '../config'
import { type Resource } from '../model'
import { TokenService } from './token'
import { handleError } from './error-handler'

export namespace ResourceService {
    interface ResourceListContainer {
        resources: Resource[]
    }

    export async function list(): Promise<Resource[]> {
        try {
            const result = await axios.get<ResourceListContainer>(`${BACKEND_URL}/system/resources`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data.resources
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function create(resource: Resource): Promise<Resource> {
        try {
            const result = await axios.post<Resource>(`${BACKEND_URL}/system/resources`, {
                resources: [resource],
                doMigration: true,
                forceMigration: true,
            }, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function update(resource: Resource): Promise<Resource> {
        try {
            const result = await axios.put<Resource>(`${BACKEND_URL}/system/resources`, {
                resources: [resource],
                doMigration: true,
                forceMigration: true,
            }, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function remove(resource: Resource, forceMigrate: boolean): Promise<void> {
        try {
            await axios.delete<void>(`${BACKEND_URL}/system/resources`, {
                data: {
                    doMigration: true,
                    forceMigration: forceMigrate,
                    ids: [resource.id]
                },
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function get(resourceId: string): Promise<Resource> {
        try {
            const result = await axios.get<Resource>(`${BACKEND_URL}/system/resources/${resourceId}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function getByName(resourceName: string, namespace?: string): Promise<Resource> {
        if (!namespace) {
            namespace = 'default'
        }

        try {
            const result = await axios.get<Resource>(`${BACKEND_URL}/system/resources/${namespace}/${resourceName}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function save(resource: Resource): Promise<Resource> {
        if (resource.id) {
            return update(resource)
        } else {
            return create(resource)
        }
    }

    export async function migrate(resource: Resource): Promise<Resource> {
        try {
            return create(resource)
        } catch (e) {
            return update(resource)
        }
    }
}
