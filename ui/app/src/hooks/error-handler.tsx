import { useContext } from "react";
import { AxiosError } from "axios";
import { Status } from "../model";
import { useNavigate } from "react-router-dom";
import { TokenService, LayoutContext } from "@apibrew/ui-lib";

export const useErrorHandler = () => {
    const layoutCtx = useContext(LayoutContext);
    const navigate = useNavigate()

    if (!layoutCtx) {
        throw new Error('LayoutContext not found');
    }

    return (err: Error) => {
        if (err instanceof TokenService.NoTokenAvailableError) {
            navigate('/login')
            return
        } else if (err instanceof AxiosError) {
            if (err.response?.status === 401) {
                // TokenService.removeToken()
                layoutCtx.showAlert({
                    severity: 'error',
                    message: 'You are not authorized to access this resource'
                })
                navigate('/login')
                return
            }
            if (err.response?.status === 403) {
                // navigate('/dashboard/')
                layoutCtx.showAlert({
                    severity: 'error',
                    message: 'You are not authorized to access this resource'
                })
                return
            }

            const responseData = err.response?.data as Status

            if (responseData.message) {
                layoutCtx.showAlert({
                    severity: 'error',
                    message: responseData.message
                })
            }
        } else {
            layoutCtx.showAlert({
                severity: 'error',
                message: err.message
            })
        }
    }
}