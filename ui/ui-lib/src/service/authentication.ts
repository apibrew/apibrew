import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type AuthenticationResponse, RenewTokenRequest, RenewTokenResponse } from '../model'
import * as TokenService from './token'
import { AuthenticationApi } from '@apibrew/client'
import { ServiceConfigProviderWithoutToken } from './service-config'

export async function authenticate(username: string, password: string): Promise<void> {
    const accessTokenResult = await AuthenticationApi.authenticate(ServiceConfigProviderWithoutToken(), username, password, 'VERY_SHORT');

    const refreshTokenResult = await AuthenticationApi.authenticate(ServiceConfigProviderWithoutToken(), username, password, 'LONG');

    TokenService.storeAccessToken(accessTokenResult)
    TokenService.storeRefreshToken(refreshTokenResult)
}

export async function refreshToken(refreshTokenContent: string): Promise<void> {
    const accessTokenResult = await AuthenticationApi.refreshToken(ServiceConfigProviderWithoutToken(), refreshTokenContent);

    TokenService.storeAccessToken(accessTokenResult.token)
}
