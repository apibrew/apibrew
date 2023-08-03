import { RenewTokenResponse, TokenTerm, Token } from '../model';
import { ServiceConfig } from './config';
export declare function authenticate(config: ServiceConfig, username: string, password: string, tokenTerm: TokenTerm): Promise<Token>;
export declare function refreshToken(config: ServiceConfig, refreshTokenContent: string): Promise<RenewTokenResponse>;
