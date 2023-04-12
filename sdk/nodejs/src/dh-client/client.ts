import {BooleanExpression} from './model/query';
import {AuthenticationClient, AuthenticationRequest} from './stub/authentication';
import {DataSourceClient} from './stub/data-source';
import {ExtensionClient} from './stub/extension';
import {GenericClient} from './stub/generic';
import {NamespaceClient} from './stub/namespace';
import {
    CreateRecordRequest,
    CreateRecordResponse,
    RecordClient,
    SearchRecordRequest,
    SearchRecordResponse,
    UpdateRecordRequest,
    UpdateRecordResponse
} from './stub/record';
import {ResourceClient} from './stub/resource';
import {UserClient} from './stub/user';
import {credentials, Server} from '@grpc/grpc-js';
import {FunctionCallRequest, FunctionCallResponse} from "./ext/function";

import * as dependency_6 from "./google/protobuf/any";
import {TokenTerm} from './model/token';
import * as dependency_1 from "./google/protobuf/struct";
import {Record} from "./model/record";


/////// #### abs #### //////

export interface Entity<T> {
    fromProperties(properties: Map<string, dependency_1.Value>): void;

    toProperties(): Map<string, dependency_1.Value>;

    getResourceName(): string;

    getNamespace(): string;

    equals(other: T): boolean;

    same(other: T): boolean;
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
    Insecure?: boolean;
    token?: string;
}

export class DhClient {
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

    public async authenticateWithUsernameAndPassword(username: string, password: string): Promise<void> {
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

    public newRepository<T extends Entity<T>>(type: new () => T): Repository<T> {
        return new RepositoryImpl<T>(this, type, {
            updateCheckVersion: false,
        });
    }

    public authenticateWithToken(token: string): void {
        this.params.token = token;
    }

    public getNamespaceClient(): NamespaceClient {
        return this.namespaceClient;
    }

    public getToken(): string {
        return this.params.token!;
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
    private client: DhClient;
    private functions: { [key: string]: ExternalFunction };

    constructor(host: string, remoteHost: string, client: DhClient) {
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
        const server = new Server();

        console.log(functionProto)

        // server.addService(functionProto.Function.service, {
        //     FunctionCall: this.functionCall.bind(this),
        // });
        // server.bindAsync(this.host, grpc.ServerCredentials.createInsecure(), () => {
        //     server.start();
        // });
    }
}

interface RepositoryParams<T extends Entity<T>> {
    updateCheckVersion: boolean;
}

export class RepositoryImpl<T extends Entity<T>> implements Repository<T> {
    private readonly client: DhClient;
    private readonly params: RepositoryParams<T>;
    private type: { new(): T };

    constructor(client: DhClient, type: new () => T, params: RepositoryParams<T>) {
        this.client = client;
        this.type = type
        this.params = params;
    }

    async create(entity: T): Promise<T> {
        const record = new Record()
        record.properties = entity.toProperties()

        const resp = await new Promise<CreateRecordResponse>((resolve, reject) => {
            this.client.getRecordClient().Create(new CreateRecordRequest({
                token: this.client.getToken(),
                namespace: entity.getNamespace(),
                resource: entity.getResourceName(),
                record: record,
            }), (err, resp) => {
                if (err) {
                    reject(err.message)
                    return
                }

                resolve(resp!)
            });
        })

        entity.fromProperties(resp.record.properties)

        return entity;
    }

    async update(entity: T): Promise<T> {
        const record = new Record()
        record.properties = entity.toProperties()

        const resp = await new Promise<UpdateRecordResponse>((resolve, reject) => {
            this.client.getRecordClient().Update(new UpdateRecordRequest({
                token: this.client.getToken(),
                namespace: entity.getNamespace(),
                resource: entity.getResourceName(),
                record: record,
                checkVersion: this.params.updateCheckVersion,
            }), (err, resp) => {
                if (err) {
                    reject(err.message)
                    return
                }

                resolve(resp!)
            })
        })

        entity.fromProperties(resp.record.properties)

        return entity;
    }

    async save(entity: T): Promise<T> {
        const resource = await this.loadResource();

        // entity.fromRecord(await this.client.applyRecord(resource, entity.toRecord()));

        return entity
    }

    async get(id: string): Promise<T> {
        const instance = new this.type()

        const resp = await this.client.getRecordClient().get({
            token: this.client.getToken(),
            namespace: instance.getNamespace(),
            resource: instance.getResourceName(),
            id: id,
        });

        instance.fromProperties(resp.record.properties)

        return instance;
    }

    async find(params: FindParams): Promise<T[]> {
        const instance = new this.type()

        if (!params.resolveReferences) {
            params.resolveReferences = ["*"];
        }

        const resp = await new Promise<SearchRecordResponse>((resolve, reject) => {
            this.client.getRecordClient().Search(new SearchRecordRequest({
                token: this.client.getToken(),
                namespace: instance.getNamespace(),
                resource: instance.getResourceName(),
                query: params.query!,
                limit: params.limit!,
                offset: params.offset!,
                useHistory: params.useHistory!,
                resolveReferences: params.resolveReferences!,
                annotations: new Map<string, string>(),
            }), (err, resp) => {
                if (err) {
                    reject(err);
                } else {
                    resolve(resp!);
                }
            });
        })

        return resp.content.map((record: Record) => {
            const newInstance = new this.type()

            newInstance.fromProperties(record.properties);

            return newInstance;
        });
    }

    extend(extensionService: ExtensionService): RepositoryExtension<T> {
        const instance = new this.type()

        return new RepositoryExtensionImpl<T>(this, extensionService, instance.getResourceName(), instance.getNamespace(), this.type, this.client);
    }

    private async loadResource() {
        const instance = new this.type()

        const resp = await this.client.getResourceClient().getByName({
            token: this.client.getToken(),
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
    private type: { new(): T };
    private client: DhClient;

    constructor(
        repository: Repository<T>,
        extension: ExtensionService,
        resourceName: string,
        namespace: string,
        type: { new(): T },
        client: DhClient
    ) {
        this.repository = repository;
        this.extension = extension;
        this.resourceName = resourceName;
        this.namespace = namespace;
        this.type = type;
        this.client = client;
    }

    async onCreate(handler: (elem: T) => Promise<T>): Promise<void> {
        // const extensionName = this.getExtensionName("OnCreate");
        //
        // this.extension.registerFunction(extensionName, CreateRecordTypedFunction(this.instanceProvider, handler));
        //
        // const ext: Extension = {
        //     name: extensionName,
        //     namespace: this.namespace,
        //     resource: this.resourceName,
        //     instead: {
        //         create: {
        //             kind: "functionCall",
        //             functionCall: {
        //                 host: this.extension.GetRemoteHost(),
        //                 functionName: extensionName,
        //             },
        //         },
        //     },
        // };
        //
        // await this.client.ApplyExtension(ext);
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
