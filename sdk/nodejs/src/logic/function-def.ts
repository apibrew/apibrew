import { Client } from '../client';
import { Argument, Function, FunctionArgsName, FunctionResource } from '../model/logic/function'
import { getModuleId, registerModuleChild } from './module-def';

export interface FunctionParams {
    [key: string]: any
}

export type FunctionDef<T> = (params: FunctionParams) => T

export interface FunctionProps {
    package: string;
    name: string;
    args?: Argument[];
}

export function defineFunction<T>(funcProps: Partial<Function>, fn: FunctionDef<T>) {
    const client = Client.getDefaultClient()

    const functionRepository = client.newRepository(FunctionResource)

    functionRepository.apply({
        ...funcProps,
        module: {
            id: getModuleId()
        },
        engine: {
            name: 'nodejs-engine'
        }
    } as Function).then(resp => {
        console.log(resp)
    }, err => {
        console.error(err)
    })

    registerModuleChild(funcProps.name!, fn)
}
