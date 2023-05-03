import {components} from "./schema";
import axios from "axios";


/////// #### abs #### //////

export interface Entity {
    id?: string

    [key: string]: any
}

interface Repository<T extends Entity> {
    create(entity: T): Promise<T>;

    update(entity: T): Promise<T>;

    loadResources(): Promise<void>;

    load(entity: T): Promise<T>;

    apply(entity: T): Promise<T>;

    get(id: string): Promise<T>;

    find(params: FindParams): Promise<{ total: number, content: { properties: T }[] }>;

    extend(extensionService: ExtensionService): RepositoryExtension<T>;
}

interface BooleanExpression {

}

interface FindParams {
    limit?: number;
    offset?: number;
    useHistory?: boolean;
    annotations?: { [key: string]: string };
    resolveReferences?: string[]; // default ["*"]
    query?: BooleanExpression | null;
}

export interface RepositoryExtension<T extends Entity> {
    onCreate(handler: (elem: T) => Promise<T>, finalize?: boolean): void;

    onUpdate(handler: (elem: T) => Promise<T>, finalize?: boolean): void;

    onDelete(handler: (elem: T) => Promise<T>, finalize?: boolean): void;

    onGet(handler: (id: string) => Promise<T>, finalize?: boolean): void;

    onList(handler: () => Promise<{ properties: T }[]>, finalize?: boolean): void;
}


/////// #### apply #### //////

export interface DhClientParams {
    Addr: string;
    Insecure?: boolean;
    token?: string;
}

export class DhClient {
    params: DhClientParams;

    constructor(params: DhClientParams) {
        this.params = params
    }

    public async authenticateWithUsernameAndPassword(username: string, password: string): Promise<void> {
        const authRequest: components["schemas"]["AuthenticationRequest"] = {
            username: username,
            password: password,
            term: "LONG"
        }

        const result = await axios.post<components["schemas"]["AuthenticationResponse"]>(`http://${this.params.Addr}/authentication/token`, authRequest);

        this.params.token = result.data.token!.content;
    }

    public newRepository<T extends Entity>(namespace: string, resource: string): Repository<T> {
        return new RepositoryImpl<T>(this, {
            namespace: namespace,
            resource: resource,
            updateCheckVersion: false,
        });
    }

    public NewExtensionService(host: string, port: number, remoteUrl?: string): ExtensionService {
        if (!remoteUrl) {
            remoteUrl = "http://" + host + ':' + port
        }
        return new ExtensionServiceImpl(host, port, remoteUrl, this);
    }
}

type ExternalFunction = (req: components["schemas"]["Event"]) => Promise<components["schemas"]["Event"]>;

interface FunctionCallRequest {
    name: string;
    event: components["schemas"]["Event"]
}

interface FunctionCallResponse {
    response: components["schemas"]["Event"]
}

interface ExtensionService {
    run(): Promise<void>;

    registerFunction(name: string, handler: ExternalFunction): void;

    getRemoteHost(): string;
}

class ExtensionServiceImpl implements ExtensionService {
    private host: string;
    private port: number;
    private remoteUrl: string;
    private client: DhClient;
    private functions: { [key: string]: ExternalFunction };

    constructor(host: string, port: number, remoteUrl: string, client: DhClient) {
        this.host = host;
        this.port = port
        this.remoteUrl = remoteUrl;
        this.client = client;
        this.functions = {};
    }

    getRemoteHost(): string {
        return this.remoteUrl;
    }

    registerFunction(name: string, handler: ExternalFunction): void {
        this.functions[name] = handler;
    }

    async run(): Promise<void> {
        const express = require('express')
        const app = express()

        app.use(express.json())

        app.get('/', (req: any, res: any) => {
            res.send('ok')
        })

        app.post('/:name', async (req: any, res: any) => {
            const name = req.params.name

            try {
                const response = await this.functions[name](req.body.content as any)
                res.send({
                    content: response
                })
            } catch (e: any) {
                console.log(e)
                res.status(400).send({
                    message: e.message
                })
            }
        })

        console.log('starting extension service')
        app.listen(this.port, this.host, () => {
            console.log(`External service is listening on ${this.host}`)
        })
    }
}

interface RepositoryParams<T extends Entity> {
    namespace: string,
    resource: string,
    updateCheckVersion: boolean;
}

export class RepositoryImpl<T extends Entity> implements Repository<T> {
    private readonly client: DhClient;
    private readonly params: RepositoryParams<T>;
    private resource?: components["schemas"]["Resource"];

    constructor(client: DhClient, params: RepositoryParams<T>) {
        this.client = client;
        this.params = params;
    }

    async loadResources(): Promise<void> {
        if (this.resource) {
            return
        }

        const result = await axios.get<components["schemas"]["GetResourceByNameResponse"]>(`http://${this.client.params.Addr}/system/resources/${this.params.namespace}/${this.params.resource}`);
        this.resource = result.data.resource;
    }

    async create(entity: T): Promise<T> {
        const result = await axios.post<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async update(entity: T): Promise<T> {
        const result = await axios.put<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}/${entity.id}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async get(id: string): Promise<T> {
        const result = await axios.get<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}/${id}`, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    public async load(entity: T): Promise<T> {
        if (entity.id) {
            return this.get(entity.id);
        } else {
            await this.loadResources();

            for (const prop of this.resource!.properties!) {
                if (prop.unique) {
                    const val = entity[prop.name!]
                    const result = await axios.get<components["schemas"]["ListRecordResponse"]>(`http://${this.client.params.Addr}/${this.params.namespace}/${this.params.resource}?filters=${prop.name}&filters=${val}&limit=1`, {
                        headers: {
                            Authorization: `Bearer ${this.client.params.token}`
                        }
                    });

                    if (!result.data.total) {
                        continue
                    }

                    return result.data.content![0].properties as T;
                }
            }
        }

        throw new Error(`Entity not found: ${this.params.namespace}/${this.params.resource}`);
    }

    public async apply<T>(entity: T): Promise<T> {
        const result = await axios.patch<T>(`http://${this.client.params.Addr}/${this.params.namespace}/${this.params.resource}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async find(params: FindParams): Promise<{ total: number, content: { properties: T }[] }> {
        if (!params.resolveReferences) {
            params.resolveReferences = ["*"];
        }

        const result = await axios.get<{
            total: number,
            content: { properties: T }[]
        }>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}`, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    extend(extensionService: ExtensionService): RepositoryExtension<T> {
        return new RepositoryExtensionImpl<T>(this, extensionService, this.params.resource, this.params.namespace, this.client);
    }
}


// ## repository extension

export class RepositoryExtensionImpl<T extends Entity> implements RepositoryExtension<T> {
    private repository: Repository<T>;
    private extension: ExtensionService;
    private resourceName: string;
    private namespace: string;
    private client: DhClient;
    private extensionRepository: RepositoryImpl<components["schemas"]["Extension"]>;

    constructor(
        repository: Repository<T>,
        extension: ExtensionService,
        resourceName: string,
        namespace: string,
        client: DhClient
    ) {
        this.repository = repository;
        this.extension = extension;
        this.resourceName = resourceName;
        this.namespace = namespace;
        this.client = client;
        this.extensionRepository = new RepositoryImpl<components["schemas"]["Extension"]>(client, {
            namespace: "system", resource: "extension", updateCheckVersion: false
        })
    }

    async onCreate(handler: (elem: T) => Promise<T>, finalize?: boolean): Promise<void> {
        const extensionName = this.getExtensionName("OnCreate");

        this.extension.registerFunction(extensionName, async function (data: components["schemas"]["Event"]) {
            const records = []

            if (data.records) {
                for (const record of data.records) {
                    const entity = await handler(record.properties as T)
                    records.push({
                        properties: entity
                    })
                }
            }

            data.records = records

            return data
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                create: {
                    kind: "httpCall",
                    uri: `${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
                finalize: finalize,
            },
        };

        await this.extensionRepository.apply(ext)
    }

    async onUpdate(handler: (elem: T) => Promise<T>, finalize?: boolean): Promise<void> {
        const extensionName = this.getExtensionName("OnUpdate");

        this.extension.registerFunction(extensionName, async function (data: components["schemas"]["Event"]) {
            const records = []

            if (data.records) {
                for (const record of data.records) {
                    const entity = await handler(record.properties as T)
                    records.push({
                        properties: entity
                    })
                }
            }

            data.records = records

            return data
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                update: {
                    kind: "httpCall",
                    uri: `${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
                finalize: finalize,
            },
        };

        await this.extensionRepository.apply(ext)
    }

    async onDelete(handler: (elem: T) => Promise<T>, finalize?: boolean): Promise<void> {
        const extensionName = this.getExtensionName("OnDelete");

        this.extension.registerFunction(extensionName, async (data: components["schemas"]["Event"]) => {
            if (data.ids) {
                for (const id of data.ids) {
                    await handler(await this.repository.get(id))
                }
            }

            return data
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                delete: {
                    kind: "httpCall",
                    uri: `${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
                finalize: finalize,
            },
        };

        await this.extensionRepository.apply(ext)
    }

    async onGet(handler: (id: string) => Promise<T>, finalize?: boolean): Promise<void> {
        const extensionName = this.getExtensionName("OnGet");

        this.extension.registerFunction(extensionName, async (data: components["schemas"]["Event"]) => {
            const id = (data.ids!)[0]
            await handler(id)

            data.records = [await handler(id)]

            return data
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                get: {
                    kind: "httpCall",
                    uri: `${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
                finalize: finalize,
            },
        };

        await this.extensionRepository.apply(ext)
    }

    async onList(handler: () => Promise<{ properties: T }[]>, finalize?: boolean): Promise<void> {
        const extensionName = this.getExtensionName("OnList");

        this.extension.registerFunction(extensionName, async function (data: components["schemas"]["Event"]) {
            const records = await handler()

            data.records = records

            return data
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                list: {
                    kind: "httpCall",
                    uri: `${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
                finalize: finalize,
            },
        };

        await this.extensionRepository.apply(ext)
    }

    private getExtensionName(action: string): string {
        return `${this.namespace}-${this.resourceName}-${action}`;
    }
}
