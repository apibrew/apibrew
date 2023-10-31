import {ExtensionInfo} from "../extension-info";
import {Operator} from "./operator";
import {EntityInfo} from "../entity-info";
import {Handler} from "./handler";
import {Entity} from "../entity";

export interface ExtensionService {
    handler<T extends Entity>(entityInfo: EntityInfo): Handler<T>;

    run(): Promise<void>;

    registerExtensionWithOperator<T extends Entity>(extensionInfo: ExtensionInfo, operator: Operator<T>): string;

    unRegisterOperator(id: string): void;

    registerPendingItems(): Promise<void>;
}

