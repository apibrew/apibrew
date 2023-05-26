export interface ActionComponent<R> {
    execute: (...args: any) => R
}
