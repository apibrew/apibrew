import { VM } from "vm2"
import { functionMap } from "./function-registry"
import { Function } from './model/function'
import axios from 'axios'

export function locateFunction(packageName: string, name: string): Function {
    return functionMap[packageName + '/' + name]
}

export function requireFn(dependency: string) {
    switch (dependency) {
        case 'axios':
            return axios
        default:
            throw new Error('Unknown dependency: ' + dependency)
    }
}

export async function executeFunction<R>(fn: Function, params: object): Promise<R> {
    const exports: {
        result?: R,
    } = {}

    new VM({
        sandbox: {
            fn: fn,
            ...params,
            params: params,
            exports: exports,
            require: requireFn,
        },
        timeout: 1000,
    }).run(`let result = (function () {
                ${fn.script}
            })();
            
            if (fn.startFunction && exports[fn.startFunction]) {
                result = exports[fn.startFunction](params)
            }

            if (!exports.result) {
                exports.result = result
            }
            `)

    return exports.result
}