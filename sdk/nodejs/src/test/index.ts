import {DhClient, Entity} from "../dh-client/client";
import * as dependency_1 from "../dh-client/google/protobuf/struct";

const client = new DhClient({
    Addr: "tisserv.net:9009",
    Insecure: true,
})

class Country implements Entity<Country> {
    id?: string
    name?: string
    description?: string

    equals(other: Country): boolean {
        if (other.id != this.id) {
            return false
        }
        if (other.name != this.name) {
            return false
        }

        if (other.description != this.description) {
            return false
        }

        return true
    }

    getNamespace(): string {
        return "default";
    }

    getResourceName(): string {
        return "country";
    }

    same(other: Country): boolean {
        if (other.id == this.id) {
            return true
        }
        if (other.name == this.name) {
            return true
        }

        return false
    }

    fromProperties(properties: Map<string, dependency_1.Value>): void {
        // set properties from record
        this.id = properties.get("id")?.stringValue
        this.name = properties.get("name")?.stringValue
        this.description = properties.get("description")?.stringValue
    }

    toProperties(): Map<string, dependency_1.Value> {
        const properties = new Map<string, dependency_1.Value>()

        properties.set("id", dependency_1.Value.fromObject({
            stringValue: this.id
        }))

        properties.set("name", dependency_1.Value.fromObject({
            stringValue: this.name
        }))

        properties.set("description", dependency_1.Value.fromObject({
            stringValue: this.description
        }))

        return properties
    }
}

async function run() {
    await client.authenticateWithUsernameAndPassword("admin", "admin")
    const repo = client.newRepository(Country)

    repo.find({}).then((result) => {
      console.log(result)
    })

    const extension = client.NewExtensionService("127.0.0.1:7691")

    await extension.run()

    // const country = new Country()
    // country.name = "India2"
    // country.description = "A country in Asia"
    //
    // repo.create(country).then((result) => {
    //     console.log(result)
    // })
}

run()