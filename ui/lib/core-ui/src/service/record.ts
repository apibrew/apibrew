import axios from 'axios'
import {BACKEND_URL} from '../config'
import {TokenService} from './token'
import {handleError} from './error-handler'
import {BooleanExpression} from "../model";

export interface Record {
    id?: string

    [key: string]: any
}

export namespace RecordService {
    interface RecordListContainer<T> {
        content: {
            properties: T
        }[]
    }

    export async function list<T>(namespace: string, resource: string): Promise<T[]> {
        try {
            const result = await axios.get<RecordListContainer<T>>(`${BACKEND_URL}/records/${namespace}/${resource}?resolveReferences=*`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            return result.data.content.map(record => record.properties)
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

    export async function update<T extends Record>(namespace: string, resource: string, record: T): Promise<T> {
        try {
            const result = await axios.put<T>(`${BACKEND_URL}/records/${namespace}/${resource}/${record.id!}`, record, {
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
            await axios.delete(`${BACKEND_URL}/records/${namespace}/${resource}/${id}`, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })
        } catch (e) {
            await handleError(e)
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

    export async function findBy<T>(namespace: string, resource: string, property: string, value: any): Promise<T | undefined> {
        return findByMulti(namespace, resource, [{
            property: property,
            value: value
        }])
    }

    export async function findByMulti<T>(namespace: string, resource: string, conditions: {
        property: string,
        value: any
    }[]): Promise<T | undefined> {
        try {
            const query: BooleanExpression = {
                and: {
                    expressions: conditions.map(condition => ({
                        equal: {
                            left: {
                                property: condition.property
                            },
                            right: {
                                value: condition.value
                            }
                        }
                    }))
                }
            } as BooleanExpression
            const result = await axios.post<RecordListContainer<T>>(`${BACKEND_URL}/records/${namespace}/${resource}/_search`, {
                query: query
            }, {
                headers: {
                    Authorization: `Bearer ${await TokenService.get()}`
                }
            })

            if (result.data.content && result.data.content.length > 0) {
                return result.data.content[0].properties
            } else {
                return undefined
            }
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
