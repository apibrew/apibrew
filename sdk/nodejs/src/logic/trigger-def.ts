import { Client } from "../client";
import { FunctionResource, FunctionTrigger, FunctionTriggerResource, Record, Function } from "../model";
import { handleError } from "../service/error";
import { getModule, registerModuleChild } from "./module-def";

export function defineTrigger<T extends Record<unknown>>(functionTrigger: FunctionTrigger, fn: (entity: T) => T) {
    const client = Client.getDefaultClient()

    const functionRepository = client.newRepository(FunctionResource)
    const triggerRepository = client.newRepository(FunctionTriggerResource)

    const module = getModule()

    async function createTrigger() {
        await functionRepository.apply({
            package: module.package,
            name: 'Trigger_' + functionTrigger.name,
            args: [{
                name: 'element'
            }],
            module: {
                id: module.id,
            },
            engine: {
                name: 'nodejs-engine'
            }
        } as Function).then(resp => {
            console.log(resp)
        }, err => {
            console.error(handleError(err))
        })

        await triggerRepository.apply({
            ...functionTrigger,
            function: {
                package: module.package,
                name: 'Trigger_' + functionTrigger.name,
            },
        } as FunctionTrigger).then(resp => {
            console.log(resp)
        }, err => {
            console.error(handleError(err))
        })
    }

    createTrigger()

    registerModuleChild('Trigger_' + functionTrigger.name, fn)
}

