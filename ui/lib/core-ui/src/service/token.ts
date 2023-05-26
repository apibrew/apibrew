import { type Token } from '../model'

export namespace TokenService {
    let _token: Token

    function init(): void {
        if (localStorage.getItem('token')) {
            setToken(JSON.parse(localStorage.getItem('token') ?? '') as Token)
        }
    }

    export function removeToken() {
        _token = undefined
        localStorage.removeItem('token')
    }

    export async function get(): Promise<string> {
        if (!(_token)) {
            return await Promise.reject(new Error('No token available'))
        }

        // implement refresh token logic here

        return _token.content
    }

    export function setToken(token: Token) {
        _token = token
    }

    export function storeToken(token: Token) {
        setToken(token)
        localStorage.setItem('token', JSON.stringify(token))
    }

    export function isLoggedIn(): boolean {
        return !!_token
    }

    init()
    console.log('TokenService initialized')
}
