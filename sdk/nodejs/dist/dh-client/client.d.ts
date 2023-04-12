export interface Entity<T> {
    id?: string;
}
interface Repository<T extends Entity<T>> {
    create(entity: T): Promise<T>;
    update(entity: T): Promise<T>;
    apply(entity: T): Promise<T>;
    get(id: string): Promise<T>;
    find(params: FindParams): Promise<{
        total: number;
        content: {
            properties: T;
        }[];
    }>;
    extend(extensionService: ExtensionService): RepositoryExtension<T>;
}
interface BooleanExpression {
}
interface FindParams {
    limit?: number;
    offset?: number;
    useHistory?: boolean;
    annotations?: {
        [key: string]: string;
    };
    resolveReferences?: string[];
    query?: BooleanExpression | null;
}
export interface RepositoryExtension<T extends Entity<T>> {
    onCreate(handler: (elem: T) => Promise<T>, finalize?: boolean): void;
    onUpdate(handler: (elem: T) => Promise<T>, finalize?: boolean): void;
    onDelete(handler: (elem: T) => Promise<T>, finalize?: boolean): void;
    onGet(handler: (id: string) => Promise<T>, finalize?: boolean): void;
    onList(handler: () => Promise<{
        properties: T;
    }[]>, finalize?: boolean): void;
}
export interface DhClientParams {
    Addr: string;
    Insecure?: boolean;
    token?: string;
}
export declare class DhClient {
    params: DhClientParams;
    constructor(params: DhClientParams);
    authenticateWithUsernameAndPassword(username: string, password: string): Promise<void>;
    newRepository<T extends Entity<T>>(namespace: string, resource: string): Repository<T>;
    NewExtensionService(host: string, port: number): ExtensionService;
}
type ExternalFunctionData = {
    [key: string]: any;
};
type ExternalFunction = (req: ExternalFunctionData) => Promise<ExternalFunctionData>;
interface ExtensionService {
    run(): Promise<void>;
    registerFunction(name: string, handler: ExternalFunction): void;
    getRemoteHost(): string;
}
interface RepositoryParams<T extends Entity<T>> {
    namespace: string;
    resource: string;
    updateCheckVersion: boolean;
}
export declare class RepositoryImpl<T extends Entity<T>> implements Repository<T> {
    private readonly client;
    private readonly params;
    constructor(client: DhClient, params: RepositoryParams<T>);
    create(entity: T): Promise<T>;
    update(entity: T): Promise<T>;
    get(id: string): Promise<T>;
    apply<T>(entity: T): Promise<T>;
    find(params: FindParams): Promise<{
        total: number;
        content: {
            properties: T;
        }[];
    }>;
    extend(extensionService: ExtensionService): RepositoryExtension<T>;
}
export declare class RepositoryExtensionImpl<T extends Entity<T>> implements RepositoryExtension<T> {
    private repository;
    private extension;
    private resourceName;
    private namespace;
    private client;
    private extensionRepository;
    constructor(repository: Repository<T>, extension: ExtensionService, resourceName: string, namespace: string, client: DhClient);
    onCreate(handler: (elem: T) => Promise<T>, finalize?: boolean): Promise<void>;
    onUpdate(handler: (elem: T) => Promise<T>): Promise<void>;
    onDelete(handler: (elem: T) => Promise<T>): Promise<void>;
    onGet(handler: (id: string) => Promise<T>): Promise<void>;
    onList(handler: () => Promise<{
        properties: T;
    }[]>): Promise<void>;
    private getExtensionName;
}
export {};
