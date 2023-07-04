export interface LambdaParams {
    name: string;
}
export interface LambdaEntity {
    id?: string;
    action: string;
}
export type LambdaFunctionDef<T extends LambdaEntity> = (element: T) => void;
export declare function defineLambda<T extends LambdaEntity>(name: string, eventSelectorPattern: string, fn: LambdaFunctionDef<T>): void;
export declare function fireLambda<T extends LambdaEntity>(trigger: string, element: Partial<T>): void;
