import {Client} from '../client';
import {Resource} from "../model/resource";
import axios, {AxiosResponse} from "axios";
import {ApiException} from "../api-exception";
import {BooleanExpression, Code, Event} from "../model/extension";
import {AuthenticationResponse, TokenTerm} from "../model/token";
import {Container} from "../container";
import {Entity} from "../entity";
import {EntityInfo} from "../entity-info";
import {BooleanExpressionBuilder} from "../boolean-expression-builder";
import {Repository} from "../repository";
import {RepositoryImpl} from "./repository-impl";
import {ConfigLoader} from "../config-loader";
import {Server} from "../config";
import {ListRecordParams} from "../list-record-params";
import {GetRecordParams} from "../get-record-params";

export class Urls {
    static resourceUrl(url: string) {
        return `${url}/resources`;
    }

    static resourceByName(url: string, namespace: string, name: string) {
        return `${Urls.resourceUrl(url)}/by-name/${namespace}/${name}`;
    }

    static resourceById(url: string, id: string) {
        return `${Urls.resourceUrl(url)}/${id}`;
    }

    static recordUrl(url: string, restPath: string) {
        return `${url}/${restPath}`;
    }

    static recordSearchUrl(url: string, restPath: string) {
        return `${url}/${restPath}/_search`;
    }

    static recordWatchUrl(url: string, restPath: string) {
        return `${url}/${restPath}/_watch`;
    }

    static recordByIdUrl(url: string, restPath: string, id: string) {
        return `${url}/${restPath}/${id}`;
    }

    static recordActionByIdUrl(url: string, restPath: string, id: string, action: string) {
        return `${url}/${restPath}/${id}/_${action}`;
    }

    static authenticate(url: string) {
        return `${url}/authentication/token`;
    }

    static eventsUrl(url: string) {
        return `${url}/_events`;
    }
}

export class ClientImpl implements Client {
    private token?: string;
    private bypassExtensionsEnabled: boolean = false;

    constructor(private url: string) {
    }

    static ensureResponseSuccess(resp: AxiosResponse) {
        if (resp.status != 200) {
            const error = resp.data;

            console.error("Error response: ", error)

            if (error.code) {
                throw ApiException.fromError(error);
            } else {
                throw new ApiException(Code.INTERNAL_ERROR, resp.statusText);
            }
        }
    }

    public async applyResource(resource: Resource): Promise<Resource> {
        const resp = await axios.post<Resource>(Urls.resourceUrl(this.url), resource, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        const existsStatus = resp.status;

        if (existsStatus == 200) {
            return this.updateResource(resource);
        } else if (existsStatus == 404) {
            return this.createResource(resource);
        } else {
            ClientImpl.ensureResponseSuccess(resp);
            throw new Error("Unreachable");
        }
    }

    public async getResourceByName(namespace: string, name: string): Promise<Resource> {
        const resp = await axios.get<Resource>(Urls.resourceByName(this.url, namespace, name), {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async listResources(): Promise<Resource[]> {
        const resp = await axios.get<Resource[]>(Urls.resourceUrl(this.url), {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async createResource(resource: Resource): Promise<Resource> {
        const resp = await axios.post<Resource>(Urls.resourceUrl(this.url), resource, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async updateResource(resource: Resource): Promise<Resource> {
        const resp = await axios.post<Resource>(Urls.resourceById(this.url, resource.id!), resource, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async deleteResource(resource: Resource): Promise<void> {
        const resp = await axios.delete(Urls.resourceById(this.url, resource.id!), {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);
    }

    public async authenticateWithToken(token: string): Promise<void> {
        this.token = token;
    }

    public async authenticateWithUsernameAndPassword(username: string, password: string, term: TokenTerm = TokenTerm.VERY_LONG): Promise<void> {
        const resp = await axios.post<AuthenticationResponse>(Urls.authenticate(this.url), {
            username: username,
            password: password,
            term: term,
        }, {
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        this.token = resp.data.token.content;
    }

    public newClientAuthenticateWithToken(token: string): Client {
        const client = new ClientImpl(this.url);
        client.bypassExtensionsEnabled = this.bypassExtensionsEnabled;
        client.authenticateWithToken(token);

        return client;
    }

    public async newClientAuthenticateWithUsernameAndPassword(username: string, password: string, term: TokenTerm = TokenTerm.VERY_LONG): Promise<Client> {
        const client = new ClientImpl(this.url);
        client.bypassExtensionsEnabled = this.bypassExtensionsEnabled;
        await client.authenticateWithUsernameAndPassword(username, password, term);

        return client;
    }

    public headers() {
        const headers: { [key: string]: string } = {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + this.token
        }

        if (this.bypassExtensionsEnabled) {
            headers["BypassExtensions"] = "true";
        }

        return headers;
    }

    public getUrl(): string {
        return this.url;
    }

    public bypassExtensions(bypassExtensions: boolean): void {
        this.bypassExtensionsEnabled = bypassExtensions;
    }

    public async listRecords<T extends Entity>(entityInfo: EntityInfo, params: ListRecordParams): Promise<Container<T>> {
        if (!params) {
            params = {}
        }

        if (!params.query) {
            const resp = await axios.get<Container<T>>(Urls.recordUrl(this.url, entityInfo.restPath), {
                headers: this.headers(),
                validateStatus: (status) => true,
            });

            ClientImpl.ensureResponseSuccess(resp)

            return resp.data;
        } else {
            const resp = await axios.post<Container<T>>(Urls.recordSearchUrl(this.url, entityInfo.restPath), params, {
                headers: this.headers(),
                validateStatus: (status) => true,
            });

            ClientImpl.ensureResponseSuccess(resp)

            return resp.data;
        }
    }

    public async applyRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T> {
        const resp = await axios.patch<T>(Urls.recordUrl(this.url, entityInfo.restPath), record, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp)

        return resp.data;
    }

    public async deleteRecord<T extends Entity>(entityInfo: EntityInfo, id: string): Promise<T> {
        const resp = await axios.delete<T>(Urls.recordByIdUrl(this.url, entityInfo.restPath, id), {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp)

        return resp.data;
    }

    public async updateRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T> {
        const resp = await axios.put<T>(Urls.recordByIdUrl(this.url, entityInfo.restPath, record.id!), record, {
            headers: this.headers(),
            validateStatus: (status) => true,
        }).catch();

        ClientImpl.ensureResponseSuccess(resp)

        return resp.data;
    }

    public async getRecord<T extends Entity>(entityInfo: EntityInfo, params: GetRecordParams): Promise<T> {
        let finalUrl = Urls.recordByIdUrl(this.url, entityInfo.restPath, params.id);

        if (params.resolveReferences && params.resolveReferences.length > 0) {
            finalUrl += "?resolve-references=" + params.resolveReferences.join(",");
        }

        const resp = await axios.get<T>(finalUrl, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async createRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T> {
        const resp = await axios.post<T>(Urls.recordUrl(this.url, entityInfo.restPath), record, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp)

        return resp.data;
    }

    public async loadRecord<T extends Entity>(entityInfo: EntityInfo, record: T): Promise<T> {
        if (record.id) {
            return this.getRecord(entityInfo, {
                id: record.id,
            });
        }

        const conditions = Object.entries(record).map(([key, value]) => BooleanExpressionBuilder.eq(key, value));

        const resp = await axios.post<Container<T>>(Urls.recordSearchUrl(this.url, entityInfo.restPath), {
            query: BooleanExpressionBuilder.and(...conditions)
        }, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp)

        if (resp.data.total == 0) {
            throw new ApiException(Code.RECORD_NOT_FOUND, "No record found for " + record);
        } else if (resp.data.total > 1) {
            throw new ApiException(Code.RECORD_VALIDATION_ERROR, "Multiple records found for " + record);
        }

        return resp.data.content[0];
    }

    public async writeEvent(channelKey: string, event: Event): Promise<void> {
        const url = Urls.eventsUrl(this.url) + "?channelKey=" + channelKey;
        const resp = await axios.post(url, event, {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp)
    }

    repo<T extends Entity>(entityInfo: EntityInfo): Repository<T> {
        return this.repository(entityInfo);
    }

    repository<T extends Entity>(entityInfo: EntityInfo): Repository<T> {
        return new RepositoryImpl(this, entityInfo)
    }

    static newClient(url: string | undefined) {
        if (!url) {
            url = "http://localhost:8080";
        }

        return new ClientImpl(url);
    }

    static newClientByServerName(serverName?: string) {
        const config = ConfigLoader.load();

        if (!serverName) {
            serverName = config.defaultServer;
        }

        const serverConfig = config.servers.find(item => item.name == serverName);

        if (!serverConfig) {
            throw new Error("Server not found: " + serverName);
        }

        return ClientImpl.newClientByServerConfig(serverConfig!);
    }

    static async newClientByServerConfig(serverConfig: Server) {
        let httpPort = serverConfig.httpPort;

        if (httpPort == 0) {
            httpPort = serverConfig.port;
        }

        let addr = serverConfig.host + ":" + httpPort;

        if (!addr.startsWith("http")) {
            if (serverConfig.insecure) {
                addr = "http://" + addr;
            } else {
                addr = "https://" + addr;
            }
        }

        if (addr.endsWith("/")) {
            addr = addr.substring(0, addr.length - 1);
        }

        const client = ClientImpl.newClient(addr);

        if (serverConfig.authentication.token) {
            await client.authenticateWithToken(serverConfig.authentication.token);
        } else {
            await client.authenticateWithUsernameAndPassword(serverConfig.authentication.username, serverConfig.authentication.password);
        }

        return client;
    }
}
