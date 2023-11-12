import React from 'react'
import { Client } from '@apibrew/client'

export const ClientContext = React.createContext<Client | undefined>(undefined)

export const ClientProvider = ClientContext.Provider

export const ClientConsumer = ClientContext.Consumer

export function useClient() {
  const context = React.useContext(ClientContext)
  if (context === undefined) {
    throw new Error('useClient must be used within a ClientProvider')
  }
  return context
}
