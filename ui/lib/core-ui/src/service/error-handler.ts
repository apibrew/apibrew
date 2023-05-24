import { AxiosError } from 'axios'
import { TokenService } from './token'

export async function handleError<T>(e: unknown): Promise<T> {
    if (e instanceof AxiosError) {
        if (e.response?.status === 401) {
            TokenService.removeToken()
            window.location.reload()
        }
    }

    return await Promise.reject<T>(e)
}
