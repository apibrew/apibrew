import { Record, RecordListContainer, RecordResourceInfo } from "./model/record";
import { RecordService } from "./service";
import { ServiceConfig, ServiceConfigProvider } from "./service/config";
import { SearchRecordParams } from "./service/record";

export class Repository<T extends Record<unknown>> {
    constructor(private configProvider: ServiceConfigProvider, private recordResourceInfo: RecordResourceInfo) {
    }

    async create(entity: T): Promise<T> {
        return RecordService.create<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async update(entity: T): Promise<T> {
        return RecordService.update<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async apply(entity: T): Promise<T> {
        return RecordService.apply<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, entity)
    }

    async list(): Promise<RecordListContainer<T>> {
        return RecordService.list<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource)
    }

    async get(id: string): Promise<T> {
        return RecordService.get<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, id)
    }

    async remove(id: string): Promise<void> {
        return RecordService.remove(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, id)
    }

    async search(params: SearchRecordParams): Promise<RecordListContainer<T>> {
        return RecordService.search<T>(this.configProvider(), params)
    }

    async findBy(property: string, value: any): Promise<T | undefined> {
        return RecordService.findBy<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, property, value)
    }

    async findByMulti(conditions: {
        property: string,
        value: any
    }[]): Promise<T | undefined> {
        return RecordService.findByMulti<T>(this.configProvider(), this.recordResourceInfo.namespace, this.recordResourceInfo.resource, conditions)
    }
}
