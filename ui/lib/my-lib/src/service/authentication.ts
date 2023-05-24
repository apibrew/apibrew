import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type AuthenticationRequest, type AuthenticationResponse } from '../model'
import { TokenService } from './token'

export async function authenticate(username: string, password: string): Promise<void> {
    const request: AuthenticationRequest = {
        username,
        password,
        term: 'LONG'
    }

    const result = await axios.post<AuthenticationResponse>(`${BACKEND_URL}/authentication/token`, request)

    TokenService.storeToken(result.data.token)
}
