export enum TokenTerm {
    // 1 minute
    VERY_SHORT = 'VERY_SHORT',
    // 20 minutes
    SHORT = 'SHORT',
    // 2 days
    MIDDLE = 'MIDDLE',
    // 60 days
    LONG = 'LONG',
    // 2 years
    VERY_LONG = 'VERY_LONG',
}

export interface Token {
    term: TokenTerm;
    content: string;
}

export interface AuthenticationRequest {
    username: string;
    password: string;
    term: TokenTerm;
}

export interface AuthenticationResponse {
    token: Token;
}
