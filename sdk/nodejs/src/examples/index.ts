import { Client } from '../client';
import { FunctionParams, defineFunction, setModuleId } from '../logic';

async function run() {
    const client = new Client('http://localhost:9009')
    await client.authenticate('admin', 'admin')
    Client.setDefaultClient(client)

    setModuleId('a0ff7cce-16ae-11ee-8665-c6aac64f19b2')

    defineFunction<number>({
        package: 'test',
        name: 'Test1',
    }, (params: FunctionParams) => {
        console.log('I am running')

        return 123
    })
}

run()