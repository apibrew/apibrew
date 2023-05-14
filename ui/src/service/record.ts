import axios, { AxiosError } from 'axios'
import { BACKEND_URL } from '../config'
import { type Resource } from '../model'
import { TokenService } from './token'
import { handleError } from './error-handler'

export namespace RecordService {
    interface RecordListContainer<T> {
        content: T[]
    }

    export async function list<T>(namespace: string, resource: string): Promise<T[]> {
        try {
            const result = await axios.get<RecordListContainer<T>>(`${BACKEND_URL}/records/${namespace}/${resource}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data.content
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function create<T>(namespace: string, resource: string, record: T): Promise<T> {
        try {
            const result = await axios.post<T>(`${BACKEND_URL}/records/${namespace}/${resource}`, record, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function update<T>(namespace: string, resource: string, record: T): Promise<T> {
        try {
            const result = await axios.put<T>(`${BACKEND_URL}/records/${namespace}/${resource}`, record, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function remove(namespace: string, resource: string, id: string): Promise<void> {
        try {
            await axios.delete<void>(`${BACKEND_URL}/records/${namespace}/${resource}/${id}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function get<T>(namespace: string, resource: string, id: string): Promise<T> {
        try {
            const result = await axios.get<T>(`${BACKEND_URL}/records/${namespace}/${resource}/${id}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }

    export async function apply<T>(namespace: string, resource: string, record: T): Promise<T> {
        try {
            const result = await axios.patch<T>(`${BACKEND_URL}/records/${namespace}/${resource}`, record, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data
        } catch (e) {
            return await handleError(e)
        }
    }
}
