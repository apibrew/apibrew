import { Struct } from 'google-protobuf/google/protobuf/struct_pb';
import { Record } from './model/record';
import { Resource } from './model/resource';
import { Namespace } from './model/namespace';
import { Extension } from './model/extension';
import { User } from './model/user';
import { DataSource } from './model/data-source';
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
import { TokenTerm } from './model/token';
import { FunctionCallRequest, FunctionCallResponse, FunctionClient } from "./ext/function";
import { Any } from "google-protobuf/google/protobuf/any_pb";
import { FunctionCall } from './model/external';
import { Server, ServerCredentials } from "@grpc/grpc-js";
import * as dependency_6 from "./google/protobuf/any";

/////// #### abs #### //////

interface Entity<T> {
    ToRecord(): Record;
    FromRecord(record: Record): void;
    FromProperties(properties: { [key: string]: Struct }): void;
    ToProperties(): { [key: string]: Struct };
    GetResourceName(): string;
    GetNamespace(): string;
    Equals(other: T): boolean;
    Same(other: T): boolean;
}

interface DhClient {
    GetAuthenticationClient(): AuthenticationClient;
    GetDataSourceClient(): DataSourceClient;
    GetResourceClient(): ResourceClient;
    GetRecordClient(): RecordClient;
    GetGenericClient(): GenericClient;
    GetNamespaceClient(): NamespaceClient;
    GetExtensionClient(): ExtensionClient;
    GetUserClient(): UserClient;
    GetToken(): string;
    AuthenticateWithToken(token: string): void;
    AuthenticateWithUsernameAndPassword(username: string, password: string): Promise<void>;
    NewExtensionService(host: string): ExtensionService;
}

interface Repository<T extends Entity<T>> {
    Create(entity: T): Promise<T>;
    Update(entity: T): Promise<T>;
    Save(entity: T): Promise<T>;
    Get(id: string): Promise<T>;
    Find(params: FindParams): Promise<T[]>;
    Extend(extension: Extension): RepositoryExtension<T>;
}

interface FindParams {
    Limit?: number;
    Offset?: number;
    UseHistory?: boolean;
    Annotations?: { [key: string]: string };
    ResolveReferences?: string[]; // default ["*"]
    Query?: BooleanExpression | null;
}

interface RepositoryExtension<T extends Entity<T>> {
    Create(entity: T): Promise<T>;
    Update(entity: T): Promise<T>;
    Save(entity: T): Promise<T>;
    Get(id: string): Promise<T>;
    Find(params: FindParams): Promise<T[]>;
}


/////// #### apply #### //////

export interface DhClientParams {
    Addr: string;
    Insecure: boolean;
    Token: string;
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

                this.params.Token = resp?.token.content as string

                resolve()
            });
        })


    }

    public AuthenticateWithToken(token: string): void {
        this.params.Token = token;
    }

    public GetNamespaceClient(): NamespaceClient {
        return this.namespaceClient;
    }

    public GetToken(): string {
        return this.params.Token;
    }

    public GetAuthenticationClient(): AuthenticationClient {
        return this.authenticationClient;
    }

    public GetDataSourceClient(): DataSourceClient {
        return this.dataSourceClient;
    }

    public GetResourceClient(): ResourceClient {
        return this.resourceClient;
    }

    public GetRecordClient(): RecordClient {
        return this.recordClient;
    }

    public GetGenericClient(): GenericClient {
        return this.genericClient;
    }

    public GetExtensionClient(): ExtensionClient {
        return this.extensionClient;
    }

    public GetUserClient(): UserClient {
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
