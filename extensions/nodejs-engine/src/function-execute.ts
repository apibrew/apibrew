import { NodeVM, VM, VMError, VMFileSystem, makeResolverFromLegacyOptions } from "vm2"
import { functionMap, functionNameIdMap } from "./function-registry"
import { Function } from './model/function'
import axios from 'axios'
import * as tar from 'tar-stream'
import { PassThrough, Stream } from 'stream'
import { PathOrFileDescriptor, statSync } from "fs"
import { FN_DIR } from "./config"

export function locateFunction(packageName: string, name: string): Function {
    return functionMap[packageName + '/' + name]
}

export async function executeFunction<R>(fn: Function, params: object): Promise<R> {
    let fnId = fn.id

    if (!fnId) {
        fnId = functionNameIdMap[fn.package + '/' + fn.name]
    }

    if (!fnId) {
        throw new Error('Function codes not found')
    }

    const vm = new NodeVM({
        sandbox: {
            fn: fn,
            ...params,
            params: params,
            exports: exports
        },
        console: 'inherit',
        require: {
            external: true,
            builtin: ['*'],
        },
        allowAsync: true,
        timeout: 1000,
    })

    const result = await (vm.runFile(FN_DIR + `/${fnId}.js`).result as Promise<R>)

    return result
}
