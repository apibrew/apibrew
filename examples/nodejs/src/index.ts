import axios from "axios";
import { test1, test2 } from "./functions.js";
import { LogicDef } from "@apibrew/client";
import { getModule } from "@apibrew/client/dist/logic/module-def.js";

LogicDef.defineFunction('Test1', [], test2)

LogicDef.defineFunction('Test2', [], test2)

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