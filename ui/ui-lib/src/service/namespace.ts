import axios from 'axios'
import {BACKEND_URL} from '../config'
import {type Namespace} from '../model'
import * as TokenService from './token'

export namespace NamespaceService {
    interface NamespaceListContainer {
        namespaces: Namespace[]
    }

    export async function list(): Promise<Namespace[]> {
        const result = await axios.get<NamespaceListContainer>(`${BACKEND_URL}/system/namespaces`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data.namespaces

    }

    export async function create(namespace: Namespace): Promise<Namespace> {
        const result = await axios.post<Namespace>(`${BACKEND_URL}/system/namespaces`, {
            Namespaces: [namespace]
        }, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data
    }

    export async function update(namespace: Namespace): Promise<Namespace> {
        const result = await axios.put<Namespace>(`${BACKEND_URL}/system/namespaces`, {
            Namespaces: [namespace]
        }, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data
    }

    export async function remove(namespace: Namespace, forceMigrate: boolean): Promise<void> {
        await axios.delete(`${BACKEND_URL}/system/namespaces`, {
            data: {
                ids: [namespace.id]
            },
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })
    }

    export async function get(namespaceId: string): Promise<Namespace> {
        const result = await axios.get<Namespace>(`${BACKEND_URL}/system/namespaces/${namespaceId}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data
    }

    export async function getByName(namespaceName: string, namespace?: string): Promise<Namespace> {
        if (!namespace) {
            namespace = 'default'
        }

        const result = await axios.get<Namespace>(`${BACKEND_URL}/system/namespaces/${namespace}/${namespaceName}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data
    }

    export async function save(namespace: Namespace): Promise<Namespace> {
        if (namespace.id) {
            return await update(namespace)
        } else {
            return await create(namespace)
        }
    }

    export async function migrate(namespace: Namespace): Promise<Namespace> {
        try {
            return await create(namespace)
        } catch (e) {
            return await update(namespace)
        }
    }
}
