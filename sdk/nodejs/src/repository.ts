import { Record, RecordListContainer, RecordResourceInfo } from "./model/record";
import { RecordApi } from "./api";
import { ServiceConfig, ServiceConfigProvider } from "./api/config";
import { SearchRecordParams } from "./api/record";

export class Repository<T extends Record<unknown>> {
    constructor(private configProvider: ServiceConfigProvider, private recordResourceInfo: RecordResourceInfo) {
    }

    async create(entity: T): Promise<T> {
        return RecordApi.create<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async update(entity: T): Promise<T> {
        return RecordApi.update<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async apply(entity: T): Promise<T> {
        return RecordApi.apply<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async list(): Promise<RecordListContainer<T>> {
        return RecordApi.list<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource)
    }

    async get(id: string): Promise<T> {
        return RecordApi.get<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, id)
    }

    async remove(id: string): Promise<void> {
        return RecordApi.remove(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, id)
    }

    async search(params: SearchRecordParams): Promise<RecordListContainer<T>> {
        return RecordApi.search<T>(this.configProvider(), params)
    }

    async findBy(property: string, value: any): Promise<T | undefined> {
        return RecordApi.findBy<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, property, value)
    }

    async findByMulti(conditions: {
        property: string,
        value: any
    }[]): Promise<T | undefined> {
        return RecordApi.findByMulti<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, conditions)
    }
}
