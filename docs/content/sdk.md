SDK
====

Introduction

Data handler SDK is a library that provides a set of functions to interact with the data handler.

Javascript SDK

# Prerequisites

```bash
npm install data-handler-client
npm install typescript --save-dev
```

Let's define out resources

country.yml

```yaml
type: resource
name: country
properties:
  - name: name
    type: STRING
    length: 255
    required: true
    unique: true
  - name: description
    type: STRING
    length: 255
```

Let's apply our resources

```bash
dhctl apply -f country.yml
```

Swagger: http://localhost:9009/docs/index.html

Now, we can generate our typescript schema

```bash
dhctl generate -p=src --platform=nodejs
```

It will generate following code for you

```typescript
export interface Country extends Entity {
    id: string;
    name: string;
    description: string;
    version: number;
}
```

Now, we can use our generated typescript schema to interact with the data handler

let's setup our client and repository

```typescript
import {DataHandlerClient} from 'data-handler-client';
import {Country} from './schema';

async function run() {
    const client = new DataHandlerClient('http://localhost:9009');

    await client.authenticateWithUsernameAndPassword("admin", "admin")

    const countryRepository = client.newRepository<Country>("default", "country")
}

run()
```

Now we can use our repository to do crud operations on our resources

```typescript
 const country = new Country();

country.name = "Turkey";
country.description = "A country in Asia";

await countryRepository.save(country)
```

We can use our repository to extend country resource

```typescript
const extension = client.NewExtensionService("127.0.0.1", 17686)
await extension.run()

const countryExtension = countryRepository.extend(extension)

countryExtension.onCreate(async (order) => {
    country.description = country.description + " - extended"

    return order
})
```

So, we are ready, we modified behaviour of our country resource. Let's test it

```bash
curl -X POST "http://localhost:9009/country" -H "Authorization: Bearer <token>" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"name\": \"Turkey\", \"description\": \"A country in Asia\"}"

curl -X GET "http://localhost:9009/country" -H "Authorization: Bearer <token>" -H "accept: application/json"
```

We will get following response

```json
{
  "content": [
    {
      "id": "5f9b9b0e-9b9b-4b9b-9b0e-9b9b9b9b9b9b",
      "name": "Turkey",
      "description": "A country in Asia - extended",
      "version": 1
    }
  ]
}

```

# Tutorials

- [Store app](https://github.com/tislib/data-handler/tree/master/examples/store-app/Readme.md)



