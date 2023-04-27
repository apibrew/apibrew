import axios from 'axios'
import { BACKEND_URL } from '../config'
import { type AuthenticationRequest, type AuthenticationResponse, type Token } from '../model'

export async function authenticate(username: string, password: string): Promise<Token> {
    const request: AuthenticationRequest = {
        username,
        password,
        term: 'LONG'
    }

    const result = await axios.post<AuthenticationResponse>(`${BACKEND_URL}/authentication/token`, request)

    return result.data.token
}
