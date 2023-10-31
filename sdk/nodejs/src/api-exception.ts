import {Code, Error as ApiBrewError} from "./model/extension";

export class ApiException extends Error {
    constructor(public code: Code = Code.INTERNAL_ERROR, message?: string) {
        super(code);
    }

    public static fromApiBrewError(error: ApiBrewError): ApiException {
        return new ApiException(error.code, error.message!);
    }

    public static fromError(error: Error): ApiException {
        return new ApiException(Code.UNKNOWN_ERROR, error.message);
    }

    public static fromCode(code: Code, message?: string): ApiException {
        return new ApiException(code, message);
    }

    public static fromMessage(message: string): ApiException {
        return new ApiException(Code.UNKNOWN_ERROR, message);
    }

}