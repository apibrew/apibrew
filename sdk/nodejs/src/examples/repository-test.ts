import { Repository } from "../repository";
import { AuthenticationService } from "../service";
import { CountryExample, CountryExampleResource } from "./country-example";

export async function run() {
    const config = {
        backendUrl: 'http://localhost:9009',
        token: ''
    }

    const configProvider = () => config

    config.token = await AuthenticationService.authenticate(config, 'admin', 'admin', 'LONG').then(result => result.content)

    const repository = new Repository<CountryExample>(configProvider, CountryExampleResource)

    const created = await repository.apply({
        name: 'Germany',
        description: 'Germany 123',
    } as CountryExample)

    const list = await repository.list()

}

run()
