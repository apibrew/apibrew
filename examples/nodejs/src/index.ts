import axios from "axios";
import { test1, test2 } from "./functions.js";
import { defineFunction } from "@apibrew/client/dist/logic/function-def.js";

defineFunction({
    package: 'test',
    name: 'Test1',
}, test1)

defineFunction({
    package: 'test',
    name: 'Test2',
}, test2)

defineFunction({
    package: 'test',
    name: 'Test3',
    args: [{
        name: 'a'
    }, {
        name: 'b'
    }]
}, ({ a, b }) => {
    return a + b  + '33xx'
})