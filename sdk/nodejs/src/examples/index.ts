import { Client } from '../client';
import { FunctionParams, defineFunction, setModule } from '../logic';
import { Module } from '../model';

async function run() {
    const client = new Client('http://localhost:9009')
    await client.authenticate('admin', 'admin')
    Client.setDefaultClient(client)

    setModule({
        id: 'a0ff7cce-16ae-11ee-8665-c6aac64f19b2'
    } as Module)

    defineFunction<number>('test', [], (params: FunctionParams) => {
        console.log('I am running')

        return 123
    })
}

run()