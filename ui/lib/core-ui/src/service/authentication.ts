import axios from 'axios'
import {BACKEND_URL} from '../config'
import {type AuthenticationResponse, RenewTokenRequest, RenewTokenResponse} from '../model'
import * as TokenService from './token'

export async function authenticate(username: string, password: string): Promise<void> {

    const accessTokenResult = await axios.post<AuthenticationResponse>(`${BACKEND_URL}/authentication/token`, {
        username,
        password,
        term: 'VERY_SHORT'
    })

    const refreshTokenResult = await axios.post<AuthenticationResponse>(`${BACKEND_URL}/authentication/token`, {
        username,
        password,
        term: 'LONG'
    })

    TokenService.storeAccessToken(accessTokenResult.data.token)
    TokenService.storeRefreshToken(refreshTokenResult.data.token)
}

export async function refreshToken(refreshTokenContent: string): Promise<void> {
    const accessTokenResult = await axios.put<RenewTokenResponse>(`${BACKEND_URL}/authentication/token`, {
        token: refreshTokenContent,
        term: 'VERY_SHORT'
    } as RenewTokenRequest)

    TokenService.storeAccessToken(accessTokenResult.data.token)
}
