import {Pet} from "./schema";
import {DhClient} from "data-handler-client";

const client = new DhClient({
    Addr: "127.0.0.1:9009", // data handler server address
    Insecure: true,
})

async function run() {
    await client.authenticateWithUsernameAndPassword("admin", "admin")
    const petRepo = client.newRepository<Pet>("default", "pet")

    const extension = client.NewExtensionService("127.0.0.1", 17686, "http://host.docker.internal:17686") // which port we will run extension

    await extension.run()

    const petExtension = petRepo.extend(extension)

    petExtension.onCreate(async (pet) => {
        if (!pet.description) {
            pet.description = pet.name + ' Pet'
        }

        return pet
    })
}

run()