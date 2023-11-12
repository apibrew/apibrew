import {Client} from '../client';
import {Entity} from "../entity";
import {Repository} from "../repository";
import {EntityInfo} from "../entity-info";
import {Container} from "../container";
import {BooleanExpression, Event} from "../model/extension";
import {Watcher} from "../watcher";
import {WatcherImpl} from "./watcher-impl";
import {ListRecordParams} from "../list-record-params";

export class RepositoryImpl<T extends Entity> implements Repository<T> {

    public constructor(private client: Client, private entityInfo: EntityInfo) {
    }

    public create(record: T): Promise<T> {
        return this.client.createRecord<T>(this.entityInfo, record);
    }

    public get(id: string, resolveReferences?: string[] ): Promise<T> {
        return this.client.getRecord<T>(this.entityInfo, {id, resolveReferences});
    }

    public update(record: Partial<T> & Entity): Promise<T> {
        return this.client.updateRecord<T>(this.entityInfo, record);
    }

    public delete(id: string): Promise<T> {
        return this.client.deleteRecord<T>(this.entityInfo, id);
    }

    public apply(record: Partial<T>): Promise<T> {
        return this.client.applyRecord<T>(this.entityInfo, record);
    }

    public load(record: Partial<T>, resolveReferences?: string[]): Promise<T> {
        return this.client.loadRecord<T>(this.entityInfo, record, resolveReferences);
    }

    public list(params: ListRecordParams = {}): Promise<Container<T>> {
        return this.client.listRecords<T>(this.entityInfo, params);
    }

    public watch(eventConsumer: (event: Event) => void): Watcher<T> {
        return new WatcherImpl<T>(this.client, this.entityInfo, eventConsumer);
    }
}
