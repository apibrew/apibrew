import { TokenStorage } from '@apibrew/client/token-storage';
export declare class LocalStorageTokenStorage implements TokenStorage {
    private prefix?;
    constructor(prefix?: string | undefined);
    clear(): void;
    private getKey;
    get(name: string): string | undefined;
    set(name: string, token: string): void;
    list(): {
        name: string;
        token: string;
    }[];
}
