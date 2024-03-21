import {Entity} from "../entity";
import {ExtensionInfo} from "../extension-info";
import {Event} from "../model/extension";
import {Condition} from "./condition";
import {Operator} from "./operator";

export interface Handler<T extends Entity> {
    when(condition: Condition<T>): Handler<T>;

    configure(configurer: (info: ExtensionInfo) => ExtensionInfo): Handler<T>;

    operate(operator: Operator<T>): string;

    localOperator(localOperator: (event: Event, entity: T) => T): string;

    unRegister(id: string): void;
}
