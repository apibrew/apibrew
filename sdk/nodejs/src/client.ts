import {Resource} from "./model/resource";
import {Entity} from "./entity";
import {EntityInfo} from "./entity-info";
import {Repository} from "./repository";
import {Event} from "./model/extension";
import {Container} from "./container";
import {ClientImpl} from "./impl/client-impl";
import {Server} from "./config";
import {GetRecordParams} from "./get-record-params";
import {ListRecordParams} from "./list-record-params";

export interface Client {
    applyResource(resource: Resource): Promise<Resource>

    getResourceByName(namespace: string, name: string): Promise<Resource>

    listResources(): Promise<Resource[]>;

    createResource(resource: Resource): Promise<Resource>;

    updateResource(resource: Resource): Promise<Resource>;

    deleteResource(resource: Resource): Promise<void>;

    authenticateWithToken(token: string): void;

    authenticateWithUsernameAndPassword(username: string, password: string): Promise<void>;

    newClientAuthenticateWithToken(token: string): Client;

    newClientAuthenticateWithUsernameAndPassword(username: string, password: string): Promise<Client>;

    repo<T extends Entity>(entityInfo: EntityInfo): Repository<T>;

    repository<T extends Entity>(entityInfo: EntityInfo): Repository<T>;

    listRecords<T extends Entity>(entityInfo: EntityInfo, params: ListRecordParams): Promise<Container<T>>;

    applyRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T>;

    deleteRecord<T extends Entity>(entityInfo: EntityInfo, id: string): Promise<T>;

    updateRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T>;

    getRecord<T extends Entity>(entityInfo: EntityInfo, params: GetRecordParams): Promise<T>;

    createRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T>;

    loadRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T>;

    writeEvent(channelKey: string, event: Event): Promise<void>;

    bypassExtensions(bypassExtensions: boolean): void;

    headers(): { [key: string]: string };

    getUrl(): string;
}

export async function newClient(serverName?: string): Promise<Client> {
    return ClientImpl.newClientByServerName(serverName)
}

export async function newClientByServerConfig(serverConfig: Server): Promise<Client> {
    return ClientImpl.newClientByServerConfig(serverConfig)
}
