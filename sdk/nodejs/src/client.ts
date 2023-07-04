import { RecordResourceInfo, Record } from "./model";
import { Repository } from "./repository";
import { AuthenticationService } from "./service";
import { ServiceConfig, ServiceConfigProvider } from "./service/config";

export class Client {
    private config: ServiceConfig
    private static defaultClient: Client

    constructor(backendUrl: string) {
        this.config = {
            backendUrl: backendUrl,
            token: ''
        }
    }

    public async authenticate(username: string, password: string) {
        this.config.token = await AuthenticationService.authenticate(this.config, username, password, 'VERY_LONG')
            .then(result => result.content)
    }

    public authenticateToken(token: string) {
        this.config.token = token
    }

    newRepository<T extends Record<unknown>>(recordResourceInfo: RecordResourceInfo): Repository<T> {
        return new Repository(this.provider(), recordResourceInfo)
    }

    provider(): ServiceConfigProvider {
        return () => this.config
    }

    static setDefaultClient(client: Client) {
        Client.defaultClient = client
    }

    static getDefaultClient(): Client {
        return Client.defaultClient
    }

}