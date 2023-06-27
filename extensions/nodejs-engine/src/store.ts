import {APBR_ADDR, TOKEN} from "./config";
import axios from "axios";

export interface Store {
    [key: string]: any
}

const store: Store = {}

export async function loadSystem<T>(namespace: string, resourceName: string) {
    console.log('loading:', namespace, resourceName)
    const response = await axios.get<{
        content: { properties: T }[]
    }>(`http://${APBR_ADDR}/${namespace}/${resourceName}`, {
        headers: {
            'Authorization': 'Bearer ' + TOKEN
        }
    })

    store[namespace + '/' + resourceName] = response.data.content

    console.log('loaded:', namespace, resourceName)
}

export async function load<T>(namespace: string, resourceName: string) {
    if (namespace == 'system') {
        return loadSystem(namespace, resourceName)
    }
    console.log('loading:', namespace, resourceName)
    const response = await axios.get<{
        content: { properties: T }[]
    }>(`http://${APBR_ADDR}/records/${namespace}/${resourceName}`, {
        headers: {
            'Authorization': 'Bearer ' + TOKEN
        }
    })

    store[namespace + '/' + resourceName] = response.data.content.map(record => record.properties)

    console.log('loaded:', namespace, resourceName, response.data.content.length)
}

export function getResourcePath(namespace: string, resourceName: string) {
    if (namespace == 'system') {
        return `http://${APBR_ADDR}/${namespace}/${resourceName}`
    }

    return `http://${APBR_ADDR}/records/${namespace}/${resourceName}`
}

export async function create<T>(namespace: string, resourceName: string, record: T): Promise<T> {
    console.log('creating:', namespace, resourceName)

    const response = await axios.post<{
        content: { properties: T }[]
    }>(getResourcePath(namespace, resourceName), record, {
        headers: {
            'Authorization': 'Bearer ' + TOKEN
        }
    })

    console.log('created:', namespace, resourceName, response.data)

    return response.data as T
}

export async function apply<T>(namespace: string, resourceName: string, record: T): Promise<T>  {
    console.log('updating:', namespace, resourceName)
    const response = await axios.patch<{
        content: { properties: T }[]
    }>(getResourcePath(namespace, resourceName), record, {
        headers: {
            'Authorization': 'Bearer ' + TOKEN
        }
    })

    console.log('applied:', namespace, resourceName)

    return response.data as T
}

export async function update<T>(namespace: string, resourceName: string, record: T) {
    console.log('updating:', namespace, resourceName)
    const response = await axios.put<{
        content: { properties: T }[]
    }>(getResourcePath(namespace, resourceName), record, {
        headers: {
            'Authorization': 'Bearer ' + TOKEN
        }
    })

    console.log('updated:', namespace, resourceName)
}

export function read<T>(namespace: string, resourceName: string): T[] {
    return store[namespace + '/' + resourceName]
}

export function filter<T>(namespace: string, resourceName: string, predicate: (record: T) => boolean) {
    store[namespace + '/' + resourceName] = store[namespace + '/' + resourceName].filter(predicate)
}
