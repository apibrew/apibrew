import {newClient} from "../client";
import {Book, BookEntityInfo} from "./model/book";
import {beforeCreate, execute, PollerExtensionService} from "../ext";

export async function run() {
    const client = await newClient("local")

    const repository = client.repository(BookEntityInfo)

    const extensionService = new PollerExtensionService("test-service-name", client, "test-channel-key")

    extensionService.handler(BookEntityInfo).when(beforeCreate()).operate(execute((event, entity) => {
        console.log('inside beforeCreate', event, entity)
    }))


    extensionService.run();

}

run();