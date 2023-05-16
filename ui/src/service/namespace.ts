import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type Namespace } from '../model'
import { TokenService } from './token'
import { handleError } from './error-handler'

export namespace NamespaceService {
    interface NamespaceListContainer {
        namespaces: Namespace[]
    }

    export async function list(): Promise<Namespace[]> {
        try {
            const result = await axios.get<NamespaceListContainer>(`${BACKEND_URL}/system/namespaces`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data.namespaces
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function create(namespace: Namespace): Promise<Namespace> {
        try {
            const result = await axios.post<Namespace>(`${BACKEND_URL}/system/namespaces`, {
                Namespaces: [namespace]
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

    export async function update(namespace: Namespace): Promise<Namespace> {
        try {
            const result = await axios.put<Namespace>(`${BACKEND_URL}/system/namespaces`, {
                Namespaces: [namespace]
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

    export async function remove(namespace: Namespace, forceMigrate: boolean): Promise<void> {
        try {
            await axios.delete(`${BACKEND_URL}/system/namespaces`, {
                data: {
                    ids: [namespace.id]
                },
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })
        } catch (e) {
            await handleError(e)
        }
    }

    export async function get(namespaceId: string): Promise<Namespace> {
        try {
            const result = await axios.get<Namespace>(`${BACKEND_URL}/system/namespaces/${namespaceId}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function getByName(namespaceName: string, namespace?: string): Promise<Namespace> {
        if (!namespace) {
            namespace = 'default'
        }

        try {
            const result = await axios.get<Namespace>(`${BACKEND_URL}/system/namespaces/${namespace}/${namespaceName}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
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
