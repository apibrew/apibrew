import { initExtensions } from "./registrator";
import { load } from "./store";
import { reloadFunction, reloadInternal, reloadModules } from "./function-registry";
import express from 'express';
import { components } from "./model/base-schema";
import { handleFunctionExecutionCall } from './handler'
import { Event } from "@apibrew/client";


function init() {
    const promises = [
        initExtensions(),
        load('logic', 'Module'),
        load('logic', 'Function'),
        load('logic', 'FunctionTrigger'),
        load('logic', 'ResourceRule'),
    ]

    Promise.all(promises).then(() => {
        console.log('All resources loaded')
        reloadInternal()
    })
}

init()

const app = express()
const port = 23619

app.use(express.json({ limit: '5000mb' }));
app.use(express.urlencoded({ limit: '5000mb' }));

app.post('/call/function', (req, res) => {
    const event = req.body as components['schemas']['Event']

    handleFunctionExecutionCall(event).then((result) => {
        res.send(result)
    })
})

app.post('/reload', (req, res) => {
    console.log('trigger reload')
    const event = req.body as Event
    // init()

    switch (`${event.resource.namespace}/${event.resource.name}`) {
        case 'logic/Function':
            load('logic', 'Function').then(() => {
                reloadFunction()
            })
            break
        case 'logic/Module':
            load('logic', 'Module').then(() => {
                reloadModules()
            })
            break
        default:
            init()

    }

    res.send(req.body)
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})