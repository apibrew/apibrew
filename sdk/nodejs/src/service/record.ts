import axios from 'axios'
import { BooleanExpression, RecordListContainer } from "../model";
import { ServiceConfig } from './config';
import { Record } from '../model';

export async function list<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string): Promise<RecordListContainer<T>> {

    const result = await axios.get<RecordListContainer<T>>(resourceUrl(config, namespace, resource), {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

function resourceUrl(config: ServiceConfig, namespace: string, resource: string): string {
    if (namespace === 'system') {
        return `${config.backendUrl}/${resource}?resolveReferences=*`;
    }

    return `${config.backendUrl}/records/${namespace}/${resource}?resolveReferences=*`;
}

export async function create<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {

    const result = await axios.post<T>(resourceUrl(config, namespace, resource), record, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function update<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {
    const result = await axios.put<T>(`${resourceUrl(config, namespace, resource)}/${record.id!}`, record, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function remove(config: ServiceConfig, namespace: string, resource: string, id: string): Promise<void> {

    await axios.delete(`${resourceUrl(config, namespace, resource)}/${id}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })
}

export async function get<T>(config: ServiceConfig, namespace: string, resource: string, id: string): Promise<T> {

    const result = await axios.get<T>(`${resourceUrl(config, namespace, resource)}/${id}`, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}

export async function findBy<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, property: string, value: any): Promise<T | undefined> {
    return findByMulti(config, namespace, resource, [{
        property: property,
        value: value
    }])
}

export async function findByMulti<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, conditions: {
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
    const result = await axios.post<RecordListContainer<T>>(`${resourceUrl(config, namespace, resource)}/_search`, {
        query: query
    }, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    if (result.data.content && result.data.content.length > 0) {
        return result.data.content[0]
    } else {
        return undefined
    }

}

export interface SearchRecordParams {
    namespace: string
    resource: string
    query?: BooleanExpression
    limit?: number
    offset?: number
    useHistory: boolean
    resolveReferences: string[]
    annotations: {
        [key: string]: string
    }
}

export async function search<T extends Record<unknown>>(config: ServiceConfig, params: SearchRecordParams): Promise<RecordListContainer<T>> {
    const result = await axios.post<RecordListContainer<T>>(`${resourceUrl(config, params.namespace, params.resource)}/_search`, params, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data
}

export async function apply<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {

    const result = await axios.patch<T>(resourceUrl(config, namespace, resource), record, {
        headers: {
            Authorization: `Bearer ${config.token}`
        }
    })

    return result.data

}