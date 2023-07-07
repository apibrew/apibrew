import { AxiosError } from "axios"
import { Client } from "../client"
import { Function, FunctionResource, RecordResourceInfo } from "../model"
import { Lambda, LambdaResource } from "../model/logic/lambda"
import { getModule, registerModuleChild } from "./module-def"
import { handleError } from "../service/error"

export interface LambdaParams {
    name: string
}

export interface LambdaEntity {
    id?: string
    action: string
}

function parseLambdaEventSelectorPattern(eventSelectorPattern: string) {
    const parts = eventSelectorPattern.split(':')

    const resourceFullName = parts[0]
    const action = parts[1]
    const subParts = resourceFullName.split('/')
    let resourceName;
    let resourceNamespace;

    if (subParts.length === 1) {
        resourceName = subParts[0]
        resourceNamespace = 'default'
    } else if (subParts.length === 2) {
        resourceNamespace = subParts[0]
        resourceName = subParts[1]
    } else {
        throw new Error('Invalid resource name: ' + resourceFullName)
    }

    return {
        resourceFullName,
        resourceName,
        resourceNamespace,
        action
    }
}

export type LambdaFunctionDef<T extends LambdaEntity> = (element: T) => void

export function defineLambda<T extends LambdaEntity>(name: string, eventSelectorPattern: string, fn: LambdaFunctionDef<T>) {
    const client = Client.getDefaultClient()

    const functionRepository = client.newRepository(FunctionResource)
    const lambdaRepository = client.newRepository(LambdaResource)

    const module = getModule()

    async function createLambda() {
        console.log('before create lambda')
        await functionRepository.apply({
            package: module.package,
            name: 'Lambda_' + name,
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

        console.log('after create lambda')

        await lambdaRepository.apply({
            package: module.package,
            name: 'Lambda_' + name,
            eventSelectorPattern: eventSelectorPattern,
            function: {
                package: module.package,
                name: 'Lambda_' + name,
            },
        } as Lambda).then(resp => {
            console.log(resp)
        }, err => {
            console.error(handleError(err))
        })

        console.log('after create lambda 2')
    }

    createLambda()

    registerModuleChild('Lambda_' + name, fn)
}

export function fireLambda<T extends LambdaEntity>(trigger: string, element: Partial<T>) {
    const module = getModule()

    const client = Client.getDefaultClient()

    const parsed = parseLambdaEventSelectorPattern(trigger)

    const repository = client.newRepository({
        namespace: parsed.resourceNamespace,
        resource: parsed.resourceName,
    })

    element.action = parsed.action

    repository.create(element).then(resp => {
        console.log('Lambda ' + trigger + ' fired')
    }, err => {
        console.error(handleError(err))
    })
}
