import {DhClient, Entity} from "../dh-client/client";

const client = new DhClient({
    Addr: "127.0.0.1:9009",
    Insecure: true,
})

interface Country extends Entity<Country> {
    id?: string
    name?: string
    description?: string
}

async function run() {
    await client.authenticateWithUsernameAndPassword("admin", "admin")
    const repo = client.newRepository<Country>("default", "country")

    const extension = client.NewExtensionService("127.0.0.1", 17686)

    await extension.run()

    repo.extend(extension).onCreate(async (entity: Country) => {
        console.log(entity)
        entity.description = 'Updated desc 123'
        return entity
    })

    repo.extend(extension).onUpdate(async (entity: Country) => {
        console.log(entity)
        entity.description = entity.description + ' Updated desc 123'
        return entity
    }, false)
}

run().then()