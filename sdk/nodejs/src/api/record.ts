import axios from 'axios'
import {BooleanExpression, RecordListContainer, Resource} from "../model";
import {ServiceConfig} from './config';
import {Record} from '../model';

export interface ListOptions {
    resolveReferences?: string[]
}

export interface GetOptions {
    resolveReferences?: string[]
}

export async function list<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, options?: ListOptions): Promise<RecordListContainer<T>> {
    let url = resourceUrl(config, namespace, resource)

    if (options) {
        if (options.resolveReferences && options.resolveReferences.length > 0) {
            url += `?resolve-references=${encodeURIComponent(options.resolveReferences.join(','))}`
        }
    }

    const result = await axios.get<RecordListContainer<T>>(url, {
        headers: prepareHeaders(config)
    })

    return result.data

}

function resourceUrl(config: ServiceConfig, namespace: string, resource: string): string {
    return `${config.backendUrl}/${namespace.toLowerCase()}-${resource.toLowerCase()}`
}

function prepareHeaders(config: ServiceConfig) {
    const headers: { [key: string]: string | string[] } = {}

    if (config.token !== '') {
        headers["Authorization"] = `Bearer ${config.token}`
    }

    return headers
}

export async function create<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {

    const result = await axios.post<T>(resourceUrl(config, namespace, resource), record, {
        headers: prepareHeaders(config)
    })

    return result.data

}

export async function update<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {
    const result = await axios.put<T>(`${resourceUrl(config, namespace, resource)}/${record.id!}`, record, {
        headers: prepareHeaders(config)
    })

    return result.data

}

export async function remove(config: ServiceConfig, namespace: string, resource: string, id: string): Promise<void> {

    await axios.delete(`${resourceUrl(config, namespace, resource)}/${id}`, {
        headers: prepareHeaders(config)
    })
}

export async function get<T>(config: ServiceConfig, namespace: string, resource: string, id: string, options?: GetOptions): Promise<T> {
    let url = `${resourceUrl(config, namespace, resource)}/${id}`

    if (options) {
        if (options.resolveReferences && options.resolveReferences.length > 0) {
            url += `?resolve-references=${encodeURIComponent(options.resolveReferences.join(','))}`
        }
    }

    const result = await axios.get<T>(url, {
        headers: prepareHeaders(config)
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
}[], options?: GetOptions): Promise<T | undefined> {

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
    const result = await axios.post<RecordListContainer<T>>(`${resourceUrl(config, namespace, resource)}/_search?resolve-references=*`, {// fixme
        query: query
    }, {
        headers: prepareHeaders(config)
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
        headers: prepareHeaders(config)
    })

    return result.data
}

export async function resource(config: ServiceConfig, namespace: string, resource: string): Promise<Resource> {
    const result = await axios.get<Resource>(`${resourceUrl(config, namespace, resource)}/_resource`, {
        headers: prepareHeaders(config)
    })

    return result.data
}

export async function apply<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T> {
    const result = await axios.patch<T>(resourceUrl(config, namespace, resource), record, {
        headers: prepareHeaders(config)
    })

    return result.data

}