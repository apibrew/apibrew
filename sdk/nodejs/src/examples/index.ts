import { LogicDef } from '..';
import { Client } from '../client';
import { FunctionParams, defineFunction, getModule, setModule } from '../logic';
import { Module } from '../model';

async function run() {
    const client = new Client('http://localhost:9009')
    await client.authenticate('admin', 'admin')
    Client.setDefaultClient(client)

    setModule({
        id: 'a0ff7cce-16ae-11ee-8665-c6aac64f19b2',
        package: 'Test'
    } as Module)

    defineFunction<number>('test', [], (params: FunctionParams) => {
        console.log('I am running')

        return 123
    })




    LogicDef.defineFunction('Test3', ['a', 'b'], ({ a, b }) => {
        return a + b + '33xx'
    })

    LogicDef.defineResource({
        name: 'SimpleEventObject',
        properties: [{
            name: 'action',
            type: 'ENUM',
            enumValues: [
                'AcceptPayment',
                'RejectPayment',
                'CancelPayment',
                'RefundPayment',
            ] as any
        }]
    })

    LogicDef.defineLambda('TestLambda', 'SimpleEventObject:AcceptPayment', (element) => {
        LogicDef.fireLambda('SimpleEventObject:RejectPayment', {})
    })

    LogicDef.defineLambda('TestLambda', 'SimpleEventObject:RejectPayment', (element) => {
        LogicDef.callFunction(getModule().package, 'CallBackForLambda', { a: 1, b: 2 })
    })

    LogicDef.defineFunction('TriggerLambda', [], ({ a, b }) => {

        LogicDef.fireLambda('SimpleEventObject:AcceptPayment', {})

        return 'ok'
    })

    LogicDef.defineFunction('CallBackForLambda', ['a', 'b'], ({ a, b }) => {

        console.log('CallBackForLambda called', a, b)

        return 'ok'
    })




}

run()