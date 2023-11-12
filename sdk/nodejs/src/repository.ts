import {Entity} from "./entity";
import {Container} from "./container";
import {BooleanExpression, Event} from "./model/extension";
import {Watcher} from "./watcher";
import {ListRecordParams} from "./list-record-params";

export interface Repository<T extends Entity> {
    create(record: T): Promise<T>;

    get(id: string, resolveReferences?: string[]): Promise<T>;

    update(record: Partial<T> & Entity): Promise<T>;

    delete(id: string): Promise<T>;

    apply(record: Partial<T>): Promise<T>;

    load(record: Partial<T>, resolveReferences?: string[]): Promise<T>;

    list(params?: ListRecordParams): Promise<Container<T>>;

    watch(eventConsumer: (event: Event) => void): Watcher<T>;
}