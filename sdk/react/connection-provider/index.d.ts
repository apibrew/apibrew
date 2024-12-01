import { Server } from '@apibrew/client/config';
export * from './cloud';
export interface Connection {
    name: string;
    title?: string;
    serverConfig: Server;
}
export interface ConnectionProvider {
    allowManageConnections?: boolean;
    connectionsPageLink?: string;
    allowUserSwitchConnections?: boolean;
    getConnection(name: string): Promise<Connection>;
    listConnections?(): Promise<Connection[]>;
    createConnection?(connection: Connection): Promise<void>;
    updateConnection?(connection: Connection): Promise<void>;
    deleteConnection?(name: string): Promise<void>;
}
