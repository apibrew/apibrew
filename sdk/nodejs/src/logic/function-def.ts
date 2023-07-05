import { Client } from '../client';
import { FunctionExecution, FunctionExecutionResource } from '../model';
import { Argument, Function, FunctionArgsName, FunctionResource } from '../model/logic/function'
import { getModule, registerModuleChild } from './module-def';

export interface FunctionParams {
    [key: string]: any
}

export type FunctionDef<T> = (params: FunctionParams) => T

export interface FunctionProps {
    package: string;
    name: string;
    args?: Argument[];
}

export function defineFunction<T>(name: string, args: string[], fn: FunctionDef<T>) {
    const client = Client.getDefaultClient()

    const functionRepository = client.newRepository(FunctionResource)

    const module = getModule()

    functionRepository.apply({
        package: module.package,
        name: name,
        args: args.map(arg => {
            return {
                name: arg
            }
        }),
        module: {
            id: module.id,
        },
        engine: {
            name: 'nodejs-engine'
        }
    } as Function).then(resp => {
        console.log(resp)
    }, err => {
        console.error(err)
    })

    registerModuleChild(name, fn)
}

export async function callFunction<T = any, R = any>(fnPackage: string, fnName: string, params: T): Promise<R> {
    const client = Client.getDefaultClient()

    const functionRepository = client.newRepository<FunctionExecution>(FunctionExecutionResource)

    const result = await functionRepository.create({
        id: '',
        version: 1,
        function: {
            package: fnPackage,
            name: fnName,
        } as Function,
        input: params,
    } as FunctionExecution)

    if (result.error) {
        throw new Error(result.error as any)
    }

    return result.output as R
}
