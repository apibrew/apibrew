import { Struct } from 'google-protobuf/google/protobuf/struct_pb';
import { Record } from './model/record';

import { Extension } from './model/extension';

import { BooleanExpression } from './model/query';
import { AuthenticationClient, AuthenticationRequest, AuthenticationResponse } from './stub/authentication';
import { DataSourceClient } from './stub/data-source';
import { ExtensionClient } from './stub/extension';
import { GenericClient } from './stub/generic';
import { NamespaceClient } from './stub/namespace';
import { RecordClient } from './stub/record';
import { ResourceClient } from './stub/resource';
import { UserClient } from './stub/user';
import { credentials } from '@grpc/grpc-js';
import { FunctionCallRequest, FunctionCallResponse, FunctionClient } from "./ext/function";

import * as dependency_6 from "./google/protobuf/any";
import { TokenTerm } from './model/token';

/////// #### abs #### //////

interface Entity<T> {
    toRecord(): Record;
    fromRecord(record: Record): void;
    fromProperties(properties: { [key: string]: Struct }): void;
    toProperties(): { [key: string]: Struct };
    getResourceName(): string;
    getNamespace(): string;
    equals(other: T): boolean;
    same(other: T): boolean;
}

interface DhClient {
    getAuthenticationClient(): AuthenticationClient;
    getDataSourceClient(): DataSourceClient;
    getResourceClient(): ResourceClient;
    getRecordClient(): RecordClient;
    getGenericClient(): GenericClient;
    getNamespaceClient(): NamespaceClient;
    getExtensionClient(): ExtensionClient;
    getUserClient(): UserClient;
    gettoken(): string;
    AuthenticateWithtoken(token: string): void;
    AuthenticateWithUsernameAndPassword(username: string, password: string): Promise<void>;
    NewExtensionService(host: string): ExtensionService;
}

interface Repository<T extends Entity<T>> {
    create(entity: T): Promise<T>;
    update(entity: T): Promise<T>;
    save(entity: T): Promise<T>;
    get(id: string): Promise<T>;
    find(params: FindParams): Promise<T[]>;
    extend(extensionService: ExtensionService): RepositoryExtension<T>;
}

interface FindParams {
    limit?: number;
    offset?: number;
    useHistory?: boolean;
    annotations?: { [key: string]: string };
    resolveReferences?: string[]; // default ["*"]
    query?: BooleanExpression | null;
}

export interface RepositoryExtension<T extends Entity<T>> {
    onCreate(handler: (elem: T) => Promise<T>): Promise<void>;
    onUpdate(handler: (elem: T) => Promise<T>): Promise<void>;
    onDelete(handler: (elem: T) => Promise<T>): Promise<void>;
    onGet(handler: (id: string) => Promise<T>): Promise<void>;
    onList(handler: () => Promise<T[]>): Promise<void>;
}


/////// #### apply #### //////

export interface DhClientParams {
    Addr: string;
    Insecure: boolean;
    token: string;
}

export class DhClientImpl implements DhClient {
    private recordClient: RecordClient;
    private authenticationClient: AuthenticationClient;
    private resourceClient: ResourceClient;
    private dataSourceClient: DataSourceClient;
    private userClient: UserClient;
    private extensionClient: ExtensionClient;
    private genericClient: GenericClient;
    private namespaceClient: NamespaceClient;
    params: DhClientParams;

    constructor(params: DhClientParams) {
        const creds = credentials.createInsecure()

        this.recordClient = new RecordClient(params.Addr, creds);
        this.authenticationClient = new AuthenticationClient(params.Addr, creds);
        this.resourceClient = new ResourceClient(params.Addr, creds);
        this.dataSourceClient = new DataSourceClient(params.Addr, creds);
        this.userClient = new UserClient(params.Addr, creds);
        this.extensionClient = new ExtensionClient(params.Addr, creds);
        this.genericClient = new GenericClient(params.Addr, creds);
        this.namespaceClient = new NamespaceClient(params.Addr, creds);

        this.params = params
    }

    public async AuthenticateWithUsernameAndPassword(username: string, password: string): Promise<void> {
        const authRequest = new AuthenticationRequest();
        authRequest.username = username
        authRequest.password = password
        authRequest.term = TokenTerm.LONG

        return new Promise((resolve, reject) => {
            this.authenticationClient.Authenticate(authRequest, (err, resp) => {
                if (err) {
                    reject(err.message)
                    return
                }

                this.params.token = resp?.token.content as string

                resolve()
            });
        })


    }

    public AuthenticateWithtoken(token: string): void {
        this.params.token = token;
    }

    public getNamespaceClient(): NamespaceClient {
        return this.namespaceClient;
    }

    public gettoken(): string {
        return this.params.token;
    }

    public getAuthenticationClient(): AuthenticationClient {
        return this.authenticationClient;
    }

    public getDataSourceClient(): DataSourceClient {
        return this.dataSourceClient;
    }

    public getResourceClient(): ResourceClient {
        return this.resourceClient;
    }

    public getRecordClient(): RecordClient {
        return this.recordClient;
    }

    public getGenericClient(): GenericClient {
        return this.genericClient;
    }

    public getExtensionClient(): ExtensionClient {
        return this.extensionClient;
    }

    public getUserClient(): UserClient {
        return this.userClient;
    }

    public NewExtensionService(host: string): ExtensionService {
        return new ExtensionServiceImpl(host, host, this);
    }
}

type ExternalFunctionData = Map<string, dependency_6.Any>

type ExternalFunction = (req: ExternalFunctionData) => Promise<[ExternalFunctionData, Error | undefined]>;


interface ExtensionService {
    run(): Promise<void>;
    registerFunction(name: string, handler: ExternalFunction): void;
    getRemoteHost(): string;
}

class ExtensionServiceImpl implements ExtensionService {
    private host: string;
    private remoteHost: string;
    private client: DhClientImpl;
    private functions: { [key: string]: ExternalFunction };

    constructor(host: string, remoteHost: string, client: DhClientImpl) {
        this.host = host;
        this.remoteHost = remoteHost;
        this.client = client;
        this.functions = {};
    }

    getRemoteHost(): string {
        return this.remoteHost;
    }

    registerFunction(name: string, handler: ExternalFunction): void {
        this.functions[name] = handler;
    }

    async functionCall(request: FunctionCallRequest): Promise<FunctionCallResponse> {
        if (!this.functions[request.name]) {
            throw new Error("External function not found")
        }

        const fn = this.functions[request.name]
        const [responseData, error] = await fn(request.request)

        if (error) {
            throw new Error(error.message)
        }

        return new FunctionCallResponse({
            response: responseData as Map<string, dependency_6.Any>
        })
    }

    async run(): Promise<void> {

    }
}

interface RepositoryParams<T extends Entity<T>> {
    updateCheckVersion: boolean;
    instanceProvider: () => T;
}

interface RecordResponse {
    record: any;
}

export class RepositoryImpl<T extends Entity<T>> implements Repository<T> {
    private readonly client: DhClient;
    private readonly params: RepositoryParams<T>;

    constructor(client: DhClient, params: RepositoryParams<T>) {
        this.client = client;
        this.params = params;
    }

    async create(entity: T): Promise<T> {
        const resp = await this.client.getRecordClient().create({
            token: this.client.gettoken(),
            namespace: entity.getNamespace(),
            resource: entity.getResourceName(),
            record: entity.toRecord(),
        });

        entity.fromRecord(resp.record);

        return entity;
    }

    async update(entity: T): Promise<T> {
        const resp = await this.client.getRecordClient().update({
            token: this.client.gettoken(),
            namespace: entity.getNamespace(),
            resource: entity.getResourceName(),
            record: entity.toRecord(),
            checkVersion: this.params.updateCheckVersion,
        });

        entity.fromRecord(resp.record);

        return entity;
    }

    async save(entity: T): Promise<T> {
        const resource = await this.loadResource();

        entity.fromRecord(await this.client.applyRecord(resource, entity.toRecord()));

        return entity
    }

    async get(id: string): Promise<T> {
        const instance = this.params.instanceProvider();

        const resp = await this.client.getRecordClient().get({
            token: this.client.gettoken(),
            namespace: instance.getNamespace(),
            resource: instance.getResourceName(),
            id: id,
        });

        instance.fromRecord(resp.record);

        return instance;
    }

    async find(params: FindParams): Promise<T[]> {
        const instance = this.params.instanceProvider();

        if (!params.resolveReferences) {
            params.resolveReferences = ["*"];
        }

        const resp = await this.client.getRecordClient().search({
            token: this.client.gettoken(),
            namespace: instance.getNamespace(),
            resource: instance.getResourceName(),
            query: params.query,
            limit: params.limit,
            offset: params.offset,
            useHistory: params.useHistory,
            resolveReferences: params.resolveReferences,
            annotations: params.annotations,
        });

        return resp.content.map((record: RecordResponse) => {
            const newInstance = this.params.instanceProvider();

            newInstance.fromRecord(record.record);

            return newInstance;
        });
    }

    extend(extensionService: ExtensionService): RepositoryExtension<T> {
        const instance = this.params.instanceProvider();

        return new RepositoryExtensionImpl<T>(this, extensionService, instance.getResourceName(), instance.getNamespace(), this.params.instanceProvider, this.client);
    }

    private async loadResource() {
        const instance = this.params.instanceProvider();

        const resp = await this.client.getResourceClient().getByName({
            token: this.client.gettoken(),
            namespace: instance.getNamespace(),
            name: instance.getResourceName(),
        });

        return resp.resource;
    }
}


// ## repository extension

export class RepositoryExtensionImpl<T extends Entity<T>> implements RepositoryExtension<T> {
    private repository: Repository<T>;
    private extension: ExtensionService;
    private resourceName: string;
    private namespace: string;
    private instanceProvider: () => T;
    private client: DhClient;

    constructor(
        repository: Repository<T>,
        extension: ExtensionService,
        resourceName: string,
        namespace: string,
        instanceProvider: () => T,
        client: DhClient
    ) {
        this.repository = repository;
        this.extension = extension;
        this.resourceName = resourceName;
        this.namespace = namespace;
        this.instanceProvider = instanceProvider;
        this.client = client;
    }

    async onCreate(handler: (elem: T) => Promise<T>): Promise<void> {
        const extensionName = this.getExtensionName("OnCreate");

        this.extension.registerFunction(extensionName, CreateRecordTypedFunction(this.instanceProvider, handler));

        const ext: Extension = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                create: {
                    kind: "functionCall",
                    functionCall: {
                        host: this.extension.GetRemoteHost(),
                        functionName: extensionName,
                    },
                },
            },
        };

        await this.client.ApplyExtension(ext);
    }

    async onUpdate(handler: (elem: T) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onDelete(handler: (elem: T) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onGet(handler: (id: string) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onList(handler: () => Promise<T[]>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    private getExtensionName(action: string): string {
        return `${this.namespace}-${this.resourceName}-${action}`;
    }
}
