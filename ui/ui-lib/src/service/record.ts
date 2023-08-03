import axios from 'axios'
import { BACKEND_URL } from '../config'
import * as TokenService from './token'
import { BooleanExpression } from "../model";
import { RecordApi, type Resource } from '@apibrew/client';
import { GetOptions, ListOptions } from '@apibrew/client/dist/api/record';
import { ServiceConfigProvider } from './service-config';

export interface Record {
    id?: string

    [key: string]: any
}

export namespace RecordService {
    interface RecordListContainer<T> {
        content: T[]
    }

    export async function list<T>(namespace: string, resource: string, options?: ListOptions): Promise<T[]> {
        const result = await RecordApi.list<T>(ServiceConfigProvider(), namespace, resource, options)

        return result.content
    }

    export async function create<T>(namespace: string, resource: string, record: T): Promise<T> {
        return RecordApi.create<T>(ServiceConfigProvider(), namespace, resource, record)
    }

    export async function update<T extends Record>(namespace: string, resource: string, record: T): Promise<T> {
        return RecordApi.update<T>(ServiceConfigProvider(), namespace, resource, record)
    }

    export async function remove(namespace: string, resource: string, id: string): Promise<void> {
        return RecordApi.remove(ServiceConfigProvider(), namespace, resource, id)
    }

    export async function get<T>(namespace: string, resource: string, id: string, options?: GetOptions): Promise<T> {
        return RecordApi.get<T>(ServiceConfigProvider(), namespace, resource, id, options)
    }

    export async function resource(namespace: string, resource: string): Promise<Resource> {
        return RecordApi.resource(ServiceConfigProvider(), namespace, resource)
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
        return RecordApi.findByMulti<T>(ServiceConfigProvider(), namespace, resource, conditions)
    }

    export async function apply<T>(namespace: string, resource: string, record: T): Promise<T> {
        return RecordApi.apply<T>(ServiceConfigProvider(), namespace, resource, record)
    }
}
