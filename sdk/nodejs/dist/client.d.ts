import { RecordResourceInfo, Record } from "./model";
import { Repository } from "./repository";
import { ServiceConfigProvider } from "./api/config";
export declare class Client {
    private config;
    private static defaultClient;
    constructor(backendUrl: string);
    authenticate(username: string, password: string): Promise<void>;
    authenticateToken(token: string): void;
    newRepository<T extends Record<unknown>>(recordResourceInfo: RecordResourceInfo): Repository<T>;
    provider(): ServiceConfigProvider;
    static setDefaultClient(client: Client): void;
    static getDefaultClient(): Client;
}
