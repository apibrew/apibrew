import { AxiosError } from "axios";

export function handleError(error: Error): Error {
    if (error instanceof AxiosError) {
        console.error('AxiosError:', error.request.url, error.request.body, error.code, error.response?.data);
        return error
    } else {
        return error
    }
}