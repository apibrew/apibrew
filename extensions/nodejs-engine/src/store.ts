import {APBR_ADDR} from "./config";
import axios from "axios";

export interface Store {
    [key: string]: any
}

const store: Store = {}

export async function load<T>(namespace: string, resourceName: string) {
    console.log('loading:', namespace, resourceName)
    const response = await axios.get<{
        content: { properties: T }[]
    }>(`http://${APBR_ADDR}/records/${namespace}/${resourceName}`)

    store[namespace + '/' + resourceName] = response.data.content.map(record => record.properties)

    console.log('loaded:', namespace, resourceName)
    console.debug(store[namespace + '/' + resourceName])
}

export function read<T>(namespace: string, resourceName: string): T[] {
    return store[namespace + '/' + resourceName]
}

export function filter<T>(namespace: string, resourceName: string, predicate: (record: T) => boolean) {
    store[namespace + '/' + resourceName] = store[namespace + '/' + resourceName].filter(predicate)
}
