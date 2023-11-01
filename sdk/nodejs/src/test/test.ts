import {StorageServiceImpl} from "../storage/impl/storage-service-impl";
import {StorageObject} from "../storage/model/storage-object";
import {newClient} from "../client";

export async function run() {
    const client = await newClient("local")

    const storageService = new StorageServiceImpl(client, "http://localhost:8080/local")

    const object = await storageService.repository().create({} as StorageObject)

    await storageService.uploadBytes(object.id, Buffer.from("Hello World 321"), "hello.txt")


    console.log((await storageService.downloadBytes(object.id)).toString())
}

run();