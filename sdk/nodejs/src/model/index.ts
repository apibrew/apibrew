export interface Token {
    /**
     * Format: enum
     * @enum {string}
     */
    term: TokenTerm;
    content: string;
    /** Format: date-time */
    expiration: string;
}

export type TokenTerm = "VERY_SHORT" | "SHORT" | "MIDDLE" | "LONG" | "VERY_LONG"

export interface AuthenticationRequest {
    username: string;
    password: string;
    /**
     * Format: enum
     * @enum {string}
     */
    term: TokenTerm;
}

export interface AuthenticationResponse {
    token?: Token;
}

export interface RenewTokenRequest {
    token: string;
    /**
     * Format: enum
     * @enum {string}
     */
    term: TokenTerm;
}

export interface RenewTokenResponse {
    token?: Token;
}



export * from './annotations'
export * from './record'

export * from './logic/function'
export * from './logic/function-execution'
export * from './logic/function-execution-engine'
export * from './logic/function-trigger'
export * from './logic/logic-code'
export * from './logic/module'
export * from './logic/resource-rule'
export * from './logic/schedule'
export * from './logic/lambda'

export * from './system'

