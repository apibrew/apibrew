import { Argument } from '../model/logic/function';
export interface FunctionParams {
    [key: string]: any;
}
export type FunctionDef<T> = (params: FunctionParams) => T;
export interface FunctionProps {
    package: string;
    name: string;
    args?: Argument[];
}
export declare function defineFunction<T>(name: string, args: string[], fn: FunctionDef<T>): void;
export declare function callFunction<T = any, R = any>(fnPackage: string, fnName: string, params: T): Promise<R>;
