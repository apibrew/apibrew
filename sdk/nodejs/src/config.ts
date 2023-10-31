export interface Config {
    type: 'server';
    defaultServer: string;
    servers: Server[];
}

export interface Server {
    name: string;
    host: string;
    port: number;
    httpPort: number;
    insecure: boolean;
    authentication: Authentication;
}

export interface Authentication {
    username: string;
    password: string;
    token: string;
}