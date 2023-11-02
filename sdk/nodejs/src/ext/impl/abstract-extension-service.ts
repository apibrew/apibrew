import {ExtensionService} from "../extension-service";
import {Client} from "../../client";
import {Entity} from "../../entity";
import {EntityInfo} from "../../entity-info";
import {Handler} from "../handler";
import {HandlerImpl} from "./handler-impl";
import {ExtensionInfo, extensionInfoToString, toExtension} from "../../extension-info";
import {Event} from "../../model";
import {Extension, ExtensionEntityInfo, ExternalCall} from "../../model/extension";
import {Repository} from "../../repository";
import {Record} from "../../model/record";
import {randomUUID} from "crypto";

export abstract class AbstractExtensionService implements ExtensionService {
    private extensionRepo: Repository<Extension>;

    protected constructor(protected serviceName: string, protected client: Client) {
        this.extensionRepo = client.repo(ExtensionEntityInfo)
    }

    readonly extensionInfoSet: Set<ExtensionInfo> = new Set();
    readonly registeredExtensionInfoSet: Set<ExtensionInfo> = new Set();
    readonly extensionInfoIdMap: Map<string, ExtensionInfo> = new Map();
    readonly extensionHandlerMap: Map<ExtensionInfo, ((event: Event, entity: Record) => Record)[]> = new Map();
    readonly operatorMap: Map<string, (event: Event, record: Record) => Record> = new Map();
    readonly operatorIdExtensionInfoMap: Map<string, ExtensionInfo> = new Map();

    handler<T extends Entity>(entityInfo: EntityInfo): Handler<T> {
        return new HandlerImpl(this.client, this, entityInfo);
    }

    protected async registerExtensions(): Promise<void> {
        console.debug(`ExtensionService: ${this.serviceName} / Registering extensions`);
        this.extensionInfoSet.forEach(async extensionInfo => {
            if (!this.registeredExtensionInfoSet.has(extensionInfo)) {
                this.registeredExtensionInfoSet.add(extensionInfo);

                await this.registerExtension(extensionInfo);
            }
        });
        console.debug(`ExtensionService: ${this.serviceName} / Registered extensions`);
    }

    async registerExtension(extensionInfo: ExtensionInfo): Promise<void> {
        console.debug(`ExtensionService: ${this.serviceName} / Registering extension: ${extensionInfoToString(extensionInfo)}`);

        let extension = toExtension(extensionInfo);
        extension.call = this.prepareExternalCall();
        extension.name = `${this.serviceName}/${extension.name}`;
        extension = await this.extensionRepo.apply(extension);

        this.extensionInfoIdMap.set(extension.id.toString(), extensionInfo);

        console.debug(`ExtensionService: ${this.serviceName} / Registered extension: ${extensionInfoToString(extensionInfo)}`);
    }

    protected abstract prepareExternalCall(): ExternalCall;

    protected processEvent(event: Event): Event {
        console.debug(`ExtensionService: ${this.serviceName} / Begin processing event: ${JSON.stringify(event)}`);

        if (event.annotations == null) {
            event.annotations = {};
        }

        const extensionId = event.annotations["ExtensionId"];
        const extensionInfo = this.extensionInfoIdMap.get(extensionId);

        console.debug(`ExtensionService: ${this.serviceName} / Event ID: ${event.id} => Extension ID: ${extensionId}`);
        console.debug(`ExtensionService: ${this.serviceName} / ExtensionInfo: ${extensionInfoToString(extensionInfo)}`);

        if (extensionInfo == null) {
            console.warn(`ExtensionInfo not found for event: ${event}`);
            return event;
        }

        const eventChain = this.processEventChain(extensionInfo, event);

        console.debug(`ExtensionService: ${this.serviceName} / End processing event: ${JSON.stringify(event)}`);

        return eventChain;
    }

    private processEventChain(extensionInfo: ExtensionInfo, eventChain: Event): Event {
        const handlers = this.extensionHandlerMap.get(extensionInfo);

        if (handlers != null) {
            for (const handler of handlers) {
                const records = eventChain.records;
                let handlerHandled = false;
                if (records != null) {
                    const eventChainRecords = eventChain.records;
                    const processedRecords: Record[] = [];

                    for (const record of eventChainRecords) {
                        console.debug(`ExtensionService: ${this.serviceName} / Processing record: ${record.id}`);
                        const processedRecord = handler(eventChain, record);
                        if (processedRecord != null) {
                            processedRecords.push(processedRecord);
                        }
                        handlerHandled = true;
                    }

                    eventChain.records = processedRecords;
                }

                if (!handlerHandled) {
                    handler(eventChain, {} as any);
                }
            }
        }

        return eventChain;
    }

    registerExtensionWithOperator(extensionInfo: ExtensionInfo, localOperator: (event: Event, entity: Record) => Record): string {
        const id = randomUUID().toString();
        this.extensionInfoSet.add(extensionInfo);

        if (!this.extensionHandlerMap.get(extensionInfo)) {
            this.extensionHandlerMap.set(extensionInfo, []);
        }

        this.extensionHandlerMap.get(extensionInfo)!.push(localOperator);

        this.operatorMap.set(id, localOperator);

        this.operatorIdExtensionInfoMap.set(id, extensionInfo);

        console.info(`ExtensionService: ${this.serviceName} / Registered operator: ${id} => ${extensionInfo}`);

        return id;
    }

    registerPendingItems(): Promise<void> {
        return this.registerExtensions();
    }

    run(): Promise<void> {
        return Promise.resolve(undefined);
    }

    unRegisterOperator(id: string): void {
        const operator = this.operatorMap.get(id);

        if (operator == null) {
            throw new Error(`Operator not found for id: ${id}`);
        }

        const extensionInfo = this.operatorIdExtensionInfoMap.get(id)!;
        this.extensionHandlerMap.get(extensionInfo)!.splice(this.extensionHandlerMap.get(extensionInfo)!.indexOf(operator), 1);

        if (this.extensionHandlerMap.get(extensionInfo)!.length === 0) {
            this.extensionHandlerMap.delete(extensionInfo);
            this.extensionInfoSet.delete(extensionInfo);
        }

        this.operatorMap.delete(id);
        this.operatorIdExtensionInfoMap.delete(id);

        console.info(`Unregistered operator: ${id} => ${extensionInfo}`);
    }

    abstract close(): void;

}
