import { cloudConnectionProvider } from './cloud'
import { Server } from '@apibrew/client/config'

export interface Connection {
  name: string
  title?: string
  serverConfig: Server
}

export interface ConnectionProvider {
  allowManageConnections?: boolean
  connectionsPageLink?: string

  allowUserSwitchConnections?: boolean

  getConnection(name: string): Promise<Connection>

  listConnections?(): Promise<Connection[]>

  createConnection?(connection: Connection): Promise<void>

  updateConnection?(connection: Connection): Promise<void>

  deleteConnection?(name: string): Promise<void>
}

const connectionProviderName = process.env.REACT_APP_CONNECTION_PROVIDER

let _connectionProvider: ConnectionProvider

if (connectionProviderName === 'LOCAL_ENV') {
  _connectionProvider = {
    getConnection(): Promise<Connection> {
      return Promise.resolve({
        name: 'local',
        serverConfig: {
          name: 'local',
          host: process.env.REACT_APP_APBR_HOST,
          httpPort: parseInt(process.env.REACT_APP_APBR_HTTP_PORT || '9009'),
          insecure: process.env.REACT_APP_APBR_INSECURE === 'true',
          authentication: {
            token: process.env.REACT_APP_APBR_AUTHENTICATION_TOKEN,
            username: process.env.REACT_APP_APBR_AUTHENTICATION_USERNAME,
            password: process.env.REACT_APP_APBR_AUTHENTICATION_PASSWORD
          }
        }
      } as Connection)
    }
  }
} else if (connectionProviderName === 'WEB') {
  _connectionProvider = {
    createConnection(connection: Connection): Promise<void> {
      localStorage.setItem(
        'connection_' + connection.name,
        JSON.stringify(connection)
      )
      return Promise.resolve()
    },
    deleteConnection(name: string): Promise<void> {
      localStorage.removeItem('connection_' + name)
      return Promise.resolve()
    },
    updateConnection(connection: Connection): Promise<void> {
      localStorage.setItem(
        'connection_' + connection.name,
        JSON.stringify(connection)
      )
      return Promise.resolve()
    },
    listConnections(): Promise<Connection[]> {
      const connections: Connection[] = []
      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i)
        if (key && key.startsWith('connection_')) {
          const connectionStr = localStorage.getItem(key)
          if (connectionStr) {
            connections.push(JSON.parse(connectionStr) as Connection)
          }
        }
      }
      return Promise.resolve(connections)
    },
    allowManageConnections: true,
    getConnection(name: string): Promise<Connection> {
      const connectionStr = localStorage.getItem('connection_' + name)

      if (!connectionStr) {
        return Promise.reject(new Error("'Connection not found: ' + name"))
      }

      return Promise.resolve(JSON.parse(connectionStr) as Connection)
    }
  }
} else if (connectionProviderName === 'WEB_CLOUD') {
  _connectionProvider = cloudConnectionProvider
} else {
  console.log('process.env', process.env)
  throw new Error('Unknown connection provider: ' + connectionProviderName)
}

export const connectionProvider = _connectionProvider
