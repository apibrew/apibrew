export interface TokenStorage {
    get(name: string): string | undefined;

    set(name: string, token: string): void;

    list(): {name: string, token: string}[];

    clear(): void;
}
