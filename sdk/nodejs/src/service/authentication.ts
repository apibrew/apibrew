import axios from 'axios'

import { type AuthenticationResponse, RenewTokenRequest, RenewTokenResponse, TokenTerm, Token } from '../model'
import { ServiceConfig } from './config'

export async function authenticate(config: ServiceConfig, username: string, password: string, tokenTerm: TokenTerm): Promise<Token> {
    const result = await axios.post<AuthenticationResponse>(`${config.backendUrl}/authentication/token`, {
        username,
        password,
        term: tokenTerm
    })

    return result.data.token!
}

export async function refreshToken(config: ServiceConfig, refreshTokenContent: string): Promise<void> {
    const accessTokenResult = await axios.put<RenewTokenResponse>(`${config.backendUrl}/authentication/token`, {
        token: refreshTokenContent,
        term: 'VERY_SHORT'
    } as RenewTokenRequest)
}
