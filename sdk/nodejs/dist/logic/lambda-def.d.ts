export interface LambdaParams {
    name: string;
}
export type LambdaFunctionDef<T> = (element: T) => T | undefined;
export declare function defineLambda<T>(params: LambdaParams, fn: LambdaFunctionDef<T>): void;
