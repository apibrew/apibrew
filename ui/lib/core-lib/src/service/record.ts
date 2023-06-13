import axios from 'axios'
import {BACKEND_URL} from '../config'
import * as TokenService from './token'
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

        const result = await axios.get<RecordListContainer<T>>(`${BACKEND_URL}/records/${namespace}/${resource}?resolveReferences=*`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data.content.map(record => record.properties)

    }

    export async function create<T>(namespace: string, resource: string, record: T): Promise<T> {

        const result = await axios.post<T>(`${BACKEND_URL}/records/${namespace}/${resource}`, record, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function update<T extends Record>(namespace: string, resource: string, record: T): Promise<T> {
        const result = await axios.put<T>(`${BACKEND_URL}/records/${namespace}/${resource}/${record.id!}`, record, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }

    export async function remove(namespace: string, resource: string, id: string): Promise<void> {

        await axios.delete(`${BACKEND_URL}/records/${namespace}/${resource}/${id}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })
    }

    export async function get<T>(namespace: string, resource: string, id: string): Promise<T> {

        const result = await axios.get<T>(`${BACKEND_URL}/records/${namespace}/${resource}/${id}`, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

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
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        if (result.data.content && result.data.content.length > 0) {
            return result.data.content[0].properties
        } else {
            return undefined
        }

    }

    export async function apply<T>(namespace: string, resource: string, record: T): Promise<T> {

        const result = await axios.patch<T>(`${BACKEND_URL}/records/${namespace}/${resource}`, record, {
            headers: {
                Authorization: `Bearer ${TokenService.get()}`
            }
        })

        return result.data

    }
}
