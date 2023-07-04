import { Argument, Function } from '../model/logic/function';
export interface FunctionParams {
    [key: string]: any;
}
export type FunctionDef<T> = (params: FunctionParams) => T;
export interface FunctionProps {
    package: string;
    name: string;
    args?: Argument[];
}
export declare function defineFunction<T>(funcProps: Partial<Function>, fn: FunctionDef<T>): void;
