export interface ActionComponent<R> {
    execute: (...args: any) => Promise<R>;
}
