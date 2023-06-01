import {useContext} from "react";
import {AlertOptions, LayoutContext} from "../context/layout-context.ts";
import {AxiosError} from "axios";
import {Resource, Status} from "../model";

export const useErrorHandler = () => {
    const layoutCtx = useContext(LayoutContext);

    if (!layoutCtx) {
        throw new Error('LayoutContext not found');
    }

    return (err) => {
        console.error(err)

        if (err instanceof AxiosError) {
            const responseData = err.response.data as Status

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