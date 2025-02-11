import {Client} from '../client';
import {Resource} from "../model/resource";
import axios, {AxiosResponse} from "axios";
import {ApiException} from "../api-exception";
import {Code, Event} from "../model/extension";
import {AuthenticationResponse, TokenTerm} from "../model/token";
import {Container} from "../container";
import {Entity} from "../entity";
import {EntityInfo} from "../entity-info";
import {BooleanExpressionBuilder} from "../boolean-expression-builder";
import {Repository} from "../repository";
import {RepositoryImpl} from "./repository-impl";
import {Server} from "../config";
import {ListRecordParams} from "../list-record-params";
import {GetRecordParams} from "../get-record-params";
import {TokenStorage} from '../token-storage';
import {TokenBody} from '../token-body';
import {decodeBase64} from "../util/base64";

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

    static recordWatchUrl(url: string, restPath: string, filters?: { [key: string]: string }) {
        let watchUrl = `${url}/${restPath}/_watch`;

        if (filters) {
            if (watchUrl) {
                watchUrl += '?' + Object.entries(filters).map(entry => encodeURIComponent(entry[0]) + '=' + encodeURIComponent(entry[1])).join('&')
            }
        }

        return watchUrl;
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

export class DefaultTokenStorage implements TokenStorage {
    private data: Map<string, string> = new Map();

    clear(): void {
        this.data.clear();
    }

    get(name: string): string | undefined {
        return this.data.get(name);
    }

    set(name: string, token: string): void {
        this.data.set(name, token);
    }

    list(): { name: string; token: string; }[] {
        const result: { name: string; token: string; }[] = [];

        this.data.forEach((value, key) => {
            result.push({
                name: key,
                token: value,
            })
        })

        return result;
    }

}

const ACCESS_TOKEN = 'ACCESS_TOKEN';
const REFRESH_TOKEN = 'REFRESH_TOKEN';

export class ClientImpl implements Client {
    private bypassExtensionsEnabled: boolean = false;
    private tokenStorage: TokenStorage = new DefaultTokenStorage();
    private tokenRefreshInterval?: NodeJS.Timer;

    constructor(private url: string) {
    }

    useTokenStorage(tokenStorage: TokenStorage): void {
        this.tokenStorage.list().forEach(token => {
            tokenStorage.set(token.name, token.token)
        })

        this.tokenStorage = tokenStorage

        this.setupTokenRefresher()
    }

    static ensureResponseSuccess(resp: AxiosResponse) {
        if (resp.status >= 400) {
            const error = resp.data;
            if (error) {
                throw ApiException.fromError(resp.data);
            } else {
                throw new ApiException(Code.INTERNAL_ERROR, resp.statusText);
            }
        }
    }

    public async applyResource(resource: Resource, forceMigrate?: boolean): Promise<Resource> {
        const resp = await axios.get<Resource>(Urls.resourceByName(this.url, resource.namespace.name, resource.name));

        const existsStatus = resp.status;

        if (existsStatus == 200) {
            return this.updateResource(resource, forceMigrate);
        } else if (existsStatus == 404) {
            return this.createResource(resource, forceMigrate);
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
        const resp = await axios.get<Container<Resource>>(Urls.resourceUrl(this.url), {
            headers: this.headers(),
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data.content;
    }

    public async createResource(resource: Resource, forceMigrate?: boolean): Promise<Resource> {
        const resp = await axios.post<Resource>(Urls.resourceUrl(this.url), resource, {
            headers: {
                ...this.headers(),
                "X-Force-Migrate": forceMigrate ? "true" : "false",
            },
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async updateResource(resource: Resource, forceMigrate?: boolean): Promise<Resource> {
        const resp = await axios.put<Resource>(Urls.resourceById(this.url, resource.id!), resource, {
            headers: {
                ...this.headers(),
                "X-Force-Migrate": forceMigrate ? "true" : "false",
            },
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);

        return resp.data;
    }

    public async deleteResource(resource: Resource, forceMigrate?: boolean): Promise<void> {
        const resp = await axios.delete(Urls.resourceById(this.url, resource.id!), {
            headers: {
                ...this.headers(),
                "X-Force-Migrate": forceMigrate ? "true" : "false",
            },
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(resp);
    }

    public async authenticateWithToken(token: string): Promise<void> {
        this.tokenStorage.set(ACCESS_TOKEN, token);
        this.tokenStorage.set(REFRESH_TOKEN, token);
        return this.refreshToken()
    }

    getCurrentToken(): string {
        return this.tokenStorage.get(ACCESS_TOKEN) || '';
    }

    public async authenticateWithUsernameAndPassword(username: string, password: string): Promise<void> {
        const refreshTokenResp = await axios.post<AuthenticationResponse>(Urls.authenticate(this.url), {
            username: username,
            password: password,
            term: TokenTerm.LONG
        }, {
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(refreshTokenResp);

        this.tokenStorage.set(REFRESH_TOKEN, refreshTokenResp.data.token.content)

        const accessTokenResp = await axios.post<AuthenticationResponse>(Urls.authenticate(this.url), {
            username: username,
            password: password,
            term: TokenTerm.LONG
        }, {
            validateStatus: (status) => true,
        });

        ClientImpl.ensureResponseSuccess(accessTokenResp);

        this.tokenStorage.set(ACCESS_TOKEN, accessTokenResp.data.token.content)

        this.setupTokenRefresher();
    }

    public newClientAuthenticateWithToken(token: string): Client {
        const client = new ClientImpl(this.url);
        client.useTokenStorage(this.tokenStorage)
        client.bypassExtensionsEnabled = this.bypassExtensionsEnabled;
        client.authenticateWithToken(token);

        return client;
    }

    public async newClientAuthenticateWithUsernameAndPassword(username: string, password: string): Promise<Client> {
        const client = new ClientImpl(this.url);
        client.useTokenStorage(this.tokenStorage)
        client.bypassExtensionsEnabled = this.bypassExtensionsEnabled;
        await client.authenticateWithUsernameAndPassword(username, password);

        return client;
    }

    public headers() {
        const headers: { [key: string]: string } = {
            "Content-Type": "application/json",
        }

        if (this.tokenStorage.get(ACCESS_TOKEN)) {
            headers["Authorization"] = "Bearer " + this.tokenStorage.get(ACCESS_TOKEN);
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

        if (!params.query && !params.sorting && !params.aggregation) {
            let url = Urls.recordUrl(this.url, entityInfo.restPath)
            const queryParams: { [key: string]: string } = {}

            if (params) {
                if (params.resolveReferences && params.resolveReferences.length > 0) {
                    queryParams['resolve-references'] = params.resolveReferences.join(',')
                }

                if (params.useHistory) {
                    queryParams['use-history'] = 'true'
                }

                if (params.limit && params.limit > 0) {
                    queryParams['limit'] = params.limit.toString()
                }

                if (params.offset && params.offset > 0) {
                    queryParams['offset'] = params.offset.toString()
                }

                if (params.filters) {
                    Object.entries(params.filters).forEach(filter => {
                        queryParams[filter[0]] = filter[1]
                    })
                }
            }

            url = url + '?' + Object.entries(queryParams).map(entry => entry[0] + '=' + entry[1]).join('&')

            const resp = await axios.get<Container<T>>(url, {
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

    public async applyRecord<T extends Entity>(entityInfo: EntityInfo, record: Partial<T>): Promise<T> {
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

    public async updateRecord<T extends Entity>(entityInfo: EntityInfo, record: Partial<T> & Entity): Promise<T> {
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

    public async loadRecord<T extends Entity>(entityInfo: EntityInfo, record: Partial<T>, resolveReferences?: string[]): Promise<T> {
        if (record.id) {
            return this.getRecord(entityInfo, {
                id: record.id,
                resolveReferences: resolveReferences,
            });
        }

        const conditions = Object.entries(record).map(([key, value]) => BooleanExpressionBuilder.eq(key, value));

        const resp = await axios.post<Container<T>>(Urls.recordSearchUrl(this.url, entityInfo.restPath), {
            query: BooleanExpressionBuilder.and(...conditions),
            resolveReferences: resolveReferences,
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

        if (serverConfig.authentication) {
            if (serverConfig.authentication.token) {
                await client.authenticateWithToken(serverConfig.authentication.token);
            } else {
                await client.authenticateWithUsernameAndPassword(serverConfig.authentication.username, serverConfig.authentication.password);
            }
        }

        return client;
    }

    private setupTokenRefresher() {
        if (this.tokenRefreshInterval) {
            clearInterval(this.tokenRefreshInterval)
        }
        this.tokenRefreshInterval = setInterval(() => this.ensureTokenFresh(), 1000 * 60 * 5);
    }

    getTokenBody(): TokenBody | undefined {
        // check if access token is expired
        const token = this.tokenStorage.get(ACCESS_TOKEN);

        if (!token) {
            return;
        }

        const tokenParts = token.split(".");

        if (tokenParts.length != 3) {
            return;
        }

        const payload = JSON.parse(decodeBase64(tokenParts[1]));

        return payload as TokenBody
    }

    private ensureTokenFresh() {
        const refreshToken = this.tokenStorage.get(REFRESH_TOKEN);

        if (!refreshToken) {
            return;
        }

        const payload = this.getTokenBody();

        if (!payload) {
            return;
        }

        if (payload.exp * 1000 > Date.now()) {
            return;
        }

        console.warn("Access token expired, refreshing...");

        this.refreshToken();
    }

    public refreshToken() {
        const refreshToken = this.tokenStorage.get(REFRESH_TOKEN);
        return axios.put<AuthenticationResponse>(Urls.authenticate(this.url), {
            token: refreshToken,
            term: TokenTerm.LONG
        }, {
            validateStatus: (status) => true,
        }).then(resp => {
            if (resp.status == 200) {
                this.tokenStorage.set(ACCESS_TOKEN, resp.data.token.content);
            }
        }).catch(e => {
            console.error("Error refreshing token", e);
        })
    }

    public isAuthenticated(): boolean {
        return !!this.tokenStorage.get(ACCESS_TOKEN);
    }

    public invalidateAuthentication(): void {
        this.tokenStorage.clear();
    }
}
