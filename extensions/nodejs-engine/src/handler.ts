import { functionIdMap, functionMap, lambdaIdMap, registerFunction, registerFunctionTrigger, registerLambda, registerModule, registerResourceRule } from "./function-registry";
import { executeFunction, executeLambda, locateFunction } from "./function-execute";

import { Event, ResourceRule, ResourceRuleName, Lambda, LambdaResource, Function, Module } from "@apibrew/client"

import { functionRepository, functionTriggerRepository, lambdaRepository, moduleRepository, resourceRuleRpository } from "./client";

const { VM } = require('vm2');

export async function handleFunctionExecutionCall(event: Event) {
    for (const record of event.records) {
        try {
            const packageName = record.properties.function.package
            const name = record.properties.function.name
            const input = record.properties.input
            const fn = locateFunction(packageName, name)

            if (!fn) {
                throw new Error('Function not found')
            }

            const output = (await executeFunction(fn, input) ?? { ok: true });

            record.properties.output = output

            record.properties.status = 'success'
        } catch (e) {
            console.error(e)

            record.properties.error = e.message
            record.properties.status = 'error'
        }
    }

    return event;
}

export async function handleFunctionCall(event: Event) {
    console.log('trigger function', (event.records ?? [{}])[0]['id'])

    return await handleFunctionExecutionCall(event)
}

export async function handleLambdaCall(event: Event, lambdaId: string) {
    console.log('trigger lambda', lambdaId)

    const lambda = lambdaIdMap[lambdaId]

    if (!lambda) {
        throw new Error('Lambda not found')
        return
    }

    for (const record of event.records) {
        try {
            await executeLambda(lambda, record.properties)
        } catch (e) {
            console.error(e)
        }
    }

}

export async function handleReload(event: Event) {
    console.log('trigger reload', event.resource.namespace, event.resource.name, (event.records ?? [{}])[0]['id'])

    for (const record of event.records ?? []) {
        switch (`${event.resource.namespace}/${event.resource.name}`) {
            case 'logic/Function':
                const fn = await functionRepository.get(record.id!)
                await registerFunction(fn)

                console.log('reloaded function', fn.id)
                break
            case 'logic/Module':
                const module = await moduleRepository.get(record.id!)
                await registerModule(module)

                console.log('reloaded module', module.id)
                break
            case 'logic/Lambda':
                const lambda = await lambdaRepository.get(record.id!)
                await registerLambda(lambda)

                console.log('reloaded lambda', lambda.id)
                break
            case 'logic/FunctionTrigger':
                const functionTrigger = await functionTriggerRepository.get(record.id!)
                await registerFunctionTrigger(functionTrigger)

                console.log('reloaded functionTrigger', functionTrigger.id)
                break
            case 'logic/ResourceRule':
                const resourceRule = await resourceRuleRpository.get(record.id!)
                await registerResourceRule(resourceRule)

                console.log('reloaded resourceRule', resourceRule.id)
                break
        }
    }
}