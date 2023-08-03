import axios from 'axios'
import { BACKEND_URL } from '../config'
import * as TokenService from './token'
import { Resource, ResourceApi } from '@apibrew/client'
import { ServiceConfigProvider } from './service-config'

export namespace ResourceService {

    export async function list(): Promise<Resource[]> {
        return ResourceApi.list(ServiceConfigProvider())

    }

    export async function create(resource: Resource): Promise<Resource> {
        return ResourceApi.create(ServiceConfigProvider(), resource)
    }

    export async function update(resource: Resource): Promise<Resource> {
        return ResourceApi.update(ServiceConfigProvider(), resource)
    }

    export async function remove(resource: Resource, forceMigrate: boolean): Promise<void> {
        return ResourceApi.remove(ServiceConfigProvider(), resource, forceMigrate)
    }

    export async function get(resourceId: string): Promise<Resource> {
        return ResourceApi.get(ServiceConfigProvider(), resourceId)
    }

    export async function getByName(resourceName: string, namespace?: string): Promise<Resource> {
        return ResourceApi.getByName(ServiceConfigProvider(), resourceName, namespace)
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
            const existingResource = await getByName(resource.name, resource.namespace.name)
            resource.id = existingResource.id
            return await update(resource)
        }
    }
}
