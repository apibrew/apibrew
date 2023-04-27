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

/*
export function login(username, password, navigate) {
    let object = {
        'Username': username,
        'Password': password,
        "term": 3
    }
    axios.post('http://tiswork.tisserv.net:9009/authentication/token',
     object).then((response) => {
        const token = response.data.token.content
        localStorage.setItem('token', token)
        navigate('/layout')
    }, err => {
        alert('username or password is incorrect')
    })
}

*/
