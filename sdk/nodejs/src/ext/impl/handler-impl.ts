import {Client} from "../../client";
import {ExtensionService} from "../extension-service";
import {ExtensionInfo} from "../../extension-info";
import {Event} from "../../model/extension";
import {Entity} from "../../entity";
import {EntityInfo} from "../../entity-info";
import {Handler} from "../handler";
import {Condition} from "../condition";
import {Operator} from "../operator";
import {Record} from "../../model/record";

type Predicate<T extends Entity> = (event: Event, entity: T) => boolean;

export class HandlerImpl<T extends Entity> implements Handler<T> {
    constructor(private client: Client,
                private extensionService: ExtensionService,
                private entityInfo: EntityInfo,
                private extensionInfo: ExtensionInfo = {
                    sync: true,
                    responds: true,
                } as ExtensionInfo,
                private predicates: Predicate<any>[] = []) {
    }

    public withExtensionInfo(extensionInfo: ExtensionInfo): Handler<T> {
        return new HandlerImpl(this.client, this.extensionService, this.entityInfo, {...this.extensionInfo}, [...this.predicates]);
    }

    configure(configurer: (info: ExtensionInfo) => ExtensionInfo): Handler<T> {
        return this.withExtensionInfo(configurer(this.extensionInfo));
    }

    localOperator(localOperator: (event: Event, entity: T) => T): string {
        return this.extensionService.registerExtensionWithOperator(this.extensionInfo, (event, record) => {
            if (!this.checkPredicates(event, record.properties as T)) {
                return record;
            }

            return {
                properties: localOperator(event, record.properties as T) as object,
            } as Record;
        });
    }

    public operate(operator: Operator<T>): string {
        return operator.operate(this);
    }

    private checkPredicates(event: Event, entity: T): boolean {
        for (const predicate of this.predicates) {
            if (!predicate(event, entity)) {
                return false;
            }
        }

        return true;
    }

    unRegister(id: string): void {
        this.extensionService.unRegisterOperator(id);
    }

    when(condition: Condition<T>): Handler<T> {
        return new HandlerImpl(this.client, this.extensionService, this.entityInfo, {...condition.configureExtensionInfo(this.extensionInfo)}, [...this.predicates, condition.eventMatches.bind(condition)]);
    }
}
