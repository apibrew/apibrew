import {Code, Error as ApiBrewError} from "./model/extension";

export class ApiException extends Error {
    constructor(public code: Code = Code.INTERNAL_ERROR, message?: string) {
        super(message);
    }

    public static fromApiBrewError(error: ApiBrewError): ApiException {
        return new ApiException(error.code, error.message!);
    }

    public static fromError(error: any): ApiException {
        if (error.code) {
            return new ApiException(error.code, error.message);
        } else {
            return new ApiException(Code.UNKNOWN_ERROR, error.message);
        }
    }

}