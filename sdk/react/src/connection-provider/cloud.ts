import { Connection, ConnectionProvider } from '../connection-provider'
import { Authentication } from '@apibrew/client/config'

export const cloudConnectionProvider: ConnectionProvider = {
  connectionsPageLink: '/cloud/instances',
  createConnection(): Promise<void> {
    throw new Error('Not implemented')
  },
  deleteConnection(): Promise<void> {
    throw new Error('Not implemented')
  },
  updateConnection(): Promise<void> {
    throw new Error('Not implemented')
  },
  async listConnections(): Promise<Connection[]> {
    return []
  },
  allowManageConnections: true,
  getConnection(name: string): Promise<Connection> {
    const localStorageItem = localStorage.getItem(
      `@apibrew/client/${name}/token`
    )

    const connection: Connection = {
      name: name,
      title: name,
      serverConfig: {
        host: `https://${name}.apibrew.io`,
        httpPort: 8443
      }
    } as Connection

    if (localStorageItem) {
      const tokenData = JSON.parse(localStorageItem)

      connection.serverConfig.authentication = {
        token: tokenData.ACCESS_TOKEN as string
      } as Authentication
    }

    return Promise.resolve(connection)
  }
}
