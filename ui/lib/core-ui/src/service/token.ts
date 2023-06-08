import {type Token} from '../model'
import axios from "axios";
import {BACKEND_URL} from "../config.ts";
import {zonedTimeToUtc, utcToZonedTime, format} from 'date-fns-tz'
import {refreshToken} from "./authentication.ts";
import {SecurityConstraint} from "../model/security-constraint.ts";


export class NoTokenAvailableError extends Error {
    constructor() {
        super('No token available')
    }
}

export interface TokenBody {
    exp: number
    uid: string
    securityConstraints: SecurityConstraint[]
}

interface Tokens {
    _access_token?: Token
    _access_tokenBody?: TokenBody
    _refresh_token?: Token
    _refresh_tokenBody?: TokenBody
}

const tokens: Tokens = {} as Tokens

function init(): void {
    if (localStorage.getItem('token')) {
        setToken(JSON.parse(localStorage.getItem('token') ?? '{}') as Token, 'access')
    }
    if (localStorage.getItem('refreshToken')) {
        setToken(JSON.parse(localStorage.getItem('refreshToken') ?? '{}') as Token, 'refresh')
        refreshToken(tokens._refresh_token.content).then()
    }
}

export function removeToken() {
    tokens._access_token = undefined
    tokens._access_tokenBody = undefined
    tokens._refresh_token = undefined
    tokens._refresh_tokenBody = undefined
    localStorage.removeItem('token')
    localStorage.removeItem('refreshToken')
}

export function get(): string {
    if (!(tokens._access_token)) {
        throw new NoTokenAvailableError()
    }

    // implement refresh token logic here

    return tokens._access_token.content
}

export function getUid() {
    if (!(tokens._access_tokenBody)) {
        throw new NoTokenAvailableError()
    }

    return tokens._access_tokenBody?.uid
}

export function getBody(): TokenBody {
    if (!(tokens._access_token)) {
        throw new NoTokenAvailableError()
    }

    return tokens._access_tokenBody
}

function setToken(token: Token, tokenType: 'access' | 'refresh') {
    tokens[`_${tokenType}_token`] = token

    if (!token.content) {
        return
    }

    const base64Url = token.content.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    tokens[`_${tokenType}_tokenBody`] = JSON.parse(jsonPayload)
}

export function storeAccessToken(accessToken: Token) {
    setToken(accessToken, 'access')
    localStorage.setItem('token', JSON.stringify(accessToken))
}

export function storeRefreshToken(refreshToken: Token) {
    setToken(refreshToken, 'refresh')
    localStorage.setItem('refreshToken', JSON.stringify(refreshToken))
}

export function isLoggedIn(): boolean {
    return tokens._access_token !== undefined
}

init()


/**
 * Configure axios to refresh token when access token is expired
 */

axios.interceptors.response.use((response) => {
    return response
}, async function (error) {
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true;
        await refreshToken(tokens._refresh_token.content);
        originalRequest.headers['Authorization'] = 'Bearer ' + tokens._access_token.content;
        return axios(originalRequest);
    }
    return Promise.reject(error);
});
