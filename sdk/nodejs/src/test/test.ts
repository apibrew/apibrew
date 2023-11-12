import {Book, BookEntityInfo} from "./model/book";
import {beforeCreate, execute, PollerExtensionService} from "../ext";
import {newClientByServerConfig} from "../client";
import {ConfigLoader} from "../config-loader";

export async function run() {
    const client = await newClientByServerConfig(ConfigLoader.loadServerConfig("local"))

    const repository = client.repository(BookEntityInfo)

    const extensionService = new PollerExtensionService("test-service-name", client, "test-channel-key")

    extensionService.handler(BookEntityInfo).when(beforeCreate()).operate(execute((event, entity) => {
        console.log('inside beforeCreate', event, entity)
    }))


    extensionService.run();

}

run();