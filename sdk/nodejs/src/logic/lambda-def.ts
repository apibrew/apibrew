export interface LambdaParams {
    name: string
}

export type LambdaFunctionDef<T> = (element: T) => T | undefined

export function defineLambda<T>(params: LambdaParams, fn: LambdaFunctionDef<T>) {
    exports.Test1 = fn
}
