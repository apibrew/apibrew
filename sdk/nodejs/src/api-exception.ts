import {Code, Error as ApiBrewError, ErrorField} from "./model/extension";

export class ApiException extends Error {
    constructor(public code: Code = Code.INTERNAL_ERROR, message?: string, public fields?: ErrorField[]) {
        super(message);
    }

    public static fromApiBrewError(error: ApiBrewError): ApiException {
        return new ApiException(error.code, error.message!, error.fields);
    }

    public static fromError(error: any): ApiException {
        if (error.code) {
            return ApiException.fromApiBrewError(error)
        } else {
            return new ApiException(Code.UNKNOWN_ERROR, error.message);
        }
    }

}