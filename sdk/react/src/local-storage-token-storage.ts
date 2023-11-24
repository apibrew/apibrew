import { TokenStorage } from '@apibrew/client/token-storage'

export class LocalStorageTokenStorage implements TokenStorage {
  public constructor(private prefix?: string) {}

  clear(): void {
    if (typeof window !== 'undefined') {
      window.localStorage.setItem(this.getKey(), '{}')
    }
  }

  private getKey() {
    if (this.prefix) {
      return `@apibrew/client/${this.prefix}/token`
    }

    return '@apibrew/client/token'
  }

  get(name: string): string | undefined {
    if (typeof window !== 'undefined') {
      const token = window.localStorage.getItem(this.getKey())
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
      const tokens = JSON.parse(
        window.localStorage.getItem(this.getKey()) || '{}'
      )
      tokens[name] = token
      window.localStorage.setItem(this.getKey(), JSON.stringify(tokens))
    } else {
      throw new Error('Cannot set token in non-browser environment')
    }
  }
}
