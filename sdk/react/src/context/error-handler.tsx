import React from 'react'
import { ApiException } from '@apibrew/client'

export interface ErrorHandlerContextProps {
  onAuthenticationFailed(err: ApiException): void
}

export type CatchArg = Parameters<Promise<any>['catch']>[0]

export const ErrorHandlerContext = React.createContext<CatchArg>(undefined)

export const useErrorHandler = () => {
  return React.useContext(ErrorHandlerContext)
}
