import { TokenStorage } from '@apibrew/client/token-storage'

export class LocalStorageTokenStorage implements TokenStorage {
  clear(): void {
    if (typeof window !== 'undefined') {
      window.localStorage.setItem('token', '{}')
    }
  }

  get(name: string): string | undefined {
    if (typeof window !== 'undefined') {
      const token = window.localStorage.getItem('token')
      if (token) {
        return JSON.parse(token)[name]
      } else {
        return undefined
      }
    } else {
      return undefined
    }
  }

  set(name: string, token: string): void {
    if (typeof window !== 'undefined') {
      const tokens = JSON.parse(window.localStorage.getItem('token') || '{}')
      tokens[name] = token
      window.localStorage.setItem('token', JSON.stringify(tokens))
    } else {
      throw new Error('Cannot set token in non-browser environment')
    }
  }
}
