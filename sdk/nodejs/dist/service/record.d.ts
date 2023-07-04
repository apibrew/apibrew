import { BooleanExpression, RecordListContainer } from "../model";
import { ServiceConfig } from './config';
import { Record } from '../model';
export declare function list<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string): Promise<RecordListContainer<T>>;
export declare function create<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T>;
export declare function update<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T>;
export declare function remove(config: ServiceConfig, namespace: string, resource: string, id: string): Promise<void>;
export declare function get<T>(config: ServiceConfig, namespace: string, resource: string, id: string): Promise<T>;
export declare function findBy<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, property: string, value: any): Promise<T | undefined>;
export declare function findByMulti<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, conditions: {
    property: string;
    value: any;
}[]): Promise<T | undefined>;
export interface SearchRecordParams {
    namespace: string;
    resource: string;
    query?: BooleanExpression;
    limit?: number;
    offset?: number;
    useHistory: boolean;
    resolveReferences: string[];
    annotations: {
        [key: string]: string;
    };
}
export declare function search<T extends Record<unknown>>(config: ServiceConfig, params: SearchRecordParams): Promise<RecordListContainer<T>>;
export declare function apply<T extends Record<unknown>>(config: ServiceConfig, namespace: string, resource: string, record: T): Promise<T>;
