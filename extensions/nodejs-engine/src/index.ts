import { loadAll } from "./function-registry";
import express from 'express';

import { handleFunctionCall, handleLambdaCall, handleReload } from './handler'
import { Event } from "@apibrew/client";
import { initExtensions } from "./registrator";
import { AxiosError } from "axios";

process.on('uncaughtException', (err) => {
    console.error('Asynchronous error caught.', err);
});


async function init() {
    await initExtensions()
    await loadAll()

    console.log('ready')
}

init()

const app = express()
const port = 23619

app.use(express.json({ limit: '5000mb' }));
app.use(express.urlencoded({ limit: '5000mb' }));

app.post('/call/function', async (req, res) => {
    const result = await handleFunctionCall(req.body as Event)

    res.send(result)
})

app.post('/call/lambda/:id', async (req, res) => {
    const event = req.body as Event
    const result = await handleLambdaCall(event, req.params.id)

    res.send(event)
})

app.post('/reload', async (req, res) => {
    handleReload(req.body as Event).then(() => {
        console.log('reloaded')
    }, console.error)

    res.send(req.body)
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})