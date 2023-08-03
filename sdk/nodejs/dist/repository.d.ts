import { Record, RecordListContainer, RecordResourceInfo } from "./model/record";
import { ServiceConfigProvider } from "./api/config";
import { SearchRecordParams } from "./api/record";
export declare class Repository<T extends Record<unknown>> {
    private configProvider;
    private recordResourceInfo;
    constructor(configProvider: ServiceConfigProvider, recordResourceInfo: RecordResourceInfo);
    create(entity: T): Promise<T>;
    update(entity: T): Promise<T>;
    apply(entity: T): Promise<T>;
    list(): Promise<RecordListContainer<T>>;
    get(id: string): Promise<T>;
    remove(id: string): Promise<void>;
    search(params: SearchRecordParams): Promise<RecordListContainer<T>>;
    findBy(property: string, value: any): Promise<T | undefined>;
    findByMulti(conditions: {
        property: string;
        value: any;
    }[]): Promise<T | undefined>;
}
