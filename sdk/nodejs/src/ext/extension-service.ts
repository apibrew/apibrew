import {ExtensionInfo} from "../extension-info";
import {Operator} from "./operator";
import {EntityInfo} from "../entity-info";
import {Handler} from "./handler";
import {Entity} from "../entity";
import {Event} from "../model";
import {Record} from "../model/record";

export interface ExtensionService {
    handler<T extends Entity>(entityInfo: EntityInfo): Handler<T>;

    run(): Promise<void>;

    close(): void;

    registerExtensionWithOperator(extensionInfo: ExtensionInfo, localOperator: (event: Event, entity: Record) => Record): string;

    unRegisterOperator(id: string): void;

    registerPendingItems(): Promise<void>;
}

