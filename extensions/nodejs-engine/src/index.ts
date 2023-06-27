import { initExtensions } from "./registrator";
import { handle } from "./handler";
import { load } from "./store";
import { reloadInternal } from "./function-registry";
import express from 'express';
import { components } from "./model/base-schema";
import { handleFunctionExecutionCall } from './handler'


const promises = [
    initExtensions(),
    load('logic', 'Function'),
    load('logic', 'FunctionTrigger'),
    load('logic', 'ResourceRule'),
]

Promise.all(promises).then(() => {
    console.log('All resources loaded')
    reloadInternal()
})

const app = express()
const port = 23619

app.use(express.json());

app.post('/call/function', (req, res) => {
    const event = req.body as components['schemas']['Event']

    handleFunctionExecutionCall(event).then((result) => {
        res.send(result)
    })
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})