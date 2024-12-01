import React from 'react';
import { ApiException } from '@apibrew/client';
export interface ErrorHandlerContextProps {
    onAuthenticationFailed(err: ApiException): void;
}
export declare type CatchArg = Parameters<Promise<any>['catch']>[0];
export declare const ErrorHandlerContext: React.Context<((reason: any) => unknown) | null | undefined>;
export declare const useErrorHandler: () => ((reason: any) => unknown) | null | undefined;
